package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	nlog "log"
	"strings"
	"time"

	"github.com/herdius/herdius-core/aws/restore"
	"github.com/herdius/herdius-core/blockchain"
	"github.com/herdius/herdius-core/blockchain/protobuf"
	blockProtobuf "github.com/herdius/herdius-core/blockchain/protobuf"
	"github.com/herdius/herdius-core/config"
	cryptoAmino "github.com/herdius/herdius-core/crypto/encoding/amino"
	"github.com/herdius/herdius-core/hbi/message"
	protoplugin "github.com/herdius/herdius-core/hbi/protobuf"
	cmn "github.com/herdius/herdius-core/libs/common"
	"github.com/herdius/herdius-core/p2p/crypto"
	keystore "github.com/herdius/herdius-core/p2p/key"
	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/herdius/herdius-core/p2p/network/discovery"
	"github.com/herdius/herdius-core/p2p/types/opcode"
	external "github.com/herdius/herdius-core/storage/exbalance"
	syncer "github.com/herdius/herdius-core/syncer"
	"github.com/herdius/herdius-core/types"

	sup "github.com/herdius/herdius-core/supervisor/service"
	amino "github.com/tendermint/go-amino"
)

var (
	cdc        = amino.NewCodec()
	nodeKeydir = "./cmd/testdata/secp205k1Accts/"
)

var (
	supsvc         *sup.Supervisor
	blockchainSvc  *blockchain.Service
	accountStorage external.BalanceStorage
)

// HerdiusMessagePlugin will receive all transmitted messages.
type HerdiusMessagePlugin struct{ *network.Plugin }

func init() {
	nlog.SetFlags(nlog.LstdFlags | nlog.Lshortfile)
	supsvc = &sup.Supervisor{}
	supsvc.SetWriteMutex()
	supsvc.ValidatorChildblock = make(map[string]*blockProtobuf.BlockID, 0)
	supsvc.ChildBlock = make([]*blockProtobuf.ChildBlock, 0)
	supsvc.VoteInfoData = make(map[string][]*blockProtobuf.VoteInfo, 0)

	// Register Amino service for message (en/de) coding
	cryptoAmino.RegisterAmino(cdc)
}

// Receive handles each received message for both Supervisor and Validator
func (state *HerdiusMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *blockProtobuf.ConnectionMessage:
		address := ctx.Client().ID.Address
		pubKey := ctx.Client().ID.PublicKey
		if err := supsvc.AddValidator(pubKey, address); err != nil {
			log.Error().Err(err).Msgf("<%s> Failed to add validator", address)
			return err
		}

		log.Info().Msgf("<%s> %s", address, msg.Message)

		// This map will be used to map validators to their respective child blocks
		mu := supsvc.GetMutex()
		mu.Lock()
		supsvc.ValidatorChildblock[address] = &blockProtobuf.BlockID{}
		mu.Unlock()

		sender, err := ctx.Network().Client(ctx.Client().Address)
		if err != nil {
			return fmt.Errorf("failed to get client network: %v", err)
		}
		nonce := 1
		if err := sender.Reply(
			network.WithSignMessage(context.Background(), true),
			uint64(nonce),
			&blockProtobuf.ConnectionMessage{Message: "Connection established with Supervisor"},
		); err != nil {
			return fmt.Errorf(fmt.Sprintf("Failed to reply to client: %v", err))
		}
	}
	return nil
}

func main() {
	// process other flags
	peersFlag := flag.String("peers", "", "peers to connect to")
	groupSizeFlag := flag.Int("groupsize", 3, "# of peers in a validator group")
	portFlag := flag.Int("port", 0, "port to bind validator to")
	envFlag := flag.String("env", "dev", "environment to build network and run process for")
	waitTimeFlag := flag.Int("waitTime", 15, "time to wait before the Memory Pool is flushed to a new block")
	restoreFlag := flag.Bool("restore", false, "restore blockchain from S3")
	backupFlag := flag.Bool("backup", false, "backup blockchain to S3")

	flag.Parse()

	noOfPeersInGroup := *groupSizeFlag
	port := *portFlag
	env := *envFlag
	waitTime := *waitTimeFlag
	restr := *restoreFlag
	backup := *backupFlag
	cfg := config.GetConfiguration(env)
	peers := []string{}
	if len(*peersFlag) > 0 {
		peers = strings.Split(*peersFlag, ",")
	}

	if port == 0 {
		port = cfg.SelfBroadcastPort
	}

	// Generate or Load Keys
	nodeAddress := cfg.SelfBroadcastIP + "_" + strconv.Itoa(port)
	nodekey, err := keystore.LoadOrGenNodeKey(nodeKeydir + nodeAddress + "_sk_peer_id.json")
	if err != nil {
		log.Error().Msgf("Failed to create or load node key: %v", err)
	}
	privKey := nodekey.PrivKey
	pubKey := privKey.PubKey()
	keys := &crypto.KeyPair{
		PublicKey:  pubKey.Bytes(),
		PrivateKey: privKey.Bytes(),
		PrivKey:    privKey,
		PubKey:     pubKey,
	}

	opcode.RegisterMessageType(types.OpcodeChildBlockMessage, &blockProtobuf.ChildBlockMessage{})
	opcode.RegisterMessageType(types.OpcodeConnectionMessage, &blockProtobuf.ConnectionMessage{})
	opcode.RegisterMessageType(types.OpcodeBlockHeightRequest, &protoplugin.BlockHeightRequest{})
	opcode.RegisterMessageType(types.OpcodeBlockResponse, &protoplugin.BlockResponse{})
	opcode.RegisterMessageType(types.OpcodeAccountRequest, &protoplugin.AccountRequest{})
	opcode.RegisterMessageType(types.OpcodeAccountResponse, &protoplugin.AccountResponse{})
	opcode.RegisterMessageType(types.OpcodeTxRequest, &protoplugin.TxRequest{})
	opcode.RegisterMessageType(types.OpcodeTxResponse, &protoplugin.TxResponse{})
	opcode.RegisterMessageType(types.OpcodeTxDetailRequest, &protoplugin.TxDetailRequest{})
	opcode.RegisterMessageType(types.OpcodeTxDetailResponse, &protoplugin.TxDetailResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByAddressRequest, &protoplugin.TxsByAddressRequest{})
	opcode.RegisterMessageType(types.OpcodeTxsResponse, &protoplugin.TxsResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByAssetAndAddressRequest, &protoplugin.TxsByAssetAndAddressRequest{})
	opcode.RegisterMessageType(types.OpcodeTxUpdateRequest, &protoplugin.TxUpdateRequest{})
	opcode.RegisterMessageType(types.OpcodeTxUpdateResponse, &protoplugin.TxUpdateResponse{})
	opcode.RegisterMessageType(types.OpcodeTxDeleteRequest, &protoplugin.TxDeleteRequest{})
	opcode.RegisterMessageType(types.OpcodeTxLockedRequest, &protoplugin.TxLockedRequest{})
	opcode.RegisterMessageType(types.OpcodeTxLockedResponse, &protoplugin.TxLockedResponse{})
	opcode.RegisterMessageType(types.OpcodePing, &protobuf.Ping{})
	opcode.RegisterMessageType(types.OpcodePong, &protobuf.Pong{})
	opcode.RegisterMessageType(types.OpcodeTxRedeemRequest, &protoplugin.TxRedeemRequest{})
	opcode.RegisterMessageType(types.OpcodeTxRedeemResponse, &protoplugin.TxRedeemResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByBlockHeightRequest, &protoplugin.TxsByBlockHeightRequest{})
	opcode.RegisterMessageType(types.OpcodeLastBlockRequest, &protoplugin.LastBlockRequest{})

	address := cfg.ConstructTCPAddress()
	builder := network.NewBuilderWithOptions(network.Address(address))
	builder.SetKeys(keys)

	builder.SetAddress(network.FormatAddress(cfg.Protocol, cfg.SelfBroadcastIP, uint16(port)))

	// Register peer discovery plugin.
	if err := builder.AddPlugin(new(discovery.Plugin)); err != nil {
		log.Fatal().Err(err).Msg("failed to add discovery plugin")
	}

	// Add custom Herdius plugin.
	if err := builder.AddPlugin(new(HerdiusMessagePlugin)); err != nil {
		log.Fatal().Err(err).Msg("failed to add herdius message plugin")
	}
	if err := builder.AddPlugin(new(message.BlockMessagePlugin)); err != nil {
		log.Fatal().Err(err).Msg("failed to add block message plugin")
	}
	if err := builder.AddPlugin(new(message.AccountMessagePlugin)); err != nil {
		log.Fatal().Err(err).Msg("failed to add account message plugin")
	}
	if err := builder.AddPlugin(message.NewTransactionMessagePlugin(env)); err != nil {
		log.Fatal().Err(err).Msg("failed to add transaction message plugin")
	}

	net, err := builder.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to build network")
	}

	go net.Listen()
	defer net.Close()

	c := new(network.ConnTester)
	go func() {
		c.IsConnected(net, peers)
	}()

	// As of now Databases will only be loaded for Supervisor.
	// Chain data and state information will be stored at supervisor's node.

	var stateRoot []byte
	accountStorage = external.New()
	if restr {
		log.Info().Msg("Restore value true: proceeding to restore from AWS S3")
		r := restore.NewRestorer(env, 3)
		if err := r.Restore(); err != nil {
			log.Error().Err(err).Msg("failed to restore from aws s3")
		}
	}
	blockchain.LoadDB()
	sup.LoadStateDB(accountStorage)
	blockchainSvc := &blockchain.Service{}

	lastBlock := blockchainSvc.GetLastBlock()

	var wg sync.WaitGroup
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	stop := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		syncer.DoSyncAllAccounts(accountStorage, env, stop)
	}()

	var lbh cmn.HexBytes
	lastBlockHash := lastBlock.GetHeader().GetBlock_ID().GetBlockHash()
	lbh = lastBlockHash
	lbHeight := lastBlock.GetHeader().GetHeight()

	log.Info().Msgf("Last Block Hash: %v", lbh)
	log.Info().Msgf("Height: %v", lbHeight)

	s := lastBlock.GetHeader().GetTime().GetSeconds()
	ts := time.Unix(s, 0)
	log.Info().Msgf("Timestamp: %v", ts)

	var stateRootHex cmn.HexBytes
	stateRoot = lastBlock.GetHeader().GetStateRoot()
	stateRootHex = stateRoot
	log.Info().Msgf("State root: %v", stateRootHex)

	supsvc.SetEnv(env)
	supsvc.SetWaitTime(waitTime)
	supsvc.SetNoOfPeersInGroup(noOfPeersInGroup)
	supsvc.SetBackup(backup)

	go func() {
		for {
			mu := supsvc.GetMutex()
			mu.Lock()
			// Check for deactivated validators and remove them from supervisor list
			for _, v := range supsvc.Validator {
				if !net.ConnectionStateExists(v.Address) {
					supsvc.RemoveValidator(v.Address)
				}
			}
			mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-stop:
				break
			default:
				supsvc.SetStateRoot(stateRoot)
				lastBlock := blockchainSvc.GetLastBlock()
				stateRoot = lastBlock.GetHeader().GetStateRoot()
				supsvc.SetStateRoot(stateRoot)
				baseBlock, err := supsvc.ProcessTxs(lastBlock, net)
				if err != nil {
					log.Error().Err(err).Msg("failed to process txs")
					continue
				}

				if err := blockchainSvc.AddBaseBlock(baseBlock); err != nil {
					log.Error().Err(err).Msg("Failed to Add Base Block")
					continue
				}

				var (
					pbbh cmn.HexBytes = baseBlock.Header.LastBlockID.BlockHash
					bbh  cmn.HexBytes = baseBlock.GetHeader().GetBlock_ID().GetBlockHash()
				)

				log.Info().Msg("New Block Added")
				log.Info().Msgf("Block Id: %v", bbh.String())
				log.Info().Msgf("Last Block Id: %v", pbbh.String())
				log.Info().Msgf("Block Height: %v", baseBlock.GetHeader().GetHeight())

				s := lastBlock.GetHeader().GetTime().GetSeconds()
				ts := time.Unix(s, 0)
				log.Info().Msgf("Timestamp : %v", ts)

				var stateRoot cmn.HexBytes
				stateRoot = baseBlock.GetHeader().GetStateRoot()
				log.Info().Msgf("State root : %v", stateRoot)
			}
		}
	}()

	sig := <-sigs
	log.Info().Msgf("Catch signal: %v", sig)
	log.Info().Msg("Notify syncer to stop")
	close(stop)
	log.Info().Msg("Wait syncer to finish.")
	wg.Wait()
	log.Info().Msg("Syncer finished, exit now.")
}
