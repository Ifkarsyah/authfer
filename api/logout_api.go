package api

import (
	"fmt"
	"github.com/Ifkarsyah/authfer/model"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	"github.com/Ifkarsyah/authfer/pkg/token"
	"github.com/Ifkarsyah/authfer/service"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

type logoutHandlerFunc func(*service.LogoutParams) error

func Logout(logoutHandler logoutHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		au, err := extractTokenMetadata(r)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		err = logoutHandler(&service.LogoutParams{AccessDetails: au})
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func extractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	prevToken, err := token.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := prevToken.Claims.(jwt.MapClaims)
	if !ok || !prevToken.Valid {
		return nil, err
	}

	accessUuid, ok := claims["access_uuid"].(string)
	if !ok {
		return nil, err
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}
	return &model.AccessDetails{
		AccessUuid: accessUuid,
		UserId:     userId,
	}, nil
}
