// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: ds_api.proto

package pb

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

// DataServClient is the client API for DataServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServClient interface {
	GetTime(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ServerTime, error)
	GetKlines(ctx context.Context, in *KlinesRequest, opts ...grpc.CallOption) (*KlinesResponse, error)
	GetFuturesPairs(ctx context.Context, in *Void, opts ...grpc.CallOption) (*FuturesPairsResponse, error)
	GetSocketData(ctx context.Context, opts ...grpc.CallOption) (DataServ_GetSocketDataClient, error)
}

type dataServClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServClient(cc grpc.ClientConnInterface) DataServClient {
	return &dataServClient{cc}
}

func (c *dataServClient) GetTime(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ServerTime, error) {
	out := new(ServerTime)
	err := c.cc.Invoke(ctx, "/pb.DataServ/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServClient) GetKlines(ctx context.Context, in *KlinesRequest, opts ...grpc.CallOption) (*KlinesResponse, error) {
	out := new(KlinesResponse)
	err := c.cc.Invoke(ctx, "/pb.DataServ/GetKlines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServClient) GetFuturesPairs(ctx context.Context, in *Void, opts ...grpc.CallOption) (*FuturesPairsResponse, error) {
	out := new(FuturesPairsResponse)
	err := c.cc.Invoke(ctx, "/pb.DataServ/GetFuturesPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServClient) GetSocketData(ctx context.Context, opts ...grpc.CallOption) (DataServ_GetSocketDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataServ_ServiceDesc.Streams[0], "/pb.DataServ/GetSocketData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServGetSocketDataClient{stream}
	return x, nil
}

type DataServ_GetSocketDataClient interface {
	Send(*SocketRequest) error
	Recv() (*SocketResponse, error)
	grpc.ClientStream
}

type dataServGetSocketDataClient struct {
	grpc.ClientStream
}

func (x *dataServGetSocketDataClient) Send(m *SocketRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServGetSocketDataClient) Recv() (*SocketResponse, error) {
	m := new(SocketResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServServer is the server API for DataServ service.
// All implementations must embed UnimplementedDataServServer
// for forward compatibility
type DataServServer interface {
	GetTime(context.Context, *Void) (*ServerTime, error)
	GetKlines(context.Context, *KlinesRequest) (*KlinesResponse, error)
	GetFuturesPairs(context.Context, *Void) (*FuturesPairsResponse, error)
	GetSocketData(DataServ_GetSocketDataServer) error
	mustEmbedUnimplementedDataServServer()
}

// UnimplementedDataServServer must be embedded to have forward compatible implementations.
type UnimplementedDataServServer struct {
}

func (UnimplementedDataServServer) GetTime(context.Context, *Void) (*ServerTime, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedDataServServer) GetKlines(context.Context, *KlinesRequest) (*KlinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKlines not implemented")
}
func (UnimplementedDataServServer) GetFuturesPairs(context.Context, *Void) (*FuturesPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuturesPairs not implemented")
}
func (UnimplementedDataServServer) GetSocketData(DataServ_GetSocketDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSocketData not implemented")
}
func (UnimplementedDataServServer) mustEmbedUnimplementedDataServServer() {}

// UnsafeDataServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServServer will
// result in compilation errors.
type UnsafeDataServServer interface {
	mustEmbedUnimplementedDataServServer()
}

func RegisterDataServServer(s grpc.ServiceRegistrar, srv DataServServer) {
	s.RegisterService(&DataServ_ServiceDesc, srv)
}

func _DataServ_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DataServ/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServServer).GetTime(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataServ_GetKlines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KlinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServServer).GetKlines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DataServ/GetKlines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServServer).GetKlines(ctx, req.(*KlinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataServ_GetFuturesPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServServer).GetFuturesPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DataServ/GetFuturesPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServServer).GetFuturesPairs(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataServ_GetSocketData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServServer).GetSocketData(&dataServGetSocketDataServer{stream})
}

type DataServ_GetSocketDataServer interface {
	Send(*SocketResponse) error
	Recv() (*SocketRequest, error)
	grpc.ServerStream
}

type dataServGetSocketDataServer struct {
	grpc.ServerStream
}

func (x *dataServGetSocketDataServer) Send(m *SocketResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServGetSocketDataServer) Recv() (*SocketRequest, error) {
	m := new(SocketRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServ_ServiceDesc is the grpc.ServiceDesc for DataServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DataServ",
	HandlerType: (*DataServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _DataServ_GetTime_Handler,
		},
		{
			MethodName: "GetKlines",
			Handler:    _DataServ_GetKlines_Handler,
		},
		{
			MethodName: "GetFuturesPairs",
			Handler:    _DataServ_GetFuturesPairs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSocketData",
			Handler:       _DataServ_GetSocketData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ds_api.proto",
}
