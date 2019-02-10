// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvpair.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Define these here for compatibility but use libs/common.KVPair.
type KVPair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVPair) Reset()         { *m = KVPair{} }
func (m *KVPair) String() string { return proto.CompactTextString(m) }
func (*KVPair) ProtoMessage()    {}
func (*KVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_kvpair_e2b4075c0b91262e, []int{0}
}
func (m *KVPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVPair.Unmarshal(m, b)
}
func (m *KVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVPair.Marshal(b, m, deterministic)
}
func (dst *KVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVPair.Merge(dst, src)
}
func (m *KVPair) XXX_Size() int {
	return xxx_messageInfo_KVPair.Size(m)
}
func (m *KVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KVPair.DiscardUnknown(m)
}

var xxx_messageInfo_KVPair proto.InternalMessageInfo

func (m *KVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*KVPair)(nil), "common.KVPair")
}

func init() { proto.RegisterFile("kvpair.proto", fileDescriptor_kvpair_e2b4075c0b91262e) }

var fileDescriptor_kvpair_e2b4075c0b91262e = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x2e, 0x2b, 0x48,
	0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf,
	0x53, 0x32, 0xe0, 0x62, 0xf3, 0x0e, 0x0b, 0x48, 0xcc, 0x2c, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e,
	0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12,
	0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x4e, 0x12, 0x1b, 0xd8, 0x00, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x19, 0x87, 0xca, 0xa1, 0x50, 0x00, 0x00, 0x00,
}
