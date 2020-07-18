package service

import (
	"github.com/Ifkarsyah/authfer/repo"
)

type IService interface {
	Login(u *LoginParams) (*LoginOutput, error)
	Logout(u *LogoutParams) error
	RefreshToken(u *RefreshTokenParams) (*LoginOutput, error)
}

type Service struct {
	Redis *repo.RedisRepo
}
