package handler

import (
	"net/http"

	"douyin/v2/api/internal/logic"
	"douyin/v2/api/internal/svc"
	"douyin/v2/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func favoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteLogic(r.Context(), svcCtx)
		resp, err := l.Favorite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
