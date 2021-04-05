package users

import (
	"kindle-highlight/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/users", create)

	router.Group("/me").Use(middlewares.AuthMiddleware(true)).GET("/", getMe)
}
