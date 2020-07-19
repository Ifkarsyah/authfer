package repo

import (
	"github.com/Ifkarsyah/authfer/model"
	"strconv"
	"time"
)

func (r *RedisRepo) RedisCreateAuth(userid uint64, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atKey := td.AccessUuid
	rtKey := td.RefreshUuid

	val := strconv.Itoa(int(userid))

	errAccess := r.client.Set(atKey, val, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := r.client.Set(rtKey, val, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (r *RedisRepo) RedisGetAuth(authD *model.AccessDetails) (uint64, error) {
	userid, err := r.client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (r *RedisRepo) RedisDeleteAuth(givenUuid string) (int64, error) {
	deleted, err := r.client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
