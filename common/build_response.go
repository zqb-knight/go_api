package common

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Errno int         `json:"errno"`
	Data  interface{} `json:"data"`
}

func (g *Gin) buildResponse(httpCode int, errno int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Errno: errno,
		Data:  data,
	})

	return
}
