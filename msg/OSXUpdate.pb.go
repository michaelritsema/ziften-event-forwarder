// Code generated by protoc-gen-go.
// source: OSXUpdate.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXUpdate_UpdateOperation int32

const (
	OSXUpdate_uoInstallation   OSXUpdate_UpdateOperation = 1
	OSXUpdate_uoUninstallation OSXUpdate_UpdateOperation = 2
)

var OSXUpdate_UpdateOperation_name = map[int32]string{
	1: "uoInstallation",
	2: "uoUninstallation",
}
var OSXUpdate_UpdateOperation_value = map[string]int32{
	"uoInstallation":   1,
	"uoUninstallation": 2,
}

func (x OSXUpdate_UpdateOperation) Enum() *OSXUpdate_UpdateOperation {
	p := new(OSXUpdate_UpdateOperation)
	*p = x
	return p
}
func (x OSXUpdate_UpdateOperation) String() string {
	return proto.EnumName(OSXUpdate_UpdateOperation_name, int32(x))
}
func (x *OSXUpdate_UpdateOperation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSXUpdate_UpdateOperation_value, data, "OSXUpdate_UpdateOperation")
	if err != nil {
		return err
	}
	*x = OSXUpdate_UpdateOperation(value)
	return nil
}
func (OSXUpdate_UpdateOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor60, []int{0, 0}
}

type OSXUpdate_OperationResultCode int32

const (
	OSXUpdate_orcNotStarted          OSXUpdate_OperationResultCode = 0
	OSXUpdate_orcInProgress          OSXUpdate_OperationResultCode = 1
	OSXUpdate_orcSucceeded           OSXUpdate_OperationResultCode = 2
	OSXUpdate_orcSucceededWithErrors OSXUpdate_OperationResultCode = 3
	OSXUpdate_orcFailed              OSXUpdate_OperationResultCode = 4
	OSXUpdate_orcAborted             OSXUpdate_OperationResultCode = 5
)

var OSXUpdate_OperationResultCode_name = map[int32]string{
	0: "orcNotStarted",
	1: "orcInProgress",
	2: "orcSucceeded",
	3: "orcSucceededWithErrors",
	4: "orcFailed",
	5: "orcAborted",
}
var OSXUpdate_OperationResultCode_value = map[string]int32{
	"orcNotStarted":          0,
	"orcInProgress":          1,
	"orcSucceeded":           2,
	"orcSucceededWithErrors": 3,
	"orcFailed":              4,
	"orcAborted":             5,
}

func (x OSXUpdate_OperationResultCode) Enum() *OSXUpdate_OperationResultCode {
	p := new(OSXUpdate_OperationResultCode)
	*p = x
	return p
}
func (x OSXUpdate_OperationResultCode) String() string {
	return proto.EnumName(OSXUpdate_OperationResultCode_name, int32(x))
}
func (x *OSXUpdate_OperationResultCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSXUpdate_OperationResultCode_value, data, "OSXUpdate_OperationResultCode")
	if err != nil {
		return err
	}
	*x = OSXUpdate_OperationResultCode(value)
	return nil
}
func (OSXUpdate_OperationResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor60, []int{0, 1}
}

type OSXUpdate struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID           *string                         `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	ClientApplicationID []string                        `protobuf:"bytes,3,rep,name=clientApplicationID" json:"clientApplicationID,omitempty"`
	DateLocal           []int64                         `protobuf:"varint,4,rep,name=dateLocal" json:"dateLocal,omitempty"`
	DateUTC             []int64                         `protobuf:"varint,5,rep,name=dateUTC" json:"dateUTC,omitempty"`
	Description         []string                        `protobuf:"bytes,6,rep,name=description" json:"description,omitempty"`
	Operation           []OSXUpdate_UpdateOperation     `protobuf:"varint,7,rep,name=operation,enum=OSXUpdate_UpdateOperation" json:"operation,omitempty"`
	ResultCode          []OSXUpdate_OperationResultCode `protobuf:"varint,8,rep,name=resultCode,enum=OSXUpdate_OperationResultCode" json:"resultCode,omitempty"`
	SupportURL          []string                        `protobuf:"bytes,9,rep,name=supportURL" json:"supportURL,omitempty"`
	Title               []string                        `protobuf:"bytes,10,rep,name=title" json:"title,omitempty"`
	CategoryName        []string                        `protobuf:"bytes,11,rep,name=categoryName" json:"categoryName,omitempty"`
	CategoryType        []string                        `protobuf:"bytes,12,rep,name=categoryType" json:"categoryType,omitempty"`
	KbID                []string                        `protobuf:"bytes,13,rep,name=kbID" json:"kbID,omitempty"`
	SiteId              *string                         `protobuf:"bytes,14,opt,name=siteId" json:"siteId,omitempty"`
	Uuid                *string                         `protobuf:"bytes,15,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized    []byte                          `json:"-"`
}

func (m *OSXUpdate) Reset()                    { *m = OSXUpdate{} }
func (m *OSXUpdate) String() string            { return proto.CompactTextString(m) }
func (*OSXUpdate) ProtoMessage()               {}
func (*OSXUpdate) Descriptor() ([]byte, []int) { return fileDescriptor60, []int{0} }

func (m *OSXUpdate) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXUpdate) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXUpdate) GetClientApplicationID() []string {
	if m != nil {
		return m.ClientApplicationID
	}
	return nil
}

func (m *OSXUpdate) GetDateLocal() []int64 {
	if m != nil {
		return m.DateLocal
	}
	return nil
}

func (m *OSXUpdate) GetDateUTC() []int64 {
	if m != nil {
		return m.DateUTC
	}
	return nil
}

func (m *OSXUpdate) GetDescription() []string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *OSXUpdate) GetOperation() []OSXUpdate_UpdateOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *OSXUpdate) GetResultCode() []OSXUpdate_OperationResultCode {
	if m != nil {
		return m.ResultCode
	}
	return nil
}

func (m *OSXUpdate) GetSupportURL() []string {
	if m != nil {
		return m.SupportURL
	}
	return nil
}

func (m *OSXUpdate) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *OSXUpdate) GetCategoryName() []string {
	if m != nil {
		return m.CategoryName
	}
	return nil
}

func (m *OSXUpdate) GetCategoryType() []string {
	if m != nil {
		return m.CategoryType
	}
	return nil
}

func (m *OSXUpdate) GetKbID() []string {
	if m != nil {
		return m.KbID
	}
	return nil
}

func (m *OSXUpdate) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXUpdate) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXUpdate)(nil), "OSXUpdate")
	proto.RegisterEnum("OSXUpdate_UpdateOperation", OSXUpdate_UpdateOperation_name, OSXUpdate_UpdateOperation_value)
	proto.RegisterEnum("OSXUpdate_OperationResultCode", OSXUpdate_OperationResultCode_name, OSXUpdate_OperationResultCode_value)
}

var fileDescriptor60 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xe2, 0xb4, 0xf5, 0x34, 0x7f, 0x8c, 0x83, 0xaa, 0x55, 0x90, 0xaa, 0x28, 0xa7,
	0x5e, 0xf0, 0xa1, 0x07, 0x2e, 0x9c, 0x9a, 0xf2, 0x47, 0x91, 0x4a, 0x5b, 0x35, 0xb5, 0xe0, 0xea,
	0xae, 0x87, 0xb0, 0x62, 0xe3, 0x59, 0xed, 0xae, 0x91, 0xca, 0x91, 0x13, 0x2f, 0xc4, 0x6b, 0xf1,
	0x0c, 0x8c, 0x37, 0x25, 0x8d, 0x4a, 0x4e, 0xab, 0xf9, 0xcd, 0xa7, 0x6f, 0x34, 0xdf, 0x0e, 0x0c,
	0xaf, 0x16, 0x9f, 0x73, 0x53, 0x16, 0x1e, 0x33, 0x63, 0xc9, 0xd3, 0x38, 0xbd, 0x26, 0xe7, 0x97,
	0x16, 0xdd, 0xed, 0xbd, 0x79, 0x60, 0xd3, 0x3f, 0x11, 0xc4, 0x1b, 0x5d, 0x3a, 0x81, 0xd8, 0xab,
	0x15, 0x2e, 0x7c, 0xb1, 0x32, 0xa2, 0x35, 0x69, 0x9f, 0x74, 0x66, 0xfd, 0x9f, 0xbf, 0x45, 0x80,
	0xae, 0x81, 0xe9, 0x4b, 0x88, 0x8b, 0x25, 0x56, 0xfe, 0x43, 0x3e, 0x7f, 0x2b, 0xda, 0xac, 0x88,
	0x67, 0x07, 0xac, 0x88, 0xea, 0x5a, 0x95, 0xdc, 0x1c, 0x49, 0xad, 0xb8, 0x7b, 0x66, 0x8c, 0x56,
	0xb2, 0xf0, 0x8a, 0x2a, 0x96, 0x75, 0x26, 0x9d, 0x93, 0xb8, 0xf1, 0x6e, 0x66, 0x5c, 0x90, 0x2c,
	0xb4, 0x88, 0x18, 0xfd, 0xe7, 0x7d, 0x0c, 0xfb, 0x8d, 0x22, 0xbf, 0x3d, 0x17, 0xdd, 0x5d, 0xfd,
	0x11, 0x1c, 0x96, 0xe8, 0xa4, 0x55, 0xa6, 0x31, 0x16, 0x7b, 0xc1, 0xf6, 0x15, 0xc4, 0x64, 0xd0,
	0x86, 0x59, 0x62, 0x9f, 0xd1, 0xe0, 0x74, 0x9c, 0x3d, 0x6e, 0xbe, 0x7e, 0xae, 0xfe, 0x29, 0xd2,
	0x53, 0x00, 0x0e, 0xa0, 0xd6, 0xfe, 0x9c, 0x4a, 0x14, 0x07, 0x41, 0x7f, 0xbc, 0xa5, 0xdf, 0x28,
	0x6f, 0x36, 0xaa, 0x34, 0x05, 0x70, 0xb5, 0x31, 0x64, 0x7d, 0x7e, 0x73, 0x21, 0xe2, 0x30, 0xb6,
	0x0f, 0x5d, 0xaf, 0xbc, 0x46, 0x01, 0xa1, 0x7c, 0x01, 0x3d, 0x5e, 0x17, 0x97, 0x64, 0xef, 0x2f,
	0x8b, 0x15, 0x8a, 0xc3, 0xa7, 0xb4, 0x89, 0x5c, 0xf4, 0x02, 0xed, 0x41, 0xf4, 0xed, 0x8e, 0x63,
	0xe9, 0x87, 0x6a, 0x00, 0x7b, 0x4e, 0x79, 0x9c, 0x97, 0x62, 0x30, 0x69, 0x71, 0x7d, 0x04, 0x21,
	0x4b, 0x31, 0x6c, 0xaa, 0xc7, 0x6c, 0xa7, 0x6f, 0x60, 0xf8, 0x74, 0x97, 0x14, 0x06, 0x35, 0xcd,
	0x2b, 0x0e, 0x47, 0xeb, 0x40, 0x92, 0x16, 0x8f, 0x4c, 0x6a, 0xca, 0x2b, 0xb5, 0x4d, 0xdb, 0xd3,
	0x5f, 0x2d, 0x18, 0xed, 0xda, 0xec, 0x39, 0xf4, 0xc9, 0xca, 0x4b, 0xf2, 0xfc, 0xe3, 0xd6, 0x63,
	0x99, 0x3c, 0x7b, 0x40, 0xf3, 0xea, 0xda, 0x52, 0x73, 0x2a, 0x8e, 0x3d, 0x13, 0xe8, 0x31, 0x5a,
	0xd4, 0x52, 0x22, 0x96, 0x2c, 0x6a, 0xa7, 0x63, 0x38, 0xda, 0x26, 0x9f, 0x94, 0xff, 0xfa, 0xce,
	0x5a, 0xb2, 0x2e, 0xe9, 0x70, 0x32, 0x31, 0xf7, 0xde, 0x17, 0x4a, 0xb3, 0x34, 0xe2, 0xfd, 0x80,
	0xcb, 0xb3, 0x3b, 0x0a, 0xfe, 0xdd, 0xd9, 0x6b, 0x98, 0x4a, 0x5a, 0x65, 0x3f, 0xd4, 0x17, 0x8f,
	0x55, 0xe6, 0xd0, 0x7e, 0x47, 0xbb, 0xbe, 0x45, 0x49, 0x3a, 0xe3, 0xaf, 0x76, 0x7c, 0x5f, 0xb3,
	0x64, 0xf3, 0x23, 0x1f, 0xd7, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xe7, 0x7e, 0x56,
	0xce, 0x02, 0x00, 0x00,
}
