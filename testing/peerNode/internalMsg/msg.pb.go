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
	Cmd_ALLIP     Cmd_Type = 0
	Cmd_ALLID     Cmd_Type = 1
	Cmd_STARTTEST Cmd_Type = 2
	Cmd_SIGNIN    Cmd_Type = 3
	Cmd_TESTDONE  Cmd_Type = 4
)

var Cmd_Type_name = map[int32]string{
	0: "ALLIP",
	1: "ALLID",
	2: "STARTTEST",
	3: "SIGNIN",
	4: "TESTDONE",
}
var Cmd_Type_value = map[string]int32{
	"ALLIP":     0,
	"ALLID":     1,
	"STARTTEST": 2,
	"SIGNIN":    3,
	"TESTDONE":  4,
}

func (x Cmd_Type) String() string {
	return proto.EnumName(Cmd_Type_name, int32(x))
}
func (Cmd_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_msg_d703ecccea598dee, []int{0, 0}
}

type Cmd struct {
	Ctype                Cmd_Type `protobuf:"varint,1,opt,name=Ctype,proto3,enum=internalMsg.Cmd_Type" json:"Ctype,omitempty"`
	Args                 []byte   `protobuf:"bytes,2,opt,name=Args,proto3" json:"Args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_d703ecccea598dee, []int{0}
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
	return Cmd_ALLIP
}

func (m *Cmd) GetArgs() []byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Cmd)(nil), "internalMsg.Cmd")
	proto.RegisterEnum("internalMsg.Cmd_Type", Cmd_Type_name, Cmd_Type_value)
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_d703ecccea598dee) }

var fileDescriptor_msg_d703ecccea598dee = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xf1,
	0x2d, 0x4e, 0x57, 0xea, 0x65, 0xe4, 0x62, 0x76, 0xce, 0x4d, 0x11, 0xd2, 0xe6, 0x62, 0x75, 0x2e,
	0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd5, 0x43, 0x52, 0xa4, 0xe7,
	0x9c, 0x9b, 0xa2, 0x17, 0x52, 0x59, 0x90, 0x1a, 0x04, 0x51, 0x23, 0x24, 0xc4, 0xc5, 0xe2, 0x58,
	0x94, 0x5e, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x2b, 0xb9, 0x72, 0xb1, 0x80,
	0x94, 0x08, 0x71, 0x72, 0xb1, 0x3a, 0xfa, 0xf8, 0x78, 0x06, 0x08, 0x30, 0xc0, 0x98, 0x2e, 0x02,
	0x8c, 0x42, 0xbc, 0x5c, 0x9c, 0xc1, 0x21, 0x8e, 0x41, 0x21, 0x21, 0xae, 0xc1, 0x21, 0x02, 0x4c,
	0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0x9e, 0xee, 0x7e, 0x9e, 0x7e, 0x02, 0xcc, 0x42, 0x3c, 0x5c, 0x1c,
	0x20, 0x51, 0x17, 0x7f, 0x3f, 0x57, 0x01, 0x96, 0x24, 0x36, 0xb0, 0x1b, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x69, 0x65, 0xba, 0xca, 0xb0, 0x00, 0x00, 0x00,
}
