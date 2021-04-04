package users

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/users", create)
}
