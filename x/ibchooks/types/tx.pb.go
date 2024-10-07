// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/ibchooks/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// MsgEmitIBCAck is the IBC Acknowledgement
type MsgEmitIBCAck struct {
	Sender         string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	PacketSequence uint64 `protobuf:"varint,2,opt,name=packet_sequence,json=packetSequence,proto3" json:"packet_sequence,omitempty"`
	Channel        string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (m *MsgEmitIBCAck) Reset()         { *m = MsgEmitIBCAck{} }
func (m *MsgEmitIBCAck) String() string { return proto.CompactTextString(m) }
func (*MsgEmitIBCAck) ProtoMessage()    {}
func (*MsgEmitIBCAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b3b50a1b9e4ff4, []int{0}
}
func (m *MsgEmitIBCAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEmitIBCAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEmitIBCAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEmitIBCAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEmitIBCAck.Merge(m, src)
}
func (m *MsgEmitIBCAck) XXX_Size() int {
	return m.Size()
}
func (m *MsgEmitIBCAck) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEmitIBCAck.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEmitIBCAck proto.InternalMessageInfo

func (m *MsgEmitIBCAck) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgEmitIBCAck) GetPacketSequence() uint64 {
	if m != nil {
		return m.PacketSequence
	}
	return 0
}

func (m *MsgEmitIBCAck) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

// MsgEmitIBCAckResponse is the IBC Acknowledgement response
type MsgEmitIBCAckResponse struct {
	ContractResult string `protobuf:"bytes,1,opt,name=contract_result,json=contractResult,proto3" json:"contract_result,omitempty"`
	IbcAck         string `protobuf:"bytes,2,opt,name=ibc_ack,json=ibcAck,proto3" json:"ibc_ack,omitempty"`
}

func (m *MsgEmitIBCAckResponse) Reset()         { *m = MsgEmitIBCAckResponse{} }
func (m *MsgEmitIBCAckResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEmitIBCAckResponse) ProtoMessage()    {}
func (*MsgEmitIBCAckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b3b50a1b9e4ff4, []int{1}
}
func (m *MsgEmitIBCAckResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEmitIBCAckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEmitIBCAckResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEmitIBCAckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEmitIBCAckResponse.Merge(m, src)
}
func (m *MsgEmitIBCAckResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEmitIBCAckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEmitIBCAckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEmitIBCAckResponse proto.InternalMessageInfo

func (m *MsgEmitIBCAckResponse) GetContractResult() string {
	if m != nil {
		return m.ContractResult
	}
	return ""
}

func (m *MsgEmitIBCAckResponse) GetIbcAck() string {
	if m != nil {
		return m.IbcAck
	}
	return ""
}

// MsgUpdateParamsRequest is a request message for the UpdateParams endpoint.
type MsgUpdateParamsRequest struct {
	// authority should be the governance module account address.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params are the new param values to set.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParamsRequest) Reset()         { *m = MsgUpdateParamsRequest{} }
func (m *MsgUpdateParamsRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsRequest) ProtoMessage()    {}
func (*MsgUpdateParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b3b50a1b9e4ff4, []int{2}
}
func (m *MsgUpdateParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsRequest.Merge(m, src)
}
func (m *MsgUpdateParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsRequest proto.InternalMessageInfo

func (m *MsgUpdateParamsRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParamsRequest) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse is a response message for the UpdateParams endpoint.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b3b50a1b9e4ff4, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEmitIBCAck)(nil), "provenance.ibchooks.v1.MsgEmitIBCAck")
	proto.RegisterType((*MsgEmitIBCAckResponse)(nil), "provenance.ibchooks.v1.MsgEmitIBCAckResponse")
	proto.RegisterType((*MsgUpdateParamsRequest)(nil), "provenance.ibchooks.v1.MsgUpdateParamsRequest")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "provenance.ibchooks.v1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("provenance/ibchooks/v1/tx.proto", fileDescriptor_f2b3b50a1b9e4ff4) }

var fileDescriptor_f2b3b50a1b9e4ff4 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6e, 0x13, 0x31,
	0x18, 0x8d, 0x69, 0x49, 0x15, 0x17, 0x82, 0x64, 0x95, 0xfc, 0xcc, 0x62, 0x5a, 0x05, 0x21, 0xaa,
	0x4a, 0x99, 0x51, 0x0b, 0x62, 0x51, 0xb1, 0x49, 0x10, 0x0b, 0x16, 0x91, 0xaa, 0xa9, 0x58, 0xc0,
	0x26, 0xf2, 0x38, 0x96, 0x63, 0x4d, 0xc7, 0x1e, 0x6c, 0x27, 0xb4, 0x3b, 0xc4, 0x09, 0x38, 0x01,
	0x67, 0xe8, 0x82, 0x43, 0x74, 0x59, 0xb1, 0x62, 0x85, 0x20, 0x59, 0xf4, 0x1a, 0x68, 0xc6, 0x0e,
	0x93, 0x8a, 0xa6, 0xea, 0xce, 0xdf, 0xf3, 0xfb, 0xbe, 0xf7, 0xbe, 0x67, 0xc3, 0xed, 0x4c, 0xc9,
	0x29, 0x15, 0x58, 0x10, 0x1a, 0xf2, 0x98, 0x8c, 0xa5, 0x4c, 0x74, 0x38, 0xdd, 0x0f, 0xcd, 0x69,
	0x90, 0x29, 0x69, 0x24, 0x6a, 0x94, 0x84, 0x60, 0x41, 0x08, 0xa6, 0xfb, 0x5e, 0x93, 0x48, 0x9d,
	0x4a, 0x1d, 0xa6, 0x9a, 0xe5, 0xfc, 0x54, 0x33, 0xdb, 0xe0, 0xb5, 0xed, 0xc5, 0xb0, 0xa8, 0x42,
	0x5b, 0xb8, 0xab, 0x2d, 0x26, 0x99, 0xb4, 0x78, 0x7e, 0x72, 0xe8, 0x93, 0x15, 0x16, 0x32, 0xac,
	0x70, 0xea, 0x5a, 0x3b, 0x9f, 0xe0, 0xc3, 0x81, 0x66, 0x6f, 0x52, 0x6e, 0xde, 0xf6, 0x5f, 0xf7,
	0x48, 0x82, 0x1a, 0xb0, 0xaa, 0xa9, 0x18, 0x51, 0xd5, 0x02, 0x3b, 0x60, 0xb7, 0x16, 0xb9, 0x0a,
	0x3d, 0x83, 0x8f, 0x32, 0x4c, 0x12, 0x6a, 0x86, 0x9a, 0x7e, 0x9c, 0x50, 0x41, 0x68, 0xeb, 0xde,
	0x0e, 0xd8, 0x5d, 0x8f, 0xea, 0x16, 0x3e, 0x76, 0x28, 0x6a, 0xc1, 0x0d, 0x32, 0xc6, 0x42, 0xd0,
	0x93, 0xd6, 0x5a, 0x31, 0x61, 0x51, 0x1e, 0x6e, 0x7e, 0xb9, 0x3a, 0xdf, 0x73, 0xf3, 0x3a, 0xef,
	0xe1, 0xe3, 0x6b, 0xc2, 0x11, 0xd5, 0x99, 0x14, 0x9a, 0xe6, 0x42, 0x44, 0x0a, 0xa3, 0x30, 0x31,
	0x43, 0x45, 0xf5, 0xe4, 0xc4, 0x38, 0x27, 0xf5, 0x05, 0x1c, 0x15, 0x28, 0x6a, 0xc2, 0x0d, 0x1e,
	0x93, 0x21, 0x26, 0x49, 0xe1, 0xa4, 0x16, 0x55, 0x79, 0x4c, 0x7a, 0x24, 0xe9, 0x7c, 0x03, 0xb0,
	0x31, 0xd0, 0xec, 0x5d, 0x36, 0xc2, 0x86, 0x1e, 0x15, 0xdb, 0x46, 0xb9, 0x3b, 0x6d, 0xd0, 0x4b,
	0x58, 0xc3, 0x13, 0x33, 0x96, 0x8a, 0x9b, 0x33, 0x3b, 0xb6, 0xdf, 0xfa, 0xf1, 0xbd, 0xbb, 0xe5,
	0xe2, 0xec, 0x8d, 0x46, 0x8a, 0x6a, 0x7d, 0x6c, 0x14, 0x17, 0x2c, 0x2a, 0xa9, 0xe8, 0x15, 0xac,
	0xda, 0xd8, 0x0a, 0xa9, 0xcd, 0x03, 0x3f, 0xb8, 0xf9, 0xf9, 0x02, 0x2b, 0xd7, 0x5f, 0xbf, 0xf8,
	0xb5, 0x5d, 0x89, 0x5c, 0xcf, 0x61, 0x3d, 0x5f, 0xbc, 0x9c, 0xd6, 0x69, 0xc3, 0xe6, 0x7f, 0xfe,
	0xec, 0xf6, 0x07, 0x7f, 0x00, 0x5c, 0x1b, 0x68, 0x86, 0x62, 0x08, 0x97, 0x1e, 0xe5, 0xe9, 0x2a,
	0xb9, 0x6b, 0x11, 0x7a, 0xdd, 0x3b, 0xd1, 0xfe, 0x25, 0x2d, 0xe1, 0x83, 0x65, 0x0f, 0x28, 0xb8,
	0xa5, 0xfd, 0x86, 0x30, 0xbd, 0xf0, 0xce, 0x7c, 0x2b, 0xe8, 0xdd, 0xff, 0x7c, 0x75, 0xbe, 0x07,
	0xfa, 0xc9, 0xc5, 0xcc, 0x07, 0x97, 0x33, 0x1f, 0xfc, 0x9e, 0xf9, 0xe0, 0xeb, 0xdc, 0xaf, 0x5c,
	0xce, 0xfd, 0xca, 0xcf, 0xb9, 0x5f, 0x81, 0x6d, 0x2e, 0x57, 0xcc, 0x3c, 0x02, 0x1f, 0x5e, 0x30,
	0x6e, 0xc6, 0x93, 0x38, 0x20, 0x32, 0x0d, 0x4b, 0x52, 0x97, 0xcb, 0xa5, 0x2a, 0x3c, 0x2d, 0xbf,
	0xba, 0x39, 0xcb, 0xa8, 0x8e, 0xab, 0xc5, 0x3f, 0x7f, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x6e, 0x18, 0xc3, 0x91, 0x03, 0x00, 0x00,
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
	// EmitIBCAck checks the sender can emit the ack and writes the IBC
	// acknowledgement
	EmitIBCAck(ctx context.Context, in *MsgEmitIBCAck, opts ...grpc.CallOption) (*MsgEmitIBCAckResponse, error)
	// UpdateParams is a governance proposal endpoint for updating the ibchooks module's params.
	UpdateParams(ctx context.Context, in *MsgUpdateParamsRequest, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EmitIBCAck(ctx context.Context, in *MsgEmitIBCAck, opts ...grpc.CallOption) (*MsgEmitIBCAckResponse, error) {
	out := new(MsgEmitIBCAckResponse)
	err := c.cc.Invoke(ctx, "/provenance.ibchooks.v1.Msg/EmitIBCAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParamsRequest, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/provenance.ibchooks.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// EmitIBCAck checks the sender can emit the ack and writes the IBC
	// acknowledgement
	EmitIBCAck(context.Context, *MsgEmitIBCAck) (*MsgEmitIBCAckResponse, error)
	// UpdateParams is a governance proposal endpoint for updating the ibchooks module's params.
	UpdateParams(context.Context, *MsgUpdateParamsRequest) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EmitIBCAck(ctx context.Context, req *MsgEmitIBCAck) (*MsgEmitIBCAckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitIBCAck not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParamsRequest) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EmitIBCAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEmitIBCAck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EmitIBCAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.ibchooks.v1.Msg/EmitIBCAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EmitIBCAck(ctx, req.(*MsgEmitIBCAck))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.ibchooks.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.ibchooks.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmitIBCAck",
			Handler:    _Msg_EmitIBCAck_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/ibchooks/v1/tx.proto",
}

func (m *MsgEmitIBCAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEmitIBCAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEmitIBCAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PacketSequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PacketSequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEmitIBCAckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEmitIBCAckResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEmitIBCAckResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcAck) > 0 {
		i -= len(m.IbcAck)
		copy(dAtA[i:], m.IbcAck)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IbcAck)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractResult) > 0 {
		i -= len(m.ContractResult)
		copy(dAtA[i:], m.ContractResult)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractResult)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgEmitIBCAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PacketSequence != 0 {
		n += 1 + sovTx(uint64(m.PacketSequence))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEmitIBCAckResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractResult)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IbcAck)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
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
func (m *MsgEmitIBCAck) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgEmitIBCAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEmitIBCAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSequence", wireType)
			}
			m.PacketSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
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
			m.Channel = string(dAtA[iNdEx:postIndex])
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
func (m *MsgEmitIBCAckResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgEmitIBCAckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEmitIBCAckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractResult", wireType)
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
			m.ContractResult = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcAck", wireType)
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
			m.IbcAck = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
