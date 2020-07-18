package api

import (
	"encoding/json"
	"github.com/Ifkarsyah/authfer/handler"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"net/http"
)

type loginHandlerFunc func(*handler.LoginParams) (*handler.LoginOutput, error)

func LoginAPI(loginHandler loginHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := new(handler.LoginParams)
		err := json.NewDecoder(r.Body).Decode(params)
		if err != nil {
			responder.ResponseError(w, errs.ErrBadRequest)
			return
		}

		ts, err := loginHandler(params)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseOK(w, ts)
	})
}
