// Code generated by protoc-gen-go.
// source: FileContents.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Send File to server
type FileContents struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Fully qualified path to file (also, what was passed in from AgentCommand)
	ImageFilePath *string `protobuf:"bytes,3,req,name=imageFilePath" json:"imageFilePath,omitempty"`
	// File Contents - can be zero length
	FileContents []byte `protobuf:"bytes,4,req,name=fileContents" json:"fileContents,omitempty"`
	// File Creation Timestamp
	CreatedTimeStamp *int64 `protobuf:"varint,5,opt,name=CreatedTimeStamp" json:"CreatedTimeStamp,omitempty"`
	// Last Modified File Timestamp
	ModifiedTimeStamp *int64 `protobuf:"varint,6,opt,name=ModifiedTimeStamp" json:"ModifiedTimeStamp,omitempty"`
	// MD5 of the File Contents to double check file validity and pattern matching
	ImageFileMD5 *string `protobuf:"bytes,7,req,name=imageFileMD5" json:"imageFileMD5,omitempty"`
	// The Command Identifer from the AgentCommand message
	CommandIdentifier *uint64 `protobuf:"varint,8,req,name=commandIdentifier" json:"commandIdentifier,omitempty"`
	// MD5 of the server name we are reporting to
	SiteId *string `protobuf:"bytes,9,opt,name=siteId" json:"siteId,omitempty"`
	Uuid   *string `protobuf:"bytes,10,opt,name=uuid" json:"uuid,omitempty"`
	// Added to count the number of times this message has been ATTEMPTED to transmit but error has occured and
	// message has been re-enqueued
	RetryCount       *int32 `protobuf:"varint,11,opt,name=retryCount,def=0" json:"retryCount,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FileContents) Reset()                    { *m = FileContents{} }
func (m *FileContents) String() string            { return proto1.CompactTextString(m) }
func (*FileContents) ProtoMessage()               {}
func (*FileContents) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

const Default_FileContents_RetryCount int32 = 0

func (m *FileContents) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *FileContents) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *FileContents) GetImageFilePath() string {
	if m != nil && m.ImageFilePath != nil {
		return *m.ImageFilePath
	}
	return ""
}

func (m *FileContents) GetFileContents() []byte {
	if m != nil {
		return m.FileContents
	}
	return nil
}

func (m *FileContents) GetCreatedTimeStamp() int64 {
	if m != nil && m.CreatedTimeStamp != nil {
		return *m.CreatedTimeStamp
	}
	return 0
}

func (m *FileContents) GetModifiedTimeStamp() int64 {
	if m != nil && m.ModifiedTimeStamp != nil {
		return *m.ModifiedTimeStamp
	}
	return 0
}

func (m *FileContents) GetImageFileMD5() string {
	if m != nil && m.ImageFileMD5 != nil {
		return *m.ImageFileMD5
	}
	return ""
}

func (m *FileContents) GetCommandIdentifier() uint64 {
	if m != nil && m.CommandIdentifier != nil {
		return *m.CommandIdentifier
	}
	return 0
}

func (m *FileContents) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *FileContents) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *FileContents) GetRetryCount() int32 {
	if m != nil && m.RetryCount != nil {
		return *m.RetryCount
	}
	return Default_FileContents_RetryCount
}

func init() {
	proto1.RegisterType((*FileContents)(nil), "FileContents")
}

var fileDescriptor25 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x50, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x9f, 0x34, 0x4b, 0x8a, 0xc0, 0x3c, 0xb4, 0x3c, 0x0e, 0x51, 0x4f, 0xe5, 0x12, 0x55,
	0x45, 0x1c, 0xe0, 0x98, 0x56, 0xa0, 0x1e, 0x2a, 0x55, 0xa2, 0x7c, 0x80, 0xa9, 0xdd, 0x62, 0xa9,
	0xb1, 0x2b, 0xdb, 0xa9, 0x14, 0x8e, 0xfc, 0x13, 0x47, 0xfe, 0x8d, 0x4d, 0x22, 0x55, 0xc1, 0x37,
	0xcf, 0xcc, 0xee, 0xec, 0x0c, 0xb0, 0x17, 0xb5, 0x95, 0x13, 0xa3, 0xbd, 0xd4, 0xde, 0xc5, 0x3b,
	0x6b, 0xbc, 0xb9, 0x61, 0x0b, 0xe3, 0xfc, 0xc6, 0x4a, 0xb7, 0xcc, 0x77, 0xb2, 0xc2, 0x06, 0xbf,
	0x4d, 0x08, 0xeb, 0x52, 0x16, 0x41, 0xe0, 0x55, 0x2a, 0xdf, 0x3c, 0x4f, 0x77, 0xd8, 0x88, 0x9a,
	0xc3, 0x56, 0xd2, 0xff, 0xfe, 0xc1, 0x12, 0x74, 0x05, 0xc8, 0x6e, 0x21, 0xe0, 0x1b, 0xd2, 0xbe,
	0xbe, 0xcf, 0xa6, 0xd8, 0x24, 0x45, 0x90, 0xf4, 0x48, 0xd1, 0xce, 0x32, 0x25, 0xd8, 0x25, 0xf4,
	0x55, 0x4a, 0x74, 0xb1, 0x73, 0xc1, 0xfd, 0x27, 0xb6, 0x0a, 0x01, 0x1b, 0x43, 0xb8, 0xae, 0xb9,
	0x60, 0x9b, 0xd0, 0x30, 0xb9, 0xa3, 0x31, 0xdc, 0x1a, 0xbd, 0x89, 0xf6, 0xdc, 0x7e, 0x28, 0xcd,
	0x6d, 0x3e, 0x7c, 0x18, 0x8f, 0xca, 0x77, 0xcf, 0x10, 0x4e, 0x27, 0x56, 0x72, 0x2f, 0xc5, 0xf2,
	0x70, 0x50, 0x27, 0x6a, 0x0c, 0x5b, 0xec, 0x1a, 0xce, 0xe6, 0x46, 0xa8, 0xb5, 0xaa, 0x53, 0xdd,
	0x92, 0xba, 0x80, 0xf0, 0xe0, 0x3f, 0x9f, 0x3e, 0xe2, 0x51, 0x69, 0x4f, 0x03, 0x2b, 0x93, 0xa6,
	0x5c, 0x8b, 0x99, 0x20, 0xff, 0x62, 0xd2, 0x62, 0x8f, 0xa8, 0x36, 0x3b, 0x81, 0xae, 0x53, 0x5e,
	0xce, 0x04, 0x06, 0xb4, 0x20, 0x60, 0x57, 0x50, 0x06, 0x41, 0x28, 0x7e, 0xff, 0x82, 0x81, 0x95,
	0xde, 0xe6, 0x13, 0x93, 0x69, 0x8f, 0xc7, 0xc4, 0x76, 0x9e, 0x1b, 0xa3, 0xe4, 0x09, 0x06, 0xb4,
	0x39, 0xfe, 0x52, 0x6b, 0x0a, 0x16, 0x3b, 0x69, 0xf7, 0xd2, 0x56, 0xd5, 0xae, 0xcc, 0x36, 0xa6,
	0xca, 0x1c, 0x1d, 0x93, 0x9c, 0xd7, 0x2b, 0x9e, 0x57, 0xe0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x43, 0xe0, 0x24, 0x85, 0xa3, 0x01, 0x00, 0x00,
}
