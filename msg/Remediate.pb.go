// Code generated by protoc-gen-go.
// source: Remediate.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Enable will only roll-back what was done by the disable command
type Remediate_Command int32

const (
	Remediate_ENABLE  Remediate_Command = 0
	Remediate_DISABLE Remediate_Command = 1
)

var Remediate_Command_name = map[int32]string{
	0: "ENABLE",
	1: "DISABLE",
}
var Remediate_Command_value = map[string]int32{
	"ENABLE":  0,
	"DISABLE": 1,
}

func (x Remediate_Command) Enum() *Remediate_Command {
	p := new(Remediate_Command)
	*p = x
	return p
}
func (x Remediate_Command) String() string {
	return proto.EnumName(Remediate_Command_name, int32(x))
}
func (x *Remediate_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Remediate_Command_value, data, "Remediate_Command")
	if err != nil {
		return err
	}
	*x = Remediate_Command(value)
	return nil
}
func (Remediate_Command) EnumDescriptor() ([]byte, []int) { return fileDescriptor74, []int{0, 0} }

type Remediate_ObjectType int32

const (
	Remediate_TYPE_SERVICE      Remediate_ObjectType = 1
	Remediate_TYPE_BHO          Remediate_ObjectType = 2
	Remediate_TYPE_AUTOSTART    Remediate_ObjectType = 3
	Remediate_TYPE_STARTUPFILE  Remediate_ObjectType = 4
	Remediate_TYPE_ACTIVEX      Remediate_ObjectType = 5
	Remediate_TYPE_OFFICEPLUGIN Remediate_ObjectType = 6
)

var Remediate_ObjectType_name = map[int32]string{
	1: "TYPE_SERVICE",
	2: "TYPE_BHO",
	3: "TYPE_AUTOSTART",
	4: "TYPE_STARTUPFILE",
	5: "TYPE_ACTIVEX",
	6: "TYPE_OFFICEPLUGIN",
}
var Remediate_ObjectType_value = map[string]int32{
	"TYPE_SERVICE":      1,
	"TYPE_BHO":          2,
	"TYPE_AUTOSTART":    3,
	"TYPE_STARTUPFILE":  4,
	"TYPE_ACTIVEX":      5,
	"TYPE_OFFICEPLUGIN": 6,
}

func (x Remediate_ObjectType) Enum() *Remediate_ObjectType {
	p := new(Remediate_ObjectType)
	*p = x
	return p
}
func (x Remediate_ObjectType) String() string {
	return proto.EnumName(Remediate_ObjectType_name, int32(x))
}
func (x *Remediate_ObjectType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Remediate_ObjectType_value, data, "Remediate_ObjectType")
	if err != nil {
		return err
	}
	*x = Remediate_ObjectType(value)
	return nil
}
func (Remediate_ObjectType) EnumDescriptor() ([]byte, []int) { return fileDescriptor74, []int{0, 1} }

// Message instructing the agent to execute a command (currently only disable/enable) on an autostart
// For example, Disable Service Ziften
type Remediate struct {
	// The UTC time that the message was produced by the server.
	TimeStamp  *int64                `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	Command    *Remediate_Command    `protobuf:"varint,2,req,name=command,enum=Remediate_Command" json:"command,omitempty"`
	ObjectType *Remediate_ObjectType `protobuf:"varint,3,req,name=objectType,enum=Remediate_ObjectType" json:"objectType,omitempty"`
	// The service name for TYPE_SERVICE (Must match 'name' field of 'services' message sent to server)
	// The fully qualified filename of the binary: i.e. 'c:\windows\system32\binary.exe' for all other types (Must match 'imageFilepath' field of message sent to server)
	ObjectName *string `protobuf:"bytes,4,req,name=objectName" json:"objectName,omitempty"`
	// The server assigns a remediation identifier to all remediation jobs
	// The agent will return the identifier in the remediation response
	RemediationIdentifier *uint64 `protobuf:"varint,5,req,name=remediationIdentifier" json:"remediationIdentifier,omitempty"`
	// Hash of server name
	SiteId *string `protobuf:"bytes,6,opt,name=siteId" json:"siteId,omitempty"`
	// string passed from server to pass back to server with remediate response
	Md5              *string `protobuf:"bytes,7,opt,name=md5" json:"md5,omitempty"`
	Uuid             *string `protobuf:"bytes,8,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Remediate) Reset()                    { *m = Remediate{} }
func (m *Remediate) String() string            { return proto.CompactTextString(m) }
func (*Remediate) ProtoMessage()               {}
func (*Remediate) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{0} }

func (m *Remediate) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *Remediate) GetCommand() Remediate_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return Remediate_ENABLE
}

func (m *Remediate) GetObjectType() Remediate_ObjectType {
	if m != nil && m.ObjectType != nil {
		return *m.ObjectType
	}
	return Remediate_TYPE_SERVICE
}

func (m *Remediate) GetObjectName() string {
	if m != nil && m.ObjectName != nil {
		return *m.ObjectName
	}
	return ""
}

func (m *Remediate) GetRemediationIdentifier() uint64 {
	if m != nil && m.RemediationIdentifier != nil {
		return *m.RemediationIdentifier
	}
	return 0
}

func (m *Remediate) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *Remediate) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *Remediate) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Remediate)(nil), "Remediate")
	proto.RegisterEnum("Remediate_Command", Remediate_Command_name, Remediate_Command_value)
	proto.RegisterEnum("Remediate_ObjectType", Remediate_ObjectType_name, Remediate_ObjectType_value)
}

var fileDescriptor74 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x4f, 0xea, 0x40,
	0x1c, 0xc4, 0x1f, 0x6d, 0x69, 0xe1, 0xff, 0x08, 0x6f, 0xd9, 0x3c, 0xcc, 0xc6, 0xc4, 0x84, 0xd4,
	0x8b, 0x5e, 0x7a, 0x30, 0xd1, 0x7b, 0x8b, 0x45, 0x9b, 0x20, 0x6d, 0x68, 0x21, 0x7a, 0x32, 0x95,
	0x2e, 0x64, 0x8d, 0xdb, 0x92, 0x76, 0x31, 0xd1, 0x8b, 0x89, 0xdf, 0xc9, 0x4f, 0xe7, 0xc5, 0xa5,
	0x85, 0xca, 0x71, 0x66, 0x7e, 0x3b, 0x99, 0xfc, 0x17, 0xfe, 0x4d, 0x29, 0xa7, 0x09, 0x8b, 0x05,
	0xb5, 0xd6, 0x79, 0x26, 0xb2, 0x63, 0x1c, 0x64, 0x85, 0x58, 0xe5, 0xb4, 0x88, 0xde, 0xd6, 0x3b,
	0xcf, 0xfc, 0x56, 0xa0, 0x5d, 0x73, 0xb8, 0x07, 0x6d, 0xc1, 0x38, 0x0d, 0x45, 0xcc, 0xd7, 0xa4,
	0x31, 0x50, 0xce, 0x54, 0x7c, 0x0a, 0xc6, 0x22, 0xe3, 0x3c, 0x4e, 0x13, 0xa2, 0x48, 0xa3, 0x7b,
	0x81, 0xad, 0xdf, 0xde, 0x61, 0x95, 0xe0, 0x73, 0x80, 0xec, 0xe9, 0x99, 0x2e, 0xc4, 0xb6, 0x99,
	0xa8, 0x25, 0xd7, 0x3f, 0xe0, 0xfc, 0x3a, 0xc4, 0x78, 0x8f, 0x4e, 0x62, 0x4e, 0x89, 0x26, 0xd1,
	0x36, 0x3e, 0x81, 0x7e, 0xbe, 0x63, 0x59, 0x96, 0x7a, 0x09, 0x4d, 0x05, 0x5b, 0x32, 0x9a, 0x93,
	0xa6, 0x8c, 0x35, 0xdc, 0x05, 0xbd, 0x60, 0x82, 0x7a, 0x09, 0xd1, 0x07, 0x0d, 0x89, 0xff, 0x05,
	0x95, 0x27, 0x97, 0xc4, 0x28, 0xc5, 0x11, 0x68, 0x9b, 0x0d, 0x4b, 0x48, 0x6b, 0xab, 0x9c, 0xd6,
	0xe7, 0x17, 0x29, 0xb5, 0x69, 0x82, 0xb1, 0x5f, 0x07, 0xa0, 0xbb, 0x13, 0xdb, 0x19, 0xbb, 0xe8,
	0x8f, 0x7c, 0x6b, 0x5c, 0x7b, 0x61, 0x29, 0x1a, 0xe6, 0x07, 0xc0, 0xc1, 0x32, 0x04, 0x9d, 0xe8,
	0x21, 0x70, 0x1f, 0x43, 0x77, 0x3a, 0xf7, 0x86, 0x32, 0xc7, 0x1d, 0x68, 0x95, 0x8e, 0x73, 0xeb,
	0x23, 0x45, 0x2e, 0xef, 0x96, 0xca, 0x9e, 0x45, 0x7e, 0x18, 0xd9, 0xd3, 0x08, 0xa9, 0xf8, 0x3f,
	0xa0, 0xea, 0xcd, 0x56, 0xcf, 0x82, 0x91, 0x27, 0x7b, 0xb5, 0xba, 0xc9, 0x1e, 0x46, 0xde, 0xdc,
	0xbd, 0x47, 0x4d, 0xdc, 0x87, 0x5e, 0xe9, 0xf8, 0xa3, 0x91, 0xac, 0x0e, 0xc6, 0xb3, 0x1b, 0x6f,
	0x82, 0x74, 0xe7, 0x0a, 0x4c, 0x79, 0x5c, 0xeb, 0x9d, 0x2d, 0x05, 0x4d, 0xad, 0x82, 0xe6, 0xaf,
	0x34, 0xaf, 0x3e, 0x66, 0x91, 0xbd, 0x58, 0x9c, 0x16, 0x45, 0xbc, 0xa2, 0x0e, 0xaa, 0x0f, 0x79,
	0x57, 0x39, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x8c, 0x14, 0x9a, 0xdb, 0x01, 0x00, 0x00,
}
