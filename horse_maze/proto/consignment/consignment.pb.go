// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package consignment

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

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Created a blank get request
type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	// Added a pluralised consignment to our generic response message
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x2d, 0xff, 0x3b, 0x25, 0x12, 0x37, 0x11, 0x36, 0x7a, 0xb0, 0xa9, 0x17, 0x4e, 0x98,
	0x60, 0xe2, 0xc1, 0x78, 0x6b, 0x22, 0xe1, 0xba, 0x3c, 0x80, 0xc1, 0x76, 0x52, 0x26, 0x91, 0xdd,
	0xba, 0xbb, 0xe0, 0xdb, 0xf0, 0x04, 0x3e, 0xa4, 0x61, 0x4b, 0x65, 0xd1, 0x70, 0xdb, 0x6f, 0xbe,
	0x6f, 0xb6, 0xbf, 0x99, 0x2e, 0xdc, 0x97, 0x5a, 0x59, 0xf5, 0x90, 0x29, 0x69, 0xa8, 0x90, 0x6b,
	0x94, 0xd6, 0x3f, 0x4f, 0x9c, 0xcb, 0x22, 0xaf, 0x94, 0x7c, 0x07, 0x10, 0xa5, 0x47, 0xcd, 0x2e,
	0xa1, 0x41, 0x39, 0x0f, 0xe2, 0x60, 0x1c, 0x8a, 0x06, 0xe5, 0x2c, 0x86, 0x28, 0x47, 0x93, 0x69,
	0x2a, 0x2d, 0x29, 0xc9, 0x1b, 0xce, 0xf0, 0x4b, 0x6c, 0x08, 0x9d, 0x2f, 0xa4, 0x62, 0x65, 0x79,
	0x33, 0x0e, 0xc6, 0x6d, 0x71, 0x50, 0xec, 0x09, 0x20, 0x53, 0xd2, 0x2e, 0x49, 0xa2, 0x36, 0xbc,
	0x15, 0x37, 0xc7, 0xd1, 0x74, 0x38, 0xf1, 0x71, 0xd2, 0xda, 0x16, 0x5e, 0x92, 0xdd, 0x42, 0xb8,
	0x45, 0x63, 0xf0, 0xe3, 0x8d, 0x72, 0xde, 0x76, 0xdf, 0xeb, 0x55, 0x85, 0x79, 0x9e, 0xac, 0x21,
	0xfc, 0xed, 0xfa, 0xc7, 0x7a, 0x07, 0x51, 0xb6, 0x31, 0x56, 0xad, 0x51, 0xef, 0x7b, 0x2b, 0x56,
	0xa8, 0x4b, 0xf3, 0x7c, 0x8f, 0xaa, 0x34, 0x15, 0x24, 0x1d, 0x6a, 0x28, 0x0e, 0x8a, 0x8d, 0xa0,
	0xbb, 0x31, 0x55, 0x53, 0xab, 0x32, 0xf6, 0x72, 0x9e, 0x27, 0x7d, 0x80, 0x19, 0x5a, 0x81, 0x9f,
	0x1b, 0x34, 0x36, 0xd9, 0x05, 0xd0, 0x13, 0x68, 0x4a, 0x25, 0x0d, 0x32, 0x0e, 0xdd, 0x4c, 0xe3,
	0xd2, 0x62, 0x45, 0xd0, 0x13, 0xb5, 0x64, 0xcf, 0xe0, 0x6f, 0xd8, 0x61, 0x44, 0x53, 0xfe, 0x77,
	0xf2, 0xfa, 0x2c, 0xfc, 0x30, 0x7b, 0x81, 0xbe, 0x27, 0x0d, 0x6f, 0xba, 0xb5, 0x9d, 0x6f, 0x3e,
	0x49, 0x4f, 0x77, 0x01, 0x0c, 0x16, 0x2b, 0x2a, 0x4b, 0x92, 0xc5, 0x02, 0xf5, 0x96, 0x32, 0x64,
	0xaf, 0x70, 0x95, 0x3a, 0x30, 0xff, 0x2f, 0x9f, 0xbd, 0xf0, 0xe6, 0xfa, 0xc4, 0xa9, 0xa7, 0x4d,
	0x2e, 0x58, 0x0a, 0x83, 0x19, 0x5a, 0x2f, 0x6a, 0xd8, 0xe8, 0x24, 0x7b, 0x5c, 0xd4, 0xd9, 0x4b,
	0xde, 0x3b, 0xee, 0x05, 0x3e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x65, 0x3e, 0x9a, 0xa6, 0xa8,
	0x02, 0x00, 0x00,
}