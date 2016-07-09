// Code generated by protoc-gen-go.
// source: ExtensionCommandResponse.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// (echoed back for server's convenience)
type ExtensionCommandResponse_ExtensionCommandType int32

const (
	ExtensionCommandResponse_Install ExtensionCommandResponse_ExtensionCommandType = 1
	ExtensionCommandResponse_Remove  ExtensionCommandResponse_ExtensionCommandType = 2
	ExtensionCommandResponse_Enable  ExtensionCommandResponse_ExtensionCommandType = 3
	ExtensionCommandResponse_Disable ExtensionCommandResponse_ExtensionCommandType = 4
	ExtensionCommandResponse_Update  ExtensionCommandResponse_ExtensionCommandType = 5
)

var ExtensionCommandResponse_ExtensionCommandType_name = map[int32]string{
	1: "Install",
	2: "Remove",
	3: "Enable",
	4: "Disable",
	5: "Update",
}
var ExtensionCommandResponse_ExtensionCommandType_value = map[string]int32{
	"Install": 1,
	"Remove":  2,
	"Enable":  3,
	"Disable": 4,
	"Update":  5,
}

func (x ExtensionCommandResponse_ExtensionCommandType) Enum() *ExtensionCommandResponse_ExtensionCommandType {
	p := new(ExtensionCommandResponse_ExtensionCommandType)
	*p = x
	return p
}
func (x ExtensionCommandResponse_ExtensionCommandType) String() string {
	return proto.EnumName(ExtensionCommandResponse_ExtensionCommandType_name, int32(x))
}
func (x *ExtensionCommandResponse_ExtensionCommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExtensionCommandResponse_ExtensionCommandType_value, data, "ExtensionCommandResponse_ExtensionCommandType")
	if err != nil {
		return err
	}
	*x = ExtensionCommandResponse_ExtensionCommandType(value)
	return nil
}
func (ExtensionCommandResponse_ExtensionCommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor22, []int{0, 0}
}

// platform independent return code of the executed command.  Use the most
// specific code possible, but avoid creating too many codes.  I suggest
// creating codes only when they are 'actionable' by programmatic response,
// or for selecting/grouping in reports.
type ExtensionCommandResponse_ExtensionCommandResponseResult int32

const (
	ExtensionCommandResponse_EXTCMD_SUCCESS                  ExtensionCommandResponse_ExtensionCommandResponseResult = 0
	ExtensionCommandResponse_EXTCMD_FAILURE                  ExtensionCommandResponse_ExtensionCommandResponseResult = 1
	ExtensionCommandResponse_EXTCMD_FAIL_UNSUPPORTED_VERSION ExtensionCommandResponse_ExtensionCommandResponseResult = 2
	ExtensionCommandResponse_EXTCMD_FAIL_TRUST               ExtensionCommandResponse_ExtensionCommandResponseResult = 3
	ExtensionCommandResponse_EXTCMD_FAIL_NOT_EXIST           ExtensionCommandResponse_ExtensionCommandResponseResult = 4
	ExtensionCommandResponse_EXTCMD_FAIL_ALREADY_EXIST       ExtensionCommandResponse_ExtensionCommandResponseResult = 5
	ExtensionCommandResponse_EXTCMD_FAIL_DISABLED            ExtensionCommandResponse_ExtensionCommandResponseResult = 6
	ExtensionCommandResponse_EXTCMD_FAIL_UNSUPPORTED_TYPE    ExtensionCommandResponse_ExtensionCommandResponseResult = 7
)

var ExtensionCommandResponse_ExtensionCommandResponseResult_name = map[int32]string{
	0: "EXTCMD_SUCCESS",
	1: "EXTCMD_FAILURE",
	2: "EXTCMD_FAIL_UNSUPPORTED_VERSION",
	3: "EXTCMD_FAIL_TRUST",
	4: "EXTCMD_FAIL_NOT_EXIST",
	5: "EXTCMD_FAIL_ALREADY_EXIST",
	6: "EXTCMD_FAIL_DISABLED",
	7: "EXTCMD_FAIL_UNSUPPORTED_TYPE",
}
var ExtensionCommandResponse_ExtensionCommandResponseResult_value = map[string]int32{
	"EXTCMD_SUCCESS":                  0,
	"EXTCMD_FAILURE":                  1,
	"EXTCMD_FAIL_UNSUPPORTED_VERSION": 2,
	"EXTCMD_FAIL_TRUST":               3,
	"EXTCMD_FAIL_NOT_EXIST":           4,
	"EXTCMD_FAIL_ALREADY_EXIST":       5,
	"EXTCMD_FAIL_DISABLED":            6,
	"EXTCMD_FAIL_UNSUPPORTED_TYPE":    7,
}

func (x ExtensionCommandResponse_ExtensionCommandResponseResult) Enum() *ExtensionCommandResponse_ExtensionCommandResponseResult {
	p := new(ExtensionCommandResponse_ExtensionCommandResponseResult)
	*p = x
	return p
}
func (x ExtensionCommandResponse_ExtensionCommandResponseResult) String() string {
	return proto.EnumName(ExtensionCommandResponse_ExtensionCommandResponseResult_name, int32(x))
}
func (x *ExtensionCommandResponse_ExtensionCommandResponseResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExtensionCommandResponse_ExtensionCommandResponseResult_value, data, "ExtensionCommandResponse_ExtensionCommandResponseResult")
	if err != nil {
		return err
	}
	*x = ExtensionCommandResponse_ExtensionCommandResponseResult(value)
	return nil
}
func (ExtensionCommandResponse_ExtensionCommandResponseResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor22, []int{0, 1}
}

// Message return the result of a previous comand to install/uninstall, etc.
type ExtensionCommandResponse struct {
	// The standard fields
	// The time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp       *int64  `protobuf:"varint,1,opt,name=timeStamp" json:"timeStamp,omitempty"`
	SiteId          *string `protobuf:"bytes,2,opt,name=siteId" json:"siteId,omitempty"`
	ServerUUID      *string `protobuf:"bytes,3,opt,name=serverUUID" json:"serverUUID,omitempty"`
	ServerJobID     *int64  `protobuf:"varint,4,opt,name=serverJobID" json:"serverJobID,omitempty"`
	CorrelationUUID *string `protobuf:"bytes,5,opt,name=correlationUUID" json:"correlationUUID,omitempty"`
	AgentGUID       *string `protobuf:"bytes,6,opt,name=agentGUID" json:"agentGUID,omitempty"`
	MessageVersion  *int32  `protobuf:"varint,7,opt,name=messageVersion,def=1" json:"messageVersion,omitempty"`
	// (echoed back for server's convenience)
	ExtensionUUID *string                                                  `protobuf:"bytes,8,opt,name=extensionUUID" json:"extensionUUID,omitempty"`
	Command       *ExtensionCommandResponse_ExtensionCommandType           `protobuf:"varint,9,opt,name=command,enum=ExtensionCommandResponse_ExtensionCommandType" json:"command,omitempty"`
	Result        *ExtensionCommandResponse_ExtensionCommandResponseResult `protobuf:"varint,10,opt,name=result,enum=ExtensionCommandResponse_ExtensionCommandResponseResult" json:"result,omitempty"`
	// if you use the generic failure code of EXTCMD_FAILURE, you are encouraged to
	// set the strResultText field to provide some context.
	StrResultText    *string `protobuf:"bytes,11,opt,name=strResultText" json:"strResultText,omitempty"`
	Uuid             *string `protobuf:"bytes,12,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExtensionCommandResponse) Reset()                    { *m = ExtensionCommandResponse{} }
func (m *ExtensionCommandResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtensionCommandResponse) ProtoMessage()               {}
func (*ExtensionCommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

const Default_ExtensionCommandResponse_MessageVersion int32 = 1

func (m *ExtensionCommandResponse) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *ExtensionCommandResponse) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *ExtensionCommandResponse) GetServerUUID() string {
	if m != nil && m.ServerUUID != nil {
		return *m.ServerUUID
	}
	return ""
}

func (m *ExtensionCommandResponse) GetServerJobID() int64 {
	if m != nil && m.ServerJobID != nil {
		return *m.ServerJobID
	}
	return 0
}

func (m *ExtensionCommandResponse) GetCorrelationUUID() string {
	if m != nil && m.CorrelationUUID != nil {
		return *m.CorrelationUUID
	}
	return ""
}

func (m *ExtensionCommandResponse) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *ExtensionCommandResponse) GetMessageVersion() int32 {
	if m != nil && m.MessageVersion != nil {
		return *m.MessageVersion
	}
	return Default_ExtensionCommandResponse_MessageVersion
}

func (m *ExtensionCommandResponse) GetExtensionUUID() string {
	if m != nil && m.ExtensionUUID != nil {
		return *m.ExtensionUUID
	}
	return ""
}

func (m *ExtensionCommandResponse) GetCommand() ExtensionCommandResponse_ExtensionCommandType {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ExtensionCommandResponse_Install
}

func (m *ExtensionCommandResponse) GetResult() ExtensionCommandResponse_ExtensionCommandResponseResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ExtensionCommandResponse_EXTCMD_SUCCESS
}

func (m *ExtensionCommandResponse) GetStrResultText() string {
	if m != nil && m.StrResultText != nil {
		return *m.StrResultText
	}
	return ""
}

func (m *ExtensionCommandResponse) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtensionCommandResponse)(nil), "ExtensionCommandResponse")
	proto.RegisterEnum("ExtensionCommandResponse_ExtensionCommandType", ExtensionCommandResponse_ExtensionCommandType_name, ExtensionCommandResponse_ExtensionCommandType_value)
	proto.RegisterEnum("ExtensionCommandResponse_ExtensionCommandResponseResult", ExtensionCommandResponse_ExtensionCommandResponseResult_name, ExtensionCommandResponse_ExtensionCommandResponseResult_value)
}

var fileDescriptor22 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0xc6, 0x65, 0x81, 0xb2, 0xbc, 0x08, 0xd6, 0x51, 0xcc, 0xa0, 0xab, 0xac, 0x78, 0xf1, 0xd4,
	0x44, 0x4f, 0xc6, 0x8b, 0xa1, 0x74, 0xd4, 0x1a, 0x16, 0x48, 0xff, 0x6c, 0x76, 0x4f, 0xa4, 0x0b,
	0xe3, 0xa6, 0x49, 0xdb, 0x69, 0x3a, 0xc3, 0x46, 0x3d, 0x7a, 0xf2, 0x0b, 0xf9, 0xe5, 0x3c, 0xf9,
	0xb6, 0x65, 0x93, 0x4a, 0x6c, 0xb2, 0x37, 0xe6, 0xf9, 0xbd, 0xcf, 0xf0, 0xcc, 0xc3, 0x0b, 0xbc,
	0x60, 0xdf, 0x14, 0x4f, 0x64, 0x28, 0x92, 0x99, 0x88, 0xe3, 0x20, 0xd9, 0x3a, 0x5c, 0xa6, 0x22,
	0x91, 0xdc, 0x48, 0x33, 0xa1, 0xc4, 0x53, 0xb2, 0x12, 0x52, 0x5d, 0x67, 0x5c, 0x7a, 0xdf, 0xd3,
	0xbd, 0x36, 0xf9, 0xa5, 0x01, 0xad, 0xb3, 0x91, 0x53, 0xe8, 0xaa, 0x30, 0xe6, 0xae, 0x0a, 0xe2,
	0x94, 0x36, 0x4e, 0x1b, 0xaf, 0x9b, 0x66, 0xff, 0xe7, 0x6f, 0x5a, 0x88, 0x32, 0x17, 0xc9, 0x00,
	0x34, 0x19, 0x2a, 0x6e, 0x6f, 0xe9, 0x11, 0xe2, 0x2e, 0x39, 0x01, 0x90, 0x3c, 0xbb, 0xe1, 0x99,
	0xef, 0xdb, 0x16, 0x6d, 0xe6, 0x9a, 0x79, 0x8c, 0x96, 0xd6, 0x6e, 0x17, 0x6e, 0xc9, 0x23, 0xe8,
	0x95, 0xf4, 0x8b, 0xb8, 0x42, 0xdc, 0xca, 0x6f, 0x24, 0x2f, 0xe1, 0xc1, 0x46, 0x64, 0x19, 0x8f,
	0x02, 0x85, 0x11, 0x0a, 0x5f, 0xfb, 0xc0, 0xf7, 0x0c, 0xba, 0xc1, 0x35, 0x4f, 0xd4, 0xa7, 0x1c,
	0x6a, 0x07, 0x70, 0x04, 0x03, 0x4c, 0x23, 0x71, 0xe0, 0x9c, 0x67, 0xf9, 0x2b, 0x68, 0x07, 0x27,
	0xda, 0xef, 0x1b, 0x6f, 0xc8, 0x18, 0xfa, 0xfc, 0xf6, 0x6d, 0xc5, 0xc5, 0xc7, 0x07, 0xde, 0x0f,
	0xd0, 0xd9, 0x94, 0x6f, 0xa6, 0x5d, 0x44, 0x83, 0xb7, 0x86, 0x51, 0xdb, 0xe1, 0x21, 0xc8, 0x4b,
	0x24, 0x9f, 0x41, 0xc3, 0x3e, 0x77, 0x91, 0xa2, 0x50, 0xf8, 0xdf, 0xdd, 0xdd, 0x7f, 0x0b, 0x9c,
	0xc2, 0x4f, 0x86, 0xd0, 0x97, 0x2a, 0x2b, 0x0f, 0x1e, 0xa6, 0xa6, 0xbd, 0xa2, 0xd0, 0x27, 0x50,
	0x24, 0xa5, 0xf7, 0xff, 0x4d, 0x3e, 0xf1, 0xe0, 0xf1, 0x7f, 0x03, 0xf5, 0xa0, 0x63, 0x27, 0xf8,
	0xdb, 0x44, 0x91, 0xde, 0x20, 0x00, 0x9a, 0xc3, 0x63, 0x71, 0xc3, 0xf5, 0xa3, 0xfc, 0x33, 0x4b,
	0x82, 0xab, 0x88, 0xeb, 0xcd, 0x7c, 0xc8, 0x0a, 0x65, 0x71, 0x68, 0xe5, 0xc0, 0x4f, 0xb7, 0x81,
	0xe2, 0x7a, 0x7b, 0xf2, 0xa7, 0x51, 0xbf, 0x44, 0xfb, 0x9c, 0x04, 0x06, 0xec, 0xc2, 0x9b, 0x9d,
	0x59, 0x6b, 0xd7, 0x9f, 0xcd, 0x98, 0xeb, 0xea, 0xf7, 0x2a, 0xda, 0xc7, 0xa9, 0x3d, 0xf7, 0x1d,
	0x86, 0xdf, 0xfd, 0x0a, 0xc6, 0x15, 0x6d, 0xed, 0x2f, 0x5c, 0x7f, 0xb5, 0x5a, 0x3a, 0x1e, 0xb3,
	0xd6, 0xe7, 0xcc, 0x71, 0xed, 0xe5, 0x02, 0x43, 0x0d, 0xe1, 0x61, 0x75, 0xc8, 0x73, 0x7c, 0xd7,
	0xc3, 0x7c, 0x23, 0x18, 0x56, 0xe5, 0xc5, 0xd2, 0x5b, 0xb3, 0x0b, 0x1b, 0x51, 0x8b, 0x3c, 0x87,
	0x51, 0x15, 0x4d, 0xe7, 0x0e, 0x9b, 0x5a, 0x97, 0x7b, 0xdc, 0x26, 0x14, 0x6b, 0xa9, 0x60, 0xcb,
	0x76, 0xa7, 0xe6, 0x9c, 0x59, 0xba, 0x86, 0xbb, 0x7c, 0x52, 0x97, 0xc7, 0xbb, 0x5c, 0x31, 0xbd,
	0x63, 0x32, 0x98, 0xe0, 0x32, 0x18, 0x3f, 0xc2, 0xaf, 0x58, 0x80, 0x51, 0x2e, 0x6a, 0xf9, 0x2f,
	0xd9, 0x88, 0xc8, 0xd8, 0xef, 0x98, 0x39, 0xae, 0xeb, 0xe7, 0xac, 0x1c, 0xf8, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x71, 0x5c, 0xc0, 0x7e, 0x86, 0x03, 0x00, 0x00,
}