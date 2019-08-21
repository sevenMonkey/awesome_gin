package controller

import (
	"awesome_gin/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Response(c *gin.Context, errCode int, data interface{}) {
	if data == nil{
		data = make(map[string]interface{})
	}
	c.JSON(http.StatusOK, Resp{Code:errCode, Data:data, Msg:e.GetMsg(errCode)})
}

func NoRoute(c *gin.Context) {
	Response(c, e.ERROR_INVALID_REQUEST, nil)
}
