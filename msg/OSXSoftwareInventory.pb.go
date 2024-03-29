// Code generated by protoc-gen-go.
// source: OSXSoftwareInventory.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OSXSoftwareInventory struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent.(e.g. 58fdc716-1f2a-3819-aaf0-972f4317e83b) This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// full path to binary file (e.g. /usr/libexec/configd)
	ImageFilepath *string `protobuf:"bytes,3,req,name=imageFilepath" json:"imageFilepath,omitempty"`
	// not used by agent anymore
	FileVersion *string `protobuf:"bytes,4,req,name=fileVersion" json:"fileVersion,omitempty"`
	// description of the product (e.g. © Copyright 2009 Apple Inc., all rights reserved.)
	FileDescription *string `protobuf:"bytes,5,req,name=fileDescription" json:"fileDescription,omitempty"`
	// name of the company (e.g. com.apple.FileSyncAgent)
	CompanyName *string `protobuf:"bytes,6,req,name=companyName" json:"companyName,omitempty"`
	// not used by agent anymore
	ProductName *string `protobuf:"bytes,7,req,name=productName" json:"productName,omitempty"`
	// not used by agent anymore
	InternalName *string `protobuf:"bytes,8,req,name=internalName" json:"internalName,omitempty"`
	// e.g. Copyright Apple Inc., 2008
	LegalCopyright *string `protobuf:"bytes,9,req,name=legalCopyright" json:"legalCopyright,omitempty"`
	// not used by agent anymore
	LegalTrademarks *string `protobuf:"bytes,10,req,name=legalTrademarks" json:"legalTrademarks,omitempty"`
	// e.g. FileSyncAgent
	OriginalFilename *string `protobuf:"bytes,11,req,name=originalFilename" json:"originalFilename,omitempty"`
	// e.g. 1.0
	ProductVersion *string `protobuf:"bytes,12,req,name=productVersion" json:"productVersion,omitempty"`
	// md5 hash of the product (e.g. c371d25d402527d3a06308689d3741b1)
	ImageFileMD5 *string `protobuf:"bytes,13,req,name=imageFileMD5" json:"imageFileMD5,omitempty"`
	// if software is deamon/service (e.g. 1)
	IsDeamon            *int32  `protobuf:"varint,14,req,name=isDeamon" json:"isDeamon,omitempty"`
	SiteId              *string `protobuf:"bytes,15,opt,name=siteId" json:"siteId,omitempty"`
	SourceModuleMessage *string `protobuf:"bytes,16,opt,name=sourceModuleMessage" json:"sourceModuleMessage,omitempty"`
	IsDisabled          *bool   `protobuf:"varint,17,opt,name=isDisabled" json:"isDisabled,omitempty"`
	// Length of the file in bytes
	FileLength       *int64  `protobuf:"varint,18,opt,name=fileLength" json:"fileLength,omitempty"`
	Uuid             *string `protobuf:"bytes,19,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OSXSoftwareInventory) Reset()                    { *m = OSXSoftwareInventory{} }
func (m *OSXSoftwareInventory) String() string            { return proto.CompactTextString(m) }
func (*OSXSoftwareInventory) ProtoMessage()               {}
func (*OSXSoftwareInventory) Descriptor() ([]byte, []int) { return fileDescriptor55, []int{0} }

func (m *OSXSoftwareInventory) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *OSXSoftwareInventory) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *OSXSoftwareInventory) GetImageFilepath() string {
	if m != nil && m.ImageFilepath != nil {
		return *m.ImageFilepath
	}
	return ""
}

func (m *OSXSoftwareInventory) GetFileVersion() string {
	if m != nil && m.FileVersion != nil {
		return *m.FileVersion
	}
	return ""
}

func (m *OSXSoftwareInventory) GetFileDescription() string {
	if m != nil && m.FileDescription != nil {
		return *m.FileDescription
	}
	return ""
}

func (m *OSXSoftwareInventory) GetCompanyName() string {
	if m != nil && m.CompanyName != nil {
		return *m.CompanyName
	}
	return ""
}

func (m *OSXSoftwareInventory) GetProductName() string {
	if m != nil && m.ProductName != nil {
		return *m.ProductName
	}
	return ""
}

func (m *OSXSoftwareInventory) GetInternalName() string {
	if m != nil && m.InternalName != nil {
		return *m.InternalName
	}
	return ""
}

func (m *OSXSoftwareInventory) GetLegalCopyright() string {
	if m != nil && m.LegalCopyright != nil {
		return *m.LegalCopyright
	}
	return ""
}

func (m *OSXSoftwareInventory) GetLegalTrademarks() string {
	if m != nil && m.LegalTrademarks != nil {
		return *m.LegalTrademarks
	}
	return ""
}

func (m *OSXSoftwareInventory) GetOriginalFilename() string {
	if m != nil && m.OriginalFilename != nil {
		return *m.OriginalFilename
	}
	return ""
}

func (m *OSXSoftwareInventory) GetProductVersion() string {
	if m != nil && m.ProductVersion != nil {
		return *m.ProductVersion
	}
	return ""
}

func (m *OSXSoftwareInventory) GetImageFileMD5() string {
	if m != nil && m.ImageFileMD5 != nil {
		return *m.ImageFileMD5
	}
	return ""
}

func (m *OSXSoftwareInventory) GetIsDeamon() int32 {
	if m != nil && m.IsDeamon != nil {
		return *m.IsDeamon
	}
	return 0
}

func (m *OSXSoftwareInventory) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *OSXSoftwareInventory) GetSourceModuleMessage() string {
	if m != nil && m.SourceModuleMessage != nil {
		return *m.SourceModuleMessage
	}
	return ""
}

func (m *OSXSoftwareInventory) GetIsDisabled() bool {
	if m != nil && m.IsDisabled != nil {
		return *m.IsDisabled
	}
	return false
}

func (m *OSXSoftwareInventory) GetFileLength() int64 {
	if m != nil && m.FileLength != nil {
		return *m.FileLength
	}
	return 0
}

func (m *OSXSoftwareInventory) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*OSXSoftwareInventory)(nil), "OSXSoftwareInventory")
}

var fileDescriptor55 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xb5, 0x9b, 0xdd, 0x25, 0x99, 0xdd, 0x76, 0x8b, 0xbb, 0x80, 0xb5, 0xbd, 0x54, 0x7b,
	0xe2, 0x94, 0x1b, 0x0f, 0x40, 0x88, 0x40, 0x95, 0x58, 0x40, 0x6a, 0x41, 0x5c, 0x4d, 0x32, 0x4d,
	0x2d, 0x62, 0x3b, 0xb2, 0x9d, 0xa2, 0x70, 0xe4, 0x9d, 0x78, 0x11, 0x9e, 0x88, 0xb1, 0xdb, 0xc0,
	0xa5, 0xd7, 0x6f, 0xfe, 0xf9, 0xe7, 0xf7, 0x6f, 0xb8, 0xff, 0xb8, 0xfe, 0xba, 0x36, 0x5b, 0xff,
	0x43, 0x58, 0x5c, 0xe9, 0x3d, 0x6a, 0x6f, 0xec, 0x90, 0x77, 0xd6, 0x78, 0x73, 0xcf, 0x3e, 0x19,
	0xe7, 0x1b, 0x8b, 0x6e, 0x33, 0x74, 0x78, 0x60, 0x0f, 0x7f, 0x12, 0xb8, 0x3b, 0xb5, 0xc2, 0x96,
	0x90, 0x79, 0xa9, 0x70, 0xed, 0x85, 0xea, 0xf8, 0xd9, 0xf2, 0xfc, 0x65, 0x52, 0x4c, 0x7e, 0xfd,
	0xe6, 0x11, 0xba, 0x00, 0xd9, 0x02, 0x32, 0xd1, 0x90, 0xfa, 0xdd, 0xe7, 0x55, 0xc9, 0xcf, 0x49,
	0x91, 0x15, 0x29, 0x29, 0x2e, 0xfa, 0x5e, 0xd6, 0xec, 0x19, 0x4c, 0xa4, 0xa2, 0xf1, 0x5b, 0xd9,
	0x62, 0x27, 0xfc, 0x8e, 0x27, 0x41, 0xc0, 0xe6, 0x70, 0xbd, 0x25, 0xf2, 0x05, 0xad, 0x93, 0x46,
	0xf3, 0x8b, 0x08, 0x5f, 0xc0, 0x6d, 0x80, 0x25, 0xba, 0xca, 0xca, 0xce, 0x87, 0xc1, 0xe5, 0xa8,
	0xae, 0x8c, 0xea, 0x84, 0x1e, 0x3e, 0x08, 0x85, 0xfc, 0x6a, 0x84, 0x14, 0xbd, 0xee, 0x2b, 0x1f,
	0xe1, 0x93, 0x08, 0xef, 0xe0, 0x46, 0x6a, 0x8f, 0x56, 0x8b, 0x36, 0xd2, 0x34, 0xd2, 0xe7, 0x30,
	0x6d, 0xb1, 0x11, 0xed, 0x1b, 0xd3, 0x0d, 0x56, 0x36, 0x3b, 0xcf, 0xb3, 0xf1, 0x60, 0xe4, 0x1b,
	0x2b, 0x6a, 0x54, 0xc2, 0x7e, 0x77, 0x1c, 0xe2, 0x80, 0xc3, 0xcc, 0x90, 0x50, 0x92, 0x4d, 0x08,
	0xae, 0x83, 0xd5, 0xf5, 0x68, 0x75, 0xbc, 0x3a, 0x66, 0xbf, 0xf9, 0x77, 0x78, 0x7c, 0xe7, 0x63,
	0xf9, 0x8a, 0x4f, 0x22, 0x9d, 0x41, 0x2a, 0x5d, 0x89, 0x42, 0x91, 0x6e, 0x4a, 0xe4, 0x92, 0x4d,
	0xe1, 0xca, 0x49, 0x8f, 0xab, 0x9a, 0xdf, 0x2e, 0xcf, 0x48, 0xb1, 0x80, 0xb9, 0x33, 0xbd, 0xad,
	0xf0, 0x91, 0x4c, 0x69, 0x15, 0x9d, 0x23, 0x17, 0x3e, 0x8b, 0x43, 0x06, 0x40, 0xeb, 0xd2, 0x89,
	0x6f, 0x2d, 0xd6, 0xfc, 0x29, 0xb1, 0x34, 0xb0, 0x50, 0xd2, 0x7b, 0xd4, 0x0d, 0xb5, 0xc9, 0x88,
	0x25, 0x14, 0x2a, 0x96, 0xcd, 0xe7, 0x61, 0xeb, 0x7f, 0xf9, 0xc5, 0x6b, 0x78, 0xa0, 0xde, 0xf2,
	0x9f, 0x72, 0xeb, 0x51, 0xe7, 0x0e, 0xed, 0x1e, 0xed, 0xe1, 0xbf, 0x2b, 0xd3, 0xe6, 0xea, 0x70,
	0xab, 0x58, 0x9c, 0xfa, 0xf7, 0x63, 0x90, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x57, 0xd9,
	0x61, 0x48, 0x02, 0x00, 0x00,
}
