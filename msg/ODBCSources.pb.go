// Code generated by protoc-gen-go.
// source: ODBCSources.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message reporting list of ODBC Data Sources
type ODBCSources struct {
	// The time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Each entry in 'User DSN' or 'SystemDSN' will have a name, description and server string entry
	// Names of ODBC Source's
	Name []string `protobuf:"bytes,3,rep,name=name" json:"name,omitempty"`
	// Description of ODBC Source's
	Description []string `protobuf:"bytes,4,rep,name=description" json:"description,omitempty"`
	// Server definition for ODBC Source's
	Server []string `protobuf:"bytes,5,rep,name=server" json:"server,omitempty"`
	// Is this a USER defined source (HKEY_CURRENT_USER\Software\ODBC\ODBC.ini) = TRUE, or 'User DSN'
	// or a SYSTEM defined source (HKEY_LOCAL_MACHINE\Software\ODBC\ODBC.ini) = FALSE, or 'System DSN'
	UserDSN []bool `protobuf:"varint,6,rep,name=userDSN" json:"userDSN,omitempty"`
	// 'Trusted_Connection', With Integrated Windows authentication, or With SQL Server authentication
	TrustedConnection []bool `protobuf:"varint,7,rep,name=TrustedConnection" json:"TrustedConnection,omitempty"`
	// 'LastUser', only if NOT Trusted Connection
	SQLUserName []string `protobuf:"bytes,8,rep,name=SQLUserName" json:"SQLUserName,omitempty"`
	// 'ServerSPN', only if Trusted Connection
	ServerSPN []string `protobuf:"bytes,9,rep,name=ServerSPN" json:"ServerSPN,omitempty"`
	// 'Database'
	Database []string `protobuf:"bytes,10,rep,name=Database" json:"Database,omitempty"`
	// Name of the ODBCDriver
	ODBCDriverName   []string `protobuf:"bytes,11,rep,name=ODBCDriverName" json:"ODBCDriverName,omitempty"`
	SiteId           *string  `protobuf:"bytes,12,opt,name=siteId" json:"siteId,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ODBCSources) Reset()                    { *m = ODBCSources{} }
func (m *ODBCSources) String() string            { return proto.CompactTextString(m) }
func (*ODBCSources) ProtoMessage()               {}
func (*ODBCSources) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{0} }

func (m *ODBCSources) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *ODBCSources) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *ODBCSources) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ODBCSources) GetDescription() []string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ODBCSources) GetServer() []string {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *ODBCSources) GetUserDSN() []bool {
	if m != nil {
		return m.UserDSN
	}
	return nil
}

func (m *ODBCSources) GetTrustedConnection() []bool {
	if m != nil {
		return m.TrustedConnection
	}
	return nil
}

func (m *ODBCSources) GetSQLUserName() []string {
	if m != nil {
		return m.SQLUserName
	}
	return nil
}

func (m *ODBCSources) GetServerSPN() []string {
	if m != nil {
		return m.ServerSPN
	}
	return nil
}

func (m *ODBCSources) GetDatabase() []string {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ODBCSources) GetODBCDriverName() []string {
	if m != nil {
		return m.ODBCDriverName
	}
	return nil
}

func (m *ODBCSources) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func init() {
	proto.RegisterType((*ODBCSources)(nil), "ODBCSources")
}

var fileDescriptor40 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0xd5, 0x1f, 0xda, 0x78, 0x5b, 0x0a, 0x31, 0x12, 0x32, 0x70, 0x89, 0x7a, 0xea, 0x29,
	0x67, 0xce, 0x69, 0x24, 0x54, 0x09, 0x42, 0x91, 0xdb, 0x07, 0x30, 0xc9, 0x52, 0x59, 0x22, 0x76,
	0x64, 0x3b, 0x95, 0xe0, 0xc8, 0x9d, 0xc7, 0xe1, 0xfd, 0x70, 0x1c, 0x84, 0x7a, 0xf4, 0xe7, 0xd9,
	0xd9, 0x9d, 0x81, 0xf8, 0x39, 0xcf, 0xd6, 0x5c, 0xb7, 0xa6, 0x44, 0x9b, 0x36, 0x46, 0x3b, 0x7d,
	0x4b, 0xb7, 0xda, 0xba, 0x83, 0x41, 0xbb, 0xfb, 0x68, 0xb0, 0x67, 0xcb, 0xef, 0x21, 0xcc, 0x4e,
	0x94, 0x34, 0x01, 0xe2, 0x64, 0x8d, 0xdc, 0x89, 0xba, 0x61, 0x83, 0x64, 0xb8, 0x1a, 0x65, 0xe7,
	0x5f, 0x3f, 0x2c, 0x40, 0xdb, 0x41, 0x7a, 0x07, 0x44, 0x1c, 0x50, 0xb9, 0x87, 0xfd, 0x26, 0x67,
	0x43, 0xaf, 0x20, 0x59, 0xe4, 0x15, 0xe3, 0xb6, 0x95, 0x15, 0x9d, 0xc3, 0x58, 0x89, 0x1a, 0xd9,
	0x28, 0x19, 0xad, 0x08, 0xbd, 0x82, 0x59, 0x85, 0xb6, 0x34, 0xb2, 0x71, 0x52, 0x2b, 0x36, 0x0e,
	0x70, 0x01, 0x13, 0x8b, 0xe6, 0x88, 0x86, 0x9d, 0x85, 0xf7, 0x05, 0x4c, 0x5b, 0x0f, 0x72, 0x5e,
	0xb0, 0x89, 0x07, 0x11, 0xbd, 0x81, 0x78, 0x67, 0x5a, 0xeb, 0xb0, 0x5a, 0x6b, 0xa5, 0xb0, 0x0c,
	0xb3, 0xd3, 0xf0, 0xe5, 0x0d, 0xf9, 0xcb, 0xe3, 0xde, 0xcb, 0x8b, 0x6e, 0x4b, 0x14, 0x0c, 0x62,
	0x20, 0x3c, 0x18, 0xf2, 0x6d, 0xc1, 0x48, 0x40, 0x97, 0x10, 0xe5, 0xc2, 0x89, 0x57, 0x61, 0x91,
	0x41, 0x20, 0xd7, 0xb0, 0xe8, 0x62, 0xe6, 0x46, 0x1e, 0xff, 0x86, 0x67, 0xff, 0xd7, 0x48, 0x87,
	0x9b, 0x8a, 0xcd, 0x93, 0x81, 0x8f, 0x72, 0x0f, 0xcb, 0x52, 0xd7, 0xe9, 0xa7, 0x7c, 0x73, 0xa8,
	0xd2, 0xfe, 0xd0, 0xbe, 0xaa, 0x52, 0xbf, 0xa7, 0xbe, 0x03, 0xeb, 0xc3, 0x67, 0xf4, 0xa4, 0xb2,
	0xa7, 0x9e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x7e, 0x91, 0x41, 0x71, 0x01, 0x00, 0x00,
}
