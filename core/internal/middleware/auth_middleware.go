package middleware

import (
	"core/core/define"
	"core/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := helper.ParseToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		// 往请求头里放置基础信息
		r.Header.Set(define.HKOBJ.Uid, string(rune(uc.Id)))
		r.Header.Set(define.HKOBJ.UserIdentity, uc.Identify)
		r.Header.Set(define.HKOBJ.UserName, uc.Name)
		// Passthrough to next handler if need
		next(w, r)
	}
}