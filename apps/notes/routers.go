package notes

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	notes := router.Group("/notes")
	notes.POST("/", create)
	notes.GET("/", getAllByUserAndBook)
}
