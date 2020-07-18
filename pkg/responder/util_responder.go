package responder

import (
	"encoding/json"
	"errors"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"net/http"
)

func ResponseOK(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err error) error {
	httpCode, resp := errToHttpFormat(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	return json.NewEncoder(w).Encode(resp)
}

func checkErrorType(err error, target ...error) bool {
	for _, t := range target {
		if errors.Is(err, t) {
			return true
		}
	}

	return false
}

func errToHttpFormat(err error) (int, interface{}) {
	resp := CommonResponse{Status: 1, Description: err.Error()}

	switch {
	case checkErrorType(err, errs.ErrAuth):
		return http.StatusUnauthorized, resp
	case checkErrorType(err, errs.ErrBadRequest):
		return http.StatusUnauthorized, resp
	default:
		resp.Description = "err internal"
		return http.StatusInternalServerError, resp
	}
}

type CommonResponse struct { // for consumer
	Status      int    `json:"status"`
	Description string `json:"description"`
}
