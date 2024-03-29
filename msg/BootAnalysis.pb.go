// Code generated by protoc-gen-go.
// source: BootAnalysis.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Report information read from Global Logger Boot Trace after a boot occurred
type BootAnalysis struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID    *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	AnalysisGUID *string `protobuf:"bytes,3,req,name=analysisGUID" json:"analysisGUID,omitempty"`
	// All of these fields are read from the Global Logger Boot Trace.
	// The ETW Process Start/End documentation provides more information. See http://msdn.microsoft.com/en-us/magazine/ee412263.aspx
	EventTimeStamp []int64 `protobuf:"varint,4,rep,name=eventTimeStamp" json:"eventTimeStamp,omitempty"`
	// Process Start = 0
	// Process End = 1
	EventType       []int32  `protobuf:"varint,5,rep,name=eventType" json:"eventType,omitempty"`
	ProcessId       []int32  `protobuf:"varint,6,rep,name=processId" json:"processId,omitempty"`
	ParentProcessId []int32  `protobuf:"varint,7,rep,name=parentProcessId" json:"parentProcessId,omitempty"`
	ImageFileName   []string `protobuf:"bytes,8,rep,name=imageFileName" json:"imageFileName,omitempty"`
	CommandLine     []string `protobuf:"bytes,9,rep,name=commandLine" json:"commandLine,omitempty"`
	// The fully qualified filename of the binary: i.e. 'c:\windows\system32\binary.exe'. This field is derived from the command line.
	ImageFilepath []string `protobuf:"bytes,10,rep,name=imageFilepath" json:"imageFilepath,omitempty"`
	// MD5 for imageFileName
	ImageFileMD5 []string `protobuf:"bytes,11,rep,name=imageFileMD5" json:"imageFileMD5,omitempty"`
	// Who is loading this control
	AccountName       []string `protobuf:"bytes,12,rep,name=accountName" json:"accountName,omitempty"`
	DomainName        []string `protobuf:"bytes,13,rep,name=domainName" json:"domainName,omitempty"`
	SiteId            *string  `protobuf:"bytes,14,opt,name=siteId" json:"siteId,omitempty"`
	ExecutionCount    []int32  `protobuf:"varint,15,rep,name=executionCount" json:"executionCount,omitempty"`
	Uuid              *string  `protobuf:"bytes,16,opt,name=uuid" json:"uuid,omitempty"`
	ProcessUUID       []string `protobuf:"bytes,17,rep,name=processUUID" json:"processUUID,omitempty"`
	ParentProcessUUID []string `protobuf:"bytes,18,rep,name=parentProcessUUID" json:"parentProcessUUID,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *BootAnalysis) Reset()                    { *m = BootAnalysis{} }
func (m *BootAnalysis) String() string            { return proto.CompactTextString(m) }
func (*BootAnalysis) ProtoMessage()               {}
func (*BootAnalysis) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *BootAnalysis) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *BootAnalysis) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *BootAnalysis) GetAnalysisGUID() string {
	if m != nil && m.AnalysisGUID != nil {
		return *m.AnalysisGUID
	}
	return ""
}

func (m *BootAnalysis) GetEventTimeStamp() []int64 {
	if m != nil {
		return m.EventTimeStamp
	}
	return nil
}

func (m *BootAnalysis) GetEventType() []int32 {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *BootAnalysis) GetProcessId() []int32 {
	if m != nil {
		return m.ProcessId
	}
	return nil
}

func (m *BootAnalysis) GetParentProcessId() []int32 {
	if m != nil {
		return m.ParentProcessId
	}
	return nil
}

func (m *BootAnalysis) GetImageFileName() []string {
	if m != nil {
		return m.ImageFileName
	}
	return nil
}

func (m *BootAnalysis) GetCommandLine() []string {
	if m != nil {
		return m.CommandLine
	}
	return nil
}

func (m *BootAnalysis) GetImageFilepath() []string {
	if m != nil {
		return m.ImageFilepath
	}
	return nil
}

func (m *BootAnalysis) GetImageFileMD5() []string {
	if m != nil {
		return m.ImageFileMD5
	}
	return nil
}

func (m *BootAnalysis) GetAccountName() []string {
	if m != nil {
		return m.AccountName
	}
	return nil
}

func (m *BootAnalysis) GetDomainName() []string {
	if m != nil {
		return m.DomainName
	}
	return nil
}

func (m *BootAnalysis) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *BootAnalysis) GetExecutionCount() []int32 {
	if m != nil {
		return m.ExecutionCount
	}
	return nil
}

func (m *BootAnalysis) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *BootAnalysis) GetProcessUUID() []string {
	if m != nil {
		return m.ProcessUUID
	}
	return nil
}

func (m *BootAnalysis) GetParentProcessUUID() []string {
	if m != nil {
		return m.ParentProcessUUID
	}
	return nil
}

func init() {
	proto.RegisterType((*BootAnalysis)(nil), "BootAnalysis")
}

var fileDescriptor12 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x86, 0x03, 0x03, 0x64, 0x87, 0x01, 0x52, 0x14, 0xab, 0x26, 0x66, 0x21, 0x31, 0xf1, 0x6a,
	0x77, 0x5e, 0x78, 0xe9, 0x24, 0x1a, 0x12, 0x31, 0x24, 0xc2, 0x03, 0x34, 0x5b, 0xc5, 0x26, 0xac,
	0x5d, 0xd6, 0x8e, 0x88, 0x97, 0xbe, 0x93, 0x8f, 0xe0, 0x7b, 0x79, 0xd6, 0x21, 0x99, 0x7a, 0xb9,
	0xef, 0xff, 0x77, 0xce, 0xff, 0x9f, 0x02, 0x09, 0x95, 0x32, 0xb7, 0x92, 0xad, 0xb7, 0x5a, 0xe8,
	0x20, 0xcd, 0x94, 0x51, 0x67, 0x64, 0xae, 0xb4, 0x59, 0x65, 0x5c, 0x2f, 0xb6, 0x29, 0x2f, 0xd9,
	0xf8, 0xcb, 0x01, 0xaf, 0x6a, 0x25, 0x3e, 0xb8, 0x46, 0x24, 0xfc, 0xd9, 0xb0, 0x24, 0xa5, 0x35,
	0xbf, 0x7e, 0xe5, 0x84, 0xdd, 0x8f, 0x4f, 0x6a, 0xa1, 0x2e, 0x20, 0x39, 0x07, 0x97, 0xad, 0xb8,
	0x34, 0x0f, 0xcb, 0xe9, 0x84, 0xd6, 0xd1, 0xe1, 0x86, 0x6d, 0x74, 0x34, 0xf2, 0x5c, 0xc4, 0xe4,
	0x02, 0x3c, 0xb6, 0x1b, 0x65, 0x75, 0xe7, 0x8f, 0x7e, 0x09, 0x3d, 0xbe, 0xc1, 0x9f, 0x17, 0xfb,
	0x1d, 0x0d, 0xdf, 0xf9, 0xbf, 0x63, 0x00, 0x6e, 0x69, 0xc3, 0xa4, 0xb4, 0x89, 0x8e, 0x66, 0x81,
	0x30, 0x72, 0xc4, 0xb5, 0x9e, 0xc6, 0xb4, 0x65, 0xd1, 0x09, 0xf4, 0x53, 0x96, 0xa1, 0x6d, 0xbe,
	0x17, 0x0e, 0xac, 0x70, 0x0c, 0x5d, 0x91, 0x60, 0xc8, 0x7b, 0xb1, 0xe6, 0x4f, 0x2c, 0xe1, 0xb4,
	0x8d, 0xd8, 0x25, 0x43, 0xe8, 0x44, 0x2a, 0x49, 0x98, 0x8c, 0x1f, 0x85, 0xe4, 0xd4, 0xb5, 0xb0,
	0xea, 0x4d, 0x99, 0x79, 0xa5, 0x60, 0xf1, 0x11, 0x78, 0x7b, 0x3c, 0x9b, 0x5c, 0xd3, 0xce, 0xcf,
	0x04, 0x16, 0x45, 0x2a, 0x97, 0xc6, 0x8e, 0xf5, 0x2c, 0x24, 0x00, 0xb1, 0x4a, 0x98, 0x90, 0x96,
	0x75, 0x2d, 0xeb, 0x41, 0x4b, 0x0b, 0xc3, 0x31, 0x51, 0xcf, 0xaf, 0xe1, 0xf7, 0x08, 0x7b, 0xbf,
	0xf1, 0x28, 0x37, 0x42, 0xc9, 0xbb, 0xe2, 0x7f, 0xda, 0xb7, 0x49, 0x47, 0x60, 0xef, 0x42, 0x0f,
	0x0b, 0x57, 0xe5, 0x4e, 0xb8, 0x68, 0xd7, 0x76, 0x59, 0x9c, 0x71, 0x60, 0x87, 0x9e, 0xc2, 0xe0,
	0x57, 0x5f, 0x2b, 0x91, 0x42, 0x0a, 0x6f, 0x60, 0x8c, 0xd5, 0x82, 0x77, 0xf1, 0x62, 0xb8, 0x0c,
	0x34, 0xcf, 0x36, 0x3c, 0x2b, 0x9f, 0x38, 0x52, 0xeb, 0x00, 0xcf, 0xaa, 0xb1, 0x4b, 0x38, 0xac,
	0x3e, 0xf5, 0xac, 0x84, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc1, 0x55, 0xe1, 0x2b, 0x02,
	0x00, 0x00,
}
