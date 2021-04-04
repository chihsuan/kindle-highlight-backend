package books

import (
	"github.com/gin-gonic/gin"
)

type BookSerializer struct {
	c *gin.Context
}

type BookResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (serializer *BookSerializer) Serialize(bookModel BookModel) BookResponse {
	book := BookResponse(bookModel)
	return book
}
