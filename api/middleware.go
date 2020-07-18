package api

import (
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"github.com/Ifkarsyah/authfer/pkg/token"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := token.IsTokenValid(r)
		if err != nil {
			responder.ResponseError(w, errs.ErrAuth)
			return
		}
		next.ServeHTTP(w, r)
	}
}
