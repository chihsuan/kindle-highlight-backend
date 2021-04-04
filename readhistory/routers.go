package readhistory

import (
	"kindle-highlight/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	readHistory := router.Group("/readhistory")
	readHistory.POST("/", create)
	users := readHistory.Group("/user").Use(middlewares.AuthMiddleware(true))
	users.GET("/", getAllByUser)
}
