package notes

import (
	"kindle-notes/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	notes := router.Group("/notes")
	notes.POST("/", create)

	notes.Use(middlewares.AuthMiddleware(true)).GET("/", getAllByUserAndBook)

}
