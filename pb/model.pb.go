// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	ModelDef
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModelDef struct {
	Size    int32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Layers  int32 `protobuf:"varint,2,opt,name=layers" json:"layers,omitempty"`
	Kernel  int32 `protobuf:"varint,3,opt,name=kernel" json:"kernel,omitempty"`
	Filters int32 `protobuf:"varint,4,opt,name=filters" json:"filters,omitempty"`
	Hidden  int32 `protobuf:"varint,5,opt,name=hidden" json:"hidden,omitempty"`
}

func (m *ModelDef) Reset()                    { *m = ModelDef{} }
func (m *ModelDef) String() string            { return proto.CompactTextString(m) }
func (*ModelDef) ProtoMessage()               {}
func (*ModelDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*ModelDef)(nil), "tak.proto.ModelDef")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x49, 0xcc, 0x86, 0x30, 0x95, 0x1a,
	0x18, 0xb9, 0x38, 0x7c, 0x41, 0x52, 0x2e, 0xa9, 0x69, 0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x99, 0x55,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0xb6, 0x90, 0x18, 0x17, 0x5b, 0x4e, 0x62,
	0x65, 0x6a, 0x51, 0xb1, 0x04, 0x13, 0x58, 0x14, 0xca, 0x03, 0x89, 0x67, 0xa7, 0x16, 0xe5, 0xa5,
	0xe6, 0x48, 0x30, 0x43, 0xc4, 0x21, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0xb4, 0xcc, 0x9c, 0x12, 0x90,
	0x06, 0x16, 0xb0, 0x04, 0x8c, 0x0b, 0xd2, 0x91, 0x91, 0x99, 0x92, 0x92, 0x9a, 0x27, 0xc1, 0x0a,
	0xd1, 0x01, 0xe1, 0x39, 0xb1, 0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0xdd, 0x63, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x99, 0xb4, 0x3a, 0xa9, 0x00, 0x00, 0x00,
}
