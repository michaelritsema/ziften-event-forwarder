// Code generated by protoc-gen-go.
// source: OSXBootProfile.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXBootProfile struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Time required to boot to the system. The time the login panel was shown is subtracted to the total boot time reported by the boot trace.
	BootTime *int64 `protobuf:"varint,3,req,name=bootTime" json:"bootTime,omitempty"`
	// Time required to login to the system
	LoginTime *int64 `protobuf:"varint,4,opt,name=loginTime" json:"loginTime,omitempty"`
	// GUID that ties this message to the boot analysis message containing the info from the boot trace
	AnalysisGUID     *string `protobuf:"bytes,5,opt,name=analysisGUID" json:"analysisGUID,omitempty"`
	PresmssTime      *int64  `protobuf:"varint,6,opt,name=presmssTime" json:"presmssTime,omitempty"`
	SiteId           *string `protobuf:"bytes,7,opt,name=siteId" json:"siteId,omitempty"`
	BootStartTime    *int64  `protobuf:"varint,8,opt,name=bootStartTime" json:"bootStartTime,omitempty"`
	Uuid             *string `protobuf:"bytes,9,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSXBootProfile) Reset()                    { *m = OSXBootProfile{} }
func (m *OSXBootProfile) String() string            { return proto.CompactTextString(m) }
func (*OSXBootProfile) ProtoMessage()               {}
func (*OSXBootProfile) Descriptor() ([]byte, []int) { return fileDescriptor45, []int{0} }

func (m *OSXBootProfile) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXBootProfile) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXBootProfile) GetBootTime() int64 {
	if m != nil && m.BootTime != nil {
		return *m.BootTime
	}
	return 0
}

func (m *OSXBootProfile) GetLoginTime() int64 {
	if m != nil && m.LoginTime != nil {
		return *m.LoginTime
	}
	return 0
}

func (m *OSXBootProfile) GetAnalysisGUID() string {
	if m != nil && m.AnalysisGUID != nil {
		return *m.AnalysisGUID
	}
	return ""
}

func (m *OSXBootProfile) GetPresmssTime() int64 {
	if m != nil && m.PresmssTime != nil {
		return *m.PresmssTime
	}
	return 0
}

func (m *OSXBootProfile) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXBootProfile) GetBootStartTime() int64 {
	if m != nil && m.BootStartTime != nil {
		return *m.BootStartTime
	}
	return 0
}

func (m *OSXBootProfile) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXBootProfile)(nil), "OSXBootProfile")
}

var fileDescriptor45 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0x5b, 0x6b, 0xe6, 0x6a, 0x8b, 0x8e, 0x3f, 0x0c, 0x0a, 0x12, 0x8a, 0x0b, 0x57,
	0x79, 0x01, 0x77, 0x83, 0x20, 0x5d, 0x88, 0x85, 0x56, 0x70, 0x1b, 0xdb, 0x69, 0x18, 0x48, 0x72,
	0xc3, 0xdc, 0xa9, 0x50, 0x97, 0xbe, 0x93, 0xcf, 0xe6, 0xd6, 0x3b, 0x13, 0x17, 0xc6, 0xed, 0xc7,
	0x77, 0xce, 0x3d, 0x33, 0x70, 0xfe, 0xbc, 0x7c, 0xd5, 0x88, 0x7e, 0xe1, 0x70, 0x6b, 0x2b, 0x93,
	0xb7, 0x0e, 0x3d, 0x5e, 0xc9, 0x05, 0x92, 0x2f, 0x9d, 0xa1, 0xd5, 0xbe, 0xfd, 0x65, 0xb3, 0xef,
	0x04, 0xa6, 0x7d, 0x59, 0x66, 0x20, 0xbc, 0xad, 0xcd, 0xd2, 0x17, 0x75, 0xab, 0x92, 0x6c, 0x70,
	0x37, 0xd4, 0x93, 0xcf, 0x2f, 0x15, 0x21, 0x05, 0x28, 0xaf, 0x41, 0x14, 0xa5, 0x69, 0xfc, 0xe3,
	0xcb, 0xfc, 0x41, 0x0d, 0xd8, 0x10, 0x3a, 0x65, 0x63, 0xb4, 0xdb, 0xd9, 0x8d, 0x3c, 0x81, 0xf4,
	0x8d, 0xdb, 0x56, 0x6c, 0xab, 0x61, 0x48, 0xcb, 0x53, 0x10, 0x15, 0x96, 0xb6, 0x89, 0x68, 0x94,
	0x25, 0x8c, 0x6e, 0xe0, 0xb8, 0x68, 0x8a, 0x6a, 0x4f, 0x96, 0x62, 0xc9, 0x01, 0xd3, 0xbf, 0x25,
	0x67, 0x70, 0xd4, 0xf2, 0xd0, 0x9a, 0x28, 0x86, 0xc6, 0x31, 0x34, 0x85, 0x31, 0x59, 0x6f, 0xe6,
	0x1b, 0x75, 0x18, 0x74, 0x79, 0x0b, 0x93, 0x70, 0x89, 0x87, 0xba, 0xee, 0x5c, 0x1a, 0xb4, 0xff,
	0x63, 0x2f, 0x21, 0x56, 0x2a, 0xd1, 0x3f, 0xa1, 0xef, 0x61, 0xb6, 0xc6, 0x3a, 0xff, 0xb0, 0x5b,
	0x6f, 0x9a, 0x9c, 0x8c, 0x7b, 0x37, 0xae, 0xfb, 0x94, 0x35, 0x56, 0x39, 0xa7, 0x89, 0xdf, 0xa8,
	0x2f, 0xfa, 0x9f, 0xf3, 0xd4, 0xe1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xb1, 0x59, 0x06,
	0x61, 0x01, 0x00, 0x00,
}