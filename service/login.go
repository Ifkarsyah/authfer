package service

import (
	"github.com/Ifkarsyah/authfer/model"
	"github.com/Ifkarsyah/authfer/pkg/errs"
	"github.com/Ifkarsyah/authfer/pkg/token"
)

type LoginParams struct {
	Username string
	Password string
}

type LoginOutput struct {
	Ts *model.TokenDetails
}

type LoginHandlerFunc func(*LoginParams) (*LoginOutput, error)

func (h *Service) Login(u *LoginParams) (*LoginOutput, error) {
	userid, err := h.DB.GetUseridByUsernamePassword(u.Username, u.Password)
	if err != nil {
		return nil, errs.ErrAuth
	}
	ts, err := token.CreateToken(userid)
	if err != nil {
		return nil, err
	}

	err = h.Redis.RedisCreateAuth(userid, ts)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{Ts: ts}, nil
}
