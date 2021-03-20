package responese

import (
	"github.com/gin-gonic/gin"
	"go_api/common/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func (g *Gin) BuildResponse(httpCode int, errno int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Errno: errno,
		Msg:   e.GetMsg(errno),
		Data:  data,
	})

	return
}
