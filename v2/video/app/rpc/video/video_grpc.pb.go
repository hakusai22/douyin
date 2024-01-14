// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: video.proto

package video

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

const (
	VideoRpc_GetPublishList_FullMethodName  = "/video.VideoRpc/GetPublishList"
	VideoRpc_GetFeed_FullMethodName         = "/video.VideoRpc/GetFeed"
	VideoRpc_CommentAction_FullMethodName   = "/video.VideoRpc/CommentAction"
	VideoRpc_GetCommentList_FullMethodName  = "/video.VideoRpc/GetCommentList"
	VideoRpc_FavoriteAction_FullMethodName  = "/video.VideoRpc/FavoriteAction"
	VideoRpc_GetFavoriteList_FullMethodName = "/video.VideoRpc/GetFavoriteList"
)

// VideoRpcClient is the client API for VideoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRpcClient interface {
	GetPublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
	GetFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error)
	CommentAction(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error)
	GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	FavoriteAction(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error)
	GetFavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
}

type videoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRpcClient(cc grpc.ClientConnInterface) VideoRpcClient {
	return &videoRpcClient{cc}
}

func (c *videoRpcClient) GetPublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	out := new(PublishListResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetPublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) CommentAction(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error) {
	out := new(CommentResp)
	err := c.cc.Invoke(ctx, VideoRpc_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	out := new(CommentListResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FavoriteAction(ctx context.Context, in *FavoriteReq, opts ...grpc.CallOption) (*FavoriteResp, error) {
	out := new(FavoriteResp)
	err := c.cc.Invoke(ctx, VideoRpc_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetFavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetFavoriteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoRpcServer is the server API for VideoRpc service.
// All implementations must embed UnimplementedVideoRpcServer
// for forward compatibility
type VideoRpcServer interface {
	GetPublishList(context.Context, *PublishListReq) (*PublishListResp, error)
	GetFeed(context.Context, *FeedReq) (*FeedResp, error)
	CommentAction(context.Context, *CommentReq) (*CommentResp, error)
	GetCommentList(context.Context, *CommentListReq) (*CommentListResp, error)
	FavoriteAction(context.Context, *FavoriteReq) (*FavoriteResp, error)
	GetFavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	mustEmbedUnimplementedVideoRpcServer()
}

// UnimplementedVideoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRpcServer struct {
}

func (UnimplementedVideoRpcServer) GetPublishList(context.Context, *PublishListReq) (*PublishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishList not implemented")
}
func (UnimplementedVideoRpcServer) GetFeed(context.Context, *FeedReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedVideoRpcServer) CommentAction(context.Context, *CommentReq) (*CommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedVideoRpcServer) GetCommentList(context.Context, *CommentListReq) (*CommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedVideoRpcServer) FavoriteAction(context.Context, *FavoriteReq) (*FavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedVideoRpcServer) GetFavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteList not implemented")
}
func (UnimplementedVideoRpcServer) mustEmbedUnimplementedVideoRpcServer() {}

// UnsafeVideoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRpcServer will
// result in compilation errors.
type UnsafeVideoRpcServer interface {
	mustEmbedUnimplementedVideoRpcServer()
}

func RegisterVideoRpcServer(s grpc.ServiceRegistrar, srv VideoRpcServer) {
	s.RegisterService(&VideoRpc_ServiceDesc, srv)
}

func _VideoRpc_GetPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetPublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetPublishList(ctx, req.(*PublishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetFeed(ctx, req.(*FeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).CommentAction(ctx, req.(*CommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetCommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FavoriteAction(ctx, req.(*FavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetFavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetFavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetFavoriteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetFavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoRpc_ServiceDesc is the grpc.ServiceDesc for VideoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoRpc",
	HandlerType: (*VideoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublishList",
			Handler:    _VideoRpc_GetPublishList_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _VideoRpc_GetFeed_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _VideoRpc_CommentAction_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _VideoRpc_GetCommentList_Handler,
		},
		{
			MethodName: "FavoriteAction",
			Handler:    _VideoRpc_FavoriteAction_Handler,
		},
		{
			MethodName: "GetFavoriteList",
			Handler:    _VideoRpc_GetFavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
