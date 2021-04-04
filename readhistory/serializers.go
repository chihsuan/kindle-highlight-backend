package readhistory

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ReadHistorySerializer struct {
	c *gin.Context
}

type ReadHistoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

func (serializer *ReadHistorySerializer) Serialize(readHistoryModels []ReadHistoryModel) interface{} {
	response := []ReadHistoryResponse{}
	for _, r := range readHistoryModels {
		response = append(response, ReadHistoryResponse{
			ID:        r.Book.ID,
			Name:      r.Book.Name,
			Author:    r.Book.Author,
			CreatedAt: r.CreatedAt,
		})
	}
	return response
}
