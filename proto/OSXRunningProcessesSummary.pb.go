// Code generated by protoc-gen-go.
// source: OSXRunningProcessesSummary.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This message is used to send a summary of running processes within a specified period of time to the server.
// Processes are grouped by the combination of full imageFilepath, account name and domain (given user).
type OSXRunningProcessesSummary struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// MD5 of the server name we are reporting to
	SiteId *string `protobuf:"bytes,3,opt,name=siteId" json:"siteId,omitempty"`
	// Timestamp for beginning of process monitoring for this message
	StartMonitorPeriod *int64 `protobuf:"varint,4,req,name=startMonitorPeriod" json:"startMonitorPeriod,omitempty"`
	// Timestamp for end of process monitoring for this message
	EndMonitorPeriod *int64 `protobuf:"varint,5,req,name=endMonitorPeriod" json:"endMonitorPeriod,omitempty"`
	// Fully qualified path to file
	ImageFilepath []string `protobuf:"bytes,6,rep,name=imageFilepath" json:"imageFilepath,omitempty"`
	// MD5 of the File Contents to double check file validity and pattern matching
	ImageFileMD5 []string `protobuf:"bytes,7,rep,name=imageFileMD5" json:"imageFileMD5,omitempty"`
	// Timestamp when first instance of process seen OR first launched (whichever is later)
	FirstSeen []int64 `protobuf:"varint,8,rep,name=FirstSeen" json:"FirstSeen,omitempty"`
	// Timestamp when the last instance of the process terminated OR end of time period
	LastSeen []int64 `protobuf:"varint,9,rep,name=LastSeen" json:"LastSeen,omitempty"`
	// How many instances of this process were already running at the start of the time period
	NumInstancesSeenAtStart []int32 `protobuf:"varint,10,rep,name=NumInstancesSeenAtStart" json:"NumInstancesSeenAtStart,omitempty"`
	// How many instances of this process were still running at the end of the time period
	NumInstancesStillRunningAtEnd []int32 `protobuf:"varint,11,rep,name=NumInstancesStillRunningAtEnd" json:"NumInstancesStillRunningAtEnd,omitempty"`
	// How many instances of this process were created during this time period
	NumInstancesCreatedDuring []int32 `protobuf:"varint,12,rep,name=NumInstancesCreatedDuring" json:"NumInstancesCreatedDuring,omitempty"`
	// How many instances of this process were terminated during this time period
	NumInstancesTerminatedDuring []int32 `protobuf:"varint,13,rep,name=NumInstancesTerminatedDuring" json:"NumInstancesTerminatedDuring,omitempty"`
	// User name tied to the process creation
	AccountName []string `protobuf:"bytes,14,rep,name=accountName" json:"accountName,omitempty"`
	// Domain name tied to the process creation
	Domain []string `protobuf:"bytes,15,rep,name=domain" json:"domain,omitempty"`
	Uuid   *string  `protobuf:"bytes,16,opt,name=uuid" json:"uuid,omitempty"`
	// This UUID is to group OSXRunningProcessSummary messages in queries on the server
	RpsUUID          *string `protobuf:"bytes,17,opt,name=rpsUUID" json:"rpsUUID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSXRunningProcessesSummary) Reset()                    { *m = OSXRunningProcessesSummary{} }
func (m *OSXRunningProcessesSummary) String() string            { return proto1.CompactTextString(m) }
func (*OSXRunningProcessesSummary) ProtoMessage()               {}
func (*OSXRunningProcessesSummary) Descriptor() ([]byte, []int) { return fileDescriptor54, []int{0} }

func (m *OSXRunningProcessesSummary) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXRunningProcessesSummary) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXRunningProcessesSummary) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXRunningProcessesSummary) GetStartMonitorPeriod() int64 {
	if m != nil && m.StartMonitorPeriod != nil {
		return *m.StartMonitorPeriod
	}
	return 0
}

func (m *OSXRunningProcessesSummary) GetEndMonitorPeriod() int64 {
	if m != nil && m.EndMonitorPeriod != nil {
		return *m.EndMonitorPeriod
	}
	return 0
}

func (m *OSXRunningProcessesSummary) GetImageFilepath() []string {
	if m != nil {
		return m.ImageFilepath
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetImageFileMD5() []string {
	if m != nil {
		return m.ImageFileMD5
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetFirstSeen() []int64 {
	if m != nil {
		return m.FirstSeen
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetLastSeen() []int64 {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetNumInstancesSeenAtStart() []int32 {
	if m != nil {
		return m.NumInstancesSeenAtStart
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetNumInstancesStillRunningAtEnd() []int32 {
	if m != nil {
		return m.NumInstancesStillRunningAtEnd
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetNumInstancesCreatedDuring() []int32 {
	if m != nil {
		return m.NumInstancesCreatedDuring
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetNumInstancesTerminatedDuring() []int32 {
	if m != nil {
		return m.NumInstancesTerminatedDuring
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetAccountName() []string {
	if m != nil {
		return m.AccountName
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *OSXRunningProcessesSummary) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *OSXRunningProcessesSummary) GetRpsUUID() string {
	if m != nil && m.RpsUUID != nil {
		return *m.RpsUUID
	}
	return ""
}

func init() {
	proto1.RegisterType((*OSXRunningProcessesSummary)(nil), "OSXRunningProcessesSummary")
}

var fileDescriptor54 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0xd2, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x07, 0x70, 0xed, 0xb6, 0xbb, 0xdb, 0xcc, 0x6e, 0x97, 0x25, 0x7c, 0x79, 0x0b, 0x88, 0xb4,
	0x02, 0x51, 0x2e, 0xb9, 0xf1, 0x00, 0x0d, 0xa5, 0x55, 0x25, 0x5a, 0x2a, 0xd2, 0x4a, 0x5c, 0xad,
	0x64, 0x28, 0x96, 0x62, 0x3b, 0xb2, 0x27, 0x48, 0xe5, 0xc8, 0x53, 0xf0, 0x22, 0xbc, 0x1f, 0x8e,
	0x5b, 0xa1, 0x54, 0xb4, 0x47, 0xff, 0xf3, 0xb3, 0xe3, 0x99, 0x31, 0x44, 0x9f, 0xd3, 0xaf, 0x5f,
	0x2a, 0xa5, 0x84, 0xda, 0x2c, 0x8d, 0xce, 0xd0, 0x5a, 0xb4, 0x69, 0x25, 0x25, 0x37, 0xdb, 0xb8,
	0x34, 0x9a, 0x74, 0x2f, 0x5c, 0x6a, 0x4b, 0x1b, 0x83, 0x76, 0xb5, 0x2d, 0x71, 0x97, 0x0d, 0x7e,
	0xb7, 0xa1, 0x77, 0x7a, 0x63, 0x18, 0x41, 0x40, 0x42, 0x62, 0x4a, 0x5c, 0x96, 0xec, 0x2c, 0x3a,
	0x1f, 0xb6, 0x92, 0xee, 0xaf, 0x3f, 0xcc, 0x87, 0xb6, 0x0e, 0xc3, 0xe7, 0x10, 0xf0, 0x0d, 0x2a,
	0x9a, 0xae, 0x67, 0x63, 0x76, 0xee, 0x44, 0x90, 0x74, 0x9c, 0x68, 0x57, 0x95, 0xc8, 0xc3, 0x5b,
	0xb8, 0xb4, 0x82, 0x70, 0x96, 0xb3, 0x56, 0x74, 0x36, 0x0c, 0xc2, 0x77, 0x10, 0xba, 0x5d, 0x86,
	0xe6, 0x5a, 0x09, 0xd2, 0x66, 0x89, 0x46, 0xe8, 0x9c, 0xb5, 0x8f, 0x9d, 0xfb, 0x16, 0xee, 0x50,
	0xe5, 0x87, 0xf0, 0xe2, 0x18, 0x7c, 0x02, 0x5d, 0x21, 0xdd, 0x15, 0x26, 0xa2, 0xc0, 0x92, 0xd3,
	0x77, 0x76, 0x19, 0xb5, 0xdc, 0xaf, 0x1e, 0xc3, 0xcd, 0xbf, 0x78, 0x3e, 0x7e, 0xcf, 0xae, 0x7c,
	0xea, 0xea, 0x99, 0x08, 0x63, 0x29, 0x45, 0x54, 0xac, 0xe3, 0xa2, 0xff, 0x8e, 0x7b, 0x05, 0x9d,
	0x4f, 0x7c, 0x0f, 0x82, 0xe3, 0xe0, 0xd9, 0xa2, 0x92, 0x33, 0xe5, 0x56, 0xca, 0x35, 0xab, 0x86,
	0x23, 0x4a, 0xeb, 0xaa, 0x18, 0x38, 0x7f, 0x11, 0xbe, 0x81, 0x97, 0x07, 0x80, 0x44, 0x51, 0xec,
	0xfb, 0x3b, 0xa2, 0x8f, 0x2a, 0x67, 0xd7, 0x9e, 0xf5, 0xe1, 0xbe, 0xc9, 0x3e, 0x18, 0xe4, 0x84,
	0xf9, 0xb8, 0x32, 0xce, 0xb1, 0x1b, 0x4f, 0x5e, 0xc3, 0x8b, 0x26, 0x59, 0xa1, 0x91, 0x42, 0x35,
	0x54, 0xd7, 0xab, 0x47, 0x70, 0xcd, 0xb3, 0x4c, 0x57, 0x8a, 0x16, 0x5c, 0x22, 0xbb, 0xf5, 0x85,
	0xba, 0xce, 0xe7, 0x5a, 0x72, 0xa1, 0xd8, 0x03, 0xbf, 0x7e, 0x0a, 0x7e, 0x22, 0xec, 0xae, 0x9e,
	0x43, 0x63, 0x42, 0xf7, 0x70, 0x65, 0x4a, 0xbb, 0xae, 0x87, 0xf7, 0xf0, 0xf0, 0x53, 0x32, 0x85,
	0x41, 0xa6, 0x65, 0xfc, 0x53, 0x7c, 0x23, 0x54, 0xb1, 0x45, 0xf3, 0x03, 0xcd, 0xee, 0xd5, 0x64,
	0xba, 0x88, 0x5d, 0x3b, 0xac, 0x6b, 0x6f, 0xd2, 0x3f, 0xfd, 0x7a, 0xe6, 0x3b, 0xf2, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0x82, 0x52, 0xa2, 0x9a, 0x02, 0x00, 0x00,
}
