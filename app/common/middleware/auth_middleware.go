package middleware

import (
	"cloud_disk/app/user/cmd/rpc/user_helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := user_helper.ParseToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		// fmt.Println(uc.Id, uc.Identity, uc.Name)
		// 往请求头里放置基础信息
		r.Header.Set(define.HKOBJ.Uid, string(rune(uc.Id)))
		r.Header.Set(define.HKOBJ.UserIdentity, uc.Identity)
		r.Header.Set(define.HKOBJ.UserName, uc.Name)
		// Passthrough to next handler if need
		next(w, r)
	}
}
