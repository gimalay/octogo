// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gimalay/octogo/todoapp/binding/viewModel.proto

package binding

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LocationType int32

const (
	LocationType_Unknown LocationType = 0
	LocationType_Home    LocationType = 1457
	LocationType_Project LocationType = 1571
	LocationType_Task    LocationType = 1648
)

var LocationType_name = map[int32]string{
	0:    "Unknown",
	1457: "Home",
	1571: "Project",
	1648: "Task",
}

var LocationType_value = map[string]int32{
	"Unknown": 0,
	"Home":    1457,
	"Project": 1571,
	"Task":    1648,
}

func (x LocationType) String() string {
	return proto.EnumName(LocationType_name, int32(x))
}

func (LocationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{0}
}

type Settings struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{0}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

type Location struct {
	Type                 LocationType `protobuf:"varint,1,opt,name=type,proto3,enum=binding.LocationType" json:"type,omitempty"`
	Payload              []byte       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetType() LocationType {
	if m != nil {
		return m.Type
	}
	return LocationType_Unknown
}

func (m *Location) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Location_Home struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location_Home) Reset()         { *m = Location_Home{} }
func (m *Location_Home) String() string { return proto.CompactTextString(m) }
func (*Location_Home) ProtoMessage()    {}
func (*Location_Home) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{1, 0}
}

func (m *Location_Home) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location_Home.Unmarshal(m, b)
}
func (m *Location_Home) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location_Home.Marshal(b, m, deterministic)
}
func (m *Location_Home) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location_Home.Merge(m, src)
}
func (m *Location_Home) XXX_Size() int {
	return xxx_messageInfo_Location_Home.Size(m)
}
func (m *Location_Home) XXX_DiscardUnknown() {
	xxx_messageInfo_Location_Home.DiscardUnknown(m)
}

var xxx_messageInfo_Location_Home proto.InternalMessageInfo

type Location_Project struct {
	ProjectID            []byte   `protobuf:"bytes,2697,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location_Project) Reset()         { *m = Location_Project{} }
func (m *Location_Project) String() string { return proto.CompactTextString(m) }
func (*Location_Project) ProtoMessage()    {}
func (*Location_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{1, 1}
}

func (m *Location_Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location_Project.Unmarshal(m, b)
}
func (m *Location_Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location_Project.Marshal(b, m, deterministic)
}
func (m *Location_Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location_Project.Merge(m, src)
}
func (m *Location_Project) XXX_Size() int {
	return xxx_messageInfo_Location_Project.Size(m)
}
func (m *Location_Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Location_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Location_Project proto.InternalMessageInfo

func (m *Location_Project) GetProjectID() []byte {
	if m != nil {
		return m.ProjectID
	}
	return nil
}

type Location_Task struct {
	TaskID               []byte   `protobuf:"bytes,1717,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location_Task) Reset()         { *m = Location_Task{} }
func (m *Location_Task) String() string { return proto.CompactTextString(m) }
func (*Location_Task) ProtoMessage()    {}
func (*Location_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{1, 2}
}

func (m *Location_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location_Task.Unmarshal(m, b)
}
func (m *Location_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location_Task.Marshal(b, m, deterministic)
}
func (m *Location_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location_Task.Merge(m, src)
}
func (m *Location_Task) XXX_Size() int {
	return xxx_messageInfo_Location_Task.Size(m)
}
func (m *Location_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Location_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Location_Task proto.InternalMessageInfo

func (m *Location_Task) GetTaskID() []byte {
	if m != nil {
		return m.TaskID
	}
	return nil
}

type Location_AddTask struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location_AddTask) Reset()         { *m = Location_AddTask{} }
func (m *Location_AddTask) String() string { return proto.CompactTextString(m) }
func (*Location_AddTask) ProtoMessage()    {}
func (*Location_AddTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{1, 3}
}

func (m *Location_AddTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location_AddTask.Unmarshal(m, b)
}
func (m *Location_AddTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location_AddTask.Marshal(b, m, deterministic)
}
func (m *Location_AddTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location_AddTask.Merge(m, src)
}
func (m *Location_AddTask) XXX_Size() int {
	return xxx_messageInfo_Location_AddTask.Size(m)
}
func (m *Location_AddTask) XXX_DiscardUnknown() {
	xxx_messageInfo_Location_AddTask.DiscardUnknown(m)
}

var xxx_messageInfo_Location_AddTask proto.InternalMessageInfo

type ViewModel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewModel) Reset()         { *m = ViewModel{} }
func (m *ViewModel) String() string { return proto.CompactTextString(m) }
func (*ViewModel) ProtoMessage()    {}
func (*ViewModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2}
}

func (m *ViewModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel.Unmarshal(m, b)
}
func (m *ViewModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel.Marshal(b, m, deterministic)
}
func (m *ViewModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel.Merge(m, src)
}
func (m *ViewModel) XXX_Size() int {
	return xxx_messageInfo_ViewModel.Size(m)
}
func (m *ViewModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel proto.InternalMessageInfo

type ViewModel_Project struct {
	ID                   []byte                    `protobuf:"bytes,4947,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string                    `protobuf:"bytes,4032,opt,name=name,proto3" json:"name,omitempty"`
	Tasks                []*ViewModel_Project_Task `protobuf:"bytes,8856,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ViewModel_Project) Reset()         { *m = ViewModel_Project{} }
func (m *ViewModel_Project) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Project) ProtoMessage()    {}
func (*ViewModel_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 0}
}

func (m *ViewModel_Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Project.Unmarshal(m, b)
}
func (m *ViewModel_Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Project.Marshal(b, m, deterministic)
}
func (m *ViewModel_Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Project.Merge(m, src)
}
func (m *ViewModel_Project) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Project.Size(m)
}
func (m *ViewModel_Project) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Project.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Project proto.InternalMessageInfo

func (m *ViewModel_Project) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ViewModel_Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewModel_Project) GetTasks() []*ViewModel_Project_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ViewModel_Project_Task struct {
	ID                   []byte   `protobuf:"bytes,5946,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,8336,opt,name=name,proto3" json:"name,omitempty"`
	Emoji                string   `protobuf:"bytes,9790,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewModel_Project_Task) Reset()         { *m = ViewModel_Project_Task{} }
func (m *ViewModel_Project_Task) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Project_Task) ProtoMessage()    {}
func (*ViewModel_Project_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 0, 0}
}

func (m *ViewModel_Project_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Project_Task.Unmarshal(m, b)
}
func (m *ViewModel_Project_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Project_Task.Marshal(b, m, deterministic)
}
func (m *ViewModel_Project_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Project_Task.Merge(m, src)
}
func (m *ViewModel_Project_Task) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Project_Task.Size(m)
}
func (m *ViewModel_Project_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Project_Task.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Project_Task proto.InternalMessageInfo

func (m *ViewModel_Project_Task) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ViewModel_Project_Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewModel_Project_Task) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

type ViewModel_Task struct {
	ID                   []byte   `protobuf:"bytes,4987,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,4932,opt,name=name,proto3" json:"name,omitempty"`
	Emoji                string   `protobuf:"bytes,2651,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewModel_Task) Reset()         { *m = ViewModel_Task{} }
func (m *ViewModel_Task) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Task) ProtoMessage()    {}
func (*ViewModel_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 1}
}

func (m *ViewModel_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Task.Unmarshal(m, b)
}
func (m *ViewModel_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Task.Marshal(b, m, deterministic)
}
func (m *ViewModel_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Task.Merge(m, src)
}
func (m *ViewModel_Task) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Task.Size(m)
}
func (m *ViewModel_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Task.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Task proto.InternalMessageInfo

func (m *ViewModel_Task) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ViewModel_Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewModel_Task) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

type ViewModel_Home struct {
	Projects             []*ViewModel_Home_Project `protobuf:"bytes,6910,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ViewModel_Home) Reset()         { *m = ViewModel_Home{} }
func (m *ViewModel_Home) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Home) ProtoMessage()    {}
func (*ViewModel_Home) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 2}
}

func (m *ViewModel_Home) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Home.Unmarshal(m, b)
}
func (m *ViewModel_Home) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Home.Marshal(b, m, deterministic)
}
func (m *ViewModel_Home) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Home.Merge(m, src)
}
func (m *ViewModel_Home) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Home.Size(m)
}
func (m *ViewModel_Home) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Home.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Home proto.InternalMessageInfo

func (m *ViewModel_Home) GetProjects() []*ViewModel_Home_Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ViewModel_Home_Project struct {
	ID                   []byte                         `protobuf:"bytes,4947,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string                         `protobuf:"bytes,4032,opt,name=name,proto3" json:"name,omitempty"`
	Tasks                []*ViewModel_Home_Project_Task `protobuf:"bytes,8856,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ViewModel_Home_Project) Reset()         { *m = ViewModel_Home_Project{} }
func (m *ViewModel_Home_Project) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Home_Project) ProtoMessage()    {}
func (*ViewModel_Home_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 2, 0}
}

func (m *ViewModel_Home_Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Home_Project.Unmarshal(m, b)
}
func (m *ViewModel_Home_Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Home_Project.Marshal(b, m, deterministic)
}
func (m *ViewModel_Home_Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Home_Project.Merge(m, src)
}
func (m *ViewModel_Home_Project) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Home_Project.Size(m)
}
func (m *ViewModel_Home_Project) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Home_Project.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Home_Project proto.InternalMessageInfo

func (m *ViewModel_Home_Project) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ViewModel_Home_Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewModel_Home_Project) GetTasks() []*ViewModel_Home_Project_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ViewModel_Home_Project_Task struct {
	ID                   []byte   `protobuf:"bytes,2946,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,4336,opt,name=name,proto3" json:"name,omitempty"`
	Emoji                string   `protobuf:"bytes,2990,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewModel_Home_Project_Task) Reset()         { *m = ViewModel_Home_Project_Task{} }
func (m *ViewModel_Home_Project_Task) String() string { return proto.CompactTextString(m) }
func (*ViewModel_Home_Project_Task) ProtoMessage()    {}
func (*ViewModel_Home_Project_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c2ead79da5e18c, []int{2, 2, 0, 0}
}

func (m *ViewModel_Home_Project_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewModel_Home_Project_Task.Unmarshal(m, b)
}
func (m *ViewModel_Home_Project_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewModel_Home_Project_Task.Marshal(b, m, deterministic)
}
func (m *ViewModel_Home_Project_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewModel_Home_Project_Task.Merge(m, src)
}
func (m *ViewModel_Home_Project_Task) XXX_Size() int {
	return xxx_messageInfo_ViewModel_Home_Project_Task.Size(m)
}
func (m *ViewModel_Home_Project_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewModel_Home_Project_Task.DiscardUnknown(m)
}

var xxx_messageInfo_ViewModel_Home_Project_Task proto.InternalMessageInfo

func (m *ViewModel_Home_Project_Task) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ViewModel_Home_Project_Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ViewModel_Home_Project_Task) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func init() {
	proto.RegisterEnum("binding.LocationType", LocationType_name, LocationType_value)
	proto.RegisterType((*Settings)(nil), "binding.Settings")
	proto.RegisterType((*Location)(nil), "binding.Location")
	proto.RegisterType((*Location_Home)(nil), "binding.Location.Home")
	proto.RegisterType((*Location_Project)(nil), "binding.Location.Project")
	proto.RegisterType((*Location_Task)(nil), "binding.Location.Task")
	proto.RegisterType((*Location_AddTask)(nil), "binding.Location.AddTask")
	proto.RegisterType((*ViewModel)(nil), "binding.ViewModel")
	proto.RegisterType((*ViewModel_Project)(nil), "binding.ViewModel.Project")
	proto.RegisterType((*ViewModel_Project_Task)(nil), "binding.ViewModel.Project.Task")
	proto.RegisterType((*ViewModel_Task)(nil), "binding.ViewModel.Task")
	proto.RegisterType((*ViewModel_Home)(nil), "binding.ViewModel.Home")
	proto.RegisterType((*ViewModel_Home_Project)(nil), "binding.ViewModel.Home.Project")
	proto.RegisterType((*ViewModel_Home_Project_Task)(nil), "binding.ViewModel.Home.Project.Task")
}

func init() {
	proto.RegisterFile("github.com/gimalay/octogo/todoapp/binding/viewModel.proto", fileDescriptor_c9c2ead79da5e18c)
}

var fileDescriptor_c9c2ead79da5e18c = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcc, 0x4d, 0xcc, 0x49, 0xac, 0xd4, 0xcf,
	0x4f, 0x2e, 0xc9, 0x4f, 0xcf, 0xd7, 0x2f, 0xc9, 0x4f, 0xc9, 0x4f, 0x2c, 0x28, 0xd0, 0x4f, 0xca,
	0xcc, 0x4b, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0xcb, 0x4c, 0x2d, 0xf7, 0xcd, 0x4f, 0x49, 0xcd, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x28, 0x71, 0x71, 0x71, 0x04, 0xa7, 0x96,
	0x94, 0x64, 0xe6, 0xa5, 0x17, 0x2b, 0xad, 0x65, 0xe4, 0xe2, 0xf0, 0xc9, 0x4f, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x13, 0xd2, 0xe4, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x33, 0x12, 0xd5, 0x83, 0x6a, 0xd0, 0x83, 0x29, 0x08, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x2b, 0x11,
	0x92, 0xe0, 0x62, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x09, 0x82, 0x71, 0xa5, 0xd8, 0xb8, 0x58, 0x3c, 0xf2, 0x73, 0x53, 0xa5, 0x34, 0xb8, 0xd8, 0x03,
	0x8a, 0xf2, 0xb3, 0x52, 0x93, 0x4b, 0x84, 0x64, 0xb9, 0x38, 0xa1, 0x4c, 0x4f, 0x17, 0x89, 0x4e,
	0x51, 0xb0, 0x7a, 0x84, 0x88, 0x94, 0x3c, 0x17, 0x4b, 0x48, 0x62, 0x71, 0xb6, 0x90, 0x38, 0x17,
	0x1b, 0x88, 0xf6, 0x74, 0x91, 0xd8, 0xca, 0x0b, 0x56, 0x03, 0xe5, 0x4a, 0x71, 0x72, 0xb1, 0x3b,
	0xa6, 0xa4, 0x80, 0x38, 0x4a, 0xbf, 0x98, 0xb9, 0x38, 0xc3, 0x60, 0x1e, 0x93, 0x5a, 0xc3, 0x88,
	0xb0, 0x84, 0x9f, 0x8b, 0xc9, 0xd3, 0x45, 0xe2, 0xb2, 0x1a, 0x58, 0x27, 0x93, 0xa7, 0x8b, 0x90,
	0x30, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0xc4, 0x01, 0x79, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30,
	0x47, 0xc8, 0x8c, 0x8b, 0xb5, 0x24, 0xb1, 0x38, 0xbb, 0x58, 0x62, 0x86, 0xab, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0x3c, 0xdc, 0x93, 0x70, 0x53, 0xf5, 0xa0, 0x26, 0xea, 0x81, 0x2c, 0x0c, 0x82, 0x28,
	0x97, 0x72, 0x86, 0xba, 0x11, 0x62, 0xcb, 0x2e, 0x3d, 0x0c, 0x5b, 0x26, 0x38, 0x22, 0xd9, 0x22,
	0xca, 0xc5, 0x9a, 0x9a, 0x9b, 0x9f, 0x95, 0x29, 0xb1, 0xcf, 0x07, 0x2c, 0x0a, 0xe1, 0xa1, 0x19,
	0xf2, 0x1b, 0xd3, 0xa9, 0x47, 0xd4, 0xb0, 0x19, 0x72, 0x5b, 0x04, 0xd9, 0x90, 0x1f, 0x8c, 0x90,
	0x00, 0x16, 0xb2, 0xe1, 0xe2, 0x28, 0x80, 0xb8, 0xb4, 0x58, 0xe2, 0x9f, 0x29, 0x4e, 0xdf, 0x80,
	0xd4, 0xc2, 0xbc, 0x14, 0x04, 0xd7, 0x21, 0xb5, 0x91, 0xe4, 0xa0, 0xb3, 0x46, 0x0b, 0x3a, 0x15,
	0x02, 0x96, 0xe1, 0x09, 0xbf, 0x26, 0x71, 0x0c, 0xab, 0x3e, 0x28, 0x62, 0xf3, 0xfa, 0x3a, 0x71,
	0x24, 0xaf, 0x6b, 0xd9, 0x73, 0xf1, 0x20, 0x27, 0x45, 0x21, 0x6e, 0x2e, 0xf6, 0xd0, 0xbc, 0xec,
	0xbc, 0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x4e, 0x48, 0xb0, 0x08, 0x6c, 0xe4, 0x16, 0xe2, 0x81,
	0x7b, 0x4d, 0x60, 0x31, 0x0f, 0x48, 0x02, 0x64, 0xb5, 0xc0, 0x07, 0x1e, 0x27, 0xe6, 0x5d, 0x4c,
	0x0c, 0x49, 0x6c, 0xe0, 0xec, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x26, 0x8b, 0x85,
	0x4b, 0x03, 0x00, 0x00,
}
