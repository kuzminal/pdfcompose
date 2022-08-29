// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/pdf_compose_service.proto

package pdfcompose

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PdfComposeServiceClient is the client API for PdfComposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PdfComposeServiceClient interface {
	Convert(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type pdfComposeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPdfComposeServiceClient(cc grpc.ClientConnInterface) PdfComposeServiceClient {
	return &pdfComposeServiceClient{cc}
}

func (c *pdfComposeServiceClient) Convert(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/service.PdfComposeService/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PdfComposeServiceServer is the server API for PdfComposeService service.
// All implementations must embed UnimplementedPdfComposeServiceServer
// for forward compatibility
type PdfComposeServiceServer interface {
	Convert(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPdfComposeServiceServer()
}

// UnimplementedPdfComposeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPdfComposeServiceServer struct {
}

func (UnimplementedPdfComposeServiceServer) Convert(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedPdfComposeServiceServer) mustEmbedUnimplementedPdfComposeServiceServer() {}

// UnsafePdfComposeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PdfComposeServiceServer will
// result in compilation errors.
type UnsafePdfComposeServiceServer interface {
	mustEmbedUnimplementedPdfComposeServiceServer()
}

func RegisterPdfComposeServiceServer(s grpc.ServiceRegistrar, srv PdfComposeServiceServer) {
	s.RegisterService(&PdfComposeService_ServiceDesc, srv)
}

func _PdfComposeService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PdfComposeServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PdfComposeService/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PdfComposeServiceServer).Convert(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PdfComposeService_ServiceDesc is the grpc.ServiceDesc for PdfComposeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PdfComposeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.PdfComposeService",
	HandlerType: (*PdfComposeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _PdfComposeService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pdf_compose_service.proto",
}