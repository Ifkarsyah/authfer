package service

import (
	"fmt"
	"github.com/Ifkarsyah/authfer/pkg/config"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/token"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type RefreshTokenParams struct {
	PreviousToken string
}

func (h *Service) RefreshToken(u *RefreshTokenParams) (*LoginOutput, error) {
	prevToken, err := jwt.Parse(u.PreviousToken, token.CheckConformHMAC(config.AppConfig.Secret))
	if err != nil {
		return nil, err
	}

	//is prevToken valid?
	if _, ok := prevToken.Claims.(jwt.Claims); !ok && !prevToken.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		//json.NewEncoder(w).Encode("Refresh prevToken expired")
		return nil, errs.ErrAuth
	}

	claims, ok := prevToken.Claims.(jwt.MapClaims)
	if !ok || !prevToken.Valid {
		return nil, errs.ErrAuth
	}

	//Since prevToken is valid, get the uuid:
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
