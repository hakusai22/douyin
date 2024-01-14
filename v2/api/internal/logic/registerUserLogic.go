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

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	// 用户名和密码长度不能超过 32 字符（不知道为什么客户端没做这种校验）
	if len(req.Username) > 32 || len(req.Password) > 32 {
		return &types.RegisterResp{
			Response: types.Response{
				StatusCode: STATUS_FAIL,
				StatusMsg:  STATUS_FAIL_TOOLONG_MSG,
			},
			Token:  nil,
			UserID: 0,
		}, err
	}

	r, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
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

	return &types.RegisterResp{
		Response: types.Response{
			StatusCode: r.StatusCode,
			StatusMsg:  r.StatusMsg,
		},
		Token:  token,
		UserID: r.UserID,
	}, nil
}
