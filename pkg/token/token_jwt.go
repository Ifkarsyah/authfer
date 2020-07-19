package token

import (
	"github.com/Ifkarsyah/authfer/model"
	"github.com/Ifkarsyah/authfer/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"time"
)

func CreateToken(userid uint64) (*model.TokenDetails, error) {
	cfgAccessExp := time.Duration(config.AppConfig.JwtAccessExpires)
	cfgRefreshExp := time.Duration(config.AppConfig.JwtRefreshExpires)

	td := &model.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * cfgAccessExp).Unix(),
		RtExpires:   time.Now().Add(time.Hour * 24 * cfgRefreshExp).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RefreshUuid: uuid.NewV4().String(),
	}

	accessToken, err := createAccessToken(userid, td)
	refreshToken, err := createRefreshToken(userid, td)

	if err != nil {
		return nil, err
	}

	td.AccessToken = accessToken
	td.RefreshToken = refreshToken

	return td, nil
}

func createRefreshToken(userid uint64, td *model.TokenDetails) (string, error) {
	rtClaims := jwt.MapClaims{
		"refresh_uuid": td.RefreshUuid,
		"user_id":      userid,
		"exp":          td.RtExpires,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return rt.SignedString([]byte(config.AppConfig.Secret))
}

func createAccessToken(userid uint64, td *model.TokenDetails) (string, error) {
	atClaims := jwt.MapClaims{
		"access_uuid": td.AccessUuid,
		"user_id":     userid,
		"exp":         td.AtExpires,
		"authorized":  true,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte(config.AppConfig.Secret))
}
