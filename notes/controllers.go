package notes

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

	note := NoteModel{BookID: validator.BookID, UserID: validator.UserID, Content: validator.Content}
	_, err := FindOneOrCreate(&note, note)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, "ok")
}

func getAllByUserAndBook(c *gin.Context) {
	validator := GetAllValidator{}
	if err := c.BindQuery(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := NoteModel{BookID: validator.BookID, UserID: validator.UserID}

	notesInDB, err := FindAll(&note)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	noteListSerializer := NoteListSerializer{c}
	c.JSON(http.StatusOK, gin.H{"notes": noteListSerializer.Serialize(notesInDB)})
}
