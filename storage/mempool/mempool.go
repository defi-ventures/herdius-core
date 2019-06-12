package mempool

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/herdius/herdius-core/accounts/account"
	"github.com/herdius/herdius-core/hbi/protobuf"
	"github.com/herdius/herdius-core/libs/common"
	"github.com/herdius/herdius-core/tx"
	"github.com/tendermint/go-amino"
)

// Service ...
type Service interface {
	AddTx(tx.Tx) int
	GetTxs() *tx.Txs
	RemoveTxs(int)
}

// MemPool ...
type MemPool struct {
	pending []mempoolTx
	queue   []mempoolTx
}

// Only one instance of MemPool will be instantiated.
var memPool *MemPool
var once sync.Once

// GetMemPool ..
func GetMemPool() *MemPool {
	once.Do(func() {
		memPool = &MemPool{}
	})
	return memPool
}

// mempoolTx is a transaction that successfully ran
type mempoolTx struct {
	height int64 // height that this tx had been validated in
	tx     *protobuf.Tx
}

// Height returns the height for this transaction
func (memTx *mempoolTx) Height() int64 {
	return atomic.LoadInt64(&memTx.height)
}

// AddTx adds the tx Transaction to the MemPool and returns the total
// number of current Transactions within the MemPool
func (m *MemPool) AddTx(tx *protobuf.Tx, accSrv account.ServiceI) (int, int) {
	account, _ := accSrv.GetAccountByAddress(tx.GetSenderAddress())
	mpSize := len(m.pending)
	mt := mempoolTx{
		tx:     tx,
		height: int64(mpSize) + 1,
	}

	if account != nil {
		if tx.GetAsset().Nonce == account.Nonce+1 {
			log.Println("First time tx Add to pending")
			m.pending = append(m.pending, mt)
			return len(m.pending), len(m.queue)
		}
	}
	if tx.GetAsset().Nonce == 0 {
		log.Println("First time tx Add to pending")
		m.pending = append(m.pending, mt)
		return len(m.pending), len(m.queue)

	}
	log.Println(" Add tx to queue")
	m.queue = append(m.queue, mt)
	return len(m.pending), len(m.queue)
}

func (m *MemPool) processQueue(accountService account.ServiceI) {
	/*
		1) get all queue tx
		2) compare account nonce with tx nonce

	*/
	log.Printf("Processing queue and pending txs, Size of pending %d, Size of queue: %d", len(m.pending), len(m.queue))

	for i, mTx := range m.queue {
		account, _ := accountService.GetAccountByAddress(mTx.tx.GetSenderAddress())
		if account != nil {
			if account.Nonce+1 == mTx.tx.GetAsset().GetNonce() {
				m.pending = append(m.pending, mTx)
				m.queue = append(m.queue[:i], m.queue[i+1:]...)
			}
		}
	}
	log.Printf("Finish Processing queue and pending txs, Size of pending %d, Size of queue: %d", len(m.pending), len(m.queue))

}

// GetTxs gets all transactions from the MemPool
func (m *MemPool) GetTxs() *tx.Txs {
	accSrv := account.NewAccountService()
	m.processQueue(accSrv)
	txs := &tx.Txs{}
	m.processQueue(accSrv)
	var cdc = amino.NewCodec()
	for _, mt := range m.pending {
		tx, _ := cdc.MarshalJSON(mt.tx)
		*txs = append(*txs, tx)
	}
	return txs
}

// GetTx returns a Tx for the given ID or nil if the corresponding TX exists
// Returns empty if Tx not found
func (m *MemPool) GetTx(id string) (int, *protobuf.Tx, error) {
	log.Println("Retrieving MemPool Tx's")
	for i, txQ := range m.pending {
		var cdc = amino.NewCodec()
		// txStr := &protobuf.Tx{}
		// err := cdc.UnmarshalJSON(txbz.tx, txStr)
		txbz, err := cdc.MarshalJSON(txQ.tx)
		if err != nil {
			return 0, nil, fmt.Errorf("unable to unmarshal tx bytes to txStr: %v", err)
		}

		txbzID := common.CreateTxID(txbz)
		if txbzID == id {
			log.Println("Matching transaction found for Tx ID:", id)
			return i, txQ.tx, nil
		}
	}
	return 0, nil, nil
}

// DeleteTx deletes a transaction currently in the MemPool by the transaction ID
// Returns true if successfully cancelled, false if can't find or cancel the transaction
func (m *MemPool) DeleteTx(id string) bool {
	log.Println("Beginning attempted removal from memory pool of Tx w/ ID:", id)
	for i, txStr := range m.pending {
		var cdc = amino.NewCodec()

		txbz, _ := cdc.MarshalJSON(txStr)
		mTxID := common.CreateTxID(txbz)
		if mTxID == id {
			log.Printf("Matched Tx ID (%v), removing from memory memory pool", id)
			m.pending = append(m.pending[:i], m.pending[i+1:]...)
			return true
		}
	}
	log.Printf("Unable to find Tx (id: %v) in memory pool", id)
	return false
}

// RemoveTxs transactions from the MemPool
func (m *MemPool) RemoveTxs(i int) {
	log.Println("Removing tx from mempool", i)
	if len(m.pending) < 1 {
		m.pending = m.pending[len(m.pending):]
		return
	}
	m.pending = m.pending[i:]
}
