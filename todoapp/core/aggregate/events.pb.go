// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/gimalay/octogo/todoapp/core/aggregate/events.proto

package aggregate

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

type EventType int32

const (
	EventType_Unknown EventType = 0
	// Project
	EventType_ProjectCreated EventType = 8812
	EventType_ProjectRenamed EventType = 8489
	EventType_TaskAdded      EventType = 1541
	EventType_TaskRemoved    EventType = 4982
	// Task
	EventType_TaskCreated EventType = 8694
	EventType_TaskRenamed EventType = 5239
)

var EventType_name = map[int32]string{
	0:    "Unknown",
	8812: "ProjectCreated",
	8489: "ProjectRenamed",
	1541: "TaskAdded",
	4982: "TaskRemoved",
	8694: "TaskCreated",
	5239: "TaskRenamed",
}

var EventType_value = map[string]int32{
	"Unknown":        0,
	"ProjectCreated": 8812,
	"ProjectRenamed": 8489,
	"TaskAdded":      1541,
	"TaskRemoved":    4982,
	"TaskCreated":    8694,
	"TaskRenamed":    5239,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8635284dd4c81cf9, []int{0}
}

func init() {
	proto.RegisterEnum("aggregate.EventType", EventType_name, EventType_value)
}

func init() {
	proto.RegisterFile("github.com/gimalay/octogo/todoapp/core/aggregate/events.proto", fileDescriptor_8635284dd4c81cf9)
}

var fileDescriptor_8635284dd4c81cf9 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcc, 0x4d, 0xcc, 0x49, 0xac, 0xd4, 0xcf,
	0x4f, 0x2e, 0xc9, 0x4f, 0xcf, 0xd7, 0x2f, 0xc9, 0x4f, 0xc9, 0x4f, 0x2c, 0x28, 0xd0, 0x4f, 0xce,
	0x2f, 0x4a, 0xd5, 0x4f, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0xd5, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x8b, 0x6b, 0x75, 0x30,
	0x72, 0x71, 0xba, 0x82, 0xe4, 0x42, 0x2a, 0x0b, 0x52, 0x85, 0xb8, 0xb9, 0xd8, 0x43, 0xf3, 0xb2,
	0xf3, 0xf2, 0xcb, 0xf3, 0x04, 0x18, 0x84, 0x84, 0xb9, 0xf8, 0x02, 0x8a, 0xf2, 0xb3, 0x52, 0x93,
	0x4b, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53, 0x04, 0xde, 0xb8, 0x20, 0x09, 0x06, 0xa5, 0xe6,
	0x25, 0xe6, 0xa6, 0xa6, 0x08, 0xac, 0x74, 0x12, 0xe2, 0xe3, 0xe2, 0x0c, 0x49, 0x2c, 0xce, 0x76,
	0x4c, 0x49, 0x49, 0x4d, 0x11, 0x68, 0xe5, 0x11, 0x12, 0xe0, 0xe2, 0x06, 0xf1, 0x83, 0x52, 0x73,
	0xf3, 0xcb, 0x52, 0x53, 0x04, 0xbe, 0xa9, 0xc1, 0x44, 0x60, 0x06, 0x7d, 0x73, 0x46, 0xa8, 0x81,
	0x98, 0xf2, 0x5d, 0xc3, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x5c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x93, 0x4c, 0xb1, 0xf6, 0xef, 0x00, 0x00, 0x00,
}
