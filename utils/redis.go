package utils

import (
	"os"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("aws.redis"),
	})
	return rdb
}

func RSet(rd *redis.Client, k string, a string) (string, error) {
	return rd.Set(k, a, 0).Result()
}

func RGet(rd *redis.Client, k string) (int64, error) {
	return rd.Get(k).Int64()
}
func RDel(rd *redis.Client, k []string) *redis.IntCmd {
	return rd.Del(k...)
}

func RGetAll(rd *redis.Client, k string) (interface{}, error) {
	var cursor uint64
	b, _, err := rd.Scan(cursor, k, 0).Result()
	var arr []interface{}
	if err != nil {
		return nil, err
	}

	for _, key := range b {
		i, _ := RGet(rd, key)
		arr = append(arr, i)
	}

	return arr, nil
}

// RDelAll ...ƒ
func RDelAll(rd *redis.Client, k string) error {
	var cursor uint64
	b, _, err := rd.Scan(cursor, k, 0).Result()
	if err != nil {
		return err
	}

	RDel(rd, b)
	return nil
}
