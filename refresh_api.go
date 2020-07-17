package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

func RefreshAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mapToken := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&mapToken)
		if err != nil {
			ResponseError(w, ErrBadRequest)
			return
		}
		refreshToken := mapToken["refresh_token"]

		token, err := jwt.Parse(refreshToken, checkConformHMAC("REFRESH_TOKEN"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Refresh token expired")
			return
		}

		//is token valid?
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Refresh token expired")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Since token is valid, get the uuid:
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Delete the previous Refresh Token
		deleted, delErr := RedisDeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := CreateToken(userId)
		if createErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//save the tokens metadata to redis
		saveErr := RedisCreateAuth(userId, ts)
		if saveErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ResponseOK(w, map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		})
	})
}
