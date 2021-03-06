// Code generated by protoc-gen-go. DO NOT EDIT.
// source: library.proto

/*
Package api_pb is a generated protocol buffer package.

It is generated from these files:
	library.proto

It has these top-level messages:
	Book
	ListBooksRequest
	ListBooksResponse
	GetBookRequest
	CreateBookRequest
	UpdateBookRequest
	DeleteBookRequest
*/
package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type Book struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Book) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type ListBooksRequest struct {
}

func (m *ListBooksRequest) Reset()                    { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()               {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListBooksResponse struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *ListBooksResponse) Reset()                    { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()               {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type CreateBookRequest struct {
	Book *Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
}

func (m *CreateBookRequest) Reset()                    { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()               {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type UpdateBookRequest struct {
	Book *Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
}

func (m *UpdateBookRequest) Reset()                    { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()               {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type DeleteBookRequest struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *DeleteBookRequest) Reset()                    { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()               {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "com.github.izumin5210.grapi.testing.api.Book")
	proto.RegisterType((*ListBooksRequest)(nil), "com.github.izumin5210.grapi.testing.api.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "com.github.izumin5210.grapi.testing.api.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "com.github.izumin5210.grapi.testing.api.GetBookRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "com.github.izumin5210.grapi.testing.api.CreateBookRequest")
	proto.RegisterType((*UpdateBookRequest)(nil), "com.github.izumin5210.grapi.testing.api.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "com.github.izumin5210.grapi.testing.api.DeleteBookRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LibraryService service

type LibraryServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type libraryServiceClient struct {
	cc *grpc.ClientConn
}

func NewLibraryServiceClient(cc *grpc.ClientConn) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := grpc.Invoke(ctx, "/com.github.izumin5210.grapi.testing.api.LibraryService/ListBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.izumin5210.grapi.testing.api.LibraryService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.izumin5210.grapi.testing.api.LibraryService/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.izumin5210.grapi.testing.api.LibraryService/UpdateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/com.github.izumin5210.grapi.testing.api.LibraryService/DeleteBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LibraryService service

type LibraryServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*google_protobuf1.Empty, error)
}

func RegisterLibraryServiceServer(s *grpc.Server, srv LibraryServiceServer) {
	s.RegisterService(&_LibraryService_serviceDesc, srv)
}

func _LibraryService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.izumin5210.grapi.testing.api.LibraryService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.izumin5210.grapi.testing.api.LibraryService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.izumin5210.grapi.testing.api.LibraryService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.izumin5210.grapi.testing.api.LibraryService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.izumin5210.grapi.testing.api.LibraryService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LibraryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.izumin5210.grapi.testing.api.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _LibraryService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _LibraryService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _LibraryService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}

func init() { proto.RegisterFile("library.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x1c, 0xc5, 0x89, 0xb6, 0x69, 0xfb, 0x15, 0x43, 0x3b, 0xf8, 0xa3, 0x44, 0xc5, 0x12, 0x0f, 0x56,
	0xd0, 0x89, 0x46, 0x44, 0xec, 0xcd, 0x56, 0x11, 0xa1, 0xa7, 0x8a, 0x22, 0x5e, 0xca, 0xa4, 0x1d,
	0xe3, 0xd0, 0x36, 0x33, 0x66, 0x26, 0x82, 0x15, 0x2f, 0x7b, 0x5d, 0xf6, 0xb2, 0x7b, 0xdd, 0xff,
	0x6a, 0xef, 0x7b, 0xda, 0x3f, 0x64, 0xc9, 0x24, 0xdd, 0xb6, 0x1b, 0x0a, 0xc9, 0xb2, 0xc7, 0x64,
	0xe6, 0xfb, 0xe6, 0x33, 0xef, 0x3d, 0x06, 0x6e, 0xcf, 0x99, 0x1f, 0x91, 0xe8, 0x2f, 0x16, 0x11,
	0x57, 0x1c, 0x3d, 0x9d, 0xf0, 0x05, 0x0e, 0x98, 0xfa, 0x15, 0xfb, 0x98, 0x2d, 0xe3, 0x05, 0x0b,
	0xdf, 0x78, 0xaf, 0x5e, 0xe2, 0x20, 0x22, 0x82, 0x61, 0x45, 0xa5, 0x62, 0x61, 0x80, 0x89, 0x60,
	0xf6, 0xc3, 0x80, 0xf3, 0x60, 0x4e, 0x5d, 0x22, 0x98, 0x4b, 0xc2, 0x90, 0x2b, 0xa2, 0x18, 0x0f,
	0x65, 0x2a, 0x63, 0x3f, 0xc8, 0x56, 0xf5, 0x97, 0x1f, 0xff, 0x74, 0xe9, 0x42, 0xa8, 0xec, 0x0c,
	0xe7, 0x31, 0x54, 0xfa, 0x9c, 0xcf, 0xd0, 0x7d, 0xa8, 0xf9, 0x9c, 0xcf, 0xc6, 0x6c, 0xda, 0x36,
	0x3a, 0x46, 0xb7, 0x31, 0x32, 0x93, 0xcf, 0xcf, 0x53, 0x07, 0x41, 0x73, 0xc8, 0xa4, 0x4a, 0x36,
	0xc9, 0x11, 0xfd, 0x1d, 0x53, 0xa9, 0x9c, 0xef, 0xd0, 0xda, 0xf8, 0x27, 0x05, 0x0f, 0x25, 0x45,
	0x03, 0xa8, 0x26, 0x23, 0xb2, 0x6d, 0x74, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x70, 0x41, 0x7a, 0x9c,
	0xc8, 0x8c, 0xd2, 0x59, 0xe7, 0x19, 0x58, 0x9f, 0xa8, 0x16, 0xce, 0xce, 0xda, 0x0d, 0xf6, 0x0d,
	0x5a, 0x83, 0x88, 0x12, 0x45, 0x37, 0x77, 0xbf, 0x87, 0x4a, 0xb2, 0xac, 0xb7, 0x96, 0x66, 0xd0,
	0xa3, 0x89, 0xee, 0x57, 0x31, 0xbd, 0x7e, 0xdd, 0xe7, 0xd0, 0xfa, 0x40, 0xe7, 0x74, 0x5b, 0x77,
	0xd7, 0xed, 0xbc, 0xd3, 0x2a, 0x58, 0xc3, 0xb4, 0x0d, 0x5f, 0x68, 0xf4, 0x87, 0x4d, 0x28, 0x3a,
	0x34, 0xa0, 0x71, 0x61, 0x3b, 0x7a, 0x57, 0x98, 0xe1, 0x72, 0x7c, 0x76, 0xef, 0x2a, 0xa3, 0x69,
	0xca, 0x8e, 0xb5, 0x77, 0x72, 0x76, 0x74, 0xa3, 0x8e, 0x4c, 0x57, 0x07, 0x86, 0xf6, 0x0d, 0xa8,
	0x65, 0x89, 0xa1, 0xb7, 0x85, 0x75, 0xb7, 0x33, 0xb6, 0xcb, 0xf9, 0xe9, 0xb4, 0x35, 0x03, 0x42,
	0xcd, 0x94, 0xc1, 0xfd, 0x97, 0x59, 0xf8, 0x1f, 0x1d, 0x18, 0x00, 0xeb, 0x52, 0xa0, 0xe2, 0x17,
	0xcd, 0x35, 0xa9, 0x2c, 0xd3, 0x1d, 0xcd, 0x64, 0x39, 0x99, 0x2f, 0x3d, 0x9d, 0x39, 0x3a, 0x36,
	0x00, 0xd6, 0x65, 0x2a, 0xc1, 0x93, 0x6b, 0x60, 0x59, 0x9e, 0x27, 0x9a, 0xe7, 0x91, 0x77, 0x77,
	0xd3, 0x23, 0xbc, 0x32, 0x2a, 0xc3, 0x5b, 0x02, 0xac, 0x2b, 0x59, 0x82, 0x2e, 0xd7, 0x63, 0xfb,
	0x1e, 0x4e, 0x1f, 0x19, 0xbc, 0x7a, 0x64, 0xf0, 0xc7, 0xe4, 0x91, 0x59, 0x45, 0xe5, 0xe5, 0xa2,
	0xea, 0xd7, 0x7f, 0x98, 0x44, 0xb0, 0xb1, 0xf0, 0x7d, 0x53, 0xcf, 0xbc, 0x3e, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0xe2, 0x36, 0x91, 0xfe, 0x04, 0x00, 0x00,
}
