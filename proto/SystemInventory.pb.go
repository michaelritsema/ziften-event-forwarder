// Code generated by protoc-gen-go.
// source: SystemInventory.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message reporting hardware information of a client
type SystemInventory struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// Windows computer name
	ComputerName *string `protobuf:"bytes,3,req,name=computerName" json:"computerName,omitempty"`
	// Simple minded CPU benchmark. It will contains the time required in milliseconds to run a simple CPU-intensive piece of code.
	CpuScore *float64 `protobuf:"fixed64,4,req,name=cpuScore" json:"cpuScore,omitempty"`
	// See WMI's Win32_Processor MOF for a detailed description the fields below
	AddressWidth              *int32  `protobuf:"varint,5,req,name=addressWidth" json:"addressWidth,omitempty"`
	DataWidth                 *int32  `protobuf:"varint,6,req,name=dataWidth" json:"dataWidth,omitempty"`
	NumberOfCores             *int32  `protobuf:"varint,7,req,name=numberOfCores" json:"numberOfCores,omitempty"`
	NumberOfLogicalProcessors *int32  `protobuf:"varint,8,req,name=numberOfLogicalProcessors" json:"numberOfLogicalProcessors,omitempty"`
	ProcessorManufacturer     *string `protobuf:"bytes,9,req,name=processorManufacturer" json:"processorManufacturer,omitempty"`
	ProcessorDescription      *string `protobuf:"bytes,10,req,name=processorDescription" json:"processorDescription,omitempty"`
	MaxClockSpeed             *int32  `protobuf:"varint,11,req,name=maxClockSpeed" json:"maxClockSpeed,omitempty"`
	ProcessorName             *string `protobuf:"bytes,12,req,name=processorName" json:"processorName,omitempty"`
	// See WMI's Win32_OperatingSystem MOF for a detailed description the fields below
	TotalVisibleMemorySize *string `protobuf:"bytes,13,req,name=totalVisibleMemorySize" json:"totalVisibleMemorySize,omitempty"`
	TotalVirtualMemorySize *string `protobuf:"bytes,14,req,name=totalVirtualMemorySize" json:"totalVirtualMemorySize,omitempty"`
	// See WMI's Win32_PageFileUsage MOF for a detailed description the fields below
	PageFileAllocatedBaseSize *int32 `protobuf:"varint,15,req,name=pageFileAllocatedBaseSize" json:"pageFileAllocatedBaseSize,omitempty"`
	// See WMI's Win32_OperatingSystem MOF for a detailed description the fields below
	OSname                    *string `protobuf:"bytes,16,req,name=OSname" json:"OSname,omitempty"`
	OSversion                 *string `protobuf:"bytes,17,req,name=OSversion" json:"OSversion,omitempty"`
	OSservicePackMajorVersion *int32  `protobuf:"varint,18,req,name=OSservicePackMajorVersion" json:"OSservicePackMajorVersion,omitempty"`
	OSservicePackMinorVersion *int32  `protobuf:"varint,19,req,name=OSservicePackMinorVersion" json:"OSservicePackMinorVersion,omitempty"`
	OSCSDVersion              *string `protobuf:"bytes,20,req,name=OSCSDVersion" json:"OSCSDVersion,omitempty"`
	OSsuiteMask               *int32  `protobuf:"varint,21,req,name=OSsuiteMask" json:"OSsuiteMask,omitempty"`
	OSProductType             *int32  `protobuf:"varint,22,req,name=OSProductType" json:"OSProductType,omitempty"`
	OSSKU                     *int32  `protobuf:"varint,23,req,name=OSSKU" json:"OSSKU,omitempty"`
	OSInstallDate             *string `protobuf:"bytes,24,req,name=OSInstallDate" json:"OSInstallDate,omitempty"`
	Locale                    *string `protobuf:"bytes,25,req,name=locale" json:"locale,omitempty"`
	// See WMI's Win32_ComputerSystem MOF for a detailed description the fields below
	DnsHostName          *string `protobuf:"bytes,26,req,name=dnsHostName" json:"dnsHostName,omitempty"`
	Domain               *string `protobuf:"bytes,27,req,name=domain" json:"domain,omitempty"`
	UserName             *string `protobuf:"bytes,28,req,name=userName" json:"userName,omitempty"`
	ComputerManufacturer *string `protobuf:"bytes,29,req,name=computerManufacturer" json:"computerManufacturer,omitempty"`
	ComputerModel        *string `protobuf:"bytes,30,req,name=computerModel" json:"computerModel,omitempty"`
	// See WMI's WMI_Win32_BaseBoard MOF for a detailed description the fields below
	BaseBoardManufacturer *string `protobuf:"bytes,31,req,name=baseBoardManufacturer" json:"baseBoardManufacturer,omitempty"`
	BaseBoardProduct      *string `protobuf:"bytes,32,req,name=baseBoardProduct" json:"baseBoardProduct,omitempty"`
	// Win32_VideoController MOF for a detailed description the fields below
	VideoControllerName *string `protobuf:"bytes,33,req,name=videoControllerName" json:"videoControllerName,omitempty"`
	// GUID that helps join this message with associated SystemInventoryNetworkAdapters message
	NetworkAdaptersGUID *string `protobuf:"bytes,34,req,name=networkAdaptersGUID" json:"networkAdaptersGUID,omitempty"`
	// GUID that helps join this message with associated SystemInventoryDiskDrives message
	DiskdrivesGUID *string `protobuf:"bytes,35,req,name=diskdrivesGUID" json:"diskdrivesGUID,omitempty"`
	// The version of the ziften agent i.e. 4.1.1.123
	ZiftenAgentVersion *string `protobuf:"bytes,36,opt,name=ziftenAgentVersion" json:"ziftenAgentVersion,omitempty"`
	SiteId             *string `protobuf:"bytes,37,opt,name=siteId" json:"siteId,omitempty"`
	// On Dell This is the ServiceTag: Win32_BIOS::SerialNumber
	SerialNumber *string `protobuf:"bytes,38,opt,name=SerialNumber" json:"SerialNumber,omitempty"`
	// Get OU name for system
	OUSystemName *string `protobuf:"bytes,39,opt,name=OUSystemName" json:"OUSystemName,omitempty"`
	Uuid         *string `protobuf:"bytes,40,opt,name=uuid" json:"uuid,omitempty"`
	// This is used by Lenovo for Asset Tags - Win32_SystemEnclosure::SMBIOSAssetTag
	BIOSAssetTag     *string `protobuf:"bytes,41,opt,name=BIOSAssetTag" json:"BIOSAssetTag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SystemInventory) Reset()                    { *m = SystemInventory{} }
func (m *SystemInventory) String() string            { return proto1.CompactTextString(m) }
func (*SystemInventory) ProtoMessage()               {}
func (*SystemInventory) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{0} }

func (m *SystemInventory) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *SystemInventory) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *SystemInventory) GetComputerName() string {
	if m != nil && m.ComputerName != nil {
		return *m.ComputerName
	}
	return ""
}

func (m *SystemInventory) GetCpuScore() float64 {
	if m != nil && m.CpuScore != nil {
		return *m.CpuScore
	}
	return 0
}

func (m *SystemInventory) GetAddressWidth() int32 {
	if m != nil && m.AddressWidth != nil {
		return *m.AddressWidth
	}
	return 0
}

func (m *SystemInventory) GetDataWidth() int32 {
	if m != nil && m.DataWidth != nil {
		return *m.DataWidth
	}
	return 0
}

func (m *SystemInventory) GetNumberOfCores() int32 {
	if m != nil && m.NumberOfCores != nil {
		return *m.NumberOfCores
	}
	return 0
}

func (m *SystemInventory) GetNumberOfLogicalProcessors() int32 {
	if m != nil && m.NumberOfLogicalProcessors != nil {
		return *m.NumberOfLogicalProcessors
	}
	return 0
}

func (m *SystemInventory) GetProcessorManufacturer() string {
	if m != nil && m.ProcessorManufacturer != nil {
		return *m.ProcessorManufacturer
	}
	return ""
}

func (m *SystemInventory) GetProcessorDescription() string {
	if m != nil && m.ProcessorDescription != nil {
		return *m.ProcessorDescription
	}
	return ""
}

func (m *SystemInventory) GetMaxClockSpeed() int32 {
	if m != nil && m.MaxClockSpeed != nil {
		return *m.MaxClockSpeed
	}
	return 0
}

func (m *SystemInventory) GetProcessorName() string {
	if m != nil && m.ProcessorName != nil {
		return *m.ProcessorName
	}
	return ""
}

func (m *SystemInventory) GetTotalVisibleMemorySize() string {
	if m != nil && m.TotalVisibleMemorySize != nil {
		return *m.TotalVisibleMemorySize
	}
	return ""
}

func (m *SystemInventory) GetTotalVirtualMemorySize() string {
	if m != nil && m.TotalVirtualMemorySize != nil {
		return *m.TotalVirtualMemorySize
	}
	return ""
}

func (m *SystemInventory) GetPageFileAllocatedBaseSize() int32 {
	if m != nil && m.PageFileAllocatedBaseSize != nil {
		return *m.PageFileAllocatedBaseSize
	}
	return 0
}

func (m *SystemInventory) GetOSname() string {
	if m != nil && m.OSname != nil {
		return *m.OSname
	}
	return ""
}

func (m *SystemInventory) GetOSversion() string {
	if m != nil && m.OSversion != nil {
		return *m.OSversion
	}
	return ""
}

func (m *SystemInventory) GetOSservicePackMajorVersion() int32 {
	if m != nil && m.OSservicePackMajorVersion != nil {
		return *m.OSservicePackMajorVersion
	}
	return 0
}

func (m *SystemInventory) GetOSservicePackMinorVersion() int32 {
	if m != nil && m.OSservicePackMinorVersion != nil {
		return *m.OSservicePackMinorVersion
	}
	return 0
}

func (m *SystemInventory) GetOSCSDVersion() string {
	if m != nil && m.OSCSDVersion != nil {
		return *m.OSCSDVersion
	}
	return ""
}

func (m *SystemInventory) GetOSsuiteMask() int32 {
	if m != nil && m.OSsuiteMask != nil {
		return *m.OSsuiteMask
	}
	return 0
}

func (m *SystemInventory) GetOSProductType() int32 {
	if m != nil && m.OSProductType != nil {
		return *m.OSProductType
	}
	return 0
}

func (m *SystemInventory) GetOSSKU() int32 {
	if m != nil && m.OSSKU != nil {
		return *m.OSSKU
	}
	return 0
}

func (m *SystemInventory) GetOSInstallDate() string {
	if m != nil && m.OSInstallDate != nil {
		return *m.OSInstallDate
	}
	return ""
}

func (m *SystemInventory) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *SystemInventory) GetDnsHostName() string {
	if m != nil && m.DnsHostName != nil {
		return *m.DnsHostName
	}
	return ""
}

func (m *SystemInventory) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *SystemInventory) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *SystemInventory) GetComputerManufacturer() string {
	if m != nil && m.ComputerManufacturer != nil {
		return *m.ComputerManufacturer
	}
	return ""
}

func (m *SystemInventory) GetComputerModel() string {
	if m != nil && m.ComputerModel != nil {
		return *m.ComputerModel
	}
	return ""
}

func (m *SystemInventory) GetBaseBoardManufacturer() string {
	if m != nil && m.BaseBoardManufacturer != nil {
		return *m.BaseBoardManufacturer
	}
	return ""
}

func (m *SystemInventory) GetBaseBoardProduct() string {
	if m != nil && m.BaseBoardProduct != nil {
		return *m.BaseBoardProduct
	}
	return ""
}

func (m *SystemInventory) GetVideoControllerName() string {
	if m != nil && m.VideoControllerName != nil {
		return *m.VideoControllerName
	}
	return ""
}

func (m *SystemInventory) GetNetworkAdaptersGUID() string {
	if m != nil && m.NetworkAdaptersGUID != nil {
		return *m.NetworkAdaptersGUID
	}
	return ""
}

func (m *SystemInventory) GetDiskdrivesGUID() string {
	if m != nil && m.DiskdrivesGUID != nil {
		return *m.DiskdrivesGUID
	}
	return ""
}

func (m *SystemInventory) GetZiftenAgentVersion() string {
	if m != nil && m.ZiftenAgentVersion != nil {
		return *m.ZiftenAgentVersion
	}
	return ""
}

func (m *SystemInventory) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *SystemInventory) GetSerialNumber() string {
	if m != nil && m.SerialNumber != nil {
		return *m.SerialNumber
	}
	return ""
}

func (m *SystemInventory) GetOUSystemName() string {
	if m != nil && m.OUSystemName != nil {
		return *m.OUSystemName
	}
	return ""
}

func (m *SystemInventory) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *SystemInventory) GetBIOSAssetTag() string {
	if m != nil && m.BIOSAssetTag != nil {
		return *m.BIOSAssetTag
	}
	return ""
}

func init() {
	proto1.RegisterType((*SystemInventory)(nil), "SystemInventory")
}

var fileDescriptor80 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x54, 0xcb, 0x72, 0x13, 0x3b,
	0x10, 0xad, 0x3c, 0xaf, 0xad, 0xc4, 0x79, 0x8c, 0x63, 0x5f, 0xe5, 0x75, 0xaf, 0x13, 0x08, 0x04,
	0x16, 0xfe, 0x02, 0x36, 0x7e, 0x14, 0xe0, 0x02, 0x67, 0x52, 0xa5, 0x24, 0xac, 0xe5, 0x51, 0xc7,
	0x08, 0x6b, 0x46, 0x53, 0x92, 0xc6, 0x90, 0x2c, 0xf9, 0x25, 0x8a, 0xff, 0xa3, 0x47, 0xe3, 0x71,
	0x62, 0xb3, 0xd4, 0xe9, 0xa3, 0xd6, 0xe9, 0xd3, 0xdd, 0x22, 0x0d, 0xf6, 0x60, 0x1d, 0xc4, 0x83,
	0x64, 0x0a, 0x89, 0xd3, 0xe6, 0xa1, 0x9d, 0x1a, 0xed, 0xf4, 0x51, 0x70, 0xad, 0xad, 0x1b, 0x1b,
	0xb0, 0x37, 0x0f, 0x29, 0x14, 0xd8, 0xf9, 0xaf, 0x0a, 0xd9, 0x5d, 0x62, 0x07, 0x2d, 0x52, 0x75,
	0x32, 0x06, 0xe6, 0x78, 0x9c, 0xd2, 0x95, 0xd6, 0xea, 0xe5, 0x5a, 0xb7, 0xf6, 0xf3, 0x37, 0xf5,
	0xa0, 0xcd, 0xc1, 0xe0, 0x98, 0x54, 0xf9, 0x18, 0xd9, 0x1f, 0x6e, 0x07, 0x7d, 0xba, 0x8a, 0x8c,
	0x6a, 0xb7, 0x82, 0x8c, 0xf5, 0x2c, 0x93, 0x22, 0x38, 0x20, 0xdb, 0x91, 0x8e, 0xd3, 0xcc, 0x81,
	0xb9, 0xe2, 0x31, 0xd0, 0xb5, 0x3c, 0x1e, 0xec, 0x91, 0x4a, 0x94, 0x66, 0x2c, 0xd2, 0x06, 0xe8,
	0x3a, 0x22, 0x2b, 0x39, 0x8f, 0x0b, 0x81, 0x72, 0xec, 0x17, 0x29, 0xdc, 0x57, 0xba, 0x81, 0xe8,
	0x46, 0xb0, 0x4f, 0xaa, 0x82, 0x3b, 0x5e, 0x40, 0x9b, 0x1e, 0x6a, 0x90, 0x5a, 0x92, 0xc5, 0x23,
	0x30, 0xe1, 0x7d, 0x0f, 0xaf, 0x5b, 0xfa, 0x8f, 0x87, 0xcf, 0xc8, 0x61, 0x09, 0x7f, 0xd6, 0x63,
	0x19, 0x71, 0x75, 0x6d, 0x74, 0x84, 0xe9, 0xb4, 0xb1, 0xb4, 0xe2, 0x29, 0xa7, 0xa4, 0x91, 0x96,
	0xd8, 0x90, 0x27, 0xd9, 0x3d, 0x8f, 0x5c, 0x66, 0xc0, 0xd0, 0xaa, 0xd7, 0x74, 0x42, 0x0e, 0xe6,
	0xe1, 0x3e, 0xd8, 0xc8, 0xc8, 0xd4, 0x49, 0x9d, 0x50, 0xe2, 0xa3, 0xf8, 0x6c, 0xcc, 0x7f, 0xf4,
	0x94, 0x8e, 0x26, 0x2c, 0x05, 0x10, 0x74, 0xab, 0x54, 0x33, 0xbf, 0xe4, 0xeb, 0xdb, 0xf6, 0xec,
	0xb7, 0xa4, 0xe9, 0xb4, 0xe3, 0xea, 0x4e, 0x5a, 0x39, 0x52, 0x30, 0x84, 0x18, 0xad, 0x64, 0xf2,
	0x11, 0x68, 0xcd, 0xfb, 0x43, 0xd0, 0x9f, 0xcd, 0x91, 0x1c, 0xcb, 0xc4, 0x3d, 0xe3, 0x1a, 0x97,
	0x71, 0xf5, 0x8c, 0xbb, 0xf3, 0x17, 0x17, 0xab, 0x4c, 0xd1, 0xeb, 0xf7, 0x52, 0x41, 0x47, 0xa1,
	0x16, 0xee, 0x40, 0x74, 0xb9, 0x05, 0x4f, 0xdf, 0xf5, 0x8a, 0x76, 0xc8, 0x66, 0xc8, 0x92, 0x5c,
	0xca, 0x9e, 0x97, 0x82, 0x16, 0x86, 0x6c, 0x0a, 0xc6, 0xe6, 0xb5, 0xec, 0x7b, 0x08, 0xb3, 0x84,
	0xcc, 0x82, 0x99, 0xca, 0x08, 0xae, 0x79, 0x34, 0x19, 0xf2, 0x6f, 0xda, 0xdc, 0xcd, 0x28, 0x41,
	0x69, 0xe7, 0x22, 0x45, 0x26, 0x4f, 0x94, 0xba, 0xa7, 0x60, 0xc7, 0x42, 0xd6, 0x63, 0xfd, 0x12,
	0x3d, 0xf0, 0xb9, 0xeb, 0x64, 0x0b, 0x2f, 0x66, 0xd2, 0xc1, 0x90, 0xdb, 0x09, 0x6d, 0x94, 0x2e,
	0x85, 0x0c, 0xfb, 0x21, 0xb2, 0xc8, 0xe5, 0xe3, 0x46, 0x9b, 0x1e, 0xae, 0x91, 0x8d, 0x90, 0xb1,
	0x4f, 0xb7, 0xf4, 0xdf, 0x27, 0xd6, 0x20, 0xc1, 0xa1, 0x52, 0xaa, 0x8f, 0x95, 0x51, 0xea, 0x33,
	0x62, 0x41, 0x79, 0xa5, 0x0a, 0xe8, 0x61, 0xf9, 0x82, 0x48, 0xec, 0x47, 0x9c, 0x5e, 0x6f, 0xf8,
	0x51, 0x49, 0x12, 0x3a, 0xe6, 0x32, 0xa1, 0xc7, 0xe5, 0x80, 0x65, 0x76, 0x36, 0x72, 0x27, 0x65,
	0x7b, 0xcb, 0x41, 0x5c, 0x68, 0xfe, 0x69, 0xd9, 0xde, 0x79, 0x54, 0x0b, 0x50, 0xf4, 0x3f, 0x0f,
	0xe3, 0xc8, 0x8c, 0xd0, 0xde, 0xae, 0xe6, 0x46, 0x2c, 0xdc, 0xfa, 0xdf, 0x87, 0x29, 0xd9, 0x9b,
	0x87, 0x67, 0xe5, 0xd1, 0x96, 0x8f, 0x1c, 0x93, 0xfa, 0x54, 0x0a, 0xd0, 0x3d, 0x9d, 0x38, 0xa3,
	0x95, 0x9a, 0x49, 0x39, 0xf3, 0xc1, 0x0b, 0x52, 0x4f, 0xc0, 0x7d, 0xd7, 0x66, 0xd2, 0x11, 0x3c,
	0xc5, 0x27, 0xad, 0x5f, 0x9d, 0xf3, 0xa5, 0xd5, 0x69, 0x91, 0x1d, 0x21, 0xed, 0x44, 0x18, 0x39,
	0x85, 0x82, 0xf1, 0x62, 0x89, 0x71, 0x44, 0x82, 0x47, 0x79, 0xef, 0x20, 0xe9, 0xe4, 0xfb, 0x57,
	0x36, 0xe2, 0x65, 0x6b, 0xa5, 0x70, 0xc4, 0x62, 0x17, 0x06, 0x82, 0x5e, 0xf8, 0x33, 0xb6, 0x8b,
	0x81, 0x91, 0x5c, 0x5d, 0xf9, 0x35, 0xa1, 0xaf, 0x4a, 0x34, 0xbc, 0x2d, 0x56, 0xde, 0x0b, 0x7c,
	0xed, 0xd1, 0x26, 0xf1, 0xf9, 0xe9, 0x65, 0x7e, 0x5a, 0x5c, 0xe6, 0xee, 0x20, 0x64, 0x1d, 0x6b,
	0xc1, 0xdd, 0xf0, 0x31, 0x7d, 0xe3, 0xe3, 0xef, 0xc8, 0x39, 0x7a, 0xd7, 0x2e, 0x94, 0xb4, 0xf3,
	0xa1, 0x01, 0x53, 0x7c, 0x28, 0x91, 0x56, 0x6d, 0xfc, 0x25, 0x2c, 0x8e, 0x6c, 0xb7, 0xb9, 0xf4,
	0xb1, 0x0c, 0x0b, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xdb, 0xc6, 0x11, 0x9f, 0x04,
	0x00, 0x00,
}
