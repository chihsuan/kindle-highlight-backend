package books

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	books := router.Group("/books")
	books.POST("/", create)
}
