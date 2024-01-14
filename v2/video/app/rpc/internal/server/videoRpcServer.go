// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"douyin/v2/video/app/rpc/internal/logic"
	"douyin/v2/video/app/rpc/internal/svc"
	"douyin/v2/video/app/rpc/video"
)

type VideoRpcServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoRpcServer
}

func NewVideoRpcServer(svcCtx *svc.ServiceContext) *VideoRpcServer {
	return &VideoRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoRpcServer) GetPublishList(ctx context.Context, in *video.PublishListReq) (*video.PublishListResp, error) {
	l := logic.NewGetPublishListLogic(ctx, s.svcCtx)
	return l.GetPublishList(in)
}

func (s *VideoRpcServer) GetFeed(ctx context.Context, in *video.FeedReq) (*video.FeedResp, error) {
	l := logic.NewGetFeedLogic(ctx, s.svcCtx)
	return l.GetFeed(in)
}

func (s *VideoRpcServer) CommentAction(ctx context.Context, in *video.CommentReq) (*video.CommentResp, error) {
	l := logic.NewCommentActionLogic(ctx, s.svcCtx)
	return l.CommentAction(in)
}

func (s *VideoRpcServer) GetCommentList(ctx context.Context, in *video.CommentListReq) (*video.CommentListResp, error) {
	l := logic.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}

func (s *VideoRpcServer) FavoriteAction(ctx context.Context, in *video.FavoriteReq) (*video.FavoriteResp, error) {
	l := logic.NewFavoriteActionLogic(ctx, s.svcCtx)
	return l.FavoriteAction(in)
}

func (s *VideoRpcServer) GetFavoriteList(ctx context.Context, in *video.FavoriteListReq) (*video.FavoriteListResp, error) {
	l := logic.NewGetFavoriteListLogic(ctx, s.svcCtx)
	return l.GetFavoriteList(in)
}