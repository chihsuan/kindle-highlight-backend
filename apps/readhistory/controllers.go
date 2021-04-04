package readhistory

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

	readHistory := ReadHistoryModel{BookID: validator.BookID, UserID: validator.UserID}

	if _, err := FindOneOrCreate(&readHistory, readHistory); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, "ok")
}

func getAllByUser(c *gin.Context) {
	validator := GetAllValidator{}
	if err := c.Bind(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readHistory := ReadHistoryModel{UserID: validator.UserID}

	readHistoriesInDB, err := FindAll(&readHistory)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	readHistorySerializer := ReadHistorySerializer{c}
	c.JSON(http.StatusOK, gin.H{"readHistories": readHistorySerializer.Serialize(readHistoriesInDB)})
}
