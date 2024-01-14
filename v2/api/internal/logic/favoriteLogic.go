package logic

import (
	"context"
	"douyin/v2/api/internal/svc"
	"douyin/v2/api/internal/types"
	"douyin/v2/jwt/app/rpc/Jwt"
	"douyin/v2/video/app/rpc/videorpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	token, err := l.svcCtx.JwtRpc.ParseToken(l.ctx, &Jwt.ParseTokenReq{Token: req.Token})
	if err != nil {
		return &types.FavoriteResp{
			Response: types.Response{
				StatusCode: STATUS_FAIL,
				StatusMsg:  STATUS_FAIL_TOKEN_MSG,
			},
		}, nil
	}
	userid := token.UserID
	r, err := l.svcCtx.VideoRpc.FavoriteAction(l.ctx, &videorpc.FavoriteReq{
		VideoId:    req.VideoId,
		UserId:     userid,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}

	return &types.FavoriteResp{
		Response: types.Response{
			StatusCode: r.StatusCode,
			StatusMsg:  r.StatusMsg,
		},
	}, nil
}
