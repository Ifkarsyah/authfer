package main

import "errors"

func LoginHandler(u *User) (*TokenDetails, error) {
	anyError := errors.New("any")
	if user.Username != u.Username || user.Password != u.Password {
		return nil, anyError
	}

	ts, err := CreateToken(user.ID)
	if err != nil {
		return nil, anyError
	}

	err = RedisCreateAuth(user.ID, ts)
	if err != nil {
		return nil, anyError
	}

	return ts, nil
}
