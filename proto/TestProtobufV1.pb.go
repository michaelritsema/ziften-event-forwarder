// Code generated by protoc-gen-go.
// source: TestProtobufV1.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This message is used to test the handling of a newer message with an optional member that goes to an older client that
// does not support it (older client with newer server)
type TestProtobufV1 struct {
	Num1             *int64  `protobuf:"varint,1,req,name=num1" json:"num1,omitempty"`
	Str1             *string `protobuf:"bytes,2,req,name=str1" json:"str1,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestProtobufV1) Reset()                    { *m = TestProtobufV1{} }
func (m *TestProtobufV1) String() string            { return proto1.CompactTextString(m) }
func (*TestProtobufV1) ProtoMessage()               {}
func (*TestProtobufV1) Descriptor() ([]byte, []int) { return fileDescriptor85, []int{0} }

func (m *TestProtobufV1) GetNum1() int64 {
	if m != nil && m.Num1 != nil {
		return *m.Num1
	}
	return 0
}

func (m *TestProtobufV1) GetStr1() string {
	if m != nil && m.Str1 != nil {
		return *m.Str1
	}
	return ""
}

func init() {
	proto1.RegisterType((*TestProtobufV1)(nil), "TestProtobufV1")
}

var fileDescriptor85 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x09, 0x49, 0x2d, 0x2e,
	0x09, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0x0b, 0x33, 0xd4, 0x2b, 0x00, 0x31, 0x95, 0x74,
	0xb8, 0xf8, 0x50, 0xc5, 0x85, 0x78, 0xb8, 0x58, 0xf2, 0x4a, 0x73, 0x0d, 0x25, 0x18, 0x15, 0x98,
	0x34, 0x98, 0x41, 0xbc, 0xe2, 0x92, 0x22, 0x43, 0x09, 0x26, 0x20, 0x8f, 0xd3, 0xc9, 0x9a, 0x4b,
	0x29, 0x39, 0x3f, 0x57, 0xaf, 0x2a, 0x33, 0xad, 0x24, 0x35, 0x4f, 0xaf, 0x38, 0xb5, 0xa8, 0x2c,
	0xb5, 0x08, 0x62, 0x50, 0x72, 0x7e, 0x8e, 0x5e, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x93,
	0x28, 0xaa, 0x89, 0xbe, 0x10, 0x61, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0xa0, 0x32, 0xdf,
	0x81, 0x00, 0x00, 0x00,
}
