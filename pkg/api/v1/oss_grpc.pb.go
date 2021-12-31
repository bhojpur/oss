// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// OssServiceClient is the client API for OssService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OssServiceClient interface {
	// StartLocalNfvo starts an NFV Orchestrator on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the oss/config.yaml
	//   3. all bytes constituting the NFV Orchestraor YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalNfvo(ctx context.Context, opts ...grpc.CallOption) (OssService_StartLocalNfvoClient, error)
	// StartFromPreviousNfvo starts a new NFV Orchestrator based on a previous one.
	// If the previous NFV Orchestrator does not have the can-replay condition set this call will result in an error.
	StartFromPreviousNfvo(ctx context.Context, in *StartFromPreviousNfvoRequest, opts ...grpc.CallOption) (*StartNfvoResponse, error)
	// StartNfvoRequest starts a new NFV Orchestrator based on its specification.
	StartNfvo(ctx context.Context, in *StartNfvoRequest, opts ...grpc.CallOption) (*StartNfvoResponse, error)
	// Searches for NFV Orchestrator(s) known to this instance
	ListNfvos(ctx context.Context, in *ListNfvosRequest, opts ...grpc.CallOption) (*ListNfvosResponse, error)
	// Subscribe listens to new NFV Orchestrator(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (OssService_SubscribeClient, error)
	// GetNfvo retrieves details of a single NFV Orchestrator
	GetNfvo(ctx context.Context, in *GetNfvoRequest, opts ...grpc.CallOption) (*GetNfvoResponse, error)
	// Listen listens to NFV Orchestrator updates and log output of a running NFV Orchestrator
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (OssService_ListenClient, error)
	// StopNfvo stops a currently running NFV Orchestrator
	StopNfvo(ctx context.Context, in *StopNfvoRequest, opts ...grpc.CallOption) (*StopNfvoResponse, error)
}

type ossServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOssServiceClient(cc grpc.ClientConnInterface) OssServiceClient {
	return &ossServiceClient{cc}
}

func (c *ossServiceClient) StartLocalNfvo(ctx context.Context, opts ...grpc.CallOption) (OssService_StartLocalNfvoClient, error) {
	stream, err := c.cc.NewStream(ctx, &OssService_ServiceDesc.Streams[0], "/v1.OssService/StartLocalNfvo", opts...)
	if err != nil {
		return nil, err
	}
	x := &ossServiceStartLocalNfvoClient{stream}
	return x, nil
}

type OssService_StartLocalNfvoClient interface {
	Send(*StartLocalNfvoRequest) error
	CloseAndRecv() (*StartNfvoResponse, error)
	grpc.ClientStream
}

type ossServiceStartLocalNfvoClient struct {
	grpc.ClientStream
}

func (x *ossServiceStartLocalNfvoClient) Send(m *StartLocalNfvoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ossServiceStartLocalNfvoClient) CloseAndRecv() (*StartNfvoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartNfvoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ossServiceClient) StartFromPreviousNfvo(ctx context.Context, in *StartFromPreviousNfvoRequest, opts ...grpc.CallOption) (*StartNfvoResponse, error) {
	out := new(StartNfvoResponse)
	err := c.cc.Invoke(ctx, "/v1.OssService/StartFromPreviousNfvo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossServiceClient) StartNfvo(ctx context.Context, in *StartNfvoRequest, opts ...grpc.CallOption) (*StartNfvoResponse, error) {
	out := new(StartNfvoResponse)
	err := c.cc.Invoke(ctx, "/v1.OssService/StartNfvo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossServiceClient) ListNfvos(ctx context.Context, in *ListNfvosRequest, opts ...grpc.CallOption) (*ListNfvosResponse, error) {
	out := new(ListNfvosResponse)
	err := c.cc.Invoke(ctx, "/v1.OssService/ListNfvos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (OssService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &OssService_ServiceDesc.Streams[1], "/v1.OssService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &ossServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OssService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type ossServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *ossServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ossServiceClient) GetNfvo(ctx context.Context, in *GetNfvoRequest, opts ...grpc.CallOption) (*GetNfvoResponse, error) {
	out := new(GetNfvoResponse)
	err := c.cc.Invoke(ctx, "/v1.OssService/GetNfvo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (OssService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &OssService_ServiceDesc.Streams[2], "/v1.OssService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &ossServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OssService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type ossServiceListenClient struct {
	grpc.ClientStream
}

func (x *ossServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ossServiceClient) StopNfvo(ctx context.Context, in *StopNfvoRequest, opts ...grpc.CallOption) (*StopNfvoResponse, error) {
	out := new(StopNfvoResponse)
	err := c.cc.Invoke(ctx, "/v1.OssService/StopNfvo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OssServiceServer is the server API for OssService service.
// All implementations must embed UnimplementedOssServiceServer
// for forward compatibility
type OssServiceServer interface {
	// StartLocalNfvo starts an NFV Orchestrator on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the oss/config.yaml
	//   3. all bytes constituting the NFV Orchestraor YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalNfvo(OssService_StartLocalNfvoServer) error
	// StartFromPreviousNfvo starts a new NFV Orchestrator based on a previous one.
	// If the previous NFV Orchestrator does not have the can-replay condition set this call will result in an error.
	StartFromPreviousNfvo(context.Context, *StartFromPreviousNfvoRequest) (*StartNfvoResponse, error)
	// StartNfvoRequest starts a new NFV Orchestrator based on its specification.
	StartNfvo(context.Context, *StartNfvoRequest) (*StartNfvoResponse, error)
	// Searches for NFV Orchestrator(s) known to this instance
	ListNfvos(context.Context, *ListNfvosRequest) (*ListNfvosResponse, error)
	// Subscribe listens to new NFV Orchestrator(s) updates
	Subscribe(*SubscribeRequest, OssService_SubscribeServer) error
	// GetNfvo retrieves details of a single NFV Orchestrator
	GetNfvo(context.Context, *GetNfvoRequest) (*GetNfvoResponse, error)
	// Listen listens to NFV Orchestrator updates and log output of a running NFV Orchestrator
	Listen(*ListenRequest, OssService_ListenServer) error
	// StopNfvo stops a currently running NFV Orchestrator
	StopNfvo(context.Context, *StopNfvoRequest) (*StopNfvoResponse, error)
	mustEmbedUnimplementedOssServiceServer()
}

// UnimplementedOssServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOssServiceServer struct {
}

func (UnimplementedOssServiceServer) StartLocalNfvo(OssService_StartLocalNfvoServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalNfvo not implemented")
}
func (UnimplementedOssServiceServer) StartFromPreviousNfvo(context.Context, *StartFromPreviousNfvoRequest) (*StartNfvoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousNfvo not implemented")
}
func (UnimplementedOssServiceServer) StartNfvo(context.Context, *StartNfvoRequest) (*StartNfvoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNfvo not implemented")
}
func (UnimplementedOssServiceServer) ListNfvos(context.Context, *ListNfvosRequest) (*ListNfvosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNfvos not implemented")
}
func (UnimplementedOssServiceServer) Subscribe(*SubscribeRequest, OssService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedOssServiceServer) GetNfvo(context.Context, *GetNfvoRequest) (*GetNfvoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNfvo not implemented")
}
func (UnimplementedOssServiceServer) Listen(*ListenRequest, OssService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedOssServiceServer) StopNfvo(context.Context, *StopNfvoRequest) (*StopNfvoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopNfvo not implemented")
}
func (UnimplementedOssServiceServer) mustEmbedUnimplementedOssServiceServer() {}

// UnsafeOssServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OssServiceServer will
// result in compilation errors.
type UnsafeOssServiceServer interface {
	mustEmbedUnimplementedOssServiceServer()
}

func RegisterOssServiceServer(s grpc.ServiceRegistrar, srv OssServiceServer) {
	s.RegisterService(&OssService_ServiceDesc, srv)
}

func _OssService_StartLocalNfvo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OssServiceServer).StartLocalNfvo(&ossServiceStartLocalNfvoServer{stream})
}

type OssService_StartLocalNfvoServer interface {
	SendAndClose(*StartNfvoResponse) error
	Recv() (*StartLocalNfvoRequest, error)
	grpc.ServerStream
}

type ossServiceStartLocalNfvoServer struct {
	grpc.ServerStream
}

func (x *ossServiceStartLocalNfvoServer) SendAndClose(m *StartNfvoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ossServiceStartLocalNfvoServer) Recv() (*StartLocalNfvoRequest, error) {
	m := new(StartLocalNfvoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OssService_StartFromPreviousNfvo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousNfvoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).StartFromPreviousNfvo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OssService/StartFromPreviousNfvo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).StartFromPreviousNfvo(ctx, req.(*StartFromPreviousNfvoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssService_StartNfvo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNfvoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).StartNfvo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OssService/StartNfvo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).StartNfvo(ctx, req.(*StartNfvoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssService_ListNfvos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNfvosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).ListNfvos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OssService/ListNfvos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).ListNfvos(ctx, req.(*ListNfvosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OssServiceServer).Subscribe(m, &ossServiceSubscribeServer{stream})
}

type OssService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type ossServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *ossServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OssService_GetNfvo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNfvoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).GetNfvo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OssService/GetNfvo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).GetNfvo(ctx, req.(*GetNfvoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OssServiceServer).Listen(m, &ossServiceListenServer{stream})
}

type OssService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type ossServiceListenServer struct {
	grpc.ServerStream
}

func (x *ossServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OssService_StopNfvo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNfvoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServiceServer).StopNfvo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.OssService/StopNfvo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServiceServer).StopNfvo(ctx, req.(*StopNfvoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OssService_ServiceDesc is the grpc.ServiceDesc for OssService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OssService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OssService",
	HandlerType: (*OssServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousNfvo",
			Handler:    _OssService_StartFromPreviousNfvo_Handler,
		},
		{
			MethodName: "StartNfvo",
			Handler:    _OssService_StartNfvo_Handler,
		},
		{
			MethodName: "ListNfvos",
			Handler:    _OssService_ListNfvos_Handler,
		},
		{
			MethodName: "GetNfvo",
			Handler:    _OssService_GetNfvo_Handler,
		},
		{
			MethodName: "StopNfvo",
			Handler:    _OssService_StopNfvo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalNfvo",
			Handler:       _OssService_StartLocalNfvo_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _OssService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _OssService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "oss.proto",
}
