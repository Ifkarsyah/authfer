package token

import (
	"github.com/Ifkarsyah/authfer/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r.Header.Get("Authorization"))
	token, err := jwt.Parse(tokenString, CheckConformHMAC(config.AppConfig.Secret))
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
