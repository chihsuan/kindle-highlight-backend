package notes

import "github.com/gin-gonic/gin"

type NoteListSerializer struct {
	c *gin.Context
}

type NoteResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func (serializer *NoteListSerializer) Serialize(noteModels []NoteModel) []NoteResponse {
	response := []NoteResponse{}
	for _, n := range noteModels {
		response = append(response, NoteResponse{
			ID:      n.ID,
			Content: n.Content,
		})
	}
	return response
}
