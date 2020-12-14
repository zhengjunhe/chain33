// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package accounts is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	Account
	MinerAccount
	Accounts
	MinerAccounts
	Config
*/
package accounts

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Addr    string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Frozen  string `protobuf:"bytes,2,opt,name=frozen" json:"frozen,omitempty"`
	Balance string `protobuf:"bytes,3,opt,name=balance" json:"balance,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Account) GetFrozen() string {
	if m != nil {
		return m.Frozen
	}
	return ""
}

func (m *Account) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

type MinerAccount struct {
	Addr              string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Total             string `protobuf:"bytes,2,opt,name=total" json:"total,omitempty"`
	Increase          string `protobuf:"bytes,3,opt,name=increase" json:"increase,omitempty"`
	Frozen            string `protobuf:"bytes,4,opt,name=frozen" json:"frozen,omitempty"`
	ExpectIncrease    string `protobuf:"bytes,5,opt,name=expectIncrease" json:"expectIncrease,omitempty"`
	MinerDpomDuring    string `protobuf:"bytes,6,opt,name=minerDpomDuring" json:"minerDpomDuring,omitempty"`
	ExpectMinerBlocks string `protobuf:"bytes,7,opt,name=expectMinerBlocks" json:"expectMinerBlocks,omitempty"`
}

func (m *MinerAccount) Reset()                    { *m = MinerAccount{} }
func (m *MinerAccount) String() string            { return proto.CompactTextString(m) }
func (*MinerAccount) ProtoMessage()               {}
func (*MinerAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MinerAccount) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *MinerAccount) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *MinerAccount) GetIncrease() string {
	if m != nil {
		return m.Increase
	}
	return ""
}

func (m *MinerAccount) GetFrozen() string {
	if m != nil {
		return m.Frozen
	}
	return ""
}

func (m *MinerAccount) GetExpectIncrease() string {
	if m != nil {
		return m.ExpectIncrease
	}
	return ""
}

func (m *MinerAccount) GetMinerDpomDuring() string {
	if m != nil {
		return m.MinerDpomDuring
	}
	return ""
}

func (m *MinerAccount) GetExpectMinerBlocks() string {
	if m != nil {
		return m.ExpectMinerBlocks
	}
	return ""
}

type Accounts struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type MinerAccounts struct {
	MinerAccounts       []*MinerAccount `protobuf:"bytes,1,rep,name=minerAccounts" json:"minerAccounts,omitempty"`
	Seconds             int64           `protobuf:"varint,2,opt,name=seconds" json:"seconds,omitempty"`
	TotalIncrease       string          `protobuf:"bytes,3,opt,name=totalIncrease" json:"totalIncrease,omitempty"`
	Blocks              int64           `protobuf:"varint,4,opt,name=blocks" json:"blocks,omitempty"`
	ExpectBlocks        int64           `protobuf:"varint,5,opt,name=expectBlocks" json:"expectBlocks,omitempty"`
	ExpectTotalIncrease string          `protobuf:"bytes,6,opt,name=expectTotalIncrease" json:"expectTotalIncrease,omitempty"`
}

func (m *MinerAccounts) Reset()                    { *m = MinerAccounts{} }
func (m *MinerAccounts) String() string            { return proto.CompactTextString(m) }
func (*MinerAccounts) ProtoMessage()               {}
func (*MinerAccounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MinerAccounts) GetMinerAccounts() []*MinerAccount {
	if m != nil {
		return m.MinerAccounts
	}
	return nil
}

func (m *MinerAccounts) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *MinerAccounts) GetTotalIncrease() string {
	if m != nil {
		return m.TotalIncrease
	}
	return ""
}

func (m *MinerAccounts) GetBlocks() int64 {
	if m != nil {
		return m.Blocks
	}
	return 0
}

func (m *MinerAccounts) GetExpectBlocks() int64 {
	if m != nil {
		return m.ExpectBlocks
	}
	return 0
}

func (m *MinerAccounts) GetExpectTotalIncrease() string {
	if m != nil {
		return m.ExpectTotalIncrease
	}
	return ""
}

type Config struct {
	Whitelist     []string `protobuf:"bytes,1,rep,name=whitelist" json:"whitelist,omitempty"`
	JrpcBindAddr  string   `protobuf:"bytes,2,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	DataDir       string   `protobuf:"bytes,3,opt,name=dataDir" json:"dataDir,omitempty"`
	MinerAddr     []string `protobuf:"bytes,4,rep,name=minerAddr" json:"minerAddr,omitempty"`
	DplatformOSHost string   `protobuf:"bytes,5,opt,name=dplatformhost" json:"dplatformhost,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Config) GetWhitelist() []string {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *Config) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Config) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

func (m *Config) GetMinerAddr() []string {
	if m != nil {
		return m.MinerAddr
	}
	return nil
}

func (m *Config) GetDplatformOSHost() string {
	if m != nil {
		return m.DplatformOSHost
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "accounts.Account")
	proto.RegisterType((*MinerAccount)(nil), "accounts.MinerAccount")
	proto.RegisterType((*Accounts)(nil), "accounts.Accounts")
	proto.RegisterType((*MinerAccounts)(nil), "accounts.MinerAccounts")
	proto.RegisterType((*Config)(nil), "accounts.Config")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x6e, 0x1a, 0x31,
	0x14, 0xd5, 0x94, 0x61, 0x80, 0x0b, 0x53, 0x09, 0xb7, 0x42, 0xa3, 0xaa, 0x0b, 0x34, 0xaa, 0x2a,
	0x16, 0x2d, 0xaa, 0xca, 0xaa, 0x52, 0x37, 0x50, 0x36, 0x2c, 0xaa, 0x48, 0xa3, 0xfc, 0x80, 0xf1,
	0x18, 0x70, 0x32, 0xd8, 0xc8, 0x36, 0xca, 0xe3, 0x7b, 0xf2, 0x75, 0xf9, 0x85, 0x6c, 0x22, 0x3f,
	0x06, 0xc6, 0x49, 0x94, 0x9d, 0xcf, 0xb9, 0xf7, 0x1e, 0xfb, 0x9c, 0xb9, 0x03, 0x29, 0x26, 0x44,
	0x1c, 0xb9, 0x9e, 0x1e, 0xa4, 0xd0, 0x02, 0x75, 0x3d, 0x54, 0xf9, 0x05, 0x74, 0xe6, 0xee, 0x8c,
	0x10, 0xc4, 0xb8, 0x2c, 0x65, 0x16, 0x8d, 0xa3, 0x49, 0xaf, 0xb0, 0x67, 0x34, 0x82, 0x64, 0x23,
	0xc5, 0x3d, 0xe5, 0xd9, 0x07, 0xcb, 0x7a, 0x84, 0x32, 0xe8, 0xac, 0x71, 0x85, 0x39, 0xa1, 0x59,
	0xcb, 0x16, 0x6a, 0x98, 0x3f, 0x46, 0x30, 0xf8, 0xcf, 0x38, 0x95, 0xef, 0xc9, 0x7e, 0x86, 0xb6,
	0x16, 0x1a, 0x57, 0x5e, 0xd5, 0x01, 0xf4, 0x05, 0xba, 0x8c, 0x13, 0x49, 0xb1, 0xaa, 0x55, 0x4f,
	0xb8, 0xf1, 0x90, 0x38, 0x78, 0xc8, 0x77, 0xf8, 0x48, 0x6f, 0x0f, 0x94, 0xe8, 0x55, 0x3d, 0xd9,
	0xb6, 0xf5, 0x17, 0xac, 0xe9, 0xdb, 0x9b, 0x57, 0x2d, 0xf4, 0xdd, 0xf2, 0x28, 0x19, 0xdf, 0x66,
	0x89, 0xeb, 0x0b, 0x59, 0xf4, 0x03, 0x86, 0x6e, 0xd2, 0x7a, 0x58, 0x54, 0x82, 0x5c, 0xab, 0xac,
	0x63, 0x5b, 0x5f, 0x17, 0xf2, 0x3f, 0xd0, 0xf5, 0x36, 0x15, 0xfa, 0x09, 0xa7, 0x54, 0xb3, 0x68,
	0xdc, 0x9a, 0xf4, 0x7f, 0x0f, 0xa7, 0x35, 0x31, 0xf5, 0x5d, 0xc5, 0x39, 0xf8, 0xa7, 0x08, 0xd2,
	0x66, 0x4e, 0x0a, 0xfd, 0x85, 0x74, 0xdf, 0x24, 0xbc, 0xca, 0xe8, 0xac, 0xd2, 0xec, 0x2f, 0xc2,
	0x66, 0xf3, 0x45, 0x14, 0x25, 0x82, 0x97, 0xca, 0x86, 0xda, 0x2a, 0x6a, 0x88, 0xbe, 0x41, 0x6a,
	0xf3, 0x5d, 0x85, 0xd9, 0x86, 0xa4, 0x09, 0x78, 0xed, 0xdc, 0xc6, 0x76, 0xdc, 0x23, 0x94, 0xc3,
	0xc0, 0xf9, 0xf6, 0x59, 0xb4, 0x6d, 0x35, 0xe0, 0xd0, 0x2f, 0xf8, 0xe4, 0xf0, 0x65, 0x70, 0x8f,
	0x4b, 0xf8, 0xad, 0x52, 0xfe, 0x10, 0x41, 0xf2, 0x4f, 0xf0, 0x0d, 0xdb, 0xa2, 0xaf, 0xd0, 0xbb,
	0xd9, 0x31, 0x4d, 0x2b, 0xa6, 0xb4, 0xb5, 0xdc, 0x2b, 0xce, 0x84, 0xb9, 0xfe, 0x4a, 0x1e, 0xc8,
	0x82, 0xf1, 0x72, 0x6e, 0xb6, 0xc8, 0x2d, 0x4c, 0xc0, 0x19, 0xeb, 0x25, 0xd6, 0x78, 0xc9, 0x64,
	0xbd, 0x8c, 0x1e, 0x1a, 0x6d, 0x97, 0x92, 0x19, 0x8d, 0x9d, 0xf6, 0x89, 0x40, 0x63, 0xe8, 0x93,
	0x1d, 0x66, 0x7c, 0x36, 0xdb, 0x09, 0xa5, 0xfd, 0xe2, 0x34, 0xa9, 0x75, 0x62, 0x7f, 0x97, 0xd9,
	0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x73, 0x90, 0xfb, 0x3f, 0x03, 0x00, 0x00,
}
