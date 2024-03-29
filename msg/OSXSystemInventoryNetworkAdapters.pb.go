// Code generated by protoc-gen-go.
// source: OSXSystemInventoryNetworkAdapters.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXSystemInventoryNetworkAdapters struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// GUID that helps join this message with associated SystemInventory message
	NetworkAdaptersGUID *string `protobuf:"bytes,3,req,name=networkAdaptersGUID" json:"networkAdaptersGUID,omitempty"`
	// See WMI's Win32_NetworkAdapterConfiguration and Win32_NetworkAdapter MOF for a detailed description the fields below
	AdapterName []string `protobuf:"bytes,4,rep,name=adapterName" json:"adapterName,omitempty"`
	// e.g. Apple Wireless Network Adapter (802.11 a/b/g/n)
	AdapterDescription []string `protobuf:"bytes,5,rep,name=adapterDescription" json:"adapterDescription,omitempty"`
	// e.g. 00:17:F2:0D:1A:D8
	MACAddress []string `protobuf:"bytes,6,rep,name=MACAddress" json:"MACAddress,omitempty"`
	// e.g. 10.4.122.25
	IPV4Address      []string `protobuf:"bytes,7,rep,name=IPV4Address" json:"IPV4Address,omitempty"`
	IsPrimaryAdapter []bool   `protobuf:"varint,8,rep,name=isPrimaryAdapter" json:"isPrimaryAdapter,omitempty"`
	SiteId           *string  `protobuf:"bytes,9,opt,name=siteId" json:"siteId,omitempty"`
	Uuid             *string  `protobuf:"bytes,10,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *OSXSystemInventoryNetworkAdapters) Reset()         { *m = OSXSystemInventoryNetworkAdapters{} }
func (m *OSXSystemInventoryNetworkAdapters) String() string { return proto.CompactTextString(m) }
func (*OSXSystemInventoryNetworkAdapters) ProtoMessage()    {}
func (*OSXSystemInventoryNetworkAdapters) Descriptor() ([]byte, []int) {
	return fileDescriptor59, []int{0}
}

func (m *OSXSystemInventoryNetworkAdapters) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXSystemInventoryNetworkAdapters) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXSystemInventoryNetworkAdapters) GetNetworkAdaptersGUID() string {
	if m != nil && m.NetworkAdaptersGUID != nil {
		return *m.NetworkAdaptersGUID
	}
	return ""
}

func (m *OSXSystemInventoryNetworkAdapters) GetAdapterName() []string {
	if m != nil {
		return m.AdapterName
	}
	return nil
}

func (m *OSXSystemInventoryNetworkAdapters) GetAdapterDescription() []string {
	if m != nil {
		return m.AdapterDescription
	}
	return nil
}

func (m *OSXSystemInventoryNetworkAdapters) GetMACAddress() []string {
	if m != nil {
		return m.MACAddress
	}
	return nil
}

func (m *OSXSystemInventoryNetworkAdapters) GetIPV4Address() []string {
	if m != nil {
		return m.IPV4Address
	}
	return nil
}

func (m *OSXSystemInventoryNetworkAdapters) GetIsPrimaryAdapter() []bool {
	if m != nil {
		return m.IsPrimaryAdapter
	}
	return nil
}

func (m *OSXSystemInventoryNetworkAdapters) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXSystemInventoryNetworkAdapters) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXSystemInventoryNetworkAdapters)(nil), "OSXSystemInventoryNetworkAdapters")
}

var fileDescriptor59 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x59, 0x3b, 0x67, 0xfb, 0x44, 0x91, 0x0c, 0x24, 0xcc, 0x4b, 0x1d, 0x88, 0x3d, 0xf5,
	0xe4, 0x17, 0x68, 0x1d, 0x48, 0x0f, 0xab, 0x85, 0xaa, 0x78, 0x0d, 0xed, 0x73, 0x04, 0x4d, 0x53,
	0x92, 0x6c, 0x52, 0x8f, 0x7e, 0x17, 0x3f, 0x82, 0xdf, 0xcf, 0x2c, 0x9d, 0x88, 0x5e, 0x76, 0xcc,
	0xef, 0xfd, 0xf2, 0xf8, 0xbf, 0x3f, 0x5c, 0xdd, 0x55, 0x4f, 0x55, 0xaf, 0x0d, 0x8a, 0xbc, 0xdd,
	0x60, 0x6b, 0xa4, 0xea, 0x0b, 0x34, 0x6f, 0x52, 0xbd, 0xa4, 0x0d, 0xeb, 0x0c, 0x2a, 0x9d, 0x74,
	0x4a, 0x1a, 0x39, 0x23, 0xa5, 0xd4, 0x66, 0xa5, 0x50, 0xdf, 0xf7, 0x1d, 0x0e, 0x6c, 0xfe, 0xe9,
	0xc1, 0xc5, 0xde, 0xff, 0x24, 0x82, 0xd0, 0x70, 0x81, 0x95, 0x61, 0xa2, 0xa3, 0xa3, 0xc8, 0x8b,
	0xfd, 0xec, 0xf8, 0xe3, 0x8b, 0x3a, 0xa8, 0xb7, 0x90, 0x9c, 0x43, 0xc8, 0x56, 0xf6, 0xf3, 0xed,
	0x43, 0xbe, 0xa0, 0x9e, 0x35, 0xc2, 0x2c, 0xb0, 0xc6, 0x78, 0xbd, 0xe6, 0x0d, 0xb9, 0x84, 0x69,
	0xfb, 0x77, 0xa3, 0xd3, 0xfc, 0x7f, 0xda, 0x14, 0x8e, 0xd8, 0x30, 0x2f, 0x98, 0x40, 0x3a, 0x8e,
	0xfc, 0x38, 0x24, 0x33, 0x20, 0x3b, 0xb8, 0x40, 0x5d, 0x2b, 0xde, 0x19, 0x2e, 0x5b, 0x7a, 0xe0,
	0x66, 0x04, 0x60, 0x99, 0xde, 0xa4, 0x4d, 0x63, 0x6f, 0xd2, 0x74, 0xe2, 0x98, 0x5d, 0x92, 0x97,
	0x8f, 0xd7, 0x3f, 0xf0, 0xd0, 0x41, 0x0a, 0xa7, 0x5c, 0x97, 0x8a, 0x0b, 0xa6, 0xfa, 0x5d, 0x04,
	0x1a, 0xd8, 0x49, 0x40, 0x4e, 0x60, 0xa2, 0xb9, 0xc1, 0xbc, 0xa1, 0x61, 0x34, 0xb2, 0xe6, 0x19,
	0xb8, 0x2c, 0x14, 0xb6, 0xaf, 0xdf, 0x6c, 0x59, 0x01, 0xf3, 0x5a, 0x8a, 0xe4, 0x9d, 0x3f, 0x1b,
	0x6c, 0x13, 0x8d, 0x6a, 0x83, 0x6a, 0xa8, 0xb0, 0x96, 0xaf, 0x89, 0x6d, 0x41, 0xdb, 0xf3, 0xb3,
	0x78, 0x6f, 0x95, 0xcb, 0xc1, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xb8, 0xdc, 0xf2, 0xb5,
	0x01, 0x00, 0x00,
}
