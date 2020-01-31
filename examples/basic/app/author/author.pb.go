// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author/author.proto

package author

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "graphql"
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

type Author struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListAuthorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAuthorsRequest) Reset()         { *m = ListAuthorsRequest{} }
func (m *ListAuthorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsRequest) ProtoMessage()    {}
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{1}
}

func (m *ListAuthorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsRequest.Unmarshal(m, b)
}
func (m *ListAuthorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsRequest.Marshal(b, m, deterministic)
}
func (m *ListAuthorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsRequest.Merge(m, src)
}
func (m *ListAuthorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsRequest.Size(m)
}
func (m *ListAuthorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsRequest proto.InternalMessageInfo

type ListAuthorsResponse struct {
	Authors              []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListAuthorsResponse) Reset()         { *m = ListAuthorsResponse{} }
func (m *ListAuthorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsResponse) ProtoMessage()    {}
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{2}
}

func (m *ListAuthorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsResponse.Unmarshal(m, b)
}
func (m *ListAuthorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsResponse.Marshal(b, m, deterministic)
}
func (m *ListAuthorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsResponse.Merge(m, src)
}
func (m *ListAuthorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsResponse.Size(m)
}
func (m *ListAuthorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsResponse proto.InternalMessageInfo

func (m *ListAuthorsResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

type GetAuthorRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorRequest) Reset()         { *m = GetAuthorRequest{} }
func (m *GetAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthorRequest) ProtoMessage()    {}
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{3}
}

func (m *GetAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorRequest.Unmarshal(m, b)
}
func (m *GetAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorRequest.Merge(m, src)
}
func (m *GetAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthorRequest.Size(m)
}
func (m *GetAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorRequest proto.InternalMessageInfo

func (m *GetAuthorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Author)(nil), "author.Author")
	proto.RegisterType((*ListAuthorsRequest)(nil), "author.ListAuthorsRequest")
	proto.RegisterType((*ListAuthorsResponse)(nil), "author.ListAuthorsResponse")
	proto.RegisterType((*GetAuthorRequest)(nil), "author.GetAuthorRequest")
}

func init() { proto.RegisterFile("author/author.proto", fileDescriptor_41efdaee09960382) }

var fileDescriptor_41efdaee09960382 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x95, 0xff, 0x47, 0x81, 0x4c, 0x01, 0x21, 0xb7, 0x48, 0x56, 0x60, 0x51, 0x79, 0x81,
	0xb2, 0x69, 0x2c, 0x95, 0x03, 0x20, 0xa8, 0x10, 0x1b, 0x56, 0x61, 0xc7, 0x0a, 0x27, 0x1a, 0x25,
	0x91, 0x9a, 0xda, 0xb5, 0x1d, 0xa0, 0x57, 0xe8, 0x7d, 0xca, 0xf9, 0x90, 0x62, 0x27, 0x40, 0x61,
	0x35, 0x93, 0xcc, 0xf3, 0xbc, 0xcf, 0x7e, 0x30, 0x16, 0xad, 0xad, 0xa4, 0xe6, 0xae, 0xa4, 0x4a,
	0x4b, 0x2b, 0x49, 0xe8, 0xbe, 0xe2, 0xf3, 0x52, 0x0b, 0x55, 0xad, 0x97, 0xdc, 0x57, 0x37, 0x66,
	0x97, 0x10, 0xde, 0x76, 0x02, 0x42, 0xe0, 0x60, 0x25, 0x1a, 0xa4, 0xff, 0xa6, 0x41, 0x12, 0x65,
	0x5d, 0xcf, 0x26, 0x40, 0x1e, 0x6b, 0x63, 0x9d, 0xc2, 0x64, 0xb8, 0x6e, 0xd1, 0x58, 0x76, 0x03,
	0xe3, 0x1f, 0x7f, 0x8d, 0x92, 0x2b, 0x83, 0x24, 0x81, 0x43, 0xe7, 0x65, 0x68, 0x30, 0xfd, 0x9f,
	0x8c, 0xe6, 0xa7, 0xa9, 0x27, 0x71, 0xca, 0xac, 0x1f, 0xb3, 0x2b, 0x38, 0x7b, 0x40, 0x7f, 0xde,
	0x2f, 0x1d, 0xec, 0x83, 0x2f, 0xfb, 0xf9, 0x47, 0x00, 0x27, 0x4e, 0xf5, 0x84, 0xfa, 0xb5, 0x2e,
	0x90, 0xbc, 0xc0, 0xe8, 0x9b, 0x35, 0x89, 0x7b, 0x87, 0xdf, 0x94, 0xf1, 0xc5, 0x9f, 0x33, 0xc7,
	0xca, 0xe8, 0x76, 0x47, 0x27, 0x03, 0x2f, 0x8b, 0xe2, 0xbe, 0x25, 0x0b, 0x88, 0x06, 0x36, 0x42,
	0xfb, 0x1d, 0xfb, 0xb8, 0xf1, 0xde, 0xdd, 0xd8, 0xf1, 0x76, 0x47, 0x8f, 0xc0, 0x3f, 0xf6, 0xdd,
	0xfd, 0xf3, 0xa2, 0xac, 0x6d, 0xd5, 0xe6, 0x69, 0x21, 0x1b, 0xbe, 0x31, 0x6d, 0x59, 0x37, 0xd2,
	0x4a, 0x5e, 0x6a, 0x55, 0xcc, 0x7c, 0x00, 0xb3, 0x52, 0x58, 0x7c, 0x13, 0x1b, 0x8e, 0xef, 0xa2,
	0x51, 0x4b, 0x34, 0x3c, 0x17, 0xa6, 0x2e, 0xb8, 0x50, 0xca, 0x27, 0x98, 0x87, 0x5d, 0x46, 0xd7,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x8c, 0xbc, 0xb3, 0xd9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorServiceClient interface {
	ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error)
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
}

type authorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorServiceClient(cc *grpc.ClientConn) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error) {
	out := new(ListAuthorsResponse)
	err := c.cc.Invoke(ctx, "/author.AuthorService/ListAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/author.AuthorService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
type AuthorServiceServer interface {
	ListAuthors(context.Context, *ListAuthorsRequest) (*ListAuthorsResponse, error)
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
}

// UnimplementedAuthorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (*UnimplementedAuthorServiceServer) ListAuthors(ctx context.Context, req *ListAuthorsRequest) (*ListAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthors not implemented")
}
func (*UnimplementedAuthorServiceServer) GetAuthor(ctx context.Context, req *GetAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}

func RegisterAuthorServiceServer(s *grpc.Server, srv AuthorServiceServer) {
	s.RegisterService(&_AuthorService_serviceDesc, srv)
}

func _AuthorService_ListAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).ListAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author.AuthorService/ListAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).ListAuthors(ctx, req.(*ListAuthorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author.AuthorService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "author.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuthors",
			Handler:    _AuthorService_ListAuthors_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorService_GetAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author/author.proto",
}