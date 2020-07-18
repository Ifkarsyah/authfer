package api

import (
	"encoding/json"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"github.com/Ifkarsyah/authfer/service"
	"net/http"
)

type loginHandlerFunc func(*service.LoginParams) (*service.LoginOutput, error)

func LoginAPI(loginHandler loginHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := new(service.LoginParams)
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
