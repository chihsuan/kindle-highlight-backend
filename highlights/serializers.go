package highlights

import "github.com/gin-gonic/gin"

type HighlightListSerializer struct {
	c *gin.Context
}

type HighlightResponse struct {
	ID        uint   `json:"id"`
	Highlight string `json:"content"`
}

func (serializer *HighlightListSerializer) Serialize(noteModels []HighlightModel) []HighlightResponse {
	response := []HighlightResponse{}
	for _, n := range noteModels {
		response = append(response, HighlightResponse{
			ID:        n.ID,
			Highlight: n.Highlight,
		})
	}
	return response
}
