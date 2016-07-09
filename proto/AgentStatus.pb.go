// Code generated by protoc-gen-go.
// source: AgentStatus.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AgentStatus_AgentStatusMessageType int32

const (
	AgentStatus_AGENTSTARTED    AgentStatus_AgentStatusMessageType = 0
	AgentStatus_AGENTSTOPPED    AgentStatus_AgentStatusMessageType = 1
	AgentStatus_AGENTGUIDCHANGE AgentStatus_AgentStatusMessageType = 5
	AgentStatus_AGENTHEARTBEAT  AgentStatus_AgentStatusMessageType = 6
)

var AgentStatus_AgentStatusMessageType_name = map[int32]string{
	0: "AGENTSTARTED",
	1: "AGENTSTOPPED",
	5: "AGENTGUIDCHANGE",
	6: "AGENTHEARTBEAT",
}
var AgentStatus_AgentStatusMessageType_value = map[string]int32{
	"AGENTSTARTED":    0,
	"AGENTSTOPPED":    1,
	"AGENTGUIDCHANGE": 5,
	"AGENTHEARTBEAT":  6,
}

func (x AgentStatus_AgentStatusMessageType) Enum() *AgentStatus_AgentStatusMessageType {
	p := new(AgentStatus_AgentStatusMessageType)
	*p = x
	return p
}
func (x AgentStatus_AgentStatusMessageType) String() string {
	return proto1.EnumName(AgentStatus_AgentStatusMessageType_name, int32(x))
}
func (x *AgentStatus_AgentStatusMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(AgentStatus_AgentStatusMessageType_value, data, "AgentStatus_AgentStatusMessageType")
	if err != nil {
		return err
	}
	*x = AgentStatus_AgentStatusMessageType(value)
	return nil
}
func (AgentStatus_AgentStatusMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 0}
}

type AgentStatus struct {
	// The UTC time (as known to the client computer) when the agent status message was generated
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID   *string                             `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	AgentStatus *AgentStatus_AgentStatusMessageType `protobuf:"varint,3,req,name=agentStatus,enum=AgentStatus_AgentStatusMessageType" json:"agentStatus,omitempty"`
	// Previous GUID for a AGENTGUIDCHANGE message
	OldAgentGUID     *string `protobuf:"bytes,8,opt,name=oldAgentGUID" json:"oldAgentGUID,omitempty"`
	SiteId           *string `protobuf:"bytes,9,opt,name=siteId" json:"siteId,omitempty"`
	ComputerName     *string `protobuf:"bytes,10,opt,name=computerName" json:"computerName,omitempty"`
	Uuid             *string `protobuf:"bytes,11,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AgentStatus) Reset()                    { *m = AgentStatus{} }
func (m *AgentStatus) String() string            { return proto1.CompactTextString(m) }
func (*AgentStatus) ProtoMessage()               {}
func (*AgentStatus) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *AgentStatus) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *AgentStatus) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *AgentStatus) GetAgentStatus() AgentStatus_AgentStatusMessageType {
	if m != nil && m.AgentStatus != nil {
		return *m.AgentStatus
	}
	return AgentStatus_AGENTSTARTED
}

func (m *AgentStatus) GetOldAgentGUID() string {
	if m != nil && m.OldAgentGUID != nil {
		return *m.OldAgentGUID
	}
	return ""
}

func (m *AgentStatus) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *AgentStatus) GetComputerName() string {
	if m != nil && m.ComputerName != nil {
		return *m.ComputerName
	}
	return ""
}

func (m *AgentStatus) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto1.RegisterType((*AgentStatus)(nil), "AgentStatus")
	proto1.RegisterEnum("AgentStatus_AgentStatusMessageType", AgentStatus_AgentStatusMessageType_name, AgentStatus_AgentStatusMessageType_value)
}

var fileDescriptor6 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x85, 0x6a, 0x53, 0x86, 0x8a, 0xb8, 0x9a, 0x66, 0xa3, 0x17, 0x82, 0x97, 0x9e, 0x38,
	0x78, 0xea, 0x75, 0xb1, 0x84, 0xf6, 0x20, 0x12, 0x8b, 0x0f, 0x40, 0xca, 0xd8, 0x90, 0x94, 0x2e,
	0x61, 0x17, 0x13, 0x3d, 0xfa, 0x4e, 0x3e, 0x80, 0x6f, 0xe6, 0xb2, 0x24, 0x96, 0xc4, 0x78, 0x63,
	0x3e, 0xfe, 0x99, 0xff, 0xcb, 0xc2, 0x25, 0xdb, 0xe1, 0x41, 0x6e, 0x64, 0x2e, 0x5b, 0x11, 0xd4,
	0x0d, 0x97, 0xfc, 0x86, 0xa4, 0x5c, 0xc8, 0x5d, 0x83, 0x22, 0x7b, 0xaf, 0xb1, 0x67, 0xfe, 0xb7,
	0x09, 0xf6, 0x20, 0x49, 0x3c, 0xb0, 0x64, 0x59, 0xa1, 0x9a, 0xaa, 0x9a, 0x1a, 0x9e, 0x39, 0x1f,
	0x85, 0xe7, 0x9f, 0x5f, 0x54, 0x43, 0xd1, 0x41, 0x72, 0x0b, 0x56, 0xde, 0x2d, 0xc4, 0x2f, 0xeb,
	0x25, 0x35, 0x55, 0xc2, 0x0a, 0x27, 0x2a, 0x71, 0xda, 0xb6, 0x65, 0x41, 0x16, 0x60, 0xe7, 0xc7,
	0x6b, 0x74, 0xa4, 0x7e, 0x3b, 0xf7, 0x77, 0xc1, 0xd0, 0x65, 0xf0, 0xfd, 0x88, 0x42, 0xa8, 0x78,
	0xa7, 0x43, 0xae, 0x61, 0xca, 0xf7, 0x05, 0xfb, 0xbd, 0x3c, 0xf1, 0x8c, 0xb9, 0x45, 0x1c, 0x18,
	0x8b, 0x52, 0xe2, 0xba, 0xa0, 0x96, 0x9e, 0x55, 0x6a, 0xcb, 0xab, 0xba, 0x95, 0xd8, 0x24, 0x79,
	0x85, 0x14, 0x34, 0x9d, 0x81, 0x6e, 0xa7, 0x76, 0x37, 0x1d, 0x6d, 0x7c, 0x84, 0xd9, 0x3f, 0x6d,
	0x2e, 0x4c, 0x59, 0x1c, 0x25, 0xd9, 0x26, 0x63, 0xcf, 0x59, 0xb4, 0x74, 0x4f, 0x06, 0xe4, 0x29,
	0x4d, 0x15, 0x31, 0xc8, 0x15, 0x5c, 0x68, 0xd2, 0xe9, 0x3c, 0xac, 0x58, 0x12, 0x47, 0xee, 0x19,
	0x21, 0xe0, 0x68, 0xb8, 0x8a, 0xd4, 0x62, 0x18, 0xb1, 0xcc, 0x1d, 0x87, 0x0b, 0xf0, 0x95, 0x54,
	0xf0, 0x51, 0xbe, 0x4a, 0x3c, 0x04, 0x02, 0x9b, 0x37, 0x6c, 0xfa, 0xe7, 0xdd, 0xf2, 0x7d, 0x50,
	0xf5, 0xb5, 0x21, 0xf9, 0xab, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xc7, 0x71, 0xec, 0xa5,
	0x01, 0x00, 0x00,
}
