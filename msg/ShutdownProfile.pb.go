// Code generated by protoc-gen-go.
// source: ShutdownProfile.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShutdownProfile struct {
	// The UTC time when the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// The number of milliseconds that have elapsed since the system was started.
	// GetTickCount() is called on shutdown and this field is set to the value returned by the API
	BootSpanInTicks  *int64   `protobuf:"varint,3,req,name=bootSpanInTicks" json:"bootSpanInTicks,omitempty"`
	SiteId           *string  `protobuf:"bytes,4,opt,name=siteId" json:"siteId,omitempty"`
	Uuid             *string  `protobuf:"bytes,5,opt,name=uuid" json:"uuid,omitempty"`
	Module           []string `protobuf:"bytes,6,rep,name=module" json:"module,omitempty"`
	ShutdownMS       []int32  `protobuf:"varint,7,rep,name=shutdownMS" json:"shutdownMS,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ShutdownProfile) Reset()                    { *m = ShutdownProfile{} }
func (m *ShutdownProfile) String() string            { return proto.CompactTextString(m) }
func (*ShutdownProfile) ProtoMessage()               {}
func (*ShutdownProfile) Descriptor() ([]byte, []int) { return fileDescriptor79, []int{0} }

func (m *ShutdownProfile) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *ShutdownProfile) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *ShutdownProfile) GetBootSpanInTicks() int64 {
	if m != nil && m.BootSpanInTicks != nil {
		return *m.BootSpanInTicks
	}
	return 0
}

func (m *ShutdownProfile) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *ShutdownProfile) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ShutdownProfile) GetModule() []string {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *ShutdownProfile) GetShutdownMS() []int32 {
	if m != nil {
		return m.ShutdownMS
	}
	return nil
}

func init() {
	proto.RegisterType((*ShutdownProfile)(nil), "ShutdownProfile")
}

var fileDescriptor79 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xd3, 0x56, 0x33, 0xa0, 0x85, 0x05, 0xeb, 0xa2, 0x97, 0xd0, 0x53, 0x4f, 0x79,
	0x02, 0x4f, 0x41, 0x90, 0x1c, 0x0a, 0x85, 0xd4, 0x07, 0x88, 0xc9, 0xb4, 0x2e, 0x26, 0x3b, 0x61,
	0x67, 0x56, 0xd1, 0xa3, 0xef, 0xd4, 0xf7, 0x73, 0xd3, 0xf6, 0x20, 0x39, 0xce, 0xc7, 0x37, 0x3f,
	0xff, 0x0f, 0x77, 0xe5, 0xbb, 0x97, 0x86, 0xbe, 0xec, 0xd6, 0xd1, 0xde, 0xb4, 0x98, 0xf5, 0x8e,
	0x84, 0x1e, 0xd4, 0x96, 0x58, 0x0e, 0x0e, 0x79, 0xf7, 0xdd, 0x5f, 0xd8, 0xea, 0x18, 0xc1, 0x62,
	0x64, 0xab, 0x14, 0x12, 0x31, 0x1d, 0x96, 0x52, 0x75, 0xbd, 0x8e, 0xd2, 0xc9, 0x3a, 0xce, 0x6f,
	0x7e, 0x8f, 0xfa, 0x04, 0x79, 0x80, 0xea, 0x11, 0x92, 0xea, 0x80, 0x56, 0x5e, 0x5e, 0x8b, 0x67,
	0x3d, 0x09, 0x46, 0x92, 0x5f, 0x07, 0x63, 0xea, 0xbd, 0x69, 0xd4, 0x3d, 0x2c, 0xde, 0x88, 0xa4,
	0xec, 0x2b, 0x5b, 0xd8, 0x9d, 0xa9, 0x3f, 0x58, 0xc7, 0x43, 0x88, 0xba, 0x85, 0x39, 0x1b, 0xc1,
	0xa2, 0xd1, 0xd3, 0x34, 0x5a, 0x27, 0x6a, 0x09, 0xa7, 0x07, 0x3d, 0x1b, 0xae, 0x7f, 0x01, 0xc1,
	0xeb, 0xa8, 0xf1, 0x2d, 0xea, 0x79, 0x1a, 0x07, 0x4f, 0x01, 0xf0, 0xa5, 0xe2, 0xa6, 0xd4, 0x57,
	0x81, 0xcd, 0xf2, 0x27, 0x58, 0xd5, 0xd4, 0x65, 0x3f, 0x66, 0x2f, 0x68, 0x33, 0x46, 0xf7, 0x89,
	0xee, 0x3c, 0xa9, 0xa6, 0x36, 0x0b, 0x3d, 0x39, 0x14, 0xcc, 0x97, 0xa3, 0x69, 0x9b, 0x33, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x16, 0x96, 0x68, 0x71, 0x21, 0x01, 0x00, 0x00,
}