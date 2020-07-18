package repo

import (
	"github.com/Ifkarsyah/authfer/model"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

type RedisRepo struct {
	client *redis.Client
}

type IRedisRepo interface {
	RedisCreateAuth(userid uint64, td *model.TokenDetails) error
	RedisGetAuth(authD *AccessDetails) (uint64, error)
	RedisDeleteAuth(givenUuid string) (int64, error)
}

func NewRedisConnection(host string, port string) *RedisRepo {
	redisClient := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return &RedisRepo{client: redisClient}
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func (r *RedisRepo) RedisCreateAuth(userid uint64, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := r.client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := r.client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (r *RedisRepo) RedisGetAuth(authD *AccessDetails) (uint64, error) {
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
