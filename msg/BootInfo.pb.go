// Code generated by protoc-gen-go.
// source: BootInfo.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BootInfo struct {
	// The UTC time in FILETIME format when the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to identify agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// GUID that ties this message to the boot profile and analysis messages
	AnalysisGUID *string `protobuf:"bytes,3,req,name=analysisGUID" json:"analysisGUID,omitempty"`
	// Time when ziften agent detected session start in FILETIME format
	SessionStartTime *int64 `protobuf:"varint,4,req,name=sessionStartTime" json:"sessionStartTime,omitempty"`
	// Time when ziften agent detected session start in ticks (milliseconds)
	SessionStartTimeTicks *int64 `protobuf:"varint,5,req,name=sessionStartTimeTicks" json:"sessionStartTimeTicks,omitempty"`
	// How long ziften agent has been running when the user session start in ticks (milliseconds)
	ProcessTimeAtSessionStartTicks *int64 `protobuf:"varint,6,req,name=ProcessTimeAtSessionStartTicks" json:"ProcessTimeAtSessionStartTicks,omitempty"`
	// Time when ziften agent decided that boot ended in FILETIME format
	BootEndTime *int64 `protobuf:"varint,7,req,name=bootEndTime" json:"bootEndTime,omitempty"`
	// Time when ziften agent decided that boot ended
	BootEndTimeTicks *int64  `protobuf:"varint,8,req,name=bootEndTimeTicks" json:"bootEndTimeTicks,omitempty"`
	SiteId           *string `protobuf:"bytes,9,opt,name=siteId" json:"siteId,omitempty"`
	Uuid             *string `protobuf:"bytes,10,opt,name=uuid" json:"uuid,omitempty"`
	// Tie this to the actual system boot (uses UUID)
	BootAnalysisGUID *string `protobuf:"bytes,11,opt,name=bootAnalysisGUID" json:"bootAnalysisGUID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BootInfo) Reset()                    { *m = BootInfo{} }
func (m *BootInfo) String() string            { return proto.CompactTextString(m) }
func (*BootInfo) ProtoMessage()               {}
func (*BootInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *BootInfo) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *BootInfo) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *BootInfo) GetAnalysisGUID() string {
	if m != nil && m.AnalysisGUID != nil {
		return *m.AnalysisGUID
	}
	return ""
}

func (m *BootInfo) GetSessionStartTime() int64 {
	if m != nil && m.SessionStartTime != nil {
		return *m.SessionStartTime
	}
	return 0
}

func (m *BootInfo) GetSessionStartTimeTicks() int64 {
	if m != nil && m.SessionStartTimeTicks != nil {
		return *m.SessionStartTimeTicks
	}
	return 0
}

func (m *BootInfo) GetProcessTimeAtSessionStartTicks() int64 {
	if m != nil && m.ProcessTimeAtSessionStartTicks != nil {
		return *m.ProcessTimeAtSessionStartTicks
	}
	return 0
}

func (m *BootInfo) GetBootEndTime() int64 {
	if m != nil && m.BootEndTime != nil {
		return *m.BootEndTime
	}
	return 0
}

func (m *BootInfo) GetBootEndTimeTicks() int64 {
	if m != nil && m.BootEndTimeTicks != nil {
		return *m.BootEndTimeTicks
	}
	return 0
}

func (m *BootInfo) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *BootInfo) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *BootInfo) GetBootAnalysisGUID() string {
	if m != nil && m.BootAnalysisGUID != nil {
		return *m.BootAnalysisGUID
	}
	return ""
}

func init() {
	proto.RegisterType((*BootInfo)(nil), "BootInfo")
}

var fileDescriptor13 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x5a, 0x6b, 0x32, 0xd5, 0x5a, 0x16, 0x94, 0x45, 0xb1, 0x84, 0x1c, 0xb4, 0xa7,
	0xdc, 0x7c, 0x80, 0x06, 0x45, 0x72, 0x10, 0x0a, 0x8d, 0x0f, 0x10, 0x93, 0x6d, 0x59, 0x6c, 0x76,
	0xc3, 0xce, 0x54, 0xa8, 0x47, 0xdf, 0xc9, 0x27, 0xf2, 0x45, 0xdc, 0x6c, 0x28, 0xa4, 0xa5, 0xc7,
	0xfc, 0xff, 0x37, 0xf3, 0xe7, 0xdf, 0x81, 0x71, 0xa2, 0x35, 0xa5, 0x6a, 0xa5, 0xe3, 0xda, 0x68,
	0xd2, 0xb7, 0x6c, 0xa1, 0x91, 0xd6, 0x46, 0x60, 0xb6, 0xab, 0x45, 0xab, 0x45, 0x7f, 0x1e, 0xf8,
	0x7b, 0x8c, 0x85, 0x10, 0x90, 0xac, 0xc4, 0x92, 0xf2, 0xaa, 0xe6, 0xbd, 0xd0, 0x9b, 0xf5, 0x93,
	0xcb, 0x9f, 0x5f, 0xee, 0x44, 0x6c, 0x44, 0x76, 0x07, 0x41, 0xbe, 0x16, 0x8a, 0x5e, 0xdf, 0xd3,
	0x67, 0xee, 0x59, 0x22, 0x48, 0x7c, 0x4b, 0x0c, 0xb6, 0x5b, 0x59, 0xb2, 0x29, 0x5c, 0xe4, 0x2a,
	0xdf, 0xec, 0x50, 0xa2, 0xf3, 0xfb, 0x47, 0xfe, 0x23, 0x4c, 0x50, 0x20, 0x4a, 0xad, 0x6c, 0x82,
	0xa1, 0xcc, 0x6e, 0xe5, 0x83, 0x53, 0x29, 0xf7, 0x70, 0x7d, 0x0c, 0x66, 0xb2, 0xf8, 0x44, 0x7e,
	0xd6, 0xd0, 0xec, 0x01, 0xa6, 0x0b, 0xa3, 0x0b, 0x4b, 0x34, 0xce, 0x9c, 0x96, 0x07, 0x6c, 0xc3,
	0x0d, 0x1d, 0x17, 0xc1, 0xe8, 0xc3, 0x56, 0x7b, 0x51, 0xa5, 0x8b, 0x3a, 0x3f, 0x15, 0xc5, 0x61,
	0xd2, 0x61, 0xda, 0x69, 0xdf, 0x4d, 0x8f, 0x61, 0x88, 0x92, 0x44, 0x5a, 0xf2, 0x20, 0xec, 0xcd,
	0x02, 0x76, 0x03, 0xae, 0x05, 0x87, 0xe6, 0xab, 0xd3, 0x2a, 0x6a, 0x37, 0xcc, 0xbb, 0xcd, 0x47,
	0x87, 0x4c, 0xf2, 0x04, 0x51, 0xa1, 0xab, 0xf8, 0x5b, 0xae, 0x48, 0xa8, 0x18, 0x85, 0xf9, 0x12,
	0xa6, 0x3d, 0x40, 0xa1, 0x37, 0xb1, 0xfd, 0x17, 0xb4, 0xaf, 0x9a, 0x5c, 0xed, 0x0f, 0xf1, 0xd6,
	0x0a, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x06, 0xad, 0x2f, 0xc1, 0x01, 0x00, 0x00,
}
