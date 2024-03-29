// Code generated by protoc-gen-go.
// source: OSXUpdateAvailable.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXUpdateAvailable struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// This is the data returned by IUpdateSearcher using search "IsInstalled=0"
	Description         []string `protobuf:"bytes,3,rep,name=description" json:"description,omitempty"`
	UpdateID            []string `protobuf:"bytes,4,rep,name=updateID" json:"updateID,omitempty"`
	Revision            []int32  `protobuf:"varint,5,rep,name=revision" json:"revision,omitempty"`
	IsDownloaded        []bool   `protobuf:"varint,6,rep,name=isDownloaded" json:"isDownloaded,omitempty"`
	IsHidden            []bool   `protobuf:"varint,7,rep,name=isHidden" json:"isHidden,omitempty"`
	IsInstalled         []bool   `protobuf:"varint,8,rep,name=isInstalled" json:"isInstalled,omitempty"`
	IsMandatory         []bool   `protobuf:"varint,9,rep,name=isMandatory" json:"isMandatory,omitempty"`
	IsUninstallable     []bool   `protobuf:"varint,10,rep,name=isUninstallable" json:"isUninstallable,omitempty"`
	EulaText            []string `protobuf:"bytes,11,rep,name=eulaText" json:"eulaText,omitempty"`
	Severity            []string `protobuf:"bytes,12,rep,name=severity" json:"severity,omitempty"`
	ReleaseNotes        []string `protobuf:"bytes,13,rep,name=releaseNotes" json:"releaseNotes,omitempty"`
	SecurityBulletinID  []string `protobuf:"bytes,14,rep,name=securityBulletinID" json:"securityBulletinID,omitempty"`
	SupersededUpdateIDs []string `protobuf:"bytes,15,rep,name=supersededUpdateIDs" json:"supersededUpdateIDs,omitempty"`
	SupportUrl          []string `protobuf:"bytes,16,rep,name=supportUrl" json:"supportUrl,omitempty"`
	Title               []string `protobuf:"bytes,17,rep,name=title" json:"title,omitempty"`
	KbArticleID         []int32  `protobuf:"varint,18,rep,name=kbArticleID" json:"kbArticleID,omitempty"`
	UninstallationNotes []string `protobuf:"bytes,19,rep,name=uninstallationNotes" json:"uninstallationNotes,omitempty"`
	SizeOfUpdate        []int64  `protobuf:"varint,21,rep,name=sizeOfUpdate" json:"sizeOfUpdate,omitempty"`
	SiteId              *string  `protobuf:"bytes,20,opt,name=siteId" json:"siteId,omitempty"`
	Uuid                *string  `protobuf:"bytes,22,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized    []byte   `json:"-"`
}

func (m *OSXUpdateAvailable) Reset()                    { *m = OSXUpdateAvailable{} }
func (m *OSXUpdateAvailable) String() string            { return proto.CompactTextString(m) }
func (*OSXUpdateAvailable) ProtoMessage()               {}
func (*OSXUpdateAvailable) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{0} }

func (m *OSXUpdateAvailable) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXUpdateAvailable) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXUpdateAvailable) GetDescription() []string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *OSXUpdateAvailable) GetUpdateID() []string {
	if m != nil {
		return m.UpdateID
	}
	return nil
}

func (m *OSXUpdateAvailable) GetRevision() []int32 {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *OSXUpdateAvailable) GetIsDownloaded() []bool {
	if m != nil {
		return m.IsDownloaded
	}
	return nil
}

func (m *OSXUpdateAvailable) GetIsHidden() []bool {
	if m != nil {
		return m.IsHidden
	}
	return nil
}

func (m *OSXUpdateAvailable) GetIsInstalled() []bool {
	if m != nil {
		return m.IsInstalled
	}
	return nil
}

func (m *OSXUpdateAvailable) GetIsMandatory() []bool {
	if m != nil {
		return m.IsMandatory
	}
	return nil
}

func (m *OSXUpdateAvailable) GetIsUninstallable() []bool {
	if m != nil {
		return m.IsUninstallable
	}
	return nil
}

func (m *OSXUpdateAvailable) GetEulaText() []string {
	if m != nil {
		return m.EulaText
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSeverity() []string {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *OSXUpdateAvailable) GetReleaseNotes() []string {
	if m != nil {
		return m.ReleaseNotes
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSecurityBulletinID() []string {
	if m != nil {
		return m.SecurityBulletinID
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSupersededUpdateIDs() []string {
	if m != nil {
		return m.SupersededUpdateIDs
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSupportUrl() []string {
	if m != nil {
		return m.SupportUrl
	}
	return nil
}

func (m *OSXUpdateAvailable) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *OSXUpdateAvailable) GetKbArticleID() []int32 {
	if m != nil {
		return m.KbArticleID
	}
	return nil
}

func (m *OSXUpdateAvailable) GetUninstallationNotes() []string {
	if m != nil {
		return m.UninstallationNotes
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSizeOfUpdate() []int64 {
	if m != nil {
		return m.SizeOfUpdate
	}
	return nil
}

func (m *OSXUpdateAvailable) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXUpdateAvailable) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXUpdateAvailable)(nil), "OSXUpdateAvailable")
}

var fileDescriptor61 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x86, 0x05, 0x26, 0x34, 0x1e, 0x08, 0xd0, 0x0d, 0xa5, 0x53, 0xb8, 0x44, 0x9c, 0x38, 0xe5,
	0x15, 0x2a, 0x2c, 0xa4, 0x96, 0x03, 0xa5, 0x12, 0x58, 0xea, 0xd5, 0xd8, 0x03, 0x1a, 0x75, 0xb3,
	0x6b, 0xed, 0xac, 0xd3, 0x86, 0x63, 0xdf, 0xa9, 0x8f, 0x57, 0xa9, 0xe3, 0x35, 0x81, 0x43, 0x8e,
	0xfb, 0xcd, 0xbf, 0x3b, 0xff, 0x3f, 0xb3, 0x80, 0xb7, 0x77, 0x3f, 0xca, 0xb6, 0xa9, 0x22, 0x5d,
	0x2e, 0x2b, 0xb6, 0xd5, 0x83, 0xa5, 0x79, 0x1b, 0x7c, 0xf4, 0xa7, 0xe6, 0xbb, 0x97, 0xf8, 0x14,
	0x48, 0xee, 0x57, 0xed, 0x0b, 0x3b, 0xff, 0x97, 0x81, 0xd9, 0xbc, 0x60, 0x66, 0x90, 0x47, 0x5e,
	0xd0, 0x5d, 0xac, 0x16, 0x2d, 0x6e, 0xcd, 0xb6, 0x2f, 0xb2, 0x62, 0xf2, 0xe7, 0x2f, 0x26, 0x28,
	0x3d, 0x34, 0x67, 0x90, 0x57, 0x4f, 0xe4, 0xe2, 0x97, 0xf2, 0xfa, 0x0a, 0xb7, 0x55, 0x91, 0x17,
	0x63, 0x55, 0xec, 0x74, 0x1d, 0x37, 0x66, 0x0a, 0x7b, 0x0d, 0x49, 0x1d, 0xb8, 0x8d, 0xec, 0x1d,
	0x66, 0xb3, 0xec, 0x22, 0x37, 0x47, 0x30, 0xee, 0x52, 0x1b, 0xbd, 0xb0, 0xb3, 0x26, 0x81, 0x96,
	0x2c, 0xbd, 0x66, 0xa4, 0x64, 0x64, 0x8e, 0x61, 0x9f, 0xe5, 0xca, 0xff, 0x72, 0xd6, 0x57, 0x0d,
	0x35, 0xb8, 0xab, 0x74, 0xdc, 0xeb, 0x58, 0xbe, 0x72, 0xd3, 0x90, 0xc3, 0x77, 0x89, 0x68, 0x03,
	0x96, 0x6b, 0xa7, 0x56, 0xac, 0x55, 0xd9, 0xf8, 0x0d, 0xde, 0x54, 0x4e, 0x7b, 0xf8, 0xb0, 0xc2,
	0x3c, 0xc1, 0x8f, 0x70, 0xc8, 0x52, 0x3a, 0x1e, 0xb4, 0x7d, 0x38, 0x84, 0xf5, 0xa3, 0xd4, 0xd9,
	0xea, 0x9e, 0x7e, 0x47, 0xdc, 0x5b, 0xdb, 0x11, 0x5a, 0x52, 0xe0, 0xb8, 0xc2, 0xfd, 0x44, 0xd4,
	0x4e, 0x20, 0x4b, 0x95, 0xd0, 0x37, 0x1f, 0x49, 0x70, 0x92, 0xe8, 0x29, 0x18, 0xa1, 0xba, 0xeb,
	0x75, 0x45, 0xa7, 0xfd, 0x23, 0x3b, 0x8d, 0x74, 0x90, 0x6a, 0x67, 0x30, 0x95, 0xae, 0xa5, 0x20,
	0xa4, 0xf6, 0xcb, 0x97, 0xb8, 0x82, 0x87, 0xa9, 0x68, 0x00, 0xb4, 0xd8, 0xfa, 0x10, 0xcb, 0x60,
	0xf1, 0x28, 0xb1, 0x09, 0x8c, 0x22, 0x47, 0x75, 0xf5, 0x3e, 0x1d, 0x35, 0xc3, 0xcf, 0x87, 0xcb,
	0x10, 0xb9, 0xb6, 0xfd, 0x9c, 0x4c, 0x9a, 0x8a, 0x3e, 0xda, 0xbd, 0x26, 0xe8, 0x27, 0x3a, 0xb8,
	0x99, 0xae, 0x3d, 0x0a, 0x3f, 0xd3, 0xed, 0xe3, 0xd0, 0x0d, 0x3f, 0x28, 0xcd, 0xcc, 0x01, 0xec,
	0x0a, 0x6b, 0xef, 0x06, 0x8f, 0x67, 0x5b, 0xaa, 0x3a, 0x81, 0xb4, 0x19, 0x3c, 0xe9, 0x4f, 0x6f,
	0x9b, 0x2a, 0x3e, 0xc3, 0x79, 0xed, 0x17, 0xf3, 0x67, 0x7e, 0x8c, 0xe4, 0xe6, 0x42, 0x41, 0xf3,
	0x0f, 0x5f, 0xa3, 0xf6, 0x76, 0xae, 0xcb, 0x16, 0xdd, 0x72, 0xf1, 0x69, 0xf3, 0x8b, 0xdc, 0x0c,
	0xa5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0x32, 0x63, 0x85, 0x6f, 0x02, 0x00, 0x00,
}
