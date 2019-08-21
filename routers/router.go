package routers

import (
	"awesome_gin/controller"
	"awesome_gin/controller/v1/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(controller.NoRoute)

	v1 := r.Group("/v1")
	{
		v1.POST("/user/register", user.Register)
	}
	return r
}
