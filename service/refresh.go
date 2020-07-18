package service

import (
	"fmt"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/token"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type RefreshTokenParams struct {
	PreviousToken string
}

func (h *Service) RefreshToken(u *RefreshTokenParams) (*LoginOutput, error) {
	token2, err := jwt.Parse(u.PreviousToken, token.CheckConformHMAC("REFRESH_TOKEN"))
	if err != nil {
		return nil, err
	}

	//is token2 valid?
	if _, ok := token2.Claims.(jwt.Claims); !ok && !token2.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		//json.NewEncoder(w).Encode("Refresh token2 expired")
		return nil, errs.ErrAuth
	}

	claims, ok := token2.Claims.(jwt.MapClaims)
	if !ok || !token2.Valid {
		return nil, errs.ErrAuth
	}

	//Since token2 is valid, get the uuid:
	refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
	if !ok {
		return nil, errs.ErrAuth
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, errs.ErrAuth
	}

	//Delete the previous Refresh Token
	deleted, err := h.Redis.RedisDeleteAuth(refreshUuid)
	if err != nil || deleted == 0 { //if any goes wrong
		return nil, errs.ErrAuth
	}

	//Create new pairs of refresh and access tokens
	ts, err := token.CreateToken(userId)
	if err != nil {
		return nil, err
	}

	//save the tokens metadata to redis
	err = h.Redis.RedisCreateAuth(userId, ts)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{Ts: ts}, nil
}
