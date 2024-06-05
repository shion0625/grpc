//
//  CommandService用のParam型とResult型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/command.proto

// パッケージの宣言

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CategoryCommand_Create_FullMethodName = "/proto.CategoryCommand/Create"
	CategoryCommand_Update_FullMethodName = "/proto.CategoryCommand/Update"
	CategoryCommand_Delete_FullMethodName = "/proto.CategoryCommand/Delete"
)

// CategoryCommandClient is the client API for CategoryCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *****************************************
// 商品カテゴリと商品の更新サービス型の定義
// *****************************************
//
//	商品カテゴリ更新サービス型
type CategoryCommandClient interface {
	// 商品カテゴリを追加した結果を返す
	Create(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error)
	// 商品カテゴリを更新した結果を返す
	Update(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error)
	// 商品カテゴリを削除した結果を返す
	Delete(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error)
}

type categoryCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryCommandClient(cc grpc.ClientConnInterface) CategoryCommandClient {
	return &categoryCommandClient{cc}
}

func (c *categoryCommandClient) Create(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryUpResult)
	err := c.cc.Invoke(ctx, CategoryCommand_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryCommandClient) Update(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryUpResult)
	err := c.cc.Invoke(ctx, CategoryCommand_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryCommandClient) Delete(ctx context.Context, in *CategoryUpParam, opts ...grpc.CallOption) (*CategoryUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryUpResult)
	err := c.cc.Invoke(ctx, CategoryCommand_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryCommandServer is the server API for CategoryCommand service.
// All implementations must embed UnimplementedCategoryCommandServer
// for forward compatibility
//
// *****************************************
// 商品カテゴリと商品の更新サービス型の定義
// *****************************************
//
//	商品カテゴリ更新サービス型
type CategoryCommandServer interface {
	// 商品カテゴリを追加した結果を返す
	Create(context.Context, *CategoryUpParam) (*CategoryUpResult, error)
	// 商品カテゴリを更新した結果を返す
	Update(context.Context, *CategoryUpParam) (*CategoryUpResult, error)
	// 商品カテゴリを削除した結果を返す
	Delete(context.Context, *CategoryUpParam) (*CategoryUpResult, error)
	mustEmbedUnimplementedCategoryCommandServer()
}

// UnimplementedCategoryCommandServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryCommandServer struct {
}

func (UnimplementedCategoryCommandServer) Create(context.Context, *CategoryUpParam) (*CategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryCommandServer) Update(context.Context, *CategoryUpParam) (*CategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryCommandServer) Delete(context.Context, *CategoryUpParam) (*CategoryUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryCommandServer) mustEmbedUnimplementedCategoryCommandServer() {}

// UnsafeCategoryCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryCommandServer will
// result in compilation errors.
type UnsafeCategoryCommandServer interface {
	mustEmbedUnimplementedCategoryCommandServer()
}

func RegisterCategoryCommandServer(s grpc.ServiceRegistrar, srv CategoryCommandServer) {
	s.RegisterService(&CategoryCommand_ServiceDesc, srv)
}

func _CategoryCommand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommand_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServer).Create(ctx, req.(*CategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryCommand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommand_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServer).Update(ctx, req.(*CategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryCommand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommand_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServer).Delete(ctx, req.(*CategoryUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryCommand_ServiceDesc is the grpc.ServiceDesc for CategoryCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CategoryCommand",
	HandlerType: (*CategoryCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoryCommand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryCommand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryCommand_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/command.proto",
}

const (
	ProductCommand_Create_FullMethodName = "/proto.ProductCommand/Create"
	ProductCommand_Update_FullMethodName = "/proto.ProductCommand/Update"
	ProductCommand_Delete_FullMethodName = "/proto.ProductCommand/Delete"
)

// ProductCommandClient is the client API for ProductCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
//	商品更新サービス型
type ProductCommandClient interface {
	// 商品カテゴリを追加した結果を返す
	Create(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error)
	// 商品カテゴリを更新した結果を返す
	Update(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error)
	// 商品カテゴリを削除した結果を返す
	Delete(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error)
}

type productCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCommandClient(cc grpc.ClientConnInterface) ProductCommandClient {
	return &productCommandClient{cc}
}

func (c *productCommandClient) Create(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductUpResult)
	err := c.cc.Invoke(ctx, ProductCommand_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandClient) Update(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductUpResult)
	err := c.cc.Invoke(ctx, ProductCommand_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandClient) Delete(ctx context.Context, in *ProductUpParam, opts ...grpc.CallOption) (*ProductUpResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductUpResult)
	err := c.cc.Invoke(ctx, ProductCommand_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCommandServer is the server API for ProductCommand service.
// All implementations must embed UnimplementedProductCommandServer
// for forward compatibility
//
//	商品更新サービス型
type ProductCommandServer interface {
	// 商品カテゴリを追加した結果を返す
	Create(context.Context, *ProductUpParam) (*ProductUpResult, error)
	// 商品カテゴリを更新した結果を返す
	Update(context.Context, *ProductUpParam) (*ProductUpResult, error)
	// 商品カテゴリを削除した結果を返す
	Delete(context.Context, *ProductUpParam) (*ProductUpResult, error)
	mustEmbedUnimplementedProductCommandServer()
}

// UnimplementedProductCommandServer must be embedded to have forward compatible implementations.
type UnimplementedProductCommandServer struct {
}

func (UnimplementedProductCommandServer) Create(context.Context, *ProductUpParam) (*ProductUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductCommandServer) Update(context.Context, *ProductUpParam) (*ProductUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductCommandServer) Delete(context.Context, *ProductUpParam) (*ProductUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductCommandServer) mustEmbedUnimplementedProductCommandServer() {}

// UnsafeProductCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCommandServer will
// result in compilation errors.
type UnsafeProductCommandServer interface {
	mustEmbedUnimplementedProductCommandServer()
}

func RegisterProductCommandServer(s grpc.ServiceRegistrar, srv ProductCommandServer) {
	s.RegisterService(&ProductCommand_ServiceDesc, srv)
}

func _ProductCommand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommand_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServer).Create(ctx, req.(*ProductUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommand_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServer).Update(ctx, req.(*ProductUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommand_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServer).Delete(ctx, req.(*ProductUpParam))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCommand_ServiceDesc is the grpc.ServiceDesc for ProductCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductCommand",
	HandlerType: (*ProductCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductCommand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductCommand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductCommand_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/command.proto",
}
