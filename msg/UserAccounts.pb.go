// Code generated by protoc-gen-go.
// source: UserAccounts.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserAccounts_PrivilegeType int32

const (
	UserAccounts_eUSER_PRIV_UNKNOWN UserAccounts_PrivilegeType = 0
	UserAccounts_eUSER_PRIV_GUEST   UserAccounts_PrivilegeType = 1
	UserAccounts_eUSER_PRIV_USER    UserAccounts_PrivilegeType = 2
	UserAccounts_eUSER_PRIV_ADMIN   UserAccounts_PrivilegeType = 3
)

var UserAccounts_PrivilegeType_name = map[int32]string{
	0: "eUSER_PRIV_UNKNOWN",
	1: "eUSER_PRIV_GUEST",
	2: "eUSER_PRIV_USER",
	3: "eUSER_PRIV_ADMIN",
}
var UserAccounts_PrivilegeType_value = map[string]int32{
	"eUSER_PRIV_UNKNOWN": 0,
	"eUSER_PRIV_GUEST":   1,
	"eUSER_PRIV_USER":    2,
	"eUSER_PRIV_ADMIN":   3,
}

func (x UserAccounts_PrivilegeType) Enum() *UserAccounts_PrivilegeType {
	p := new(UserAccounts_PrivilegeType)
	*p = x
	return p
}
func (x UserAccounts_PrivilegeType) String() string {
	return proto.EnumName(UserAccounts_PrivilegeType_name, int32(x))
}
func (x *UserAccounts_PrivilegeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserAccounts_PrivilegeType_value, data, "UserAccounts_PrivilegeType")
	if err != nil {
		return err
	}
	*x = UserAccounts_PrivilegeType(value)
	return nil
}
func (UserAccounts_PrivilegeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor87, []int{0, 0}
}

// This message is used by agent to report Users Logging On and Off the system
type UserAccounts struct {
	// The UTC time that the message was produced by the agent. Time is what was reported by the client clock.
	TimeStamp *int64 `protobuf:"varint,1,req,name=timeStamp" json:"timeStamp,omitempty"`
	// The unique identifier of the agent. This field is used by the server to distinguish agents.
	AgentGUID *string `protobuf:"bytes,2,req,name=agentGUID" json:"agentGUID,omitempty"`
	// The server ID for multi-site hosting
	SiteId *string `protobuf:"bytes,3,opt,name=siteId" json:"siteId,omitempty"`
	// UUID used on the server
	Uuid        *string                      `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
	Name        []string                     `protobuf:"bytes,5,rep,name=name" json:"name,omitempty"`
	PasswordAge []uint32                     `protobuf:"varint,6,rep,name=password_age" json:"password_age,omitempty"`
	Priv        []UserAccounts_PrivilegeType `protobuf:"varint,7,rep,name=priv,enum=UserAccounts_PrivilegeType" json:"priv,omitempty"`
	HomeDir     []string                     `protobuf:"bytes,8,rep,name=home_dir" json:"home_dir,omitempty"`
	Comment     []string                     `protobuf:"bytes,9,rep,name=comment" json:"comment,omitempty"`
	ScriptPath  []string                     `protobuf:"bytes,10,rep,name=script_path" json:"script_path,omitempty"`
	// The following has been removed because it has been found to be unreliable
	// repeated AccountType account_type = 11; // Account type
	FullName              []string `protobuf:"bytes,12,rep,name=full_name" json:"full_name,omitempty"`
	UsrComment            []string `protobuf:"bytes,13,rep,name=usr_comment" json:"usr_comment,omitempty"`
	Parms                 []string `protobuf:"bytes,14,rep,name=parms" json:"parms,omitempty"`
	Workstations          []string `protobuf:"bytes,15,rep,name=workstations" json:"workstations,omitempty"`
	LastLogon             []uint32 `protobuf:"varint,16,rep,name=last_logon" json:"last_logon,omitempty"`
	LastLogoff            []uint32 `protobuf:"varint,17,rep,name=last_logoff" json:"last_logoff,omitempty"`
	AcctExpires           []uint32 `protobuf:"varint,18,rep,name=acct_expires" json:"acct_expires,omitempty"`
	MaxStorage            []uint32 `protobuf:"varint,19,rep,name=max_storage" json:"max_storage,omitempty"`
	BadPwCount            []uint32 `protobuf:"varint,20,rep,name=bad_pw_count" json:"bad_pw_count,omitempty"`
	NumLogons             []uint32 `protobuf:"varint,21,rep,name=num_logons" json:"num_logons,omitempty"`
	LogonServer           []string `protobuf:"bytes,22,rep,name=logon_server" json:"logon_server,omitempty"`
	CountryCode           []uint32 `protobuf:"varint,23,rep,name=country_code" json:"country_code,omitempty"`
	CodePage              []uint32 `protobuf:"varint,24,rep,name=code_page" json:"code_page,omitempty"`
	UserId                []uint32 `protobuf:"varint,25,rep,name=user_id" json:"user_id,omitempty"`
	PrimaryGroupId        []uint32 `protobuf:"varint,26,rep,name=primary_group_id" json:"primary_group_id,omitempty"`
	Profile               []string `protobuf:"bytes,27,rep,name=profile" json:"profile,omitempty"`
	HomeDirDrive          []string `protobuf:"bytes,28,rep,name=home_dir_drive" json:"home_dir_drive,omitempty"`
	PasswordExpired       []uint32 `protobuf:"varint,29,rep,name=password_expired" json:"password_expired,omitempty"`
	MemberOfPrintGroup    []bool   `protobuf:"varint,30,rep,name=member_of_print_group" json:"member_of_print_group,omitempty"`
	MemberOfCommGroup     []bool   `protobuf:"varint,31,rep,name=member_of_comm_group" json:"member_of_comm_group,omitempty"`
	MemberOfServerGroup   []bool   `protobuf:"varint,32,rep,name=member_of_server_group" json:"member_of_server_group,omitempty"`
	MemberOfAccountsGroup []bool   `protobuf:"varint,33,rep,name=member_of_accounts_group" json:"member_of_accounts_group,omitempty"`
	Sid                   []string `protobuf:"bytes,34,rep,name=sid" json:"sid,omitempty"`
	Disabled              []bool   `protobuf:"varint,35,rep,name=disabled" json:"disabled,omitempty"`
	LockedOut             []bool   `protobuf:"varint,36,rep,name=lockedOut" json:"lockedOut,omitempty"`
	XXX_unrecognized      []byte   `json:"-"`
}

func (m *UserAccounts) Reset()                    { *m = UserAccounts{} }
func (m *UserAccounts) String() string            { return proto.CompactTextString(m) }
func (*UserAccounts) ProtoMessage()               {}
func (*UserAccounts) Descriptor() ([]byte, []int) { return fileDescriptor87, []int{0} }

func (m *UserAccounts) GetTimeStamp() int64 {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return 0
}

func (m *UserAccounts) GetAgentGUID() string {
	if m != nil && m.AgentGUID != nil {
		return *m.AgentGUID
	}
	return ""
}

func (m *UserAccounts) GetSiteId() string {
	if m != nil && m.SiteId != nil {
		return *m.SiteId
	}
	return ""
}

func (m *UserAccounts) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *UserAccounts) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserAccounts) GetPasswordAge() []uint32 {
	if m != nil {
		return m.PasswordAge
	}
	return nil
}

func (m *UserAccounts) GetPriv() []UserAccounts_PrivilegeType {
	if m != nil {
		return m.Priv
	}
	return nil
}

func (m *UserAccounts) GetHomeDir() []string {
	if m != nil {
		return m.HomeDir
	}
	return nil
}

func (m *UserAccounts) GetComment() []string {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *UserAccounts) GetScriptPath() []string {
	if m != nil {
		return m.ScriptPath
	}
	return nil
}

func (m *UserAccounts) GetFullName() []string {
	if m != nil {
		return m.FullName
	}
	return nil
}

func (m *UserAccounts) GetUsrComment() []string {
	if m != nil {
		return m.UsrComment
	}
	return nil
}

func (m *UserAccounts) GetParms() []string {
	if m != nil {
		return m.Parms
	}
	return nil
}

func (m *UserAccounts) GetWorkstations() []string {
	if m != nil {
		return m.Workstations
	}
	return nil
}

func (m *UserAccounts) GetLastLogon() []uint32 {
	if m != nil {
		return m.LastLogon
	}
	return nil
}

func (m *UserAccounts) GetLastLogoff() []uint32 {
	if m != nil {
		return m.LastLogoff
	}
	return nil
}

func (m *UserAccounts) GetAcctExpires() []uint32 {
	if m != nil {
		return m.AcctExpires
	}
	return nil
}

func (m *UserAccounts) GetMaxStorage() []uint32 {
	if m != nil {
		return m.MaxStorage
	}
	return nil
}

func (m *UserAccounts) GetBadPwCount() []uint32 {
	if m != nil {
		return m.BadPwCount
	}
	return nil
}

func (m *UserAccounts) GetNumLogons() []uint32 {
	if m != nil {
		return m.NumLogons
	}
	return nil
}

func (m *UserAccounts) GetLogonServer() []string {
	if m != nil {
		return m.LogonServer
	}
	return nil
}

func (m *UserAccounts) GetCountryCode() []uint32 {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *UserAccounts) GetCodePage() []uint32 {
	if m != nil {
		return m.CodePage
	}
	return nil
}

func (m *UserAccounts) GetUserId() []uint32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *UserAccounts) GetPrimaryGroupId() []uint32 {
	if m != nil {
		return m.PrimaryGroupId
	}
	return nil
}

func (m *UserAccounts) GetProfile() []string {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *UserAccounts) GetHomeDirDrive() []string {
	if m != nil {
		return m.HomeDirDrive
	}
	return nil
}

func (m *UserAccounts) GetPasswordExpired() []uint32 {
	if m != nil {
		return m.PasswordExpired
	}
	return nil
}

func (m *UserAccounts) GetMemberOfPrintGroup() []bool {
	if m != nil {
		return m.MemberOfPrintGroup
	}
	return nil
}

func (m *UserAccounts) GetMemberOfCommGroup() []bool {
	if m != nil {
		return m.MemberOfCommGroup
	}
	return nil
}

func (m *UserAccounts) GetMemberOfServerGroup() []bool {
	if m != nil {
		return m.MemberOfServerGroup
	}
	return nil
}

func (m *UserAccounts) GetMemberOfAccountsGroup() []bool {
	if m != nil {
		return m.MemberOfAccountsGroup
	}
	return nil
}

func (m *UserAccounts) GetSid() []string {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *UserAccounts) GetDisabled() []bool {
	if m != nil {
		return m.Disabled
	}
	return nil
}

func (m *UserAccounts) GetLockedOut() []bool {
	if m != nil {
		return m.LockedOut
	}
	return nil
}

func init() {
	proto.RegisterType((*UserAccounts)(nil), "UserAccounts")
	proto.RegisterEnum("UserAccounts_PrivilegeType", UserAccounts_PrivilegeType_name, UserAccounts_PrivilegeType_value)
}

var fileDescriptor87 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x53, 0x4d, 0x53, 0xdb, 0x30,
	0x10, 0x2d, 0x24, 0x40, 0xb2, 0x24, 0xc1, 0x18, 0x48, 0xb7, 0x7c, 0x35, 0x4d, 0x7b, 0xa0, 0x97,
	0x1c, 0x7a, 0xeb, 0x91, 0x0c, 0x0c, 0x93, 0xe9, 0x10, 0x18, 0x20, 0xed, 0x51, 0x23, 0x6c, 0x25,
	0x78, 0xb0, 0x2d, 0x8f, 0x24, 0xf3, 0xd1, 0x63, 0xff, 0x40, 0x7f, 0x4d, 0xff, 0x5f, 0xd7, 0xeb,
	0xb8, 0x84, 0xde, 0xac, 0xf7, 0x9e, 0x76, 0xf7, 0xbd, 0x95, 0xc1, 0x9f, 0x58, 0x65, 0x8e, 0x83,
	0x40, 0xe7, 0xa9, 0xb3, 0x83, 0xcc, 0x68, 0xa7, 0x77, 0xfd, 0x4b, 0x6d, 0xdd, 0xcc, 0x28, 0x7b,
	0xf3, 0x9c, 0xa9, 0x12, 0xeb, 0xff, 0x5e, 0x83, 0xd6, 0xa2, 0xd4, 0xef, 0x41, 0xd3, 0x45, 0x89,
	0xba, 0x76, 0x32, 0xc9, 0x70, 0xa9, 0xb7, 0x7c, 0x54, 0x1b, 0xb6, 0x7f, 0xfd, 0x41, 0x06, 0x6d,
	0x01, 0xfa, 0x7b, 0xd0, 0x94, 0x33, 0x95, 0xba, 0xb3, 0xc9, 0xe8, 0x04, 0x97, 0x49, 0xd1, 0x1c,
	0x36, 0x48, 0x51, 0xcf, 0xf3, 0x28, 0xf4, 0x3b, 0xb0, 0x6a, 0x23, 0xa7, 0x46, 0x21, 0xd6, 0x7a,
	0x4b, 0x47, 0x4d, 0xbf, 0x0b, 0x8c, 0x63, 0xbd, 0x38, 0x2d, 0xe8, 0x5a, 0x50, 0x4f, 0x65, 0xa2,
	0x70, 0xa5, 0x57, 0x23, 0xd5, 0x36, 0xb4, 0x32, 0x69, 0xed, 0xa3, 0x36, 0xa1, 0xa0, 0xda, 0xb8,
	0x4a, 0x68, 0xdb, 0xff, 0x0c, 0xf5, 0xcc, 0x44, 0x0f, 0xb8, 0x46, 0xa7, 0xce, 0x97, 0xbd, 0xc1,
	0x2b, 0x4b, 0x97, 0xc4, 0x44, 0xb1, 0x9a, 0xa9, 0xc2, 0x8c, 0xef, 0x41, 0xe3, 0x4e, 0x27, 0x4a,
	0x84, 0x91, 0xc1, 0x06, 0x97, 0xdc, 0x80, 0xb5, 0x40, 0x27, 0x09, 0xcd, 0x89, 0x4d, 0x06, 0xb6,
	0x60, 0xdd, 0x06, 0x26, 0xca, 0x9c, 0xc8, 0xa4, 0xbb, 0x43, 0x60, 0x70, 0x13, 0x9a, 0xd3, 0x3c,
	0x8e, 0x05, 0xcf, 0xd2, 0xaa, 0x74, 0xb9, 0x35, 0xa2, 0xba, 0xdc, 0x66, 0xb0, 0x0d, 0x2b, 0x99,
	0x34, 0x89, 0xc5, 0x4e, 0x35, 0x2f, 0xcd, 0x7a, 0x4f, 0x79, 0xb8, 0x48, 0xa7, 0x16, 0x37, 0x18,
	0xf5, 0x01, 0x62, 0x69, 0x9d, 0x88, 0xf5, 0x4c, 0xa7, 0xe8, 0xb1, 0x07, 0xaa, 0xf6, 0x0f, 0x9b,
	0x4e, 0x71, 0x93, 0x41, 0xba, 0x2e, 0x83, 0xc0, 0x09, 0xf5, 0x94, 0x45, 0xb4, 0x0e, 0xf4, 0x2b,
	0x69, 0x22, 0x9f, 0x84, 0x75, 0xda, 0x14, 0x19, 0x6c, 0x55, 0xd2, 0x5b, 0x19, 0x8a, 0xec, 0x51,
	0xb0, 0x6f, 0xdc, 0x66, 0x94, 0x3a, 0xa5, 0x79, 0x52, 0x36, 0xb2, 0xb8, 0x53, 0x29, 0xf9, 0x2c,
	0x28, 0xa5, 0x07, 0x65, 0xb0, 0x5b, 0x4d, 0xca, 0x17, 0xcd, 0x33, 0x15, 0x08, 0x15, 0xbe, 0x65,
	0x2d, 0xd9, 0x2e, 0x4e, 0x94, 0x04, 0x35, 0x42, 0x86, 0x28, 0xaf, 0x9c, 0x6e, 0x0a, 0xda, 0xd5,
	0x3b, 0x06, 0x10, 0x3c, 0x4a, 0x3f, 0x91, 0x74, 0x73, 0x66, 0x74, 0x9e, 0x15, 0xcc, 0x6e, 0x25,
	0xa5, 0xc7, 0x33, 0xa5, 0xf4, 0x71, 0x8f, 0x9b, 0x74, 0xa1, 0x53, 0xa5, 0x2f, 0x42, 0xda, 0x8b,
	0xc2, 0x7d, 0xc6, 0x8b, 0x12, 0xd5, 0x5a, 0x4b, 0xaf, 0x21, 0x1e, 0x70, 0x89, 0x03, 0xd8, 0x49,
	0x54, 0x72, 0x4b, 0xfd, 0xf4, 0x54, 0x50, 0x9b, 0xd4, 0x95, 0x4d, 0xf0, 0x90, 0xe8, 0x86, 0xbf,
	0x0f, 0xdb, 0x2f, 0x74, 0xb1, 0x89, 0x39, 0xfb, 0x9e, 0xd9, 0x43, 0xe8, 0xbe, 0xb0, 0xa5, 0xdb,
	0x39, 0xdf, 0x63, 0xbe, 0x07, 0xf8, 0xc2, 0xcb, 0xf9, 0x83, 0x99, 0x2b, 0x3e, 0xb0, 0x62, 0x1d,
	0x6a, 0x96, 0xec, 0xf4, 0x79, 0x4a, 0x7a, 0x3b, 0x61, 0x64, 0xe5, 0x6d, 0x4c, 0xd3, 0x7d, 0x64,
	0x9a, 0xe2, 0x89, 0x75, 0x70, 0xaf, 0xc2, 0x8b, 0xdc, 0xe1, 0xa7, 0x02, 0xea, 0xdf, 0x41, 0xfb,
	0xf5, 0x8b, 0xeb, 0x82, 0xaf, 0x26, 0xd7, 0xa7, 0x57, 0xe2, 0xf2, 0x6a, 0xf4, 0x5d, 0x4c, 0xc6,
	0xdf, 0xc6, 0x17, 0x3f, 0xc6, 0xde, 0x1b, 0x0a, 0xdc, 0x5b, 0xc0, 0xcf, 0x26, 0xa7, 0xd7, 0x37,
	0xde, 0x12, 0xed, 0x76, 0x63, 0x51, 0x4d, 0x5f, 0xde, 0xf2, 0x7f, 0xd2, 0xe3, 0x93, 0xf3, 0xd1,
	0xd8, 0xab, 0x0d, 0xbf, 0x42, 0x9f, 0x1c, 0x0f, 0x7e, 0x46, 0x53, 0xa7, 0xd2, 0x41, 0x69, 0xaf,
	0xfc, 0x59, 0x03, 0x1d, 0x0f, 0xe8, 0x27, 0xb4, 0xb4, 0xb4, 0xe1, 0xd6, 0xe2, 0xcf, 0x70, 0x5e,
	0x82, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xee, 0xa8, 0x52, 0xf5, 0x03, 0x00, 0x00,
}
