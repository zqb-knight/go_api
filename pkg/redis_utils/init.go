package redis_utils

import "github.com/go-redis/redis"

/*
	初始化redis
*/
func InitRedis() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := cli.Ping().Result()
	if err != nil {
		return nil
	}
	return cli
}
