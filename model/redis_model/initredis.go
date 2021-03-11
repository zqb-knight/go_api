package redis_model

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_api/common"
	"log"
)

var (
	addr string
	port string
)

/*
	@ 初始化redis
	@return 返回redis.client
*/
func InitRedis() *redis.Client {
	common.InitConfig("")
	//配置文件加载失败，则使用默认配置
	if common.Viper == nil {
		log.Println("配置加载失败，使用默认配置")
		addr = "127.0.0.1"
		port = "6379"
	} else {
		addr = common.Viper.GetString("redis.addr")
		port = common.Viper.GetString("redis.port")
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: "",
		DB:       0,
	})
	_, err := cli.Ping().Result()
	if err != nil {
		return nil
	}
	return cli
}
