package highlights

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

	DeleteAll(validator.UserID, validator.BookID)

	highlights := []HighlightModel{}
	for _, highlight := range validator.Highlights {
		highlights = append(highlights, HighlightModel{BookID: validator.BookID, UserID: validator.UserID, Highlight: highlight})
	}
	_, err := SaveAll(highlights)
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
	userID := c.MustGet("current_user_id").(uint)
	highlight := HighlightModel{BookID: validator.BookID, UserID: userID}

	highlightsInDB, err := FindAll(&highlight)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	noteListSerializer := HighlightListSerializer{c}
	c.JSON(http.StatusOK, gin.H{"highlights": noteListSerializer.Serialize(highlightsInDB)})
}
