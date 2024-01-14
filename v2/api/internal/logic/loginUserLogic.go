package logic

import (
	"context"
	"douyin/v2/api/internal/svc"
	"douyin/v2/api/internal/types"
	"douyin/v2/jwt/app/rpc/Jwt"
	"douyin/v2/user/app/rpc/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginReq) (resp *types.LoginResp, err error) {
	r, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var token *string
	token = nil
	if r.StatusCode == STATUS_SUCCESS {
		tokenResp, err := l.svcCtx.JwtRpc.CreateToken(l.ctx, &Jwt.CreateTokenReq{
			UserID:       strconv.FormatUint(r.UserID, 10),
			AccessExpire: l.svcCtx.Config.JwtConfig.AccessExpire,
		})
		if err != nil {
			return nil, err
		}
		token = &tokenResp.Token
	}

	return &types.LoginResp{
		Response: types.Response{
			StatusCode: r.StatusCode,
			StatusMsg:  r.StatusMsg,
		},
		Token:  token,
		UserID: r.UserID,
	}, nil
}
