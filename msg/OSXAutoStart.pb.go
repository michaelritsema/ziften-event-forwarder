// Code generated by protoc-gen-go.
// source: OSXAutoStart.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXAutoStart struct {
	// The UTC time (as known to the client computer) when the agent status message was generated
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// e.g. com.apple.DirectoryServices
	Label *string `protobuf:"bytes,3,req,name=label" json:"label,omitempty"`
	// System/Library/LaunchDaemons/com.apple.DirectoryServices.plist
	Autostartlocation *string `protobuf:"bytes,4,req,name=autostartlocation" json:"autostartlocation,omitempty"`
	// This optional key is used as a hint to launchctl(1) that it should not submit this job to launchd when
	//     loading a job or jobs. The value of this key does NOT reflect the current state of the job on the run-
	//     ning system. If you wish to know whether a job is loaded in launchd, reading this key from a configura-
	//     tion file yourself is not a sufficient test. You should query launchd for the presence of the job using
	//     the launchctl(1) list subcommand or use the ServiceManagement framework's SMJobCopyDictionary() method.
	// e.g. 1
	Disabled *string `protobuf:"bytes,5,req,name=disabled" json:"disabled,omitempty"`
	// e.g. root
	Procuser *string `protobuf:"bytes,6,req,name=procuser" json:"procuser,omitempty"`
	// e.g. daemon
	Procgroup *string `protobuf:"bytes,7,req,name=procgroup" json:"procgroup,omitempty"`
	// identifies if process launches at startup (e.g. 1)
	Runatload *string `protobuf:"bytes,8,req,name=runatload" json:"runatload,omitempty"`
	// identifier if relaunch process in case of crash etc. (e.g. 1)
	Ondemand *string `protobuf:"bytes,9,req,name=ondemand" json:"ondemand,omitempty"`
	//   This key lets one override the default throttling policy imposed on jobs by launchd.  The value is in
	//     seconds, and by default, jobs will not be spawned more than once every 10 seconds.  The principle
	//     behind this is that jobs should linger around just in case they are needed again in the near future.
	//     This not only reduces the latency of responses, but it encourages developers to amortize the cost of
	//     program invocation
	// e.g. 10
	Throttleinterval *string `protobuf:"bytes,10,req,name=throttleinterval" json:"throttleinterval,omitempty"`
	// e.g. /usr/libexec/atrun
	Path *string `protobuf:"bytes,11,req,name=path" json:"path,omitempty"`
	// e.g. /usr/bin/sandbox-exec -f /usr/share/sandbox/cvmsCompAgent.sb /System/Library/Frameworks/OpenGL.framework/Versions/A/Libraries/cvmsComp_i386 1
	Commandline      *string `protobuf:"bytes,12,req,name=commandline" json:"commandline,omitempty"`
	SiteId           *string `protobuf:"bytes,13,opt,name=siteId" json:"siteId,omitempty"`
	Uuid             *string `protobuf:"bytes,14,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSXAutoStart) Reset()                    { *m = OSXAutoStart{} }
func (m *OSXAutoStart) String() string            { return proto.CompactTextString(m) }
func (*OSXAutoStart) ProtoMessage()               {}
func (*OSXAutoStart) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{0} }

func (m *OSXAutoStart) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXAutoStart) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXAutoStart) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *OSXAutoStart) GetAutostartlocation() string {
	if m != nil && m.Autostartlocation != nil {
		return *m.Autostartlocation
	}
	return ""
}

func (m *OSXAutoStart) GetDisabled() string {
	if m != nil && m.Disabled != nil {
		return *m.Disabled
	}
	return ""
}

func (m *OSXAutoStart) GetProcuser() string {
	if m != nil && m.Procuser != nil {
		return *m.Procuser
	}
	return ""
}

func (m *OSXAutoStart) GetProcgroup() string {
	if m != nil && m.Procgroup != nil {
		return *m.Procgroup
	}
	return ""
}

func (m *OSXAutoStart) GetRunatload() string {
	if m != nil && m.Runatload != nil {
		return *m.Runatload
	}
	return ""
}

func (m *OSXAutoStart) GetOndemand() string {
	if m != nil && m.Ondemand != nil {
		return *m.Ondemand
	}
	return ""
}

func (m *OSXAutoStart) GetThrottleinterval() string {
	if m != nil && m.Throttleinterval != nil {
		return *m.Throttleinterval
	}
	return ""
}

func (m *OSXAutoStart) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *OSXAutoStart) GetCommandline() string {
	if m != nil && m.Commandline != nil {
		return *m.Commandline
	}
	return ""
}

func (m *OSXAutoStart) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXAutoStart) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXAutoStart)(nil), "OSXAutoStart")
}

var fileDescriptor43 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xe9, 0x5f, 0x9b, 0xe9, 0x1f, 0xda, 0x14, 0x64, 0xd4, 0x4b, 0xe9, 0xc9, 0x53, 0xee,
	0x1e, 0x0d, 0x82, 0xf4, 0x20, 0x0a, 0x55, 0xf0, 0xba, 0xcd, 0x8e, 0xed, 0xc2, 0x66, 0x37, 0xec,
	0x4e, 0x0a, 0x7a, 0xf4, 0x3b, 0x09, 0x7e, 0x3c, 0x37, 0x5b, 0x82, 0x3d, 0xbe, 0xdf, 0xcc, 0x3c,
	0xe6, 0x3d, 0x48, 0x9f, 0xb7, 0xef, 0xf7, 0x35, 0xdb, 0x2d, 0x0b, 0xc7, 0x59, 0xe5, 0x2c, 0xdb,
	0xeb, 0xf4, 0xc5, 0x7a, 0xde, 0x3b, 0xf2, 0xaf, 0x9f, 0x15, 0x9d, 0xd8, 0xfa, 0xb7, 0x0b, 0x93,
	0xf3, 0xd5, 0x74, 0x05, 0x09, 0xab, 0x92, 0x82, 0x28, 0x2b, 0xec, 0xac, 0xba, 0xb7, 0xbd, 0x7c,
	0xfa, 0xfd, 0x83, 0x11, 0xfa, 0x06, 0xa6, 0x37, 0x90, 0x88, 0x3d, 0x19, 0x7e, 0x7c, 0xdb, 0x3c,
	0x60, 0x37, 0x6c, 0x24, 0xf9, 0x28, 0x6c, 0xf4, 0xeb, 0x5a, 0xc9, 0x74, 0x0a, 0x03, 0x2d, 0x76,
	0xa4, 0xb1, 0xd7, 0x0c, 0xd2, 0x2b, 0x58, 0x88, 0x60, 0xed, 0x1b, 0x6b, 0x6d, 0x0b, 0xc1, 0xca,
	0x1a, 0xec, 0xc7, 0xd1, 0x1c, 0x46, 0x52, 0x79, 0xb1, 0xd3, 0x24, 0x71, 0xd0, 0x92, 0xf0, 0x54,
	0x51, 0x7b, 0x72, 0x38, 0x8c, 0x64, 0x01, 0x49, 0x43, 0xf6, 0xce, 0xd6, 0x15, 0x5e, 0xb4, 0xc8,
	0xd5, 0x46, 0x04, 0x37, 0x21, 0x71, 0xd4, 0xde, 0x59, 0x23, 0xa9, 0x14, 0x46, 0x62, 0x12, 0x09,
	0xc2, 0x9c, 0x0f, 0x21, 0x1f, 0x6b, 0x52, 0x86, 0xc9, 0x1d, 0x85, 0x46, 0x88, 0x93, 0x09, 0xf4,
	0x2b, 0xc1, 0x07, 0x1c, 0x47, 0xb5, 0x84, 0x71, 0x61, 0xcb, 0xe6, 0x50, 0x2b, 0x43, 0x38, 0x89,
	0x70, 0x06, 0x43, 0xaf, 0x98, 0x36, 0x12, 0xa7, 0xab, 0x4e, 0xd0, 0x97, 0x10, 0xa3, 0xe1, 0xac,
	0x51, 0xff, 0x51, 0xf3, 0x3b, 0x58, 0x87, 0xe3, 0xec, 0x4b, 0x7d, 0x30, 0x99, 0x2c, 0x3c, 0x7d,
	0x24, 0x77, 0x6a, 0xb5, 0xb0, 0x3a, 0x0b, 0x6d, 0xf9, 0x50, 0x53, 0xbe, 0x3c, 0x6f, 0xf7, 0xe9,
	0x04, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x85, 0x8b, 0x9f, 0x9e, 0x01, 0x00, 0x00,
}
