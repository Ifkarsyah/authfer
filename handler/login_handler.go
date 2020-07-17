package handler

import (
	"github.com/Ifkarsyah/authfer/model"
	"github.com/Ifkarsyah/authfer/repo"
	"github.com/Ifkarsyah/authfer/util/errs"
	"github.com/Ifkarsyah/authfer/util/token"
)

type LoginParams struct {
	Username string
	Password string
}

type LoginOutput struct {
	Ts *model.TokenDetails
}

func LoginHandler(u *LoginParams) (*LoginOutput, error) {
	userid, err := UserRepoSearchWithUsernamePassword(u.Username, u.Password)
	if err != nil {
		return nil, errs.ErrAuth
	}

	ts, err := token.CreateToken(userid)
	if err != nil {
		return nil, err
	}

	err = repo.RedisCreateAuth(userid, ts)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{Ts: ts}, nil
}

func UserRepoSearchWithUsernamePassword(username string, password string) (userid uint64, err error) {
	if model.AuthSample.Username != username || model.AuthSample.Password != password {
		return 0, err
	}
	return 1, nil
}
