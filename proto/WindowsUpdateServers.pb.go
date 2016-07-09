// Code generated by protoc-gen-go.
// source: WindowsUpdateServers.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This message is tied to a WindowsUpdateSetting message to provide a list of
// Windows Update Servers present at the time of the WindowsUpdateSetting message
type WindowsUpdateServers struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Unique identifier for this message to tie it to WindowsUpdateSetting message
	WUA_ServersGUID *string  `protobuf:"bytes,3,req,name=WUA_ServersGUID" json:"WUA_ServersGUID,omitempty"`
	ServerID        []string `protobuf:"bytes,4,rep,name=serverID" json:"serverID,omitempty"`
	ServerName      []string `protobuf:"bytes,5,rep,name=serverName" json:"serverName,omitempty"`
	SiteId          *string  `protobuf:"bytes,6,opt,name=siteId" json:"siteId,omitempty"`
	Uuid            *string  `protobuf:"bytes,7,opt,name=uuid" json:"uuid,omitempty"`
	// This UUID is to group WindowsUpdate, WindowsUpdateAvailable, WindowsUpdateServers and WindowsUpdateSettings message
	WuaScanUUID      *string `protobuf:"bytes,8,opt,name=wuaScanUUID" json:"wuaScanUUID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WindowsUpdateServers) Reset()                    { *m = WindowsUpdateServers{} }
func (m *WindowsUpdateServers) String() string            { return proto1.CompactTextString(m) }
func (*WindowsUpdateServers) ProtoMessage()               {}
func (*WindowsUpdateServers) Descriptor() ([]byte, []int) { return fileDescriptor91, []int{0} }

func (m *WindowsUpdateServers) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *WindowsUpdateServers) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *WindowsUpdateServers) GetWUA_ServersGUID() string {
	if m != nil && m.WUA_ServersGUID != nil {
		return *m.WUA_ServersGUID
	}
	return ""
}

func (m *WindowsUpdateServers) GetServerID() []string {
	if m != nil {
		return m.ServerID
	}
	return nil
}

func (m *WindowsUpdateServers) GetServerName() []string {
	if m != nil {
		return m.ServerName
	}
	return nil
}

func (m *WindowsUpdateServers) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *WindowsUpdateServers) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *WindowsUpdateServers) GetWuaScanUUID() string {
	if m != nil && m.WuaScanUUID != nil {
		return *m.WuaScanUUID
	}
	return ""
}

func init() {
	proto1.RegisterType((*WindowsUpdateServers)(nil), "WindowsUpdateServers")
}

var fileDescriptor91 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0x53, 0x6b, 0x32, 0xe2, 0x07, 0x8b, 0xe8, 0xd2, 0x22, 0x84, 0x9e, 0x3c, 0xe5,
	0x1d, 0x1a, 0x04, 0xe9, 0x41, 0x11, 0x62, 0xe8, 0x51, 0x96, 0x64, 0x2c, 0x0b, 0x26, 0x1b, 0x32,
	0x13, 0x8b, 0x1e, 0x7d, 0x27, 0x9f, 0xc6, 0x97, 0x71, 0x3f, 0x3c, 0x58, 0xe8, 0x71, 0x7e, 0xf3,
	0xdb, 0xf9, 0xf3, 0x5f, 0x98, 0x6f, 0x74, 0xd7, 0x98, 0x1d, 0x55, 0x7d, 0xa3, 0x18, 0x4b, 0x1c,
	0xde, 0x71, 0xa0, 0xbc, 0x1f, 0x0c, 0x9b, 0xb9, 0x78, 0x32, 0xc4, 0xdb, 0x01, 0xe9, 0xf9, 0xa3,
	0xc7, 0xc0, 0x96, 0x3f, 0x11, 0x5c, 0x1e, 0x7a, 0x22, 0x32, 0x48, 0x59, 0xb7, 0x58, 0xb2, 0x6a,
	0x7b, 0x19, 0x65, 0x93, 0xdb, 0xb8, 0x38, 0xfd, 0xfa, 0x96, 0x1e, 0x92, 0x83, 0x62, 0x01, 0xa9,
	0xda, 0x62, 0xc7, 0xf7, 0xd5, 0xfa, 0x4e, 0x4e, 0xac, 0x91, 0x16, 0x89, 0x35, 0xa6, 0xe3, 0xa8,
	0x1b, 0x71, 0x0d, 0xe7, 0x9b, 0x6a, 0xf5, 0xf2, 0x77, 0xcd, 0x2b, 0xb1, 0x53, 0xc4, 0x05, 0x24,
	0xe4, 0xa1, 0x25, 0xd3, 0x2c, 0xb6, 0x44, 0x00, 0x04, 0xf2, 0xa8, 0x5a, 0x94, 0x47, 0x9e, 0x9d,
	0xc1, 0x8c, 0x34, 0xe3, 0xba, 0x91, 0xb3, 0x2c, 0xb2, 0xf3, 0x15, 0xf8, 0xb3, 0xf2, 0xd8, 0x4d,
	0xff, 0x62, 0x6e, 0xe0, 0x64, 0x37, 0xaa, 0xb2, 0x56, 0x5d, 0xe5, 0x22, 0x92, 0xfd, 0x75, 0xb1,
	0x82, 0x65, 0x6d, 0xda, 0xfc, 0x53, 0xbf, 0x32, 0x76, 0x79, 0x48, 0x09, 0xc5, 0x6b, 0xf3, 0x96,
	0xdb, 0x22, 0x64, 0x1b, 0x14, 0x8b, 0x43, 0x1f, 0xf0, 0x10, 0x96, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x38, 0x1a, 0xdb, 0xef, 0x51, 0x01, 0x00, 0x00,
}
