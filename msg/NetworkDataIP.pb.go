// Code generated by protoc-gen-go.
// source: NetworkDataIP.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Send performance metrics about traffic (TCP SEND/RECEIVE) to certain IP addresses.
// Data is derived by ETW Network Trace
// see http://msdn.microsoft.com/en-us/library/windows/desktop/aa364128(v=vs.85).aspx
type NetworkDataIP struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// IPV4 Address
	IP                         []byte  `protobuf:"bytes,3,req,name=IP" json:"IP,omitempty"`
	SentBytes                  *int32  `protobuf:"varint,4,req,name=sentBytes" json:"sentBytes,omitempty"`
	ReceivedBytes              *int32  `protobuf:"varint,5,req,name=receivedBytes" json:"receivedBytes,omitempty"`
	SentBytesV2                *uint64 `protobuf:"varint,10,opt,name=sentBytes_v2" json:"sentBytes_v2,omitempty"`
	ReceivedBytesV2            *uint64 `protobuf:"varint,11,opt,name=receivedBytes_v2" json:"receivedBytes_v2,omitempty"`
	SumTransactionTime         *int64  `protobuf:"varint,6,req,name=sumTransactionTime" json:"sumTransactionTime,omitempty"`
	NumberSampledResponseTimes *int32  `protobuf:"varint,7,req,name=numberSampledResponseTimes" json:"numberSampledResponseTimes,omitempty"`
	SumResponseTime            *int64  `protobuf:"varint,8,req,name=sumResponseTime" json:"sumResponseTime,omitempty"`
	SiteId                     *string `protobuf:"bytes,9,opt,name=siteId" json:"siteId,omitempty"`
	// Indexes 10-11 used above with new versions of sent/receivedBytes using int64 instead of int32 to avoid overflow
	ImageFilepath    *string `protobuf:"bytes,12,opt,name=imageFilepath" json:"imageFilepath,omitempty"`
	Uuid             *string `protobuf:"bytes,13,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetworkDataIP) Reset()                    { *m = NetworkDataIP{} }
func (m *NetworkDataIP) String() string            { return proto.CompactTextString(m) }
func (*NetworkDataIP) ProtoMessage()               {}
func (*NetworkDataIP) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{0} }

func (m *NetworkDataIP) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *NetworkDataIP) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *NetworkDataIP) GetIP() []byte {
	if m != nil {
		return m.IP
	}
	return nil
}

func (m *NetworkDataIP) GetSentBytes() int32 {
	if m != nil && m.SentBytes != nil {
		return *m.SentBytes
	}
	return 0
}

func (m *NetworkDataIP) GetReceivedBytes() int32 {
	if m != nil && m.ReceivedBytes != nil {
		return *m.ReceivedBytes
	}
	return 0
}

func (m *NetworkDataIP) GetSentBytesV2() uint64 {
	if m != nil && m.SentBytesV2 != nil {
		return *m.SentBytesV2
	}
	return 0
}

func (m *NetworkDataIP) GetReceivedBytesV2() uint64 {
	if m != nil && m.ReceivedBytesV2 != nil {
		return *m.ReceivedBytesV2
	}
	return 0
}

func (m *NetworkDataIP) GetSumTransactionTime() int64 {
	if m != nil && m.SumTransactionTime != nil {
		return *m.SumTransactionTime
	}
	return 0
}

func (m *NetworkDataIP) GetNumberSampledResponseTimes() int32 {
	if m != nil && m.NumberSampledResponseTimes != nil {
		return *m.NumberSampledResponseTimes
	}
	return 0
}

func (m *NetworkDataIP) GetSumResponseTime() int64 {
	if m != nil && m.SumResponseTime != nil {
		return *m.SumResponseTime
	}
	return 0
}

func (m *NetworkDataIP) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *NetworkDataIP) GetImageFilepath() string {
	if m != nil && m.ImageFilepath != nil {
		return *m.ImageFilepath
	}
	return ""
}

func (m *NetworkDataIP) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkDataIP)(nil), "NetworkDataIP")
}

var fileDescriptor37 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xd5, 0xf4, 0xf2, 0x37, 0xf3, 0x37, 0x5c, 0x4c, 0x01, 0xab, 0x6c, 0xaa, 0xae, 0xba,
	0xca, 0x82, 0x25, 0xcb, 0xa8, 0x02, 0x65, 0x01, 0xaa, 0x68, 0x59, 0x23, 0x93, 0x0c, 0xc5, 0xa2,
	0xb1, 0x23, 0x7b, 0x52, 0x54, 0x96, 0xbc, 0x13, 0xaf, 0xc2, 0xf3, 0xe0, 0x38, 0x12, 0x90, 0xa5,
	0xbf, 0x73, 0xc6, 0x73, 0xe6, 0xc0, 0xc9, 0x1d, 0xd2, 0x9b, 0x36, 0xaf, 0x0b, 0x41, 0x22, 0x5d,
	0xc6, 0xa5, 0xd1, 0xa4, 0x27, 0x6c, 0xa9, 0x2d, 0x6d, 0x0c, 0xda, 0xf5, 0xbe, 0xc4, 0x86, 0xcd,
	0xbe, 0x02, 0x88, 0x5a, 0x5e, 0x36, 0x85, 0x90, 0x64, 0x81, 0x2b, 0x12, 0x45, 0xc9, 0x3b, 0xd3,
	0x60, 0xde, 0x4d, 0xa2, 0x8f, 0x4f, 0xee, 0xa1, 0xad, 0x21, 0xbb, 0x80, 0x50, 0x6c, 0x50, 0xd1,
	0xcd, 0x43, 0xba, 0xe0, 0x81, 0x73, 0x84, 0xc9, 0xd0, 0x39, 0x7a, 0x55, 0x25, 0x73, 0x36, 0x86,
	0x20, 0x5d, 0xf2, 0xae, 0xa3, 0xa3, 0x86, 0x4a, 0x85, 0xc4, 0x8e, 0x21, 0xb4, 0x6e, 0x22, 0xd9,
	0x13, 0x5a, 0xde, 0x73, 0x62, 0x9f, 0x9d, 0x42, 0x64, 0x30, 0x43, 0xb9, 0xc3, 0xbc, 0xc1, 0x7d,
	0x8f, 0xc7, 0x30, 0xfa, 0x71, 0x3e, 0xee, 0x2e, 0x39, 0x4c, 0x3b, 0xf3, 0x1e, 0xe3, 0x70, 0xd4,
	0x32, 0xd7, 0xca, 0x7f, 0xaf, 0x4c, 0x80, 0xd9, 0xaa, 0x58, 0x1b, 0xa1, 0xac, 0xc8, 0x48, 0x6a,
	0xb5, 0x76, 0x39, 0xf9, 0xa0, 0xce, 0xcd, 0x66, 0x30, 0x51, 0x55, 0xf1, 0x84, 0x66, 0xe5, 0x62,
	0x6f, 0x31, 0xbf, 0x47, 0x5b, 0x6a, 0x65, 0xb1, 0xb6, 0x58, 0xfe, 0xcf, 0xef, 0x3b, 0x87, 0x43,
	0x37, 0xff, 0x57, 0xe1, 0x43, 0x3f, 0x7c, 0x00, 0x03, 0x2b, 0x09, 0xd3, 0x9c, 0x87, 0x6e, 0x51,
	0x58, 0xe7, 0x95, 0x85, 0xbb, 0xfb, 0x5a, 0x6e, 0xb1, 0x14, 0xf4, 0xc2, 0x47, 0x1e, 0x9f, 0x81,
	0xbf, 0x9b, 0x47, 0xf5, 0xeb, 0xb7, 0x87, 0xe4, 0x0a, 0x66, 0x99, 0x2e, 0xe2, 0x77, 0xf9, 0x4c,
	0xa8, 0x62, 0x8b, 0x66, 0x87, 0xa6, 0xe9, 0x3c, 0xd3, 0xdb, 0xd8, 0xed, 0xb7, 0xee, 0xaf, 0x64,
	0xdc, 0xea, 0xfe, 0xb6, 0xa1, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xfc, 0x4e, 0x61, 0xbe,
	0x01, 0x00, 0x00,
}