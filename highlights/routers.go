package highlights

import (
	"kindle-highlight/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	highlights := router.Group("/highlights")
	highlights.POST("/", create)

	highlights.Use(middlewares.AuthMiddleware(true)).GET("/", getAllByUserAndBook)

}
