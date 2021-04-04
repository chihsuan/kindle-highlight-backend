package readhistory

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	readHistory := router.Group("/readHistory")
	readHistory.POST("/", create)
	readHistory.GET("/user", getAllByUser)
}
