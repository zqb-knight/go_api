package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/pkg/redis_utils"
	"log"
)

func GetTags(c *gin.Context) {

	rd := redis_utils.InitRedis()
	if rd == nil {
		log.Println("redis初始化失败")
		return
	}
	c.JSON(200, gin.H{
		"msg": "test",
	})
	//
	//response, err := rd.Get("lm").Result()
	//if err != nil{
	//	log.Println("redis获取值失败")
	//}
	//log.Println(response)
	//c.JSON(200, gin.H{
	//	"msg": "test",
	//})

}
