package logic

import (
	"context"
	"douyin/v2/jwt/app/rpc/Jwt"
	"douyin/v2/user/app/rpc/userrpc"

	"douyin/v2/api/internal/svc"
	"douyin/v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	token, err := l.svcCtx.JwtRpc.ParseToken(l.ctx, &Jwt.ParseTokenReq{Token: req.Token})
	if err != nil {
		return &types.FollowerListResp{
			Response: types.Response{
				StatusCode: STATUS_FAIL,
				StatusMsg:  STATUS_FAIL_TOKEN_MSG,
			},
			UserList: nil,
		}, nil
	}
	userid := token.UserID

	r, err := l.svcCtx.UserRpc.FollowerList(l.ctx, &userrpc.FollowerListReq{UserId: userid, ToUserId: req.UserId})
	if err != nil {
		return nil, err
	}

	var userList []types.User
	for _, v := range r.UserList {
		u := types.User{
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			ID:            v.ID,
			IsFollow:      v.IsFollow,
			Name:          v.Name,
		}
		userList = append(userList, u)
	}

	return &types.FollowerListResp{
		Response: types.Response{
			StatusCode: STATUS_SUCCESS,
			StatusMsg:  STATUS_SUCCESS_MSG,
		},
		UserList: userList,
	}, nil
}
