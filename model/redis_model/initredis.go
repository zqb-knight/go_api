package redis_model

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_api/common/handlers/config"
	"log"
)

var (
	redisClient *redis.Client
)

/*
	@ 初始化redis
	@return 返回redis.client
*/
func InitRedis() error {
	var (
		addr string
		port string
	)
	config.InitConfig("")
	//配置文件加载失败，则使用默认配置
	if config.Viper == nil {
		log.Println("配置加载失败，使用默认配置")
		addr = "127.0.0.1"
		port = "6379"
	} else {
		addr = config.Viper.GetString("redis.addr")
		port = config.Viper.GetString("redis.port")
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
