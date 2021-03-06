// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_intr_rpc.proto

/*
Package resourceproto is a generated protocol buffer package.

It is generated from these files:
	game_intr_rpc.proto

It has these top-level messages:
	ZoneStatusGetReqProto
	ZoneStatusGetRespProto
*/
package resourceproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ZoneStatusGetReqProto struct {
	ZoneID uint32 `protobuf:"varint,1,opt,name=zoneID" json:"zoneID,omitempty"`
}

func (m *ZoneStatusGetReqProto) Reset()                    { *m = ZoneStatusGetReqProto{} }
func (m *ZoneStatusGetReqProto) String() string            { return proto.CompactTextString(m) }
func (*ZoneStatusGetReqProto) ProtoMessage()               {}
func (*ZoneStatusGetReqProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ZoneStatusGetReqProto) GetZoneID() uint32 {
	if m != nil {
		return m.ZoneID
	}
	return 0
}

type ZoneStatusGetRespProto struct {
	ZoneID     uint32 `protobuf:"varint,1,opt,name=zoneID" json:"zoneID,omitempty"`
	ZoneStatus int32  `protobuf:"varint,2,opt,name=zoneStatus" json:"zoneStatus,omitempty"`
}

func (m *ZoneStatusGetRespProto) Reset()                    { *m = ZoneStatusGetRespProto{} }
func (m *ZoneStatusGetRespProto) String() string            { return proto.CompactTextString(m) }
func (*ZoneStatusGetRespProto) ProtoMessage()               {}
func (*ZoneStatusGetRespProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ZoneStatusGetRespProto) GetZoneID() uint32 {
	if m != nil {
		return m.ZoneID
	}
	return 0
}

func (m *ZoneStatusGetRespProto) GetZoneStatus() int32 {
	if m != nil {
		return m.ZoneStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*ZoneStatusGetReqProto)(nil), "proto.resource.ZoneStatusGetReqProto")
	proto.RegisterType((*ZoneStatusGetRespProto)(nil), "proto.resource.ZoneStatusGetRespProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GameIntrService service

type GameIntrServiceClient interface {
	ZoneStatusGet(ctx context.Context, in *ZoneStatusGetReqProto, opts ...grpc.CallOption) (*ZoneStatusGetRespProto, error)
}

type gameIntrServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameIntrServiceClient(cc *grpc.ClientConn) GameIntrServiceClient {
	return &gameIntrServiceClient{cc}
}

func (c *gameIntrServiceClient) ZoneStatusGet(ctx context.Context, in *ZoneStatusGetReqProto, opts ...grpc.CallOption) (*ZoneStatusGetRespProto, error) {
	out := new(ZoneStatusGetRespProto)
	err := grpc.Invoke(ctx, "/proto.resource.GameIntrService/ZoneStatusGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameIntrService service

type GameIntrServiceServer interface {
	ZoneStatusGet(context.Context, *ZoneStatusGetReqProto) (*ZoneStatusGetRespProto, error)
}

func RegisterGameIntrServiceServer(s *grpc.Server, srv GameIntrServiceServer) {
	s.RegisterService(&_GameIntrService_serviceDesc, srv)
}

func _GameIntrService_ZoneStatusGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneStatusGetReqProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameIntrServiceServer).ZoneStatusGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.resource.GameIntrService/ZoneStatusGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameIntrServiceServer).ZoneStatusGet(ctx, req.(*ZoneStatusGetReqProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameIntrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.resource.GameIntrService",
	HandlerType: (*GameIntrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ZoneStatusGet",
			Handler:    _GameIntrService_ZoneStatusGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_intr_rpc.proto",
}

func init() { proto.RegisterFile("game_intr_rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4f, 0xcc, 0x4d,
	0x8d, 0xcf, 0xcc, 0x2b, 0x29, 0x8a, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x03, 0x53, 0x7a, 0x45, 0xa9, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x4a, 0xfa, 0x5c, 0xa2,
	0x51, 0xf9, 0x79, 0xa9, 0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0xee, 0xa9, 0x25, 0x41, 0xa9, 0x85,
	0x01, 0x60, 0x85, 0x62, 0x5c, 0x6c, 0x55, 0xf9, 0x79, 0xa9, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x50, 0x9e, 0x52, 0x00, 0x97, 0x18, 0x9a, 0x86, 0xe2, 0x02, 0xbc, 0x3a, 0x84,
	0xe4, 0xb8, 0xb8, 0xaa, 0xe0, 0x3a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x90, 0x44, 0x8c,
	0x0a, 0xb9, 0xf8, 0xdd, 0x13, 0x73, 0x53, 0x3d, 0xf3, 0x4a, 0x8a, 0x82, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x85, 0xe2, 0xb8, 0x78, 0x51, 0x2c, 0x11, 0x52, 0xd5, 0x43, 0x75, 0xb7, 0x1e, 0x56,
	0x47, 0x4b, 0xa9, 0x11, 0x50, 0x06, 0x75, 0xaa, 0x13, 0x7f, 0x14, 0x2f, 0x4c, 0x09, 0x58, 0x43,
	0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0x4f, 0xd0, 0x52, 0x34, 0x01,
	0x00, 0x00,
}
