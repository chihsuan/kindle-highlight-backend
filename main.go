package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"kindle-notes/apps/books"
	"kindle-notes/apps/common"
	"kindle-notes/apps/notes"
	"kindle-notes/apps/readhistory"
	"kindle-notes/apps/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	books.AutoMigrate()
	notes.AutoMigrate()
	readhistory.AutoMigrate()
}

func SetupAPIRoutes(v1 *gin.RouterGroup) {
	v1.Use(users.AuthMiddleware(false))
	users.Route(v1)
	books.Route(v1)
	readhistory.Route(v1)
	notes.Route(v1)
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	SetupAPIRoutes(r.Group("/api/v1"))

	healthCheckAuth := r.Group("/api/healthy")
	healthCheckAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
