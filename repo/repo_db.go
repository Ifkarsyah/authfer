package repo

import (
	"github.com/Ifkarsyah/authfer/model"
)

type DBRepo struct {
}

type IDBRepo interface {
	GetUseridByUsernamePassword(username string, password string) (userid uint64, err error)
}

func NewDBConnection() *DBRepo {
	return &DBRepo{}
}

func (r *DBRepo) GetUseridByUsernamePassword(username string, password string) (userid uint64, err error) {
	if model.AuthSample.Username != username || model.AuthSample.Password != password {
		return 0, err
	}
	return 1, nil
}
