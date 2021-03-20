package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/common/e"
	"go_api/common/handlers/responese"
	"go_api/common/handlers/toekn"
	"net/http"
)

func GetAuth(c *gin.Context) {
	resp := responese.Gin{C: c}
	userName := c.Query("username")
	passWord := c.Query("password")
	tokenString, err := toekn.GenerateToken(userName, passWord)
	if err != nil {
		c.JSON(e.ERROR, gin.H{
			"errno": e.ERROR_AUTH_TOKEN,
			"msg":   e.GetMsg(e.ERROR_AUTH_TOKEN),
			"err":   err.Error(),
		})
		return
	}

	resp.BuildResponse(http.StatusOK, e.SUCCESS, map[string]string{
		"token": tokenString,
	})

}
