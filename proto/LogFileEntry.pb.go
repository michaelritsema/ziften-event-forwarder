// Code generated by protoc-gen-go.
// source: LogFileEntry.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogFileEntry struct {
	// The UTC time (as known to the client computer) when the agent status message was generated
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	Logfile   *string `protobuf:"bytes,3,req,name=logfile" json:"logfile,omitempty"`
	Severity  *int32  `protobuf:"varint,4,req,name=severity" json:"severity,omitempty"`
	// the following match regexp groups >= 0 and are reserved
	Message          *string `protobuf:"bytes,10,req,name=message" json:"message,omitempty"`
	SiteId           *string `protobuf:"bytes,11,opt,name=siteId" json:"siteId,omitempty"`
	Subsystem        *string `protobuf:"bytes,12,opt,name=subsystem" json:"subsystem,omitempty"`
	Uuid             *string `protobuf:"bytes,13,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogFileEntry) Reset()                    { *m = LogFileEntry{} }
func (m *LogFileEntry) String() string            { return proto1.CompactTextString(m) }
func (*LogFileEntry) ProtoMessage()               {}
func (*LogFileEntry) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

func (m *LogFileEntry) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *LogFileEntry) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *LogFileEntry) GetLogfile() string {
	if m != nil && m.Logfile != nil {
		return *m.Logfile
	}
	return ""
}

func (m *LogFileEntry) GetSeverity() int32 {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return 0
}

func (m *LogFileEntry) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *LogFileEntry) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *LogFileEntry) GetSubsystem() string {
	if m != nil && m.Subsystem != nil {
		return *m.Subsystem
	}
	return ""
}

func (m *LogFileEntry) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto1.RegisterType((*LogFileEntry)(nil), "LogFileEntry")
}

var fileDescriptor32 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x5b, 0xb5, 0x19, 0x5b, 0x3f, 0x56, 0x90, 0x41, 0x2f, 0xa1, 0x27, 0x4f, 0xb9,
	0x7b, 0x0d, 0x7e, 0x50, 0x50, 0x10, 0xd4, 0x1f, 0x50, 0xdb, 0x69, 0x58, 0x48, 0xb2, 0x61, 0x67,
	0x22, 0xc4, 0xa3, 0xff, 0xc9, 0x9b, 0x3f, 0xce, 0xd9, 0x8d, 0xd0, 0x1e, 0xe7, 0x99, 0x87, 0xf7,
	0xe5, 0x05, 0xf3, 0xe4, 0xca, 0x07, 0x5b, 0xd1, 0x7d, 0x23, 0xbe, 0xcf, 0x5b, 0xef, 0xc4, 0x5d,
	0x99, 0x17, 0xc7, 0x52, 0x7a, 0xe2, 0xb7, 0xbe, 0xa5, 0x81, 0x2d, 0x7e, 0x13, 0x98, 0xed, 0xab,
	0x26, 0x83, 0x54, 0x6c, 0x4d, 0xaf, 0xb2, 0xaa, 0x5b, 0x4c, 0xb2, 0xd1, 0xcd, 0xb8, 0x98, 0x7f,
	0xff, 0x60, 0x84, 0x1c, 0xa0, 0xb9, 0x86, 0x74, 0x55, 0x52, 0x23, 0x8f, 0xef, 0xcb, 0x3b, 0x1c,
	0xa9, 0x91, 0x16, 0x53, 0x35, 0x26, 0x5d, 0x67, 0x37, 0xe6, 0x14, 0x8e, 0x2a, 0x57, 0x6e, 0x35,
	0x0e, 0xc7, 0xe1, 0x65, 0xce, 0x60, 0xca, 0xf4, 0x49, 0xde, 0x4a, 0x8f, 0x13, 0x25, 0x07, 0x41,
	0xd1, 0x28, 0xd6, 0x0c, 0x84, 0xa8, 0x9c, 0xc0, 0x21, 0x5b, 0xa1, 0xe5, 0x06, 0x8f, 0xb3, 0x44,
	0xef, 0x73, 0x48, 0xb9, 0xfb, 0xe0, 0x9e, 0x85, 0x6a, 0x9c, 0x45, 0x74, 0x09, 0x31, 0x1e, 0xe7,
	0xe1, 0xda, 0xd5, 0x15, 0xb7, 0xb0, 0x58, 0xbb, 0x3a, 0xff, 0xb2, 0x5b, 0xa1, 0x26, 0x67, 0xf2,
	0xda, 0x34, 0x2c, 0x5b, 0xbb, 0x2a, 0xff, 0xaf, 0x29, 0x2e, 0xf6, 0x17, 0x3e, 0x0f, 0xf0, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0xee, 0x7b, 0xc0, 0x26, 0x22, 0x01, 0x00, 0x00,
}