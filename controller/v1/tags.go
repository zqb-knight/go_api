package v1

import (
	"github.com/gin-gonic/gin"
	//"go_api/model/redis_model"
	"go_api/model/mysql_model"
)

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	rep := mysql_model.GetTags()
	//rd := redis_model.InitRedis()
	//if rd == nil {
	//	log.Println("redis初始化失败")
	//	return
	//}
	//
	//response, err := rd.Get("xj").Result()
	//if err != nil {
	//	log.Println("redis获取值失败")
	//	c.JSON(200, gin.H{
	//		"msg": "请重试",
	//	})
	//	return
	//
	//}
	//log.Println(response)

	c.JSON(200, gin.H{
		"msg": rep.Name,
	})

}

//添加文章标签
func AddTag(c *gin.Context) {

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
