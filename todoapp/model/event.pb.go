// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gimalay/octogo/todoapp/model/event.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventType int32

const (
	EventType_Unknown EventType = 0
	// Project
	EventType_ProjectCreated EventType = 8812
	EventType_ProjectRenamed EventType = 8489
	EventType_TaskAdded      EventType = 1541
	EventType_TaskRemoved    EventType = 4982
	// Task
	EventType_TaskImported EventType = 1036
	EventType_TaskCreated  EventType = 8694
	EventType_TaskRenamed  EventType = 5239
)

var EventType_name = map[int32]string{
	0:    "Unknown",
	8812: "ProjectCreated",
	8489: "ProjectRenamed",
	1541: "TaskAdded",
	4982: "TaskRemoved",
	1036: "TaskImported",
	8694: "TaskCreated",
	5239: "TaskRenamed",
}

var EventType_value = map[string]int32{
	"Unknown":        0,
	"ProjectCreated": 8812,
	"ProjectRenamed": 8489,
	"TaskAdded":      1541,
	"TaskRemoved":    4982,
	"TaskImported":   1036,
	"TaskCreated":    8694,
	"TaskRenamed":    5239,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0}
}

type AggregateType int32

const (
	AggregateType_UnknownAggregate AggregateType = 0
	AggregateType_Project          AggregateType = 3756
	AggregateType_Task             AggregateType = 8068
)

var AggregateType_name = map[int32]string{
	0:    "UnknownAggregate",
	3756: "Project",
	8068: "Task",
}

var AggregateType_value = map[string]int32{
	"UnknownAggregate": 0,
	"Project":          3756,
	"Task":             8068,
}

func (x AggregateType) String() string {
	return proto.EnumName(AggregateType_name, int32(x))
}

func (AggregateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{1}
}

type Event struct {
	ID                   uint64               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PreviousMessageHash  []byte               `protobuf:"bytes,2,opt,name=previousMessageHash,proto3" json:"previousMessageHash,omitempty"`
	AggregateID          []byte               `protobuf:"bytes,3,opt,name=aggregateID,proto3" json:"aggregateID,omitempty"`
	AggregateType        AggregateType        `protobuf:"varint,4,opt,name=aggregateType,proto3,enum=model.AggregateType" json:"aggregateType,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SourceID             []byte               `protobuf:"bytes,6,opt,name=sourceID,proto3" json:"sourceID,omitempty"`
	SourceVersion        int32                `protobuf:"varint,7,opt,name=sourceVersion,proto3" json:"sourceVersion,omitempty"`
	Type                 EventType            `protobuf:"varint,8,opt,name=type,proto3,enum=model.EventType" json:"type,omitempty"`
	Payload              []byte               `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Event) GetPreviousMessageHash() []byte {
	if m != nil {
		return m.PreviousMessageHash
	}
	return nil
}

func (m *Event) GetAggregateID() []byte {
	if m != nil {
		return m.AggregateID
	}
	return nil
}

func (m *Event) GetAggregateType() AggregateType {
	if m != nil {
		return m.AggregateType
	}
	return AggregateType_UnknownAggregate
}

func (m *Event) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetSourceID() []byte {
	if m != nil {
		return m.SourceID
	}
	return nil
}

func (m *Event) GetSourceVersion() int32 {
	if m != nil {
		return m.SourceVersion
	}
	return 0
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Event_ProjectCreated struct {
	Name                 string   `protobuf:"bytes,2825,opt,name=name,proto3" json:"name,omitempty"`
	Tasks                [][]byte `protobuf:"bytes,9477,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_ProjectCreated) Reset()         { *m = Event_ProjectCreated{} }
func (m *Event_ProjectCreated) String() string { return proto.CompactTextString(m) }
func (*Event_ProjectCreated) ProtoMessage()    {}
func (*Event_ProjectCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 0}
}

func (m *Event_ProjectCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_ProjectCreated.Unmarshal(m, b)
}
func (m *Event_ProjectCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_ProjectCreated.Marshal(b, m, deterministic)
}
func (m *Event_ProjectCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_ProjectCreated.Merge(m, src)
}
func (m *Event_ProjectCreated) XXX_Size() int {
	return xxx_messageInfo_Event_ProjectCreated.Size(m)
}
func (m *Event_ProjectCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_ProjectCreated.DiscardUnknown(m)
}

var xxx_messageInfo_Event_ProjectCreated proto.InternalMessageInfo

func (m *Event_ProjectCreated) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event_ProjectCreated) GetTasks() [][]byte {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Event_TaskRemoved struct {
	TaskID               []byte   `protobuf:"bytes,2296,opt,name=taskID,proto3" json:"taskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_TaskRemoved) Reset()         { *m = Event_TaskRemoved{} }
func (m *Event_TaskRemoved) String() string { return proto.CompactTextString(m) }
func (*Event_TaskRemoved) ProtoMessage()    {}
func (*Event_TaskRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 1}
}

func (m *Event_TaskRemoved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_TaskRemoved.Unmarshal(m, b)
}
func (m *Event_TaskRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_TaskRemoved.Marshal(b, m, deterministic)
}
func (m *Event_TaskRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_TaskRemoved.Merge(m, src)
}
func (m *Event_TaskRemoved) XXX_Size() int {
	return xxx_messageInfo_Event_TaskRemoved.Size(m)
}
func (m *Event_TaskRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_TaskRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_Event_TaskRemoved proto.InternalMessageInfo

func (m *Event_TaskRemoved) GetTaskID() []byte {
	if m != nil {
		return m.TaskID
	}
	return nil
}

type Event_TaskAdded struct {
	TaskID               []byte   `protobuf:"bytes,6339,opt,name=taskID,proto3" json:"taskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_TaskAdded) Reset()         { *m = Event_TaskAdded{} }
func (m *Event_TaskAdded) String() string { return proto.CompactTextString(m) }
func (*Event_TaskAdded) ProtoMessage()    {}
func (*Event_TaskAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 2}
}

func (m *Event_TaskAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_TaskAdded.Unmarshal(m, b)
}
func (m *Event_TaskAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_TaskAdded.Marshal(b, m, deterministic)
}
func (m *Event_TaskAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_TaskAdded.Merge(m, src)
}
func (m *Event_TaskAdded) XXX_Size() int {
	return xxx_messageInfo_Event_TaskAdded.Size(m)
}
func (m *Event_TaskAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_TaskAdded.DiscardUnknown(m)
}

var xxx_messageInfo_Event_TaskAdded proto.InternalMessageInfo

func (m *Event_TaskAdded) GetTaskID() []byte {
	if m != nil {
		return m.TaskID
	}
	return nil
}

type Event_ProjectRenamed struct {
	Name                 string   `protobuf:"bytes,8825,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_ProjectRenamed) Reset()         { *m = Event_ProjectRenamed{} }
func (m *Event_ProjectRenamed) String() string { return proto.CompactTextString(m) }
func (*Event_ProjectRenamed) ProtoMessage()    {}
func (*Event_ProjectRenamed) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 3}
}

func (m *Event_ProjectRenamed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_ProjectRenamed.Unmarshal(m, b)
}
func (m *Event_ProjectRenamed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_ProjectRenamed.Marshal(b, m, deterministic)
}
func (m *Event_ProjectRenamed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_ProjectRenamed.Merge(m, src)
}
func (m *Event_ProjectRenamed) XXX_Size() int {
	return xxx_messageInfo_Event_ProjectRenamed.Size(m)
}
func (m *Event_ProjectRenamed) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_ProjectRenamed.DiscardUnknown(m)
}

var xxx_messageInfo_Event_ProjectRenamed proto.InternalMessageInfo

func (m *Event_ProjectRenamed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Event_TaskCreated struct {
	Name                 string   `protobuf:"bytes,3500,opt,name=name,proto3" json:"name,omitempty"`
	Emoji                string   `protobuf:"bytes,4282,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_TaskCreated) Reset()         { *m = Event_TaskCreated{} }
func (m *Event_TaskCreated) String() string { return proto.CompactTextString(m) }
func (*Event_TaskCreated) ProtoMessage()    {}
func (*Event_TaskCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 4}
}

func (m *Event_TaskCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_TaskCreated.Unmarshal(m, b)
}
func (m *Event_TaskCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_TaskCreated.Marshal(b, m, deterministic)
}
func (m *Event_TaskCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_TaskCreated.Merge(m, src)
}
func (m *Event_TaskCreated) XXX_Size() int {
	return xxx_messageInfo_Event_TaskCreated.Size(m)
}
func (m *Event_TaskCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_TaskCreated.DiscardUnknown(m)
}

var xxx_messageInfo_Event_TaskCreated proto.InternalMessageInfo

func (m *Event_TaskCreated) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event_TaskCreated) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

type Event_TaskRenamed struct {
	Name                 string   `protobuf:"bytes,1770,opt,name=name,proto3" json:"name,omitempty"`
	Emoji                string   `protobuf:"bytes,1957,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_TaskRenamed) Reset()         { *m = Event_TaskRenamed{} }
func (m *Event_TaskRenamed) String() string { return proto.CompactTextString(m) }
func (*Event_TaskRenamed) ProtoMessage()    {}
func (*Event_TaskRenamed) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f057e723007719, []int{0, 5}
}

func (m *Event_TaskRenamed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_TaskRenamed.Unmarshal(m, b)
}
func (m *Event_TaskRenamed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_TaskRenamed.Marshal(b, m, deterministic)
}
func (m *Event_TaskRenamed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_TaskRenamed.Merge(m, src)
}
func (m *Event_TaskRenamed) XXX_Size() int {
	return xxx_messageInfo_Event_TaskRenamed.Size(m)
}
func (m *Event_TaskRenamed) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_TaskRenamed.DiscardUnknown(m)
}

var xxx_messageInfo_Event_TaskRenamed proto.InternalMessageInfo

func (m *Event_TaskRenamed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event_TaskRenamed) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func init() {
	proto.RegisterEnum("model.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("model.AggregateType", AggregateType_name, AggregateType_value)
	proto.RegisterType((*Event)(nil), "model.Event")
	proto.RegisterType((*Event_ProjectCreated)(nil), "model.Event.ProjectCreated")
	proto.RegisterType((*Event_TaskRemoved)(nil), "model.Event.TaskRemoved")
	proto.RegisterType((*Event_TaskAdded)(nil), "model.Event.TaskAdded")
	proto.RegisterType((*Event_ProjectRenamed)(nil), "model.Event.ProjectRenamed")
	proto.RegisterType((*Event_TaskCreated)(nil), "model.Event.TaskCreated")
	proto.RegisterType((*Event_TaskRenamed)(nil), "model.Event.TaskRenamed")
}

func init() {
	proto.RegisterFile("github.com/gimalay/octogo/todoapp/model/event.proto", fileDescriptor_55f057e723007719)
}

var fileDescriptor_55f057e723007719 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xeb, 0x36, 0x4e, 0xe2, 0xc9, 0x9f, 0x77, 0xbb, 0xe9, 0x0b, 0x2b, 0x23, 0x84, 0x41,
	0xa1, 0xb2, 0x7a, 0xb0, 0xa1, 0xbd, 0x00, 0x42, 0xa0, 0x52, 0x23, 0x11, 0x24, 0x24, 0x64, 0x05,
	0xee, 0x9b, 0x78, 0x71, 0xdd, 0xc4, 0x5e, 0xcb, 0xde, 0x04, 0xe5, 0x4e, 0x0f, 0x48, 0xfd, 0x06,
	0x9c, 0x39, 0x20, 0xf5, 0x13, 0x70, 0xe5, 0xdb, 0x20, 0xce, 0x05, 0x6e, 0xc8, 0xeb, 0xd8, 0x89,
	0xa3, 0xde, 0xfc, 0x3c, 0xf3, 0xcc, 0xcc, 0xcf, 0xa3, 0x85, 0x23, 0x3f, 0x10, 0xa7, 0xb3, 0x91,
	0x35, 0xe6, 0xa1, 0xed, 0x07, 0x21, 0x9d, 0xd2, 0x85, 0xcd, 0xc7, 0x82, 0xfb, 0xdc, 0x16, 0xdc,
	0xe3, 0x34, 0x8e, 0xed, 0x90, 0x7b, 0x6c, 0x6a, 0xb3, 0x39, 0x8b, 0x84, 0x15, 0x27, 0x5c, 0x70,
	0xac, 0x4a, 0x4b, 0xbf, 0xe3, 0x73, 0xee, 0x4f, 0x99, 0x2d, 0xcd, 0xd1, 0xec, 0x83, 0x2d, 0x82,
	0x90, 0xa5, 0x82, 0x86, 0x71, 0x9e, 0xbb, 0x77, 0xa1, 0x82, 0xfa, 0x32, 0xeb, 0xc3, 0x5d, 0xd8,
	0x1e, 0x38, 0x44, 0x31, 0x14, 0xb3, 0xe6, 0x6e, 0x0f, 0x1c, 0xfc, 0x00, 0x7a, 0x71, 0xc2, 0xe6,
	0x01, 0x9f, 0xa5, 0x6f, 0x58, 0x9a, 0x52, 0x9f, 0xbd, 0xa2, 0xe9, 0x29, 0xd9, 0x36, 0x14, 0xb3,
	0xed, 0x5e, 0x57, 0xc2, 0x06, 0xb4, 0xa8, 0xef, 0x27, 0xcc, 0xa7, 0x82, 0x0d, 0x1c, 0xb2, 0x23,
	0x93, 0xeb, 0x16, 0x7e, 0x02, 0x9d, 0x52, 0x0e, 0x17, 0x31, 0x23, 0x35, 0x43, 0x31, 0xbb, 0x87,
	0x7b, 0x96, 0xa4, 0xb5, 0x8e, 0xd7, 0x6b, 0x6e, 0x35, 0x8a, 0x1f, 0x81, 0x56, 0xc2, 0x13, 0xd5,
	0x50, 0xcc, 0xd6, 0xa1, 0x6e, 0xe5, 0xbf, 0x67, 0x15, 0xbf, 0x67, 0x0d, 0x8b, 0x84, 0xbb, 0x0a,
	0x63, 0x1d, 0x9a, 0x29, 0x9f, 0x25, 0xe3, 0x0c, 0xaa, 0x2e, 0xa1, 0x4a, 0x8d, 0xfb, 0xd0, 0xc9,
	0xbf, 0xdf, 0xb3, 0x24, 0x0d, 0x78, 0x44, 0x1a, 0x86, 0x62, 0xaa, 0x6e, 0xd5, 0xc4, 0x7d, 0xa8,
	0x89, 0x0c, 0xb7, 0x29, 0x71, 0xd1, 0x12, 0x57, 0xde, 0x4d, 0xa2, 0xca, 0x2a, 0x26, 0xd0, 0x88,
	0xe9, 0x62, 0xca, 0xa9, 0x47, 0x34, 0xb9, 0xa6, 0x90, 0xfa, 0x53, 0xe8, 0xbe, 0x4d, 0xf8, 0x19,
	0x1b, 0x8b, 0x93, 0x84, 0x51, 0xc1, 0x3c, 0xdc, 0x83, 0x5a, 0x44, 0x43, 0x46, 0x3e, 0xdf, 0x30,
	0x14, 0x53, 0x73, 0xa5, 0xc0, 0xff, 0x83, 0x2a, 0x68, 0x3a, 0x49, 0xc9, 0xf9, 0x6b, 0x63, 0xc7,
	0x6c, 0xbb, 0xb9, 0xd2, 0xf7, 0xa1, 0x35, 0xa4, 0xe9, 0xc4, 0x65, 0x21, 0x9f, 0x33, 0x0f, 0xdf,
	0x84, 0x7a, 0xe6, 0x0f, 0x1c, 0xf2, 0x67, 0x57, 0xae, 0x59, 0x4a, 0xbd, 0x0f, 0x5a, 0x96, 0x3b,
	0xf6, 0xbc, 0x4a, 0xea, 0xc7, 0xc3, 0x4a, 0xea, 0x7e, 0xc9, 0xe2, 0xb2, 0x6c, 0xeb, 0x8a, 0xe5,
	0xaf, 0xb3, 0x62, 0xd1, 0x1f, 0xe7, 0x4b, 0x37, 0x79, 0x2f, 0x6f, 0x55, 0x79, 0x59, 0xc8, 0xcf,
	0x02, 0xf2, 0xfd, 0xae, 0x74, 0x73, 0x55, 0xb4, 0x6e, 0x8e, 0xff, 0xd9, 0xb9, 0xae, 0xf5, 0xeb,
	0x7f, 0x6b, 0xad, 0x07, 0x5f, 0x14, 0xd0, 0xca, 0xb3, 0xe2, 0x16, 0x34, 0xde, 0x45, 0x93, 0x88,
	0x7f, 0x8c, 0xd0, 0x16, 0xee, 0x6d, 0xde, 0x10, 0xfd, 0x72, 0xd6, 0xcc, 0xe5, 0x36, 0xf4, 0xed,
	0x05, 0xee, 0xae, 0xdd, 0x01, 0x9d, 0xb7, 0x31, 0xaa, 0xdc, 0x0f, 0x5d, 0xed, 0xe3, 0x5d, 0x68,
	0x67, 0xce, 0x20, 0x8c, 0x79, 0x92, 0x4d, 0xba, 0x68, 0x16, 0xa1, 0x62, 0xf6, 0xd5, 0xc9, 0xaa,
	0x2d, 0x1f, 0xfc, 0xdb, 0x3c, 0x78, 0x06, 0x9d, 0xca, 0x13, 0xc5, 0x7b, 0x80, 0x96, 0x80, 0xa5,
	0x8f, 0xb6, 0x70, 0x1b, 0x1a, 0x4b, 0x28, 0x74, 0x79, 0x1b, 0x6b, 0x50, 0xcb, 0xc6, 0xa0, 0x4f,
	0xcf, 0x47, 0x75, 0xf9, 0x4e, 0x8f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x77, 0xe2, 0xcd, 0x65,
	0xd2, 0x03, 0x00, 0x00,
}