// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: dict/dict.proto

package dict

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

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Add(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Sub(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Mult(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Div(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Sqrt(ctx context.Context, in *SqrtNum, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Pow(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
	Min(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Add(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Sub(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Mult(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Mult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Div(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Sqrt(ctx context.Context, in *SqrtNum, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Pow(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Pow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Min(ctx context.Context, in *TwoNumber, opts ...grpc.CallOption) (*ResultTwoNumber, error) {
	out := new(ResultTwoNumber)
	err := c.cc.Invoke(ctx, "/dict.CalcService/Min", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	Add(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	Sub(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	Mult(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	Div(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	Sqrt(context.Context, *SqrtNum) (*ResultTwoNumber, error)
	Pow(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	Min(context.Context, *TwoNumber) (*ResultTwoNumber, error)
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) Add(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalcServiceServer) Sub(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (UnimplementedCalcServiceServer) Mult(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mult not implemented")
}
func (UnimplementedCalcServiceServer) Div(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (UnimplementedCalcServiceServer) Sqrt(context.Context, *SqrtNum) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalcServiceServer) Pow(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pow not implemented")
}
func (UnimplementedCalcServiceServer) Min(context.Context, *TwoNumber) (*ResultTwoNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Min not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Add(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sub(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Mult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Mult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Mult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Mult(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Div(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtNum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sqrt(ctx, req.(*SqrtNum))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Pow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Pow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Pow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Pow(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Min_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Min(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.CalcService/Min",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Min(ctx, req.(*TwoNumber))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dict.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalcService_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _CalcService_Sub_Handler,
		},
		{
			MethodName: "Mult",
			Handler:    _CalcService_Mult_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _CalcService_Div_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalcService_Sqrt_Handler,
		},
		{
			MethodName: "Pow",
			Handler:    _CalcService_Pow_Handler,
		},
		{
			MethodName: "Min",
			Handler:    _CalcService_Min_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dict/dict.proto",
}
