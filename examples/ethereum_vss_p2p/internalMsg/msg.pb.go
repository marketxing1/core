// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package internalMsg

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

type Cmd_Type int32

const (
	Cmd_CHECKURL Cmd_Type = 0
)

var Cmd_Type_name = map[int32]string{
	0: "CHECKURL",
}
var Cmd_Type_value = map[string]int32{
	"CHECKURL": 0,
}

func (x Cmd_Type) String() string {
	return proto.EnumName(Cmd_Type_name, int32(x))
}
func (Cmd_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{8, 0}
}

type EncryptedDeal struct {
	DHKey                []byte   `protobuf:"bytes,1,opt,name=DHKey,json=dHKey,proto3" json:"DHKey,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=Signature,json=signature,proto3" json:"Signature,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=Nonce,json=nonce,proto3" json:"Nonce,omitempty"`
	Cipher               []byte   `protobuf:"bytes,4,opt,name=Cipher,json=cipher,proto3" json:"Cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedDeal) Reset()         { *m = EncryptedDeal{} }
func (m *EncryptedDeal) String() string { return proto.CompactTextString(m) }
func (*EncryptedDeal) ProtoMessage()    {}
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{0}
}
func (m *EncryptedDeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedDeal.Unmarshal(m, b)
}
func (m *EncryptedDeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedDeal.Marshal(b, m, deterministic)
}
func (dst *EncryptedDeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedDeal.Merge(dst, src)
}
func (m *EncryptedDeal) XXX_Size() int {
	return xxx_messageInfo_EncryptedDeal.Size(m)
}
func (m *EncryptedDeal) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedDeal.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedDeal proto.InternalMessageInfo

func (m *EncryptedDeal) GetDHKey() []byte {
	if m != nil {
		return m.DHKey
	}
	return nil
}

func (m *EncryptedDeal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EncryptedDeal) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptedDeal) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type EncryptedDeals struct {
	Deals                []*EncryptedDeal `protobuf:"bytes,1,rep,name=Deals,json=deals,proto3" json:"Deals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EncryptedDeals) Reset()         { *m = EncryptedDeals{} }
func (m *EncryptedDeals) String() string { return proto.CompactTextString(m) }
func (*EncryptedDeals) ProtoMessage()    {}
func (*EncryptedDeals) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{1}
}
func (m *EncryptedDeals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedDeals.Unmarshal(m, b)
}
func (m *EncryptedDeals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedDeals.Marshal(b, m, deterministic)
}
func (dst *EncryptedDeals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedDeals.Merge(dst, src)
}
func (m *EncryptedDeals) XXX_Size() int {
	return xxx_messageInfo_EncryptedDeals.Size(m)
}
func (m *EncryptedDeals) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedDeals.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedDeals proto.InternalMessageInfo

func (m *EncryptedDeals) GetDeals() []*EncryptedDeal {
	if m != nil {
		return m.Deals
	}
	return nil
}

type Response struct {
	SessionID            []byte   `protobuf:"bytes,1,opt,name=SessionID,json=sessionID,proto3" json:"SessionID,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=Index,json=index,proto3" json:"Index,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=Signature,json=signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSessionID() []byte {
	if m != nil {
		return m.SessionID
	}
	return nil
}

func (m *Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Responses struct {
	Responses            []*Response `protobuf:"bytes,1,rep,name=Responses,json=responses,proto3" json:"Responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Responses) Reset()         { *m = Responses{} }
func (m *Responses) String() string { return proto.CompactTextString(m) }
func (*Responses) ProtoMessage()    {}
func (*Responses) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{3}
}
func (m *Responses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Responses.Unmarshal(m, b)
}
func (m *Responses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Responses.Marshal(b, m, deterministic)
}
func (dst *Responses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Responses.Merge(dst, src)
}
func (m *Responses) XXX_Size() int {
	return xxx_messageInfo_Responses.Size(m)
}
func (m *Responses) XXX_DiscardUnknown() {
	xxx_messageInfo_Responses.DiscardUnknown(m)
}

var xxx_messageInfo_Responses proto.InternalMessageInfo

func (m *Responses) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Justification struct {
	SessionID            []byte   `protobuf:"bytes,1,opt,name=SessionID,json=sessionID,proto3" json:"SessionID,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=Index,json=index,proto3" json:"Index,omitempty"`
	Deal                 []byte   `protobuf:"bytes,3,opt,name=Deal,json=deal,proto3" json:"Deal,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=Signature,json=signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Justification) Reset()         { *m = Justification{} }
func (m *Justification) String() string { return proto.CompactTextString(m) }
func (*Justification) ProtoMessage()    {}
func (*Justification) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{4}
}
func (m *Justification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Justification.Unmarshal(m, b)
}
func (m *Justification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Justification.Marshal(b, m, deterministic)
}
func (dst *Justification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Justification.Merge(dst, src)
}
func (m *Justification) XXX_Size() int {
	return xxx_messageInfo_Justification.Size(m)
}
func (m *Justification) XXX_DiscardUnknown() {
	xxx_messageInfo_Justification.DiscardUnknown(m)
}

var xxx_messageInfo_Justification proto.InternalMessageInfo

func (m *Justification) GetSessionID() []byte {
	if m != nil {
		return m.SessionID
	}
	return nil
}

func (m *Justification) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Justification) GetDeal() []byte {
	if m != nil {
		return m.Deal
	}
	return nil
}

func (m *Justification) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PublicKey struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=PublicKey,json=publicKey,proto3" json:"PublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{5}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type PublicKeys struct {
	PublicKey            []*PublicKey `protobuf:"bytes,1,rep,name=PublicKey,json=publicKey,proto3" json:"PublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PublicKeys) Reset()         { *m = PublicKeys{} }
func (m *PublicKeys) String() string { return proto.CompactTextString(m) }
func (*PublicKeys) ProtoMessage()    {}
func (*PublicKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{6}
}
func (m *PublicKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeys.Unmarshal(m, b)
}
func (m *PublicKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeys.Marshal(b, m, deterministic)
}
func (dst *PublicKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeys.Merge(dst, src)
}
func (m *PublicKeys) XXX_Size() int {
	return xxx_messageInfo_PublicKeys.Size(m)
}
func (m *PublicKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeys.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeys proto.InternalMessageInfo

func (m *PublicKeys) GetPublicKey() []*PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type Signature struct {
	Index                uint32   `protobuf:"varint,1,opt,name=Index,json=index,proto3" json:"Index,omitempty"`
	QueryId              string   `protobuf:"bytes,2,opt,name=QueryId,json=queryId,proto3" json:"QueryId,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=Content,json=content,proto3" json:"Content,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=Signature,json=signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{7}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Signature) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Signature) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Cmd struct {
	Ctype                Cmd_Type `protobuf:"varint,1,opt,name=Ctype,json=ctype,proto3,enum=internalMsg.Cmd_Type" json:"Ctype,omitempty"`
	Args                 string   `protobuf:"bytes,2,opt,name=Args,json=args,proto3" json:"Args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e22fc611fbeebddd, []int{8}
}
func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (dst *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(dst, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetCtype() Cmd_Type {
	if m != nil {
		return m.Ctype
	}
	return Cmd_CHECKURL
}

func (m *Cmd) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func init() {
	proto.RegisterType((*EncryptedDeal)(nil), "internalMsg.EncryptedDeal")
	proto.RegisterType((*EncryptedDeals)(nil), "internalMsg.EncryptedDeals")
	proto.RegisterType((*Response)(nil), "internalMsg.Response")
	proto.RegisterType((*Responses)(nil), "internalMsg.Responses")
	proto.RegisterType((*Justification)(nil), "internalMsg.Justification")
	proto.RegisterType((*PublicKey)(nil), "internalMsg.PublicKey")
	proto.RegisterType((*PublicKeys)(nil), "internalMsg.PublicKeys")
	proto.RegisterType((*Signature)(nil), "internalMsg.Signature")
	proto.RegisterType((*Cmd)(nil), "internalMsg.Cmd")
	proto.RegisterEnum("internalMsg.Cmd_Type", Cmd_Type_name, Cmd_Type_value)
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_e22fc611fbeebddd) }

var fileDescriptor_msg_e22fc611fbeebddd = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0xaf, 0xd2, 0x30,
	0x14, 0xc6, 0x9d, 0x6c, 0x83, 0x9e, 0x7b, 0xb9, 0x31, 0xcd, 0x95, 0x2c, 0xc6, 0x07, 0xb2, 0x27,
	0x8c, 0xc9, 0x62, 0xc4, 0x3f, 0x40, 0x19, 0x24, 0x20, 0x6a, 0xb4, 0xe8, 0x9b, 0x2f, 0x63, 0xab,
	0xb3, 0x09, 0x74, 0xa3, 0xa7, 0x4b, 0xdc, 0x7f, 0x6f, 0x5a, 0x36, 0xdc, 0x78, 0x21, 0xf1, 0x6d,
	0xdf, 0x39, 0x5f, 0x77, 0x7e, 0xdf, 0x69, 0x0a, 0xe4, 0x88, 0x79, 0x54, 0xaa, 0x42, 0x17, 0xf4,
	0x4e, 0x48, 0xcd, 0x95, 0x4c, 0x0e, 0x9f, 0x31, 0x0f, 0x4f, 0x30, 0x5e, 0xc9, 0x54, 0xd5, 0xa5,
	0xe6, 0xd9, 0x92, 0x27, 0x07, 0xfa, 0x08, 0xde, 0x72, 0xbd, 0xe5, 0x75, 0xe0, 0x4c, 0x9d, 0xd9,
	0x3d, 0xf3, 0x32, 0x23, 0xe8, 0x4b, 0x20, 0x3b, 0x91, 0xcb, 0x44, 0x57, 0x8a, 0x07, 0x4f, 0x6d,
	0x87, 0x60, 0x5b, 0x30, 0x67, 0xbe, 0x14, 0x32, 0xe5, 0xc1, 0xe0, 0x7c, 0x46, 0x1a, 0x41, 0x27,
	0xe0, 0xc7, 0xa2, 0xfc, 0xcd, 0x55, 0xe0, 0xda, 0xb2, 0x9f, 0x5a, 0x15, 0x2e, 0xe0, 0xa1, 0x37,
	0x12, 0xe9, 0x1b, 0xf0, 0xec, 0x47, 0xe0, 0x4c, 0x07, 0xb3, 0xbb, 0xb7, 0x2f, 0xa2, 0x0e, 0x61,
	0xd4, 0xf3, 0x32, 0x2f, 0x33, 0xc6, 0x50, 0xc3, 0x88, 0x71, 0x2c, 0x0b, 0x89, 0xdc, 0xb2, 0x71,
	0x44, 0x51, 0xc8, 0xcd, 0xb2, 0xa1, 0x26, 0xd8, 0x16, 0x0c, 0xdb, 0x46, 0x66, 0xfc, 0x8f, 0xa5,
	0x1e, 0x33, 0x4f, 0x18, 0x61, 0xd8, 0x76, 0x3a, 0xd1, 0x15, 0x5a, 0xe4, 0x11, 0xf3, 0xd1, 0xaa,
	0x7e, 0x4e, 0xf7, 0x2a, 0x67, 0xf8, 0x1e, 0x48, 0x3b, 0x15, 0xe9, 0xbc, 0x23, 0x1a, 0xf0, 0xe7,
	0x3d, 0xf0, 0xb6, 0xcb, 0x88, 0x6a, 0x7d, 0x61, 0x05, 0xe3, 0x8f, 0x15, 0x6a, 0xf1, 0x4b, 0xa4,
	0x89, 0x16, 0x85, 0xfc, 0x2f, 0x78, 0x0a, 0xae, 0xd9, 0x45, 0xb3, 0x6d, 0xd7, 0x6c, 0xe4, 0x06,
	0xf8, 0x2b, 0x20, 0x5f, 0xab, 0xfd, 0x41, 0xa4, 0xcd, 0x5d, 0x5e, 0x44, 0x3b, 0xb2, 0x6c, 0x0b,
	0xe1, 0x02, 0xe0, 0xd2, 0x45, 0xfa, 0xae, 0xef, 0x35, 0x21, 0x27, 0xbd, 0x90, 0x97, 0x6e, 0xf7,
	0x1f, 0xd8, 0x81, 0xf9, 0x97, 0xc1, 0xe9, 0x66, 0x08, 0x60, 0xf8, 0xad, 0xe2, 0xaa, 0xde, 0x64,
	0x36, 0x1b, 0x61, 0xc3, 0xd3, 0x59, 0x9a, 0x4e, 0x5c, 0x48, 0xcd, 0xa5, 0x6e, 0x02, 0x0e, 0xd3,
	0xb3, 0xbc, 0x91, 0xf1, 0x27, 0x0c, 0xe2, 0x63, 0x46, 0x5f, 0x83, 0x17, 0xeb, 0xba, 0xe4, 0x76,
	0xdc, 0xc3, 0xd5, 0x95, 0xc4, 0xc7, 0x2c, 0xfa, 0x5e, 0x97, 0x9c, 0x79, 0xa9, 0xf1, 0x98, 0x4d,
	0x7e, 0x50, 0x39, 0x36, 0x08, 0x6e, 0xa2, 0x72, 0x0c, 0x1f, 0xc1, 0x35, 0x16, 0x7a, 0x0f, 0xa3,
	0x78, 0xbd, 0x8a, 0xb7, 0x3f, 0xd8, 0xa7, 0x67, 0x4f, 0xf6, 0xbe, 0x7d, 0x3b, 0xf3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x94, 0xab, 0x4d, 0x48, 0x03, 0x00, 0x00,
}