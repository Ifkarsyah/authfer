package api

import (
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"github.com/Ifkarsyah/authfer/pkg/token"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currToken, err := token.VerifyToken(r)
		if err != nil {
			responder.ResponseError(w, errs.ErrAuth)
			return
		}
		if _, ok := currToken.Claims.(jwt.Claims); !ok && !currToken.Valid {
			responder.ResponseError(w, errs.ErrAuth)
			return
		}
		next.ServeHTTP(w, r)
	}
}
