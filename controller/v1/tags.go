package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/common/e"
	"go_api/common/handlers/responese"
	"go_api/model"
	"log"
	"net/http"
)

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	var res string
	name := c.Query("name")
	var maps = make(map[string]interface{})
	maps["name"] = name
	rep, err := model.GetTags(maps)
	if err != nil {
		log.Println("数据库查询失败")
		c.JSON(300, gin.H{
			"errorNo":  101,
			"errorMsg": "数据库错误",
		})
	}
	if len(rep) == 0 {
		res = "未查询到创建者"
	}
	res = rep[0].CreatedBy
	c.JSON(200, gin.H{
		"msg": res,
	})

}

//添加文章标签
func AddTag(c *gin.Context) {
	resp := responese.Gin{C: c}
	name := c.Query("name")
	createdBy := c.Query("created_by")
	err := model.AddTag(name, createdBy)
	if err != nil {
		resp.BuildResponse(http.StatusAccepted, e.ERROR_DATABASE_INSERT_FAIL, nil)
	}
	resp.BuildResponse(http.StatusOK, e.SUCCESS, nil)

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
