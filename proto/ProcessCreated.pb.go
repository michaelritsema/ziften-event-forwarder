// Code generated by protoc-gen-go.
// source: ProcessCreated.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcessCreated struct {
	TimeStamp         *int64  `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	AgentGUID         *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	PID               *int32  `protobuf:"varint,3,req,name=PID" json:"PID,omitempty"`
	ProcessUUID       *string `protobuf:"bytes,4,req,name=processUUID" json:"processUUID,omitempty"`
	ProcessImagePath  *string `protobuf:"bytes,5,req,name=processImagePath" json:"processImagePath,omitempty"`
	ProcessMD5        *string `protobuf:"bytes,6,req,name=processMD5" json:"processMD5,omitempty"`
	ParentProcessPID  *int32  `protobuf:"varint,7,req,name=parentProcessPID" json:"parentProcessPID,omitempty"`
	ParentProcessUUID *string `protobuf:"bytes,8,req,name=parentProcessUUID" json:"parentProcessUUID,omitempty"`
	AccountName       *string `protobuf:"bytes,9,req,name=accountName" json:"accountName,omitempty"`
	DomainName        *string `protobuf:"bytes,10,req,name=domainName" json:"domainName,omitempty"`
	CommandLine       *string `protobuf:"bytes,11,opt,name=commandLine" json:"commandLine,omitempty"`
	SiteId            *string `protobuf:"bytes,12,opt,name=siteId" json:"siteId,omitempty"`
	Uuid              *string `protobuf:"bytes,13,opt,name=uuid" json:"uuid,omitempty"`
	ProcessWriter     *string `protobuf:"bytes,14,opt,name=processWriter" json:"processWriter,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *ProcessCreated) Reset()                    { *m = ProcessCreated{} }
func (m *ProcessCreated) String() string            { return proto1.CompactTextString(m) }
func (*ProcessCreated) ProtoMessage()               {}
func (*ProcessCreated) Descriptor() ([]byte, []int) { return fileDescriptor69, []int{0} }

func (m *ProcessCreated) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *ProcessCreated) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *ProcessCreated) GetPID() int32 {
	if m != nil && m.PID != nil {
		return *m.PID
	}
	return 0
}

func (m *ProcessCreated) GetProcessUUID() string {
	if m != nil && m.ProcessUUID != nil {
		return *m.ProcessUUID
	}
	return ""
}

func (m *ProcessCreated) GetProcessImagePath() string {
	if m != nil && m.ProcessImagePath != nil {
		return *m.ProcessImagePath
	}
	return ""
}

func (m *ProcessCreated) GetProcessMD5() string {
	if m != nil && m.ProcessMD5 != nil {
		return *m.ProcessMD5
	}
	return ""
}

func (m *ProcessCreated) GetParentProcessPID() int32 {
	if m != nil && m.ParentProcessPID != nil {
		return *m.ParentProcessPID
	}
	return 0
}

func (m *ProcessCreated) GetParentProcessUUID() string {
	if m != nil && m.ParentProcessUUID != nil {
		return *m.ParentProcessUUID
	}
	return ""
}

func (m *ProcessCreated) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *ProcessCreated) GetDomainName() string {
	if m != nil && m.DomainName != nil {
		return *m.DomainName
	}
	return ""
}

func (m *ProcessCreated) GetCommandLine() string {
	if m != nil && m.CommandLine != nil {
		return *m.CommandLine
	}
	return ""
}

func (m *ProcessCreated) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *ProcessCreated) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ProcessCreated) GetProcessWriter() string {
	if m != nil && m.ProcessWriter != nil {
		return *m.ProcessWriter
	}
	return ""
}

func init() {
	proto1.RegisterType((*ProcessCreated)(nil), "ProcessCreated")
}

var fileDescriptor69 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xe9, 0xdf, 0xb7, 0x3b, 0xfb, 0xb6, 0x68, 0xb4, 0x12, 0x14, 0xa1, 0xd4, 0x8b, 0xa7,
	0xbd, 0x79, 0xf2, 0x56, 0x0b, 0x52, 0xb0, 0xb2, 0xa0, 0xc5, 0x73, 0xd8, 0x1d, 0x6b, 0xc0, 0x24,
	0x4b, 0x92, 0x0a, 0x7a, 0xf4, 0x3b, 0xf9, 0x91, 0xfc, 0x1e, 0x4e, 0xb2, 0x2b, 0xb8, 0x3d, 0xe6,
	0x37, 0x3f, 0x9e, 0x67, 0x76, 0x07, 0x8e, 0x73, 0x6b, 0x0a, 0x74, 0xee, 0xc6, 0xa2, 0xf0, 0x58,
	0x66, 0x95, 0x35, 0xde, 0x9c, 0xb2, 0xdc, 0x38, 0xbf, 0xb5, 0xe8, 0x1e, 0xdf, 0x2b, 0xac, 0xd9,
	0xfc, 0xbb, 0x0b, 0x93, 0xb6, 0xcc, 0x66, 0x90, 0x78, 0xa9, 0xf0, 0xc1, 0x0b, 0x55, 0xf1, 0xce,
	0xac, 0x7b, 0xd9, 0x5b, 0x8c, 0x3f, 0xbf, 0x78, 0x84, 0x2e, 0x40, 0x76, 0x06, 0x89, 0xd8, 0xa2,
	0xf6, 0xb7, 0x9b, 0xd5, 0x92, 0x77, 0xc9, 0x48, 0x16, 0x23, 0x32, 0xfa, 0xbb, 0x9d, 0x2c, 0x59,
	0x0a, 0xbd, 0x9c, 0x70, 0x8f, 0xf0, 0x80, 0x9d, 0x43, 0x5a, 0xd5, 0xe9, 0x9b, 0xe0, 0xf6, 0xf7,
	0x5c, 0x0e, 0x07, 0xcd, 0x78, 0xa5, 0x28, 0x31, 0x17, 0xfe, 0x85, 0x0f, 0x82, 0xc3, 0x18, 0x40,
	0x33, 0x59, 0x2f, 0xaf, 0xf8, 0x30, 0xb2, 0x60, 0x0b, 0x4b, 0xbd, 0xcd, 0xc2, 0xa1, 0xe6, 0x5f,
	0xac, 0xb9, 0x80, 0xc3, 0xd6, 0x24, 0x96, 0x8d, 0xf6, 0xca, 0x8e, 0x20, 0x15, 0x45, 0x61, 0x76,
	0xda, 0xdf, 0x0b, 0x85, 0x3c, 0xf9, 0xed, 0x29, 0x8d, 0x12, 0x52, 0x47, 0x06, 0x91, 0x91, 0x58,
	0x18, 0xa5, 0x84, 0x2e, 0xef, 0xa4, 0x46, 0x9e, 0xce, 0x3a, 0x04, 0x27, 0x30, 0x74, 0xd2, 0xe3,
	0xaa, 0xe4, 0xff, 0xe3, 0xfb, 0x04, 0x62, 0x2a, 0x1f, 0x87, 0xd7, 0x9f, 0x96, 0x29, 0x8c, 0x9b,
	0xc5, 0x9f, 0x2c, 0xf9, 0x96, 0x4f, 0xa2, 0x70, 0x0d, 0x73, 0xca, 0xcc, 0x3e, 0xe4, 0xb3, 0x47,
	0x9d, 0x39, 0xb4, 0x6f, 0x68, 0xeb, 0x13, 0x14, 0xe6, 0x35, 0xa3, 0x1f, 0xeb, 0xe8, 0xfb, 0x17,
	0xd3, 0xf6, 0x29, 0xd6, 0x35, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xd1, 0xfb, 0x34, 0xcf,
	0x01, 0x00, 0x00,
}
