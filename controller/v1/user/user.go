package user

import (
	"awesome_gin/controller"
	"awesome_gin/models"
	"awesome_gin/pkg/e"
	"awesome_gin/pkg/logging"
	"github.com/gin-gonic/gin"
)

type regPara struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var (
		para regPara
		err error
	)
	if err = c.Bind(&para); err != nil{
		controller.Response(c, e.ERROR_INVLIAD_PARA, nil)
		return
	}
	if err := models.AddUser(models.UsersInfo{Username:para.Username, Password:para.Password, Avatar:para.Avatar}); err != nil{
		logging.Warn("Register", err)
		controller.Response(c, e.ERROR, nil)
		return
	}
	controller.Response(c, e.SUCCESS, nil)
	return
}
