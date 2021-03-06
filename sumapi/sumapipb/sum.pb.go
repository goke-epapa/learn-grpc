// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sumapi/sumapipb/sum.proto

package sumapipb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Calculation struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculation) Reset()         { *m = Calculation{} }
func (m *Calculation) String() string { return proto.CompactTextString(m) }
func (*Calculation) ProtoMessage()    {}
func (*Calculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0485bd9a3be20f7f, []int{0}
}

func (m *Calculation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculation.Unmarshal(m, b)
}
func (m *Calculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculation.Marshal(b, m, deterministic)
}
func (m *Calculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculation.Merge(m, src)
}
func (m *Calculation) XXX_Size() int {
	return xxx_messageInfo_Calculation.Size(m)
}
func (m *Calculation) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculation.DiscardUnknown(m)
}

var xxx_messageInfo_Calculation proto.InternalMessageInfo

func (m *Calculation) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Calculation) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumRequest struct {
	Calculation          *Calculation `protobuf:"bytes,1,opt,name=calculation,proto3" json:"calculation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0485bd9a3be20f7f, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetCalculation() *Calculation {
	if m != nil {
		return m.Calculation
	}
	return nil
}

type SumResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0485bd9a3be20f7f, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculation)(nil), "sumapi.Calculation")
	proto.RegisterType((*SumRequest)(nil), "sumapi.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sumapi.SumResponse")
}

func init() { proto.RegisterFile("sumapi/sumapipb/sum.proto", fileDescriptor_0485bd9a3be20f7f) }

var fileDescriptor_0485bd9a3be20f7f = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0x3d, 0x0f, 0x17, 0x99, 0x9c, 0xcd, 0x1c, 0x88, 0x5a, 0x69, 0x44, 0xb0, 0x5a, 0xe5,
	0xc4, 0xd6, 0xc2, 0xeb, 0x2d, 0xb2, 0xd8, 0xd8, 0xc8, 0x6e, 0x1c, 0x21, 0xb0, 0xf9, 0x30, 0x1f,
	0xfe, 0x7e, 0x71, 0xe2, 0xb2, 0x6b, 0x95, 0xcc, 0xc3, 0xcb, 0xf3, 0x0e, 0x03, 0xe7, 0xa9, 0xd8,
	0x3e, 0x98, 0xbb, 0xfa, 0x84, 0xe1, 0xf7, 0xd3, 0x86, 0xe8, 0xb3, 0xc7, 0xa6, 0x32, 0xf9, 0x0a,
	0x62, 0xdf, 0x8f, 0xba, 0x8c, 0x7d, 0x36, 0xde, 0xe1, 0x15, 0x6c, 0x3e, 0x4d, 0x4c, 0xf9, 0xdd,
	0x15, 0x3b, 0x50, 0x3c, 0x5b, 0x5d, 0xae, 0x6e, 0x8f, 0x94, 0x60, 0xf6, 0xc2, 0x08, 0xaf, 0xe1,
	0x24, 0x91, 0xf6, 0xee, 0x63, 0xca, 0x1c, 0x72, 0x66, 0x53, 0x61, 0x0d, 0xc9, 0x3d, 0x40, 0x57,
	0xac, 0xa2, 0xaf, 0x42, 0x29, 0xe3, 0x23, 0x08, 0x3d, 0x97, 0xb0, 0x54, 0xec, 0xb6, 0x6d, 0x5d,
	0xa1, 0x5d, 0xf4, 0xab, 0x65, 0x4e, 0xde, 0x80, 0x60, 0x49, 0x0a, 0xde, 0x25, 0xc2, 0x53, 0x68,
	0x22, 0xa5, 0x32, 0x66, 0x16, 0xac, 0xd5, 0xdf, 0xb4, 0x7b, 0xe2, 0xae, 0x8e, 0xe2, 0xb7, 0xd1,
	0x84, 0xf7, 0xb0, 0xee, 0x8a, 0x45, 0x9c, 0xec, 0xf3, 0x1a, 0x17, 0xdb, 0x7f, 0xac, 0x5a, 0xe5,
	0xc1, 0x33, 0xbc, 0x1d, 0x4f, 0x07, 0x1a, 0x1a, 0xbe, 0xce, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0x4a, 0xdc, 0x8b, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sumapi.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumapi.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumapi.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sumapi/sumapipb/sum.proto",
}
