package main

type LoginParams struct {
	Username string
	Password string
}

type LoginOutput struct {
	Ts *TokenDetails
}

func LoginHandler(u *LoginParams) (*LoginOutput, error) {
	userid, err := UserRepoSearchWithUsernamePassword(u.Username, u.Password)
	if err != nil {
		return nil, ErrAuth
	}

	ts, err := CreateToken(userid)
	if err != nil {
		return nil, err
	}

	err = RedisCreateAuth(userid, ts)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{Ts: ts}, nil
}

func UserRepoSearchWithUsernamePassword(username string, password string) (userid uint64, err error) {
	if user.Username != username || user.Password != password {
		return 0, err
	}
	return 1, nil
}
