// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/trigger/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateTriggerRequest is the request type for creating a trigger RPC
type MsgCreateTriggerRequest struct {
	// The signing authorities for the request
	Authorities []string `protobuf:"bytes,1,rep,name=authorities,proto3" json:"authorities,omitempty"`
	// The event that must be detected for the trigger to fire.
	Event *types.Any `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	// The messages to run when the trigger fires.
	Actions []*types.Any `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (m *MsgCreateTriggerRequest) Reset()         { *m = MsgCreateTriggerRequest{} }
func (m *MsgCreateTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTriggerRequest) ProtoMessage()    {}
func (*MsgCreateTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f001c93b8aeec1f, []int{0}
}
func (m *MsgCreateTriggerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTriggerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTriggerRequest.Merge(m, src)
}
func (m *MsgCreateTriggerRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTriggerRequest proto.InternalMessageInfo

func (m *MsgCreateTriggerRequest) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *MsgCreateTriggerRequest) GetEvent() *types.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *MsgCreateTriggerRequest) GetActions() []*types.Any {
	if m != nil {
		return m.Actions
	}
	return nil
}

// MsgCreateTriggerResponse is the response type for creating a trigger RPC
type MsgCreateTriggerResponse struct {
	// trigger id that is generated on creation.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateTriggerResponse) Reset()         { *m = MsgCreateTriggerResponse{} }
func (m *MsgCreateTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTriggerResponse) ProtoMessage()    {}
func (*MsgCreateTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f001c93b8aeec1f, []int{1}
}
func (m *MsgCreateTriggerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTriggerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTriggerResponse.Merge(m, src)
}
func (m *MsgCreateTriggerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTriggerResponse proto.InternalMessageInfo

func (m *MsgCreateTriggerResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// MsgDestroyTriggerRequest is the request type for creating a trigger RPC
type MsgDestroyTriggerRequest struct {
	// the id of the trigger to destroy.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The signing authority for the request
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgDestroyTriggerRequest) Reset()         { *m = MsgDestroyTriggerRequest{} }
func (m *MsgDestroyTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*MsgDestroyTriggerRequest) ProtoMessage()    {}
func (*MsgDestroyTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f001c93b8aeec1f, []int{2}
}
func (m *MsgDestroyTriggerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDestroyTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDestroyTriggerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDestroyTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDestroyTriggerRequest.Merge(m, src)
}
func (m *MsgDestroyTriggerRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgDestroyTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDestroyTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDestroyTriggerRequest proto.InternalMessageInfo

func (m *MsgDestroyTriggerRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgDestroyTriggerRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgDestroyTriggerResponse is the response type for creating a trigger RPC
type MsgDestroyTriggerResponse struct {
}

func (m *MsgDestroyTriggerResponse) Reset()         { *m = MsgDestroyTriggerResponse{} }
func (m *MsgDestroyTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDestroyTriggerResponse) ProtoMessage()    {}
func (*MsgDestroyTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f001c93b8aeec1f, []int{3}
}
func (m *MsgDestroyTriggerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDestroyTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDestroyTriggerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDestroyTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDestroyTriggerResponse.Merge(m, src)
}
func (m *MsgDestroyTriggerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDestroyTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDestroyTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDestroyTriggerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateTriggerRequest)(nil), "provenance.trigger.v1.MsgCreateTriggerRequest")
	proto.RegisterType((*MsgCreateTriggerResponse)(nil), "provenance.trigger.v1.MsgCreateTriggerResponse")
	proto.RegisterType((*MsgDestroyTriggerRequest)(nil), "provenance.trigger.v1.MsgDestroyTriggerRequest")
	proto.RegisterType((*MsgDestroyTriggerResponse)(nil), "provenance.trigger.v1.MsgDestroyTriggerResponse")
}

func init() { proto.RegisterFile("provenance/trigger/v1/tx.proto", fileDescriptor_4f001c93b8aeec1f) }

var fileDescriptor_4f001c93b8aeec1f = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xb1, 0x6e, 0x13, 0x41,
	0x10, 0x86, 0xbd, 0x36, 0x01, 0x65, 0xa3, 0x44, 0x62, 0x65, 0x94, 0xcb, 0x21, 0x1d, 0x96, 0x2b,
	0xcb, 0x92, 0x77, 0x49, 0x22, 0x51, 0x58, 0xa2, 0x88, 0x81, 0x82, 0x22, 0x12, 0x3a, 0xa8, 0x68,
	0xd0, 0xd9, 0x5e, 0x36, 0x2b, 0xe1, 0x9b, 0x63, 0x67, 0x7d, 0xca, 0x75, 0x88, 0x27, 0xe0, 0x11,
	0x78, 0x84, 0x14, 0x79, 0x08, 0x44, 0x15, 0x51, 0x21, 0x2a, 0x64, 0x17, 0xa1, 0xe7, 0x05, 0x90,
	0x6f, 0xef, 0xb0, 0x93, 0x5c, 0xa2, 0x74, 0x37, 0x3b, 0xdf, 0xcc, 0xff, 0xcf, 0xdc, 0x2e, 0x0d,
	0x12, 0x03, 0xa9, 0x8c, 0xa3, 0x78, 0x24, 0x85, 0x35, 0x5a, 0x29, 0x69, 0x44, 0xba, 0x2b, 0xec,
	0x31, 0x4f, 0x0c, 0x58, 0x60, 0x0f, 0x96, 0x79, 0x5e, 0xe4, 0x79, 0xba, 0xeb, 0x6f, 0x8f, 0x00,
	0x27, 0x80, 0x62, 0x82, 0x6a, 0x81, 0x4f, 0x50, 0x39, 0xde, 0xdf, 0x71, 0x89, 0x77, 0x79, 0x24,
	0x5c, 0x50, 0xa4, 0x9a, 0x0a, 0x14, 0xb8, 0xf3, 0xc5, 0x57, 0x59, 0xa0, 0x00, 0xd4, 0x07, 0x29,
	0xf2, 0x68, 0x38, 0x7d, 0x2f, 0xa2, 0x38, 0x73, 0xa9, 0xf6, 0x2f, 0x42, 0xb7, 0x0f, 0x51, 0x3d,
	0x33, 0x32, 0xb2, 0xf2, 0x8d, 0x13, 0x0f, 0xe5, 0xc7, 0xa9, 0x44, 0xcb, 0xfa, 0x74, 0x23, 0x9a,
	0xda, 0x23, 0x30, 0xda, 0x6a, 0x89, 0x1e, 0x69, 0x35, 0x3a, 0xeb, 0x03, 0xef, 0xc7, 0x69, 0xaf,
	0x59, 0x68, 0x1e, 0x8c, 0xc7, 0x46, 0x22, 0xbe, 0xb6, 0x46, 0xc7, 0x2a, 0x5c, 0x85, 0xd9, 0x53,
	0xba, 0x26, 0x53, 0x19, 0x5b, 0xaf, 0xde, 0x22, 0x9d, 0x8d, 0xbd, 0x26, 0x77, 0x16, 0x78, 0x69,
	0x81, 0x1f, 0xc4, 0xd9, 0xe0, 0xfe, 0xf7, 0xd3, 0xde, 0x66, 0x21, 0xfa, 0x62, 0x41, 0xbf, 0x0c,
	0x5d, 0x15, 0xe3, 0xf4, 0x5e, 0x34, 0xb2, 0x1a, 0x62, 0xf4, 0x1a, 0xad, 0xc6, 0x75, 0x0d, 0xc2,
	0x12, 0xea, 0x37, 0xff, 0x7c, 0x7d, 0x44, 0x3e, 0x9f, 0x9f, 0x74, 0x57, 0x4d, 0xb4, 0xbb, 0xd4,
	0xbb, 0x3a, 0x1b, 0x26, 0x10, 0xa3, 0x64, 0x5b, 0xb4, 0xae, 0xc7, 0x1e, 0x69, 0x91, 0xce, 0x9d,
	0xb0, 0xae, 0xc7, 0xed, 0x34, 0x67, 0x9f, 0x4b, 0xb4, 0x06, 0xb2, 0x4b, 0x8b, 0xb8, 0xc4, 0xb2,
	0x27, 0x74, 0xbd, 0x94, 0xc9, 0xf2, 0x01, 0x6f, 0x5a, 0xcb, 0x12, 0xed, 0xb3, 0xd2, 0xe5, 0xf2,
	0xac, 0xfd, 0x90, 0xee, 0x54, 0xe8, 0x3a, 0x93, 0x7b, 0x7f, 0x09, 0x6d, 0x1c, 0xa2, 0x62, 0x09,
	0xdd, 0xbc, 0x30, 0x05, 0xe3, 0xbc, 0xf2, 0xce, 0xf0, 0x6b, 0x7e, 0xa5, 0x2f, 0x6e, 0xcd, 0x17,
	0xeb, 0x41, 0xba, 0x75, 0xd1, 0x13, 0xbb, 0xa1, 0x45, 0xe5, 0xd6, 0xfc, 0xc7, 0xb7, 0x2f, 0x70,
	0xa2, 0xfe, 0xda, 0xa7, 0xf3, 0x93, 0x2e, 0x19, 0xe8, 0x6f, 0xb3, 0x80, 0x9c, 0xcd, 0x02, 0xf2,
	0x7b, 0x16, 0x90, 0x2f, 0xf3, 0xa0, 0x76, 0x36, 0x0f, 0x6a, 0x3f, 0xe7, 0x41, 0x8d, 0x7a, 0x1a,
	0xaa, 0x9b, 0xbe, 0x22, 0x6f, 0xf7, 0x95, 0xb6, 0x47, 0xd3, 0x21, 0x1f, 0xc1, 0x44, 0x2c, 0x99,
	0x9e, 0x86, 0x95, 0x48, 0x1c, 0xff, 0x7f, 0x80, 0x36, 0x4b, 0x24, 0x0e, 0xef, 0xe6, 0xd7, 0x69,
	0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xf4, 0x1d, 0xca, 0xa3, 0x03, 0x00, 0x00,
}

func (this *MsgCreateTriggerRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateTriggerRequest)
	if !ok {
		that2, ok := that.(MsgCreateTriggerRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Authorities) != len(that1.Authorities) {
		return false
	}
	for i := range this.Authorities {
		if this.Authorities[i] != that1.Authorities[i] {
			return false
		}
	}
	if !this.Event.Equal(that1.Event) {
		return false
	}
	if len(this.Actions) != len(that1.Actions) {
		return false
	}
	for i := range this.Actions {
		if !this.Actions[i].Equal(that1.Actions[i]) {
			return false
		}
	}
	return true
}
func (this *MsgDestroyTriggerRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgDestroyTriggerRequest)
	if !ok {
		that2, ok := that.(MsgDestroyTriggerRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Authority != that1.Authority {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateTrigger is the RPC endpoint for creating a trigger
	CreateTrigger(ctx context.Context, in *MsgCreateTriggerRequest, opts ...grpc.CallOption) (*MsgCreateTriggerResponse, error)
	// DestroyTrigger is the RPC endpoint for creating a trigger
	DestroyTrigger(ctx context.Context, in *MsgDestroyTriggerRequest, opts ...grpc.CallOption) (*MsgDestroyTriggerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateTrigger(ctx context.Context, in *MsgCreateTriggerRequest, opts ...grpc.CallOption) (*MsgCreateTriggerResponse, error) {
	out := new(MsgCreateTriggerResponse)
	err := c.cc.Invoke(ctx, "/provenance.trigger.v1.Msg/CreateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DestroyTrigger(ctx context.Context, in *MsgDestroyTriggerRequest, opts ...grpc.CallOption) (*MsgDestroyTriggerResponse, error) {
	out := new(MsgDestroyTriggerResponse)
	err := c.cc.Invoke(ctx, "/provenance.trigger.v1.Msg/DestroyTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateTrigger is the RPC endpoint for creating a trigger
	CreateTrigger(context.Context, *MsgCreateTriggerRequest) (*MsgCreateTriggerResponse, error)
	// DestroyTrigger is the RPC endpoint for creating a trigger
	DestroyTrigger(context.Context, *MsgDestroyTriggerRequest) (*MsgDestroyTriggerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateTrigger(ctx context.Context, req *MsgCreateTriggerRequest) (*MsgCreateTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrigger not implemented")
}
func (*UnimplementedMsgServer) DestroyTrigger(ctx context.Context, req *MsgDestroyTriggerRequest) (*MsgDestroyTriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyTrigger not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.trigger.v1.Msg/CreateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTrigger(ctx, req.(*MsgCreateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DestroyTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDestroyTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DestroyTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.trigger.v1.Msg/DestroyTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DestroyTrigger(ctx, req.(*MsgDestroyTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.trigger.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrigger",
			Handler:    _Msg_CreateTrigger_Handler,
		},
		{
			MethodName: "DestroyTrigger",
			Handler:    _Msg_DestroyTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/trigger/v1/tx.proto",
}

func (m *MsgCreateTriggerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTriggerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTriggerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authorities) > 0 {
		for iNdEx := len(m.Authorities) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authorities[iNdEx])
			copy(dAtA[i:], m.Authorities[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Authorities[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateTriggerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTriggerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTriggerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDestroyTriggerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDestroyTriggerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDestroyTriggerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDestroyTriggerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDestroyTriggerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDestroyTriggerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateTriggerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Authorities) > 0 {
		for _, s := range m.Authorities {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateTriggerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDestroyTriggerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDestroyTriggerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateTriggerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateTriggerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTriggerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authorities = append(m.Authorities, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &types.Any{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateTriggerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateTriggerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTriggerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDestroyTriggerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDestroyTriggerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDestroyTriggerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDestroyTriggerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDestroyTriggerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDestroyTriggerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
