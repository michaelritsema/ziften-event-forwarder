// Code generated by protoc-gen-go.
// source: TestProtobufV2.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This message is used to test the handling of a newer message with an optional member that goes to an older client that
// does not support it (older client with newer server)
type TestProtobufV2 struct {
	Num1             *int64  `protobuf:"varint,1,req,name=num1" json:"num1,omitempty"`
	Str1             *string `protobuf:"bytes,2,req,name=str1" json:"str1,omitempty"`
	NewStr2          *string `protobuf:"bytes,3,opt,name=newStr2" json:"newStr2,omitempty"`
	NewNum2          *int64  `protobuf:"varint,4,opt,name=newNum2" json:"newNum2,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestProtobufV2) Reset()                    { *m = TestProtobufV2{} }
func (m *TestProtobufV2) String() string            { return proto.CompactTextString(m) }
func (*TestProtobufV2) ProtoMessage()               {}
func (*TestProtobufV2) Descriptor() ([]byte, []int) { return fileDescriptor86, []int{0} }

func (m *TestProtobufV2) GetNum1() int64 {
	if m != nil && m.Num1 != nil {
		return *m.Num1
	}
	return 0
}

func (m *TestProtobufV2) GetStr1() string {
	if m != nil && m.Str1 != nil {
		return *m.Str1
	}
	return ""
}

func (m *TestProtobufV2) GetNewStr2() string {
	if m != nil && m.NewStr2 != nil {
		return *m.NewStr2
	}
	return ""
}

func (m *TestProtobufV2) GetNewNum2() int64 {
	if m != nil && m.NewNum2 != nil {
		return *m.NewNum2
	}
	return 0
}

func init() {
	proto.RegisterType((*TestProtobufV2)(nil), "TestProtobufV2")
}

var fileDescriptor86 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x09, 0x49, 0x2d, 0x2e,
	0x09, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0x0b, 0x33, 0xd2, 0x2b, 0x00, 0x31, 0x95, 0xfc,
	0xb8, 0xf8, 0x50, 0xc5, 0x85, 0x78, 0xb8, 0x58, 0xf2, 0x4a, 0x73, 0x0d, 0x25, 0x18, 0x15, 0x98,
	0x34, 0x98, 0x41, 0xbc, 0xe2, 0x92, 0x22, 0x43, 0x09, 0x26, 0x20, 0x8f, 0x53, 0x88, 0x9f, 0x8b,
	0x3d, 0x2f, 0xb5, 0x3c, 0xb8, 0xa4, 0xc8, 0x48, 0x82, 0x59, 0x81, 0x11, 0x2e, 0xe0, 0x57, 0x9a,
	0x6b, 0x24, 0xc1, 0x02, 0x14, 0x60, 0x76, 0xb2, 0xe6, 0x52, 0x4a, 0xce, 0xcf, 0xd5, 0xab, 0xca,
	0x4c, 0x2b, 0x49, 0xcd, 0xd3, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x82, 0x58, 0x95, 0x9c, 0x9f,
	0xa3, 0x97, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0xea, 0x24, 0x8a, 0x6a, 0xa7, 0x2f, 0x44, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe9, 0x67, 0xe6, 0xa3, 0x00, 0x00, 0x00,
}