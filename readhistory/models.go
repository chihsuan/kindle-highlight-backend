package readhistory

import (
	"time"

	"kindle-notes/books"
	"kindle-notes/common"
)

type ReadHistoryModel struct {
	UserID    uint `gorm:"primaryKey"`
	BookID    uint `gorm:"primaryKey"`
	Book      books.BookModel
	CreatedAt time.Time // Set to current time if it is zero on creating
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&ReadHistoryModel{})
}

func FindAll(condition interface{}) ([]ReadHistoryModel, error) {
	db := common.GetDB()
	var models []ReadHistoryModel
	err := db.Where(condition).Preload("Book").Find(&models).Error
	return models, err
}

func FindOneOrCreate(condition interface{}, readhistory ReadHistoryModel) (ReadHistoryModel, error) {
	readhistoryInDB, err := FindOne(condition)
	if err != nil {
		readhistoryInDB, err := SaveOne(readhistory)
		return readhistoryInDB, err
	}
	return readhistoryInDB, err
}

func FindOne(condition interface{}) (ReadHistoryModel, error) {
	db := common.GetDB()
	var model ReadHistoryModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(readhistory ReadHistoryModel) (ReadHistoryModel, error) {
	db := common.GetDB()
	err := db.Save(&readhistory).Error
	return readhistory, err
}
