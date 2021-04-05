package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	validator := CreateValidator{}
	if err := c.Bind(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := BookModel{Name: validator.Name, Author: validator.Author}
	bookInDB, err := FindOneOrCreate(&book, book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	serializer := BookSerializer{c}
	c.JSON(http.StatusOK, gin.H{"book": serializer.Serialize(bookInDB)})
}
