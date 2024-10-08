// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transcribe

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EnglishTranscriptionServiceClient is the client API for EnglishTranscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnglishTranscriptionServiceClient interface {
	TranscribeAudio(ctx context.Context, in *TranscriptionRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error)
}

type englishTranscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnglishTranscriptionServiceClient(cc grpc.ClientConnInterface) EnglishTranscriptionServiceClient {
	return &englishTranscriptionServiceClient{cc}
}

func (c *englishTranscriptionServiceClient) TranscribeAudio(ctx context.Context, in *TranscriptionRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error) {
	out := new(TranscriptionResponse)
	err := c.cc.Invoke(ctx, "/transcription.EnglishTranscriptionService/TranscribeAudio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnglishTranscriptionServiceServer is the server API for EnglishTranscriptionService service.
// All implementations must embed UnimplementedEnglishTranscriptionServiceServer
// for forward compatibility
type EnglishTranscriptionServiceServer interface {
	TranscribeAudio(context.Context, *TranscriptionRequest) (*TranscriptionResponse, error)
	mustEmbedUnimplementedEnglishTranscriptionServiceServer()
}

// UnimplementedEnglishTranscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnglishTranscriptionServiceServer struct {
}

func (UnimplementedEnglishTranscriptionServiceServer) TranscribeAudio(context.Context, *TranscriptionRequest) (*TranscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscribeAudio not implemented")
}
func (UnimplementedEnglishTranscriptionServiceServer) mustEmbedUnimplementedEnglishTranscriptionServiceServer() {
}

// UnsafeEnglishTranscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnglishTranscriptionServiceServer will
// result in compilation errors.
type UnsafeEnglishTranscriptionServiceServer interface {
	mustEmbedUnimplementedEnglishTranscriptionServiceServer()
}

func RegisterEnglishTranscriptionServiceServer(s *grpc.Server, srv EnglishTranscriptionServiceServer) {
	s.RegisterService(&_EnglishTranscriptionService_serviceDesc, srv)
}

func _EnglishTranscriptionService_TranscribeAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnglishTranscriptionServiceServer).TranscribeAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transcription.EnglishTranscriptionService/TranscribeAudio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnglishTranscriptionServiceServer).TranscribeAudio(ctx, req.(*TranscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnglishTranscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transcription.EnglishTranscriptionService",
	HandlerType: (*EnglishTranscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranscribeAudio",
			Handler:    _EnglishTranscriptionService_TranscribeAudio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transcribe.proto",
}

// IndonesianTranscriptionServiceClient is the client API for IndonesianTranscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndonesianTranscriptionServiceClient interface {
	TranscribeAudio(ctx context.Context, in *TranscriptionRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error)
}

type indonesianTranscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndonesianTranscriptionServiceClient(cc grpc.ClientConnInterface) IndonesianTranscriptionServiceClient {
	return &indonesianTranscriptionServiceClient{cc}
}

func (c *indonesianTranscriptionServiceClient) TranscribeAudio(ctx context.Context, in *TranscriptionRequest, opts ...grpc.CallOption) (*TranscriptionResponse, error) {
	out := new(TranscriptionResponse)
	err := c.cc.Invoke(ctx, "/transcription.IndonesianTranscriptionService/TranscribeAudio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndonesianTranscriptionServiceServer is the server API for IndonesianTranscriptionService service.
// All implementations must embed UnimplementedIndonesianTranscriptionServiceServer
// for forward compatibility
type IndonesianTranscriptionServiceServer interface {
	TranscribeAudio(context.Context, *TranscriptionRequest) (*TranscriptionResponse, error)
	mustEmbedUnimplementedIndonesianTranscriptionServiceServer()
}

// UnimplementedIndonesianTranscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndonesianTranscriptionServiceServer struct {
}

func (UnimplementedIndonesianTranscriptionServiceServer) TranscribeAudio(context.Context, *TranscriptionRequest) (*TranscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscribeAudio not implemented")
}
func (UnimplementedIndonesianTranscriptionServiceServer) mustEmbedUnimplementedIndonesianTranscriptionServiceServer() {
}

// UnsafeIndonesianTranscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndonesianTranscriptionServiceServer will
// result in compilation errors.
type UnsafeIndonesianTranscriptionServiceServer interface {
	mustEmbedUnimplementedIndonesianTranscriptionServiceServer()
}

func RegisterIndonesianTranscriptionServiceServer(s *grpc.Server, srv IndonesianTranscriptionServiceServer) {
	s.RegisterService(&_IndonesianTranscriptionService_serviceDesc, srv)
}

func _IndonesianTranscriptionService_TranscribeAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndonesianTranscriptionServiceServer).TranscribeAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transcription.IndonesianTranscriptionService/TranscribeAudio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndonesianTranscriptionServiceServer).TranscribeAudio(ctx, req.(*TranscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndonesianTranscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transcription.IndonesianTranscriptionService",
	HandlerType: (*IndonesianTranscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranscribeAudio",
			Handler:    _IndonesianTranscriptionService_TranscribeAudio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transcribe.proto",
}
