// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/v1/query.proto

package blog

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryAllPostsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostsRequest) Reset()         { *m = QueryAllPostsRequest{} }
func (m *QueryAllPostsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostsRequest) ProtoMessage()    {}
func (*QueryAllPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{0}
}
func (m *QueryAllPostsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostsRequest.Merge(m, src)
}
func (m *QueryAllPostsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostsRequest proto.InternalMessageInfo

func (m *QueryAllPostsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPostsResponse struct {
	Posts      []*Post             `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostsResponse) Reset()         { *m = QueryAllPostsResponse{} }
func (m *QueryAllPostsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostsResponse) ProtoMessage()    {}
func (*QueryAllPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{1}
}
func (m *QueryAllPostsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostsResponse.Merge(m, src)
}
func (m *QueryAllPostsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostsResponse proto.InternalMessageInfo

func (m *QueryAllPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *QueryAllPostsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryPostCommentsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Slug       string             `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (m *QueryPostCommentsRequest) Reset()         { *m = QueryPostCommentsRequest{} }
func (m *QueryPostCommentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPostCommentsRequest) ProtoMessage()    {}
func (*QueryPostCommentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{2}
}
func (m *QueryPostCommentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPostCommentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPostCommentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPostCommentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPostCommentsRequest.Merge(m, src)
}
func (m *QueryPostCommentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPostCommentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPostCommentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPostCommentsRequest proto.InternalMessageInfo

func (m *QueryPostCommentsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryPostCommentsRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type QueryPostCommentsResponse struct {
	Comments   []*Comment          `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryPostCommentsResponse) Reset()         { *m = QueryPostCommentsResponse{} }
func (m *QueryPostCommentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPostCommentsResponse) ProtoMessage()    {}
func (*QueryPostCommentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{3}
}
func (m *QueryPostCommentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPostCommentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPostCommentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPostCommentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPostCommentsResponse.Merge(m, src)
}
func (m *QueryPostCommentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPostCommentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPostCommentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPostCommentsResponse proto.InternalMessageInfo

func (m *QueryPostCommentsResponse) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *QueryPostCommentsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllPostsRequest)(nil), "blog.v1.QueryAllPostsRequest")
	proto.RegisterType((*QueryAllPostsResponse)(nil), "blog.v1.QueryAllPostsResponse")
	proto.RegisterType((*QueryPostCommentsRequest)(nil), "blog.v1.QueryPostCommentsRequest")
	proto.RegisterType((*QueryPostCommentsResponse)(nil), "blog.v1.QueryPostCommentsResponse")
}

func init() { proto.RegisterFile("blog/v1/query.proto", fileDescriptor_eb782a0cfb4fa324) }

var fileDescriptor_eb782a0cfb4fa324 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xee, 0xdc, 0x7b, 0x7b, 0xad, 0x53, 0x05, 0x19, 0x15, 0x6a, 0xc0, 0x50, 0x53, 0xd0, 0x22,
	0x3a, 0x43, 0xea, 0xda, 0x85, 0x0a, 0x8a, 0xbb, 0x1a, 0x70, 0xe3, 0x42, 0x48, 0xc2, 0x10, 0x8b,
	0x49, 0x26, 0xcd, 0x4c, 0xa2, 0x7d, 0x00, 0xf7, 0xe2, 0x5b, 0xf8, 0x26, 0x2e, 0xbb, 0x74, 0x29,
	0xed, 0x8b, 0xc8, 0x64, 0xc6, 0xd8, 0x4a, 0xd5, 0x4d, 0x77, 0xe1, 0xcc, 0x77, 0xbe, 0xbf, 0x1c,
	0xb8, 0xea, 0x85, 0x2c, 0x20, 0xb9, 0x4d, 0xfa, 0x19, 0x4d, 0x07, 0x38, 0x49, 0x99, 0x60, 0x68,
	0x41, 0x0e, 0x71, 0x6e, 0x1b, 0xbb, 0x3e, 0xe3, 0x11, 0xe3, 0xc4, 0x73, 0x39, 0x55, 0x08, 0x92,
	0xdb, 0x1e, 0x15, 0xae, 0x4d, 0x12, 0x37, 0xe8, 0xc5, 0xae, 0xe8, 0xb1, 0x58, 0x2d, 0x19, 0x25,
	0x93, 0x18, 0x24, 0x94, 0xab, 0xa1, 0x75, 0x0d, 0xd7, 0x2e, 0xe4, 0xda, 0x51, 0x18, 0x76, 0x19,
	0x17, 0xdc, 0xa1, 0xfd, 0x8c, 0x72, 0x81, 0x4e, 0x21, 0xfc, 0x24, 0x68, 0x80, 0x26, 0x68, 0xd7,
	0x3b, 0xdb, 0x58, 0xa9, 0x61, 0xa9, 0x86, 0x95, 0x1f, 0xad, 0x86, 0xbb, 0x6e, 0x40, 0xf5, 0xae,
	0x33, 0xb1, 0x69, 0x3d, 0x00, 0xb8, 0xfe, 0x45, 0x80, 0x27, 0x2c, 0xe6, 0x14, 0xb5, 0x60, 0x35,
	0x91, 0x83, 0x06, 0x68, 0xfe, 0x6d, 0xd7, 0x3b, 0xcb, 0x58, 0x67, 0xc2, 0x12, 0xe6, 0xa8, 0x37,
	0x74, 0x36, 0x65, 0xe3, 0x4f, 0x61, 0x63, 0xe7, 0x57, 0x1b, 0x4a, 0x61, 0xca, 0x47, 0x0e, 0x1b,
	0x85, 0x0d, 0x49, 0x7e, 0xc2, 0xa2, 0x88, 0xc6, 0x73, 0xcf, 0x8a, 0x10, 0xfc, 0xc7, 0xc3, 0x2c,
	0x28, 0x6c, 0x2e, 0x3a, 0xc5, 0xb7, 0xf5, 0x04, 0xe0, 0xc6, 0x0c, 0x61, 0xdd, 0xc1, 0x1e, 0xac,
	0xf9, 0x7a, 0xa6, 0x6b, 0x58, 0x29, 0x6b, 0xd0, 0x60, 0xa7, 0x44, 0xcc, 0xad, 0x8c, 0xce, 0x33,
	0x80, 0xd5, 0xc2, 0x14, 0x3a, 0x87, 0xb5, 0x8f, 0x1f, 0x83, 0x36, 0x4b, 0xe9, 0x59, 0x17, 0x61,
	0x98, 0xdf, 0x3d, 0xeb, 0x2c, 0x97, 0x70, 0x69, 0x32, 0x23, 0xda, 0x9a, 0xc6, 0xcf, 0x28, 0xde,
	0xb0, 0x7e, 0x82, 0x28, 0xda, 0xe3, 0xc3, 0x97, 0x91, 0x09, 0x86, 0x23, 0x13, 0xbc, 0x8d, 0x4c,
	0xf0, 0x38, 0x36, 0x2b, 0xc3, 0xb1, 0x59, 0x79, 0x1d, 0x9b, 0x95, 0xab, 0x56, 0xd0, 0x13, 0x37,
	0x99, 0x87, 0x7d, 0x16, 0x91, 0x94, 0x06, 0x34, 0xde, 0x8f, 0xa9, 0xb8, 0x63, 0xe9, 0x2d, 0xf1,
	0xa8, 0x4f, 0xee, 0x89, 0xe4, 0xf6, 0xfe, 0x17, 0x67, 0x7e, 0xf0, 0x1e, 0x00, 0x00, 0xff, 0xff,
	0xfb, 0x16, 0x8f, 0xb8, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	AllPosts(ctx context.Context, in *QueryAllPostsRequest, opts ...grpc.CallOption) (*QueryAllPostsResponse, error)
	PostComments(ctx context.Context, in *QueryPostCommentsRequest, opts ...grpc.CallOption) (*QueryPostCommentsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllPosts(ctx context.Context, in *QueryAllPostsRequest, opts ...grpc.CallOption) (*QueryAllPostsResponse, error) {
	out := new(QueryAllPostsResponse)
	err := c.cc.Invoke(ctx, "/blog.v1.Query/AllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PostComments(ctx context.Context, in *QueryPostCommentsRequest, opts ...grpc.CallOption) (*QueryPostCommentsResponse, error) {
	out := new(QueryPostCommentsResponse)
	err := c.cc.Invoke(ctx, "/blog.v1.Query/PostComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	AllPosts(context.Context, *QueryAllPostsRequest) (*QueryAllPostsResponse, error)
	PostComments(context.Context, *QueryPostCommentsRequest) (*QueryPostCommentsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AllPosts(ctx context.Context, req *QueryAllPostsRequest) (*QueryAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPosts not implemented")
}
func (*UnimplementedQueryServer) PostComments(ctx context.Context, req *QueryPostCommentsRequest) (*QueryPostCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostComments not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.Query/AllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllPosts(ctx, req.(*QueryAllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PostComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPostCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PostComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.Query/PostComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PostComments(ctx, req.(*QueryPostCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllPosts",
			Handler:    _Query_AllPosts_Handler,
		},
		{
			MethodName: "PostComments",
			Handler:    _Query_PostComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/v1/query.proto",
}

func (m *QueryAllPostsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPostsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Posts) > 0 {
		for iNdEx := len(m.Posts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Posts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryPostCommentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPostCommentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPostCommentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Slug) > 0 {
		i -= len(m.Slug)
		copy(dAtA[i:], m.Slug)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Slug)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPostCommentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPostCommentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPostCommentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Comments) > 0 {
		for iNdEx := len(m.Comments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Comments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllPostsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPostsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPostCommentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Slug)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPostCommentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllPostsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPostsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllPostsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPostsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Posts = append(m.Posts, &Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryPostCommentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPostCommentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPostCommentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryPostCommentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPostCommentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPostCommentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
