package eksternal

import (
	"context"
	"log"
	"time"

	"github.com/farhaniupr/dating-api/package/library"
)

type RedisEksternal struct {
	env   library.Env
	redis library.RedisDatabase
}

func ModuleRedisEksternal(
	env library.Env,
	redis library.RedisDatabase,

) RedisEksternal {
	return RedisEksternal{
		env:   env,
		redis: redis,
	}
}

func (r RedisEksternal) Store(key string, value interface{}) bool {

	set, err := r.redis.SetNX(context.Background(), key, value, 300*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	if set {
		return true
	} else {
		return false
	}
}

func (r RedisEksternal) Get(key string) string {

	set, err := r.redis.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return set
}

func (r RedisEksternal) GetList(key string) []string {

	resultKey, err := r.redis.Keys(context.Background(), key).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return resultKey
}

func (r RedisEksternal) Delete(key string) int64 {

	set, err := r.redis.Del(context.Background(), key).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return set
}

func (r RedisEksternal) Store1Day(key string, value interface{}) bool {

	// time := 86400 * time.Second

	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	durationUntilEndOfDay := endOfDay.Sub(now)

	timeExpired := durationUntilEndOfDay

	set, err := r.redis.SetNX(context.Background(), key, value, timeExpired).Result()
	_ = err
	if set {
		return true
	} else {
		return false
	}
}
