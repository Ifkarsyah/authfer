package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

func Logout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		au, err := ExtractTokenMetadata(r)
		if err != nil {
			ResponseError(w, err)
			return
		}

		deleted, delErr := RedisDeleteAuth(au.AccessUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			ResponseError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
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
	return &AccessDetails{
		AccessUuid: accessUuid,
		UserId:     userId,
	}, nil
}
