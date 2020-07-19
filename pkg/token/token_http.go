package token

import (
	"fmt"
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

func CheckConformHMAC(secret string) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}
}
