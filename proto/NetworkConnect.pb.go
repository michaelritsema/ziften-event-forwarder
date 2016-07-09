// Code generated by protoc-gen-go.
// source: NetworkConnect.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetworkConnect_ProtocolType int32

const (
	NetworkConnect_IPVersion4 NetworkConnect_ProtocolType = 0
	NetworkConnect_IPVersion6 NetworkConnect_ProtocolType = 1
)

var NetworkConnect_ProtocolType_name = map[int32]string{
	0: "IPVersion4",
	1: "IPVersion6",
}
var NetworkConnect_ProtocolType_value = map[string]int32{
	"IPVersion4": 0,
	"IPVersion6": 1,
}

func (x NetworkConnect_ProtocolType) Enum() *NetworkConnect_ProtocolType {
	p := new(NetworkConnect_ProtocolType)
	*p = x
	return p
}
func (x NetworkConnect_ProtocolType) String() string {
	return proto1.EnumName(NetworkConnect_ProtocolType_name, int32(x))
}
func (x *NetworkConnect_ProtocolType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(NetworkConnect_ProtocolType_value, data, "NetworkConnect_ProtocolType")
	if err != nil {
		return err
	}
	*x = NetworkConnect_ProtocolType(value)
	return nil
}
func (NetworkConnect_ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor35, []int{0, 0}
}

type NetworkConnect_StreamType int32

const (
	NetworkConnect_TCP NetworkConnect_StreamType = 0
	NetworkConnect_UDP NetworkConnect_StreamType = 1
)

var NetworkConnect_StreamType_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var NetworkConnect_StreamType_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x NetworkConnect_StreamType) Enum() *NetworkConnect_StreamType {
	p := new(NetworkConnect_StreamType)
	*p = x
	return p
}
func (x NetworkConnect_StreamType) String() string {
	return proto1.EnumName(NetworkConnect_StreamType_name, int32(x))
}
func (x *NetworkConnect_StreamType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(NetworkConnect_StreamType_value, data, "NetworkConnect_StreamType")
	if err != nil {
		return err
	}
	*x = NetworkConnect_StreamType(value)
	return nil
}
func (NetworkConnect_StreamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor35, []int{0, 1}
}

// Notify server about the first connection after boot of an application to a certain remote IP address.
type NetworkConnect struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Remote IP Address Protocol Type
	IPAddressProtocol *NetworkConnect_ProtocolType `protobuf:"varint,3,req,name=IPAddressProtocol,enum=NetworkConnect_ProtocolType" json:"IPAddressProtocol,omitempty"`
	// Remote IP Address (4 bytes for IPv4 and 16 bytes for IPv6 addresses)
	IPAddress []byte `protobuf:"bytes,4,req,name=IPAddress" json:"IPAddress,omitempty"`
	// The Process ID of the application
	PID *int32 `protobuf:"varint,5,req,name=PID" json:"PID,omitempty"`
	// The fully qualified filename of the binary: i.e. 'c:\windows\system32\binary.exe'
	ImageFilepath *string `protobuf:"bytes,6,req,name=imageFilepath" json:"imageFilepath,omitempty"`
	// The MD5 of the binary connecting to the remote IP (md5 of imageFilepath)
	ImageFileMD5 *string `protobuf:"bytes,7,opt,name=imageFileMD5" json:"imageFileMD5,omitempty"`
	// User (using this app) Information
	AccountName *string `protobuf:"bytes,8,opt,name=accountName" json:"accountName,omitempty"`
	DomainName  *string `protobuf:"bytes,9,opt,name=domainName" json:"domainName,omitempty"`
	// Port used to connect to the remote IP
	Port   *int32  `protobuf:"varint,10,opt,name=port" json:"port,omitempty"`
	SiteId *string `protobuf:"bytes,11,opt,name=siteId" json:"siteId,omitempty"`
	// Source IP Address (4 bytes for IPv4 and 16 bytes for IPv6 addresses)
	IPSourceAddress []byte `protobuf:"bytes,12,opt,name=IPSourceAddress" json:"IPSourceAddress,omitempty"`
	// Port used to connect to the source IP
	SourcePort           *int32                     `protobuf:"varint,13,opt,name=sourcePort" json:"sourcePort,omitempty"`
	ConnectionStreamType *NetworkConnect_StreamType `protobuf:"varint,14,opt,name=connectionStreamType,enum=NetworkConnect_StreamType" json:"connectionStreamType,omitempty"`
	ConnectionSuccessful *bool                      `protobuf:"varint,15,opt,name=connectionSuccessful" json:"connectionSuccessful,omitempty"`
	// Flag - true if incoming message, false if outgoing
	Incoming *bool `protobuf:"varint,16,opt,name=incoming" json:"incoming,omitempty"`
	// Version =1 so that IPaddress/port is 'destination' and IPSourceAddress/sourcePort is 'source'
	// Reversed from Windows to match industry standard
	Version *int32 `protobuf:"varint,17,opt,name=version" json:"version,omitempty"`
	// host name of IPAddress if outgoing or host name of IPSourceAddress if outgoing
	// If there is more than one hostname for an IP address, they are all given with ',' between them
	Hostname *string `protobuf:"bytes,18,opt,name=hostname" json:"hostname,omitempty"`
	Uuid     *string `protobuf:"bytes,19,opt,name=uuid" json:"uuid,omitempty"`
	// The is the unique identifier for the specific process performing the network connect
	ProcessUUID *string `protobuf:"bytes,20,opt,name=processUUID" json:"processUUID,omitempty"`
	// network connect uuid as key to bind additional network connect secondary messages
	NcUuid           *string `protobuf:"bytes,21,opt,name=nc_uuid" json:"nc_uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NetworkConnect) Reset()                    { *m = NetworkConnect{} }
func (m *NetworkConnect) String() string            { return proto1.CompactTextString(m) }
func (*NetworkConnect) ProtoMessage()               {}
func (*NetworkConnect) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

func (m *NetworkConnect) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *NetworkConnect) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *NetworkConnect) GetIPAddressProtocol() NetworkConnect_ProtocolType {
	if m != nil && m.IPAddressProtocol != nil {
		return *m.IPAddressProtocol
	}
	return NetworkConnect_IPVersion4
}

func (m *NetworkConnect) GetIPAddress() []byte {
	if m != nil {
		return m.IPAddress
	}
	return nil
}

func (m *NetworkConnect) GetPID() int32 {
	if m != nil && m.PID != nil {
		return *m.PID
	}
	return 0
}

func (m *NetworkConnect) GetImageFilepath() string {
	if m != nil && m.ImageFilepath != nil {
		return *m.ImageFilepath
	}
	return ""
}

func (m *NetworkConnect) GetImageFileMD5() string {
	if m != nil && m.ImageFileMD5 != nil {
		return *m.ImageFileMD5
	}
	return ""
}

func (m *NetworkConnect) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *NetworkConnect) GetDomainName() string {
	if m != nil && m.DomainName != nil {
		return *m.DomainName
	}
	return ""
}

func (m *NetworkConnect) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *NetworkConnect) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *NetworkConnect) GetIPSourceAddress() []byte {
	if m != nil {
		return m.IPSourceAddress
	}
	return nil
}

func (m *NetworkConnect) GetSourcePort() int32 {
	if m != nil && m.SourcePort != nil {
		return *m.SourcePort
	}
	return 0
}

func (m *NetworkConnect) GetConnectionStreamType() NetworkConnect_StreamType {
	if m != nil && m.ConnectionStreamType != nil {
		return *m.ConnectionStreamType
	}
	return NetworkConnect_TCP
}

func (m *NetworkConnect) GetConnectionSuccessful() bool {
	if m != nil && m.ConnectionSuccessful != nil {
		return *m.ConnectionSuccessful
	}
	return false
}

func (m *NetworkConnect) GetIncoming() bool {
	if m != nil && m.Incoming != nil {
		return *m.Incoming
	}
	return false
}

func (m *NetworkConnect) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *NetworkConnect) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *NetworkConnect) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *NetworkConnect) GetProcessUUID() string {
	if m != nil && m.ProcessUUID != nil {
		return *m.ProcessUUID
	}
	return ""
}

func (m *NetworkConnect) GetNcUuid() string {
	if m != nil && m.NcUuid != nil {
		return *m.NcUuid
	}
	return ""
}

func init() {
	proto1.RegisterType((*NetworkConnect)(nil), "NetworkConnect")
	proto1.RegisterEnum("NetworkConnect_ProtocolType", NetworkConnect_ProtocolType_name, NetworkConnect_ProtocolType_value)
	proto1.RegisterEnum("NetworkConnect_StreamType", NetworkConnect_StreamType_name, NetworkConnect_StreamType_value)
}

var fileDescriptor35 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0x4d, 0xd2, 0x24, 0x13, 0xc7, 0x71, 0xb7, 0x09, 0x5a, 0x4a, 0x41, 0x21, 0x27, 0x4e,
	0x3e, 0x54, 0x7c, 0x49, 0x9c, 0x70, 0x23, 0x90, 0x0f, 0xad, 0x56, 0x4a, 0xc3, 0x15, 0x59, 0xeb,
	0x6d, 0xba, 0x22, 0xde, 0xb5, 0xbc, 0x6b, 0x10, 0x1c, 0x39, 0xf3, 0x77, 0xf8, 0x7f, 0xcc, 0xae,
	0xd5, 0x36, 0xce, 0xcd, 0xfb, 0xe6, 0xcd, 0xbc, 0xf7, 0x3c, 0x03, 0xb3, 0x6b, 0x61, 0x7f, 0xea,
	0xea, 0xfb, 0xa5, 0x56, 0x4a, 0x70, 0x1b, 0x97, 0x95, 0xb6, 0xfa, 0x8c, 0x30, 0x6d, 0xec, 0xb6,
	0x12, 0xe6, 0xe6, 0x57, 0x29, 0x1a, 0x6c, 0xf9, 0xb7, 0x0f, 0x61, 0x9b, 0x4c, 0x16, 0x30, 0xb2,
	0xb2, 0x10, 0x6b, 0x9b, 0x15, 0x25, 0xed, 0x2c, 0x8e, 0x5e, 0x77, 0x93, 0xc9, 0x9f, 0x7f, 0xd4,
	0x83, 0xc6, 0x81, 0xe4, 0x39, 0x8c, 0xb2, 0xad, 0x50, 0xf6, 0xcb, 0x26, 0x5d, 0xd1, 0x23, 0x64,
	0x8c, 0x92, 0x21, 0x32, 0x7a, 0x75, 0x2d, 0x73, 0xf2, 0x1e, 0x4e, 0x52, 0xf6, 0x29, 0xcf, 0x51,
	0xc7, 0x30, 0xa7, 0xc1, 0xf5, 0x8e, 0x76, 0x91, 0x14, 0x5e, 0x9c, 0xc7, 0x07, 0xbe, 0xee, 0xeb,
	0xce, 0x90, 0x9b, 0xfa, 0xd0, 0x48, 0x7b, 0xd8, 0x10, 0x34, 0x53, 0xa5, 0x12, 0x96, 0x8c, 0xa1,
	0xcb, 0x50, 0xac, 0x8f, 0x70, 0x9f, 0xcc, 0x61, 0x22, 0x0b, 0x74, 0xf0, 0x59, 0xee, 0x44, 0x99,
	0xd9, 0x3b, 0x7a, 0xec, 0x3c, 0x90, 0x19, 0x04, 0x0f, 0xf0, 0xd5, 0xea, 0x2d, 0x1d, 0x2c, 0x3a,
	0x88, 0x9e, 0xc2, 0x38, 0xe3, 0x5c, 0xd7, 0xca, 0x5e, 0x67, 0x85, 0xa0, 0x43, 0x0f, 0x12, 0x80,
	0x5c, 0x17, 0x99, 0x54, 0x1e, 0x1b, 0x79, 0x2c, 0x80, 0x5e, 0xa9, 0x2b, 0x4b, 0x01, 0x5f, 0x7d,
	0x12, 0xc2, 0xb1, 0x91, 0x56, 0xa4, 0x39, 0x1d, 0xfb, 0xea, 0x2b, 0x98, 0xa6, 0x6c, 0xad, 0xeb,
	0x8a, 0x8b, 0x7b, 0x8f, 0x01, 0x16, 0xf6, 0x3d, 0xe2, 0x50, 0xe3, 0x09, 0xcc, 0x8d, 0x99, 0xf8,
	0x31, 0x1f, 0x60, 0xc6, 0x9b, 0xb0, 0x52, 0xab, 0xb5, 0xad, 0x44, 0x56, 0xb8, 0xb0, 0x34, 0xc4,
	0x6a, 0x78, 0x71, 0x76, 0xf8, 0x43, 0x1e, 0x19, 0xe4, 0xbc, 0xd5, 0x59, 0x73, 0x8e, 0x92, 0xb7,
	0xf5, 0x8e, 0x4e, 0xb1, 0x73, 0x48, 0x22, 0x18, 0x4a, 0xc5, 0x75, 0x21, 0xd5, 0x96, 0x46, 0x1e,
	0x99, 0xc2, 0xe0, 0x87, 0xa8, 0x0c, 0x92, 0xe9, 0x89, 0x97, 0x46, 0xca, 0x1d, 0x2e, 0x5c, 0xb9,
	0x84, 0xc4, 0x67, 0x78, 0x0a, 0x7e, 0x45, 0xf4, 0xd4, 0xbd, 0xf6, 0x56, 0xf6, 0x02, 0xc6, 0x78,
	0x0d, 0x4e, 0x60, 0xe3, 0x36, 0x3a, 0x3b, 0x28, 0x3f, 0x83, 0x81, 0xe2, 0xdf, 0x7c, 0xe7, 0xbc,
	0x5d, 0x5a, 0xc6, 0x10, 0xb4, 0x76, 0x18, 0x02, 0xa4, 0xec, 0x6b, 0x63, 0xe3, 0x4d, 0xf4, 0xa4,
	0xf5, 0x7e, 0x17, 0x75, 0x96, 0x2f, 0x01, 0xf6, 0x22, 0x0e, 0xa0, 0x7b, 0x73, 0xc9, 0x90, 0x86,
	0x1f, 0x9b, 0x15, 0x8b, 0x3a, 0xc9, 0x47, 0x58, 0x62, 0xa8, 0xf8, 0xb7, 0xbc, 0xb5, 0x42, 0xc5,
	0x46, 0x54, 0x18, 0xa9, 0xb9, 0x54, 0x54, 0x88, 0xf1, 0xfe, 0x0c, 0xee, 0x37, 0x99, 0xb7, 0xff,
	0xda, 0x55, 0x03, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x24, 0xdc, 0x5d, 0xe3, 0xf6, 0x02, 0x00,
	0x00,
}
