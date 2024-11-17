package router

import (
	"github.com/gin-gonic/gin"
	"gopaste/internal/http-server/controllers/user/addUser"
	"gopaste/internal/service/user"
)

func New(userService *user.Service) *gin.Engine {
	r := gin.Default()
	r.Group("/v1")
	{
		r.Group("/user")
		{
			r.POST("user/add", addUser.New(userService))
		}

	}

	return r
}
