package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/model/redis_model"
	"log"
)

//@Summary 测试接口
//@Produce json
//@Description 从redis获取值并返回
//@Success 200 {string} json "{"code":200, "msg":"667"}"
//@Router /api/v1/tags [get]

func GetTags(c *gin.Context) {

	rd := redis_model.InitRedis()
	if rd == nil {
		log.Println("redis初始化失败")
		return
	}

	response, err := rd.Get("xj").Result()
	if err != nil {
		log.Println("redis获取值失败")
		c.JSON(200, gin.H{
			"msg": "请重试",
		})
		return

	}
	log.Println(response)
	c.JSON(200, gin.H{
		"msg": response,
	})

}
