package route

import (
	"go_web/api"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	// 用户组
	user := r.Group("/user")
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.Login)
	}

	return r
}
