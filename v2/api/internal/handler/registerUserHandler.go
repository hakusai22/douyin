package handler

import (
	"net/http"

	"douyin/v2/api/internal/logic"
	"douyin/v2/api/internal/svc"
	"douyin/v2/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterUserLogic(r.Context(), svcCtx)
		resp, err := l.RegisterUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
