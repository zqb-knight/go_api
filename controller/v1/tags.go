package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/model/mysql_model"
	"log"
)

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	var maps = make(map[string]interface{})
	maps["name"] = name
	rep := mysql_model.GetTags(maps)
	c.JSON(200, gin.H{
		"msg": rep[0].CreatedBy,
	})

}

//添加文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	createdBy := c.Query("created_by")
	err := mysql_model.AddTag(name, createdBy)
	if err != nil {
		log.Println("插入数据失败")
	}

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
