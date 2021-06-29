// Code generated by protoc-gen-go. DO NOT EDIT.
// source: evm_event.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// 一条evm event log数据
type EVMLog struct {
	Topic                [][]byte `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EVMLog) Reset()         { *m = EVMLog{} }
func (m *EVMLog) String() string { return proto.CompactTextString(m) }
func (*EVMLog) ProtoMessage()    {}
func (*EVMLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{0}
}

func (m *EVMLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMLog.Unmarshal(m, b)
}
func (m *EVMLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMLog.Marshal(b, m, deterministic)
}
func (m *EVMLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMLog.Merge(m, src)
}
func (m *EVMLog) XXX_Size() int {
	return xxx_messageInfo_EVMLog.Size(m)
}
func (m *EVMLog) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMLog.DiscardUnknown(m)
}

var xxx_messageInfo_EVMLog proto.InternalMessageInfo

func (m *EVMLog) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *EVMLog) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 多条evm event log数据
type EVMLogsPerTx struct {
	Logs                 []*EVMLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EVMLogsPerTx) Reset()         { *m = EVMLogsPerTx{} }
func (m *EVMLogsPerTx) String() string { return proto.CompactTextString(m) }
func (*EVMLogsPerTx) ProtoMessage()    {}
func (*EVMLogsPerTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{1}
}

func (m *EVMLogsPerTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMLogsPerTx.Unmarshal(m, b)
}
func (m *EVMLogsPerTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMLogsPerTx.Marshal(b, m, deterministic)
}
func (m *EVMLogsPerTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMLogsPerTx.Merge(m, src)
}
func (m *EVMLogsPerTx) XXX_Size() int {
	return xxx_messageInfo_EVMLogsPerTx.Size(m)
}
func (m *EVMLogsPerTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMLogsPerTx.DiscardUnknown(m)
}

var xxx_messageInfo_EVMLogsPerTx proto.InternalMessageInfo

func (m *EVMLogsPerTx) GetLogs() []*EVMLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type EVMTxAndLogs struct {
	Tx                   *Transaction  `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	LogsPerTx            *EVMLogsPerTx `protobuf:"bytes,2,opt,name=logsPerTx,proto3" json:"logsPerTx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EVMTxAndLogs) Reset()         { *m = EVMTxAndLogs{} }
func (m *EVMTxAndLogs) String() string { return proto.CompactTextString(m) }
func (*EVMTxAndLogs) ProtoMessage()    {}
func (*EVMTxAndLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{2}
}

func (m *EVMTxAndLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMTxAndLogs.Unmarshal(m, b)
}
func (m *EVMTxAndLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMTxAndLogs.Marshal(b, m, deterministic)
}
func (m *EVMTxAndLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMTxAndLogs.Merge(m, src)
}
func (m *EVMTxAndLogs) XXX_Size() int {
	return xxx_messageInfo_EVMTxAndLogs.Size(m)
}
func (m *EVMTxAndLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMTxAndLogs.DiscardUnknown(m)
}

var xxx_messageInfo_EVMTxAndLogs proto.InternalMessageInfo

func (m *EVMTxAndLogs) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *EVMTxAndLogs) GetLogsPerTx() *EVMLogsPerTx {
	if m != nil {
		return m.LogsPerTx
	}
	return nil
}

//一个块中包含的多条evm event log数据
type EVMTxLogPerBlk struct {
	TxAndLogs            []*EVMTxAndLogs `protobuf:"bytes,1,rep,name=txAndLogs,proto3" json:"txAndLogs,omitempty"`
	Height               int64           `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	BlockHash            []byte          `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash           []byte          `protobuf:"bytes,4,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	PreviousHash         []byte          `protobuf:"bytes,5,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	AddDelType           int32           `protobuf:"varint,6,opt,name=addDelType,proto3" json:"addDelType,omitempty"`
	SeqNum               int64           `protobuf:"varint,7,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EVMTxLogPerBlk) Reset()         { *m = EVMTxLogPerBlk{} }
func (m *EVMTxLogPerBlk) String() string { return proto.CompactTextString(m) }
func (*EVMTxLogPerBlk) ProtoMessage()    {}
func (*EVMTxLogPerBlk) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{3}
}

func (m *EVMTxLogPerBlk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMTxLogPerBlk.Unmarshal(m, b)
}
func (m *EVMTxLogPerBlk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMTxLogPerBlk.Marshal(b, m, deterministic)
}
func (m *EVMTxLogPerBlk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMTxLogPerBlk.Merge(m, src)
}
func (m *EVMTxLogPerBlk) XXX_Size() int {
	return xxx_messageInfo_EVMTxLogPerBlk.Size(m)
}
func (m *EVMTxLogPerBlk) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMTxLogPerBlk.DiscardUnknown(m)
}

var xxx_messageInfo_EVMTxLogPerBlk proto.InternalMessageInfo

func (m *EVMTxLogPerBlk) GetTxAndLogs() []*EVMTxAndLogs {
	if m != nil {
		return m.TxAndLogs
	}
	return nil
}

func (m *EVMTxLogPerBlk) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *EVMTxLogPerBlk) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *EVMTxLogPerBlk) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *EVMTxLogPerBlk) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *EVMTxLogPerBlk) GetAddDelType() int32 {
	if m != nil {
		return m.AddDelType
	}
	return 0
}

func (m *EVMTxLogPerBlk) GetSeqNum() int64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

//多个块中包含的多条evm event log数据
type EVMTxLogsInBlks struct {
	Logs4EVMPerBlk       []*EVMTxLogPerBlk `protobuf:"bytes,1,rep,name=logs4EVMPerBlk,proto3" json:"logs4EVMPerBlk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EVMTxLogsInBlks) Reset()         { *m = EVMTxLogsInBlks{} }
func (m *EVMTxLogsInBlks) String() string { return proto.CompactTextString(m) }
func (*EVMTxLogsInBlks) ProtoMessage()    {}
func (*EVMTxLogsInBlks) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{4}
}

func (m *EVMTxLogsInBlks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMTxLogsInBlks.Unmarshal(m, b)
}
func (m *EVMTxLogsInBlks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMTxLogsInBlks.Marshal(b, m, deterministic)
}
func (m *EVMTxLogsInBlks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMTxLogsInBlks.Merge(m, src)
}
func (m *EVMTxLogsInBlks) XXX_Size() int {
	return xxx_messageInfo_EVMTxLogsInBlks.Size(m)
}
func (m *EVMTxLogsInBlks) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMTxLogsInBlks.DiscardUnknown(m)
}

var xxx_messageInfo_EVMTxLogsInBlks proto.InternalMessageInfo

func (m *EVMTxLogsInBlks) GetLogs4EVMPerBlk() []*EVMTxLogPerBlk {
	if m != nil {
		return m.Logs4EVMPerBlk
	}
	return nil
}

// 创建/调用合约的请求结构
type EVMContractAction4Chain33 struct {
	// 转账金额
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// 消耗限制，默认为Transaction.Fee
	GasLimit uint64 `protobuf:"varint,2,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	// gas价格，默认为1
	GasPrice uint32 `protobuf:"varint,3,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	// 合约数据
	Code []byte `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	//交易参数
	Para []byte `protobuf:"bytes,5,opt,name=para,proto3" json:"para,omitempty"`
	// 合约别名，方便识别
	Alias string `protobuf:"bytes,6,opt,name=alias,proto3" json:"alias,omitempty"`
	// 交易备注
	Note string `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	// 调用合约地址
	ContractAddr         string   `protobuf:"bytes,8,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EVMContractAction4Chain33) Reset()         { *m = EVMContractAction4Chain33{} }
func (m *EVMContractAction4Chain33) String() string { return proto.CompactTextString(m) }
func (*EVMContractAction4Chain33) ProtoMessage()    {}
func (*EVMContractAction4Chain33) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a9a715c51188e3, []int{5}
}

func (m *EVMContractAction4Chain33) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVMContractAction4Chain33.Unmarshal(m, b)
}
func (m *EVMContractAction4Chain33) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVMContractAction4Chain33.Marshal(b, m, deterministic)
}
func (m *EVMContractAction4Chain33) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMContractAction4Chain33.Merge(m, src)
}
func (m *EVMContractAction4Chain33) XXX_Size() int {
	return xxx_messageInfo_EVMContractAction4Chain33.Size(m)
}
func (m *EVMContractAction4Chain33) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMContractAction4Chain33.DiscardUnknown(m)
}

var xxx_messageInfo_EVMContractAction4Chain33 proto.InternalMessageInfo

func (m *EVMContractAction4Chain33) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EVMContractAction4Chain33) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *EVMContractAction4Chain33) GetGasPrice() uint32 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *EVMContractAction4Chain33) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *EVMContractAction4Chain33) GetPara() []byte {
	if m != nil {
		return m.Para
	}
	return nil
}

func (m *EVMContractAction4Chain33) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *EVMContractAction4Chain33) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *EVMContractAction4Chain33) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*EVMLog)(nil), "types.EVMLog")
	proto.RegisterType((*EVMLogsPerTx)(nil), "types.EVMLogsPerTx")
	proto.RegisterType((*EVMTxAndLogs)(nil), "types.EVMTxAndLogs")
	proto.RegisterType((*EVMTxLogPerBlk)(nil), "types.EVMTxLogPerBlk")
	proto.RegisterType((*EVMTxLogsInBlks)(nil), "types.EVMTxLogsInBlks")
	proto.RegisterType((*EVMContractAction4Chain33)(nil), "types.EVMContractAction4Chain33")
}

func init() {
	proto.RegisterFile("evm_event.proto", fileDescriptor_00a9a715c51188e3)
}

var fileDescriptor_00a9a715c51188e3 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x55, 0x20, 0xd0, 0xc5, 0xb0, 0xbb, 0xaa, 0xfb, 0xa1, 0x74, 0xd5, 0x0f, 0x9a, 0x13, 0x27,
	0xd0, 0x92, 0xbd, 0xf6, 0xb0, 0x6c, 0x91, 0x5a, 0x09, 0x2a, 0x64, 0x21, 0x0e, 0xbd, 0x54, 0xc6,
	0xb1, 0x12, 0x8b, 0xc4, 0x4e, 0x6d, 0x83, 0xd8, 0x9f, 0xdb, 0x1f, 0xd1, 0x7b, 0xe5, 0x71, 0xf8,
	0xea, 0x6d, 0x66, 0xde, 0x9b, 0x99, 0x37, 0x2f, 0x0e, 0xba, 0xe5, 0xbb, 0xf2, 0x17, 0xdf, 0x71,
	0x69, 0x87, 0x95, 0x56, 0x56, 0xe1, 0x96, 0x7d, 0xae, 0xb8, 0xb9, 0x7b, 0x69, 0x35, 0x95, 0x86,
	0x32, 0x2b, 0x94, 0xf4, 0x48, 0x3c, 0x46, 0xed, 0xe9, 0x6a, 0x3e, 0x53, 0x19, 0x7e, 0x8d, 0x5a,
	0x56, 0x55, 0x82, 0x45, 0x41, 0xbf, 0x39, 0xe8, 0x11, 0x9f, 0x60, 0x8c, 0xc2, 0x94, 0x5a, 0x1a,
	0x35, 0xfa, 0xc1, 0xa0, 0x47, 0x20, 0x8e, 0xef, 0x51, 0xcf, 0xf7, 0x98, 0x05, 0xd7, 0xcb, 0x3d,
	0xfe, 0x8c, 0xc2, 0x42, 0x65, 0x06, 0x1a, 0xbb, 0xe3, 0xeb, 0x21, 0x2c, 0x1b, 0x7a, 0x0a, 0x01,
	0x28, 0xe6, 0xd0, 0xb2, 0xdc, 0x3f, 0xca, 0xd4, 0xf5, 0xe1, 0x18, 0x35, 0xec, 0x3e, 0x0a, 0xfa,
	0xc1, 0xa0, 0x3b, 0xc6, 0x75, 0xc3, 0xf2, 0x24, 0x8e, 0x34, 0xec, 0x1e, 0xdf, 0xa3, 0x4e, 0x71,
	0xd8, 0x01, 0xfb, 0xbb, 0xe3, 0x57, 0x17, 0xb3, 0x3d, 0x44, 0x4e, 0xac, 0xf8, 0x6f, 0x80, 0x6e,
	0x60, 0xcf, 0x4c, 0x65, 0x0b, 0xae, 0x27, 0xc5, 0xc6, 0x4d, 0xb1, 0x87, 0xb5, 0xb5, 0xc2, 0xb3,
	0x29, 0x47, 0x45, 0xe4, 0xc4, 0xc2, 0x6f, 0x51, 0x3b, 0xe7, 0x22, 0xcb, 0x2d, 0x6c, 0x6d, 0x92,
	0x3a, 0xc3, 0xef, 0x51, 0x67, 0x5d, 0x28, 0xb6, 0xf9, 0x46, 0x4d, 0x1e, 0x35, 0xc1, 0x90, 0x53,
	0x01, 0x7f, 0x44, 0xa8, 0xa2, 0x9a, 0x4b, 0x0b, 0x70, 0x08, 0xf0, 0x59, 0x05, 0xc7, 0xa8, 0x57,
	0x69, 0xbe, 0x13, 0x6a, 0x6b, 0x80, 0xd1, 0x02, 0xc6, 0x45, 0xcd, 0xcd, 0xa0, 0x69, 0xfa, 0x95,
	0x17, 0xcb, 0xe7, 0x8a, 0x47, 0xed, 0x7e, 0x30, 0x68, 0x91, 0xb3, 0x8a, 0x53, 0x66, 0xf8, 0xef,
	0x1f, 0xdb, 0x32, 0x7a, 0xe1, 0x95, 0xf9, 0x2c, 0x5e, 0xa0, 0xdb, 0xc3, 0xd9, 0xe6, 0xbb, 0x9c,
	0x14, 0x1b, 0x83, 0xbf, 0xa0, 0x1b, 0xe7, 0xcb, 0xc3, 0x74, 0x35, 0xf7, 0x4e, 0xd4, 0xc7, 0xbf,
	0x39, 0x3f, 0xfe, 0x68, 0x13, 0xf9, 0x8f, 0x1c, 0xff, 0x09, 0xd0, 0xbb, 0xe9, 0x6a, 0xfe, 0xa4,
	0xa4, 0xd5, 0x94, 0xd9, 0x47, 0xf8, 0x2c, 0x0f, 0x4f, 0x39, 0x15, 0x32, 0x49, 0x9c, 0x0e, 0x5a,
	0xaa, 0xad, 0xb4, 0xf0, 0x09, 0x43, 0x52, 0x67, 0xf8, 0x0e, 0x5d, 0x65, 0xd4, 0xcc, 0x44, 0x29,
	0xbc, 0x77, 0x21, 0x39, 0xe6, 0x35, 0xb6, 0xd0, 0x82, 0x71, 0x30, 0xef, 0x9a, 0x1c, 0x73, 0xf7,
	0xca, 0x98, 0x4a, 0x79, 0xed, 0x1a, 0xc4, 0xae, 0x56, 0x51, 0x4d, 0x6b, 0x9f, 0x20, 0x76, 0x6f,
	0x94, 0x16, 0x82, 0x1a, 0xb0, 0xa6, 0x43, 0x7c, 0xe2, 0x98, 0x52, 0x59, 0x0e, 0x9e, 0x74, 0x08,
	0xc4, 0xce, 0x6d, 0x76, 0xd0, 0x9e, 0xa6, 0x3a, 0xba, 0x02, 0xec, 0xa2, 0x36, 0xf9, 0xf4, 0xf3,
	0x43, 0x26, 0x6c, 0xbe, 0x5d, 0x0f, 0x99, 0x2a, 0x47, 0x49, 0xc2, 0xe4, 0x88, 0xf9, 0x03, 0x47,
	0xe0, 0xd1, 0xba, 0x0d, 0xff, 0x48, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x16, 0x20, 0x16,
	0x50, 0x03, 0x00, 0x00,
}
