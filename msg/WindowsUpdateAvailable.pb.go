// Code generated by protoc-gen-go.
// source: WindowsUpdateAvailable.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WindowsUpdateAvailable struct {
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
	// This UUID is to group WindowsUpdate, WindowsUpdateAvailable, WindowsUpdateServers and WindowsUpdateSettings message
	WuaScanUUID      *string `protobuf:"bytes,23,opt,name=wuaScanUUID" json:"wuaScanUUID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WindowsUpdateAvailable) Reset()                    { *m = WindowsUpdateAvailable{} }
func (m *WindowsUpdateAvailable) String() string            { return proto.CompactTextString(m) }
func (*WindowsUpdateAvailable) ProtoMessage()               {}
func (*WindowsUpdateAvailable) Descriptor() ([]byte, []int) { return fileDescriptor90, []int{0} }

func (m *WindowsUpdateAvailable) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *WindowsUpdateAvailable) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *WindowsUpdateAvailable) GetDescription() []string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetUpdateID() []string {
	if m != nil {
		return m.UpdateID
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetRevision() []int32 {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetIsDownloaded() []bool {
	if m != nil {
		return m.IsDownloaded
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetIsHidden() []bool {
	if m != nil {
		return m.IsHidden
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetIsInstalled() []bool {
	if m != nil {
		return m.IsInstalled
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetIsMandatory() []bool {
	if m != nil {
		return m.IsMandatory
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetIsUninstallable() []bool {
	if m != nil {
		return m.IsUninstallable
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetEulaText() []string {
	if m != nil {
		return m.EulaText
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSeverity() []string {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetReleaseNotes() []string {
	if m != nil {
		return m.ReleaseNotes
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSecurityBulletinID() []string {
	if m != nil {
		return m.SecurityBulletinID
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSupersededUpdateIDs() []string {
	if m != nil {
		return m.SupersededUpdateIDs
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSupportUrl() []string {
	if m != nil {
		return m.SupportUrl
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetKbArticleID() []int32 {
	if m != nil {
		return m.KbArticleID
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetUninstallationNotes() []string {
	if m != nil {
		return m.UninstallationNotes
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSizeOfUpdate() []int64 {
	if m != nil {
		return m.SizeOfUpdate
	}
	return nil
}

func (m *WindowsUpdateAvailable) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *WindowsUpdateAvailable) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *WindowsUpdateAvailable) GetWuaScanUUID() string {
	if m != nil && m.WuaScanUUID != nil {
		return *m.WuaScanUUID
	}
	return ""
}

func init() {
	proto.RegisterType((*WindowsUpdateAvailable)(nil), "WindowsUpdateAvailable")
}

var fileDescriptor90 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x4e, 0x1b, 0x31,
	0x10, 0x86, 0x05, 0x49, 0x68, 0x76, 0x42, 0x80, 0x3a, 0x34, 0x58, 0x50, 0xa4, 0x88, 0x13, 0xa7,
	0xbc, 0x03, 0x5b, 0xa4, 0x96, 0x03, 0x6d, 0x25, 0x58, 0xf5, 0x6c, 0x76, 0x07, 0x34, 0xaa, 0x63,
	0xaf, 0x3c, 0xde, 0xa4, 0xe1, 0xd8, 0x27, 0xe9, 0x4b, 0xf4, 0xfd, 0x3a, 0xeb, 0x25, 0x6d, 0x91,
	0x72, 0xf4, 0x37, 0xbf, 0xfd, 0xff, 0x33, 0x1e, 0x78, 0xff, 0x8d, 0x5c, 0xe5, 0x57, 0x5c, 0xd4,
	0x95, 0x89, 0x78, 0xb5, 0x34, 0x64, 0xcd, 0x83, 0xc5, 0x79, 0x1d, 0x7c, 0xf4, 0xa7, 0xea, 0xab,
	0xe7, 0xf8, 0x14, 0x90, 0xef, 0xd7, 0xf5, 0x0b, 0xbb, 0xf8, 0xd5, 0x87, 0xe9, 0xf6, 0x4b, 0x6a,
	0x06, 0x59, 0xa4, 0x05, 0xde, 0x45, 0xb3, 0xa8, 0xf5, 0xce, 0x6c, 0xf7, 0xb2, 0x97, 0x8f, 0x7f,
	0xfe, 0xd6, 0x09, 0x72, 0x0b, 0xd5, 0x19, 0x64, 0xe6, 0x09, 0x5d, 0xfc, 0x58, 0xdc, 0x5c, 0xeb,
	0x5d, 0x51, 0x64, 0xf9, 0x50, 0x14, 0xfd, 0xa6, 0xa1, 0x4a, 0x4d, 0x60, 0x54, 0x21, 0x97, 0x81,
	0xea, 0x48, 0xde, 0xe9, 0xde, 0xac, 0x77, 0x99, 0xa9, 0x23, 0x18, 0x36, 0xc9, 0x46, 0x2e, 0xf4,
	0x37, 0x24, 0xe0, 0x92, 0xb8, 0xd5, 0x0c, 0x84, 0x0c, 0xd4, 0x31, 0xec, 0x13, 0x5f, 0xfb, 0x95,
	0xb3, 0xde, 0x54, 0x58, 0xe9, 0x3d, 0xa1, 0xc3, 0x56, 0x47, 0xfc, 0x89, 0xaa, 0x0a, 0x9d, 0x7e,
	0x93, 0x88, 0x18, 0x10, 0xdf, 0x38, 0x89, 0x62, 0xad, 0xc8, 0x86, 0xff, 0xe0, 0xad, 0x71, 0xe2,
	0xe1, 0xc3, 0x5a, 0x67, 0x09, 0x9e, 0xc0, 0x21, 0x71, 0xe1, 0xa8, 0xd3, 0xb6, 0xcd, 0x69, 0xd8,
	0x3c, 0x8a, 0x8d, 0x35, 0xf7, 0xf8, 0x23, 0xea, 0xd1, 0x26, 0x0e, 0xe3, 0x12, 0x03, 0xc5, 0xb5,
	0xde, 0x4f, 0x44, 0xe2, 0x04, 0xb4, 0x68, 0x18, 0x3f, 0xfb, 0x88, 0xac, 0xc7, 0x89, 0x9e, 0x82,
	0x62, 0x2c, 0x9b, 0x56, 0x97, 0x37, 0xe2, 0x1f, 0xc9, 0x49, 0x4b, 0x07, 0xa9, 0x76, 0x06, 0x13,
	0x6e, 0x6a, 0x0c, 0x8c, 0x12, 0xbf, 0x78, 0x69, 0x97, 0xf5, 0x61, 0x2a, 0x2a, 0x00, 0x29, 0xd6,
	0x3e, 0xc4, 0x22, 0x58, 0x7d, 0x94, 0xd8, 0x18, 0x06, 0x91, 0xa2, 0xa4, 0x7a, 0x9b, 0x8e, 0xd2,
	0xc3, 0xf7, 0x87, 0xab, 0x10, 0xa9, 0xb4, 0xed, 0x9c, 0x54, 0x9a, 0x8a, 0x3c, 0xda, 0xfc, 0xed,
	0xa0, 0x9d, 0x68, 0x97, 0x66, 0xb2, 0xc9, 0xc8, 0xf4, 0x8c, 0x5f, 0x1e, 0x3b, 0x37, 0xfd, 0x4e,
	0x68, 0x4f, 0x1d, 0xc0, 0x1e, 0x93, 0x78, 0x57, 0xfa, 0x78, 0xb6, 0x23, 0xaa, 0x29, 0xa4, 0x9f,
	0xd1, 0xd3, 0xf6, 0xf4, 0xdf, 0x4f, 0x9d, 0xc3, 0x68, 0xd5, 0x98, 0xbb, 0xd2, 0xb8, 0xa2, 0xfd,
	0xc8, 0x93, 0xd7, 0xe5, 0xfc, 0x03, 0x5c, 0x94, 0x7e, 0x31, 0x7f, 0xa6, 0xc7, 0x88, 0x6e, 0xce,
	0x18, 0x64, 0x3c, 0xdd, 0xf6, 0x94, 0xde, 0xce, 0x65, 0x17, 0x58, 0x96, 0x20, 0x3f, 0xdf, 0xbe,
	0x45, 0xb7, 0x5d, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xcc, 0xe8, 0xb4, 0x9a, 0x02,
	0x00, 0x00,
}
