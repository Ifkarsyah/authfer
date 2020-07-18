package api

import (
	"fmt"
	"github.com/Ifkarsyah/authfer/pkg/responder"
	token2 "github.com/Ifkarsyah/authfer/pkg/token"
	"github.com/Ifkarsyah/authfer/repo"
	"github.com/Ifkarsyah/authfer/service"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

type logoutHandlerFunc func(*service.LogoutParams) error

func Logout(logoutHandler logoutHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		au, err := ExtractTokenMetadata(r)
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

func ExtractTokenMetadata(r *http.Request) (*repo.AccessDetails, error) {
	token, err := token2.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
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
	return &repo.AccessDetails{
		AccessUuid: accessUuid,
		UserId:     userId,
	}, nil
}
