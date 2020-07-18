package repo

import (
	"github.com/Ifkarsyah/authfer/model"
	"github.com/go-redis/redis/v7"
)

type RedisRepo struct {
	client *redis.Client
}

type IRedisRepo interface {
	RedisCreateAuth(userid uint64, td *model.TokenDetails) error
	RedisGetAuth(authD *model.AccessDetails) (uint64, error)
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
