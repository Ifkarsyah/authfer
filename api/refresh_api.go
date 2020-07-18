package api

import (
	"encoding/json"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"github.com/Ifkarsyah/authfer/service"
	"net/http"
)

type refreshTokenHandlerFunc func(u *service.RefreshTokenParams) (*service.LoginOutput, error)

func RefreshAPI(handlerFunc refreshTokenHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mapToken := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&mapToken)
		if err != nil {
			responder.ResponseError(w, errs.ErrBadRequest)
			return
		}
		refreshToken := mapToken["refresh_token"]

		ts, err := handlerFunc(&service.RefreshTokenParams{
			PreviousToken: refreshToken,
		})
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseOK(w, ts)
	})
}
