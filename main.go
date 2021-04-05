package main

import (
	"github.com/gin-gonic/gin"

	"kindle-highlight/books"
	"kindle-highlight/common"
	"kindle-highlight/highlights"
	"kindle-highlight/middlewares"
	"kindle-highlight/readhistory"
	"kindle-highlight/users"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	books.AutoMigrate()
	highlights.AutoMigrate()
	readhistory.AutoMigrate()
}

func SetupAPIRoutes(v1 *gin.RouterGroup) {
	v1.Use(middlewares.AuthMiddleware(false))
	users.Route(v1)
	books.Route(v1)
	readhistory.Route(v1)
	highlights.Route(v1)
}

func main() {
	db := common.Init()
	Migrate(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "content-type"},
	}))

	SetupAPIRoutes(r.Group("/api/v1"))

	healthCheckAuth := r.Group("/api/healthy")
	healthCheckAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
