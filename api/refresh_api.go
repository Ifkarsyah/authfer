package api

import (
	"encoding/json"
	"fmt"
	"github.com/Ifkarsyah/authfer/repo"
	"github.com/Ifkarsyah/authfer/util/errs"
	"github.com/Ifkarsyah/authfer/util/responder"
	"github.com/Ifkarsyah/authfer/util/token"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

func RefreshAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mapToken := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&mapToken)
		if err != nil {
			responder.ResponseError(w, errs.ErrBadRequest)
			return
		}
		refreshToken := mapToken["refresh_token"]

		token2, err := jwt.Parse(refreshToken, token.CheckConformHMAC("REFRESH_TOKEN"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Refresh token2 expired")
			return
		}

		//is token2 valid?
		if _, ok := token2.Claims.(jwt.Claims); !ok && !token2.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Refresh token2 expired")
			return
		}

		claims, ok := token2.Claims.(jwt.MapClaims)
		if !ok || !token2.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Since token2 is valid, get the uuid:
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
		deleted, delErr := repo.RedisDeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := token.CreateToken(userId)
		if createErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//save the tokens metadata to redis
		saveErr := repo.RedisCreateAuth(userId, ts)
		if saveErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responder.ResponseOK(w, map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		})
	})
}
