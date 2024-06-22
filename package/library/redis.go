package library

import (
	redisv8 "github.com/go-redis/redis/v8"
)

type RedisDatabase struct {
	*redisv8.Client
}

func ModuleRedis() RedisDatabase {
	rdb := redisv8.NewClient(&redisv8.Options{
		Addr:     EnvGlobal.Redis.Url,
		Password: EnvGlobal.Redis.Password,
		DB:       0,
	})

	return RedisDatabase{
		rdb,
	}
}
