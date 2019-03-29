// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hbi/protobuf/service.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Timestamp struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int64    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type BlockHeightRequest struct {
	BlockHeight          int64    `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeightRequest) Reset()         { *m = BlockHeightRequest{} }
func (m *BlockHeightRequest) String() string { return proto.CompactTextString(m) }
func (*BlockHeightRequest) ProtoMessage()    {}
func (*BlockHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{1}
}

func (m *BlockHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeightRequest.Unmarshal(m, b)
}
func (m *BlockHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeightRequest.Marshal(b, m, deterministic)
}
func (m *BlockHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightRequest.Merge(m, src)
}
func (m *BlockHeightRequest) XXX_Size() int {
	return xxx_messageInfo_BlockHeightRequest.Size(m)
}
func (m *BlockHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightRequest proto.InternalMessageInfo

func (m *BlockHeightRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type BlockResponse struct {
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Time of block intialization
	Time         *Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Transactions int32      `protobuf:"varint,3,opt,name=transactions,proto3" json:"transactions,omitempty"`
	// Supervisor herdius token address who created the block
	SupervisorAddress    string   `protobuf:"bytes,4,opt,name=supervisor_address,json=supervisorAddress,proto3" json:"supervisor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockResponse) Reset()         { *m = BlockResponse{} }
func (m *BlockResponse) String() string { return proto.CompactTextString(m) }
func (*BlockResponse) ProtoMessage()    {}
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{2}
}

func (m *BlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockResponse.Unmarshal(m, b)
}
func (m *BlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockResponse.Marshal(b, m, deterministic)
}
func (m *BlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponse.Merge(m, src)
}
func (m *BlockResponse) XXX_Size() int {
	return xxx_messageInfo_BlockResponse.Size(m)
}
func (m *BlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponse proto.InternalMessageInfo

func (m *BlockResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BlockResponse) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *BlockResponse) GetTransactions() int32 {
	if m != nil {
		return m.Transactions
	}
	return 0
}

func (m *BlockResponse) GetSupervisorAddress() string {
	if m != nil {
		return m.SupervisorAddress
	}
	return ""
}

type AccountRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{3}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AccountResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Nonce                uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              uint64   `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	StorageRoot          string   `protobuf:"bytes,4,opt,name=storage_root,json=storageRoot,proto3" json:"storage_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{4}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AccountResponse) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountResponse) GetStorageRoot() string {
	if m != nil {
		return m.StorageRoot
	}
	return ""
}

type Asset struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Network              string   `protobuf:"bytes,2,opt,name=Network,proto3" json:"Network,omitempty"`
	Value                uint64   `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
	Fee                  uint64   `protobuf:"varint,4,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Nonce                uint64   `protobuf:"varint,5,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{5}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Asset) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Asset) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Asset) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Asset) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Transaction struct {
	Senderpubkey         []byte   `protobuf:"bytes,1,opt,name=senderpubkey,proto3" json:"senderpubkey,omitempty"`
	Signature            string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Recaddress           string   `protobuf:"bytes,3,opt,name=recaddress,proto3" json:"recaddress,omitempty"`
	Asset                *Asset   `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{6}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetSenderpubkey() []byte {
	if m != nil {
		return m.Senderpubkey
	}
	return nil
}

func (m *Transaction) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Transaction) GetRecaddress() string {
	if m != nil {
		return m.Recaddress
	}
	return ""
}

func (m *Transaction) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type TransactionRequest struct {
	Tx                   *Transaction `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{7}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

type TransactionResponse struct {
	TxId                 string   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Pending              int64    `protobuf:"varint,2,opt,name=pending,proto3" json:"pending,omitempty"`
	Queued               int64    `protobuf:"varint,3,opt,name=queued,proto3" json:"queued,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebece8ff681ed6b0, []int{8}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TransactionResponse) GetPending() int64 {
	if m != nil {
		return m.Pending
	}
	return 0
}

func (m *TransactionResponse) GetQueued() int64 {
	if m != nil {
		return m.Queued
	}
	return 0
}

func (m *TransactionResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "protobuf.Timestamp")
	proto.RegisterType((*BlockHeightRequest)(nil), "protobuf.BlockHeightRequest")
	proto.RegisterType((*BlockResponse)(nil), "protobuf.BlockResponse")
	proto.RegisterType((*AccountRequest)(nil), "protobuf.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "protobuf.AccountResponse")
	proto.RegisterType((*Asset)(nil), "protobuf.Asset")
	proto.RegisterType((*Transaction)(nil), "protobuf.Transaction")
	proto.RegisterType((*TransactionRequest)(nil), "protobuf.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "protobuf.TransactionResponse")
}

func init() { proto.RegisterFile("hbi/protobuf/service.proto", fileDescriptor_ebece8ff681ed6b0) }

var fileDescriptor_ebece8ff681ed6b0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0xeb, 0xb8, 0xfd, 0x65, 0x9c, 0x1f, 0xa5, 0x1b, 0x40, 0x16, 0x42, 0x28, 0x58, 0xaa,
	0x88, 0x90, 0x9a, 0x4a, 0xe1, 0xc0, 0xa1, 0xa7, 0xe4, 0x80, 0xca, 0x81, 0xaa, 0x5a, 0x45, 0x1c,
	0xb8, 0x44, 0x6b, 0x7b, 0x48, 0xac, 0x24, 0xbb, 0xee, 0xce, 0x1a, 0xd2, 0x0b, 0x1f, 0x83, 0x8f,
	0xc1, 0x67, 0x44, 0xbb, 0xf6, 0xe6, 0xcf, 0x01, 0x89, 0x9b, 0xdf, 0x9b, 0x37, 0x9e, 0x37, 0x6f,
	0x77, 0xe1, 0xe5, 0x32, 0x2b, 0xaf, 0x2b, 0xad, 0x8c, 0xca, 0xea, 0x6f, 0xd7, 0x84, 0xfa, 0x7b,
	0x99, 0xe3, 0xc8, 0x11, 0xec, 0x3f, 0xcf, 0xa7, 0x37, 0xd0, 0x9d, 0x95, 0x1b, 0x24, 0x23, 0x36,
	0x15, 0x4b, 0xe0, 0x8c, 0x30, 0x57, 0xb2, 0xa0, 0x24, 0x18, 0x04, 0xc3, 0x90, 0x7b, 0xc8, 0x9e,
	0x41, 0x24, 0x85, 0x54, 0x94, 0x9c, 0x38, 0xbe, 0x01, 0xe9, 0x07, 0x60, 0xd3, 0xb5, 0xca, 0x57,
	0xb7, 0x58, 0x2e, 0x96, 0x86, 0xe3, 0x43, 0x8d, 0x64, 0xd8, 0x1b, 0xe8, 0x65, 0x96, 0x9d, 0x2f,
	0x1d, 0xdd, 0xfe, 0x2a, 0xce, 0xf6, 0xca, 0xf4, 0x77, 0x00, 0xff, 0xbb, 0x4e, 0x8e, 0x54, 0x29,
	0x49, 0xf8, 0x0f, 0x4d, 0xec, 0x2d, 0x74, 0x4c, 0xb9, 0x41, 0x67, 0x21, 0x1e, 0xf7, 0x47, 0x7e,
	0x87, 0xd1, 0x6e, 0x01, 0xee, 0x04, 0x2c, 0x85, 0x9e, 0xd1, 0x42, 0x92, 0xc8, 0x4d, 0xa9, 0x24,
	0x25, 0xe1, 0x20, 0x18, 0x46, 0xfc, 0x88, 0x63, 0x57, 0xc0, 0xa8, 0xae, 0x6c, 0x28, 0xa4, 0xf4,
	0x5c, 0x14, 0x85, 0x46, 0xa2, 0xa4, 0x33, 0x08, 0x86, 0x5d, 0x7e, 0xb1, 0xaf, 0x4c, 0x9a, 0x42,
	0xfa, 0x0e, 0x9e, 0x4c, 0xf2, 0x5c, 0xd5, 0x72, 0xb7, 0x65, 0x02, 0x67, 0xbe, 0x2b, 0x70, 0x5d,
	0x1e, 0xa6, 0x3f, 0xe1, 0x7c, 0xa7, 0x6d, 0xb7, 0xfb, 0xab, 0xd8, 0x05, 0xab, 0x64, 0xde, 0x6c,
	0xd5, 0xe1, 0x0d, 0xb0, 0xfa, 0x4c, 0xac, 0x85, 0xe5, 0x43, 0xc7, 0x7b, 0x68, 0x73, 0x22, 0xa3,
	0xb4, 0x58, 0xe0, 0x5c, 0x2b, 0x65, 0x5a, 0xc7, 0x71, 0xcb, 0x71, 0xa5, 0x4c, 0xfa, 0x08, 0xd1,
	0x84, 0x08, 0x9d, 0xc5, 0xcf, 0x48, 0x24, 0x16, 0xe8, 0xa7, 0xb6, 0xd0, 0x56, 0xee, 0xd0, 0xfc,
	0x50, 0x7a, 0xe5, 0xe6, 0x76, 0xb9, 0x87, 0xd6, 0xcf, 0x17, 0xb1, 0xae, 0xfd, 0xdc, 0x06, 0xb0,
	0xa7, 0x10, 0x7e, 0x44, 0x74, 0xc3, 0x3a, 0xdc, 0x7e, 0x5a, 0xdd, 0x9d, 0xf3, 0x1d, 0x35, 0x3a,
	0x07, 0xd2, 0x5f, 0x01, 0xc4, 0xb3, 0x7d, 0xcc, 0xf6, 0x24, 0x08, 0x65, 0x81, 0xba, 0xaa, 0xb3,
	0x15, 0x3e, 0x3a, 0x1b, 0x3d, 0x7e, 0xc4, 0xb1, 0x57, 0xd0, 0xa5, 0x72, 0x21, 0x85, 0xa9, 0x35,
	0xb6, 0x6e, 0xf6, 0x04, 0x7b, 0x0d, 0xa0, 0x31, 0xf7, 0xe1, 0x85, 0xae, 0x7c, 0xc0, 0xb0, 0x4b,
	0x88, 0x84, 0x5d, 0xd6, 0x79, 0x8b, 0xc7, 0xe7, 0xfb, 0x5b, 0xe1, 0x32, 0xe0, 0x4d, 0x35, 0xbd,
	0x01, 0x76, 0xe0, 0xcb, 0x9f, 0xe1, 0x25, 0x9c, 0xcc, 0xb6, 0xce, 0x54, 0x3c, 0x7e, 0x7e, 0x70,
	0x9f, 0x0e, 0x94, 0x27, 0xb3, 0x6d, 0x6a, 0xa0, 0x7f, 0xd4, 0xdc, 0x1e, 0x6a, 0x1f, 0x22, 0xb3,
	0x9d, 0x97, 0x45, 0x1b, 0x6e, 0xc7, 0x6c, 0x3f, 0x15, 0x36, 0xd9, 0x0a, 0x65, 0x51, 0xca, 0x45,
	0xfb, 0x54, 0x3c, 0x64, 0x2f, 0xe0, 0xf4, 0xa1, 0xc6, 0x1a, 0x0b, 0xb7, 0x45, 0xc8, 0x5b, 0x64,
	0x79, 0x32, 0xc2, 0xd4, 0xfe, 0xf6, 0xb5, 0x68, 0x7a, 0x05, 0x17, 0xb9, 0xda, 0x8c, 0x96, 0xa8,
	0x8b, 0xb2, 0xa6, 0xc6, 0xdd, 0xb4, 0x77, 0xdb, 0xc0, 0x7b, 0x8b, 0xee, 0x83, 0xaf, 0xbb, 0x87,
	0x9c, 0x9d, 0xba, 0xaf, 0xf7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1b, 0xad, 0x37, 0xf7,
	0x03, 0x00, 0x00,
}
