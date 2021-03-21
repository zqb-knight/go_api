package v1

import (
	"github.com/gin-gonic/gin"
	"go_api/common/e"
	"go_api/common/handlers/responese"
	"go_api/common/handlers/token"
	"go_api/model"
	"net/http"
)

func GetAuth(c *gin.Context) {
	resp := responese.Gin{C: c}
	username := c.Query("username")
	password := c.Query("password")
	if ok := model.CheckAuth(username, password); !ok {
		resp.BuildResponse(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, nil)
	} else {
		tokenString, err := token.GenerateToken(username, password)
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
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

}
