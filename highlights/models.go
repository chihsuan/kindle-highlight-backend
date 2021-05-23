package highlights

import (
	"kindle-highlight/books"
	"kindle-highlight/common"
)

type HighlightModel struct {
	ID        uint `gorm:"primary_key"`
	Highlight string
	BookID    uint
	UserID    uint
	Book      books.BookModel
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&HighlightModel{})
}

func FindAll(condition interface{}) ([]HighlightModel, error) {
	db := common.GetDB()
	var models []HighlightModel
	err := db.Where(condition).Find(&models).Error
	return models, err
}

func DeleteAll(userID uint, bookID uint) {
	db := common.GetDB()
	db.Debug().Delete(HighlightModel{}, "user_id = ? and book_id = ?", userID, bookID)
}

func FindOneOrCreate(condition interface{}, highlight HighlightModel) (HighlightModel, error) {
	highlightInDB, err := FindOne(condition)
	if err != nil {
		highlightInDB, err := SaveOne(highlight)
		return highlightInDB, err
	}
	return highlightInDB, err
}

func FindOne(condition interface{}) (HighlightModel, error) {
	db := common.GetDB()
	var model HighlightModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(highlight HighlightModel) (HighlightModel, error) {
	db := common.GetDB()
	err := db.Save(&highlight).Error
	return highlight, err
}

func SaveAll(highlights []HighlightModel) ([]HighlightModel, error) {
	db := common.GetDB()
	err := db.Create(&highlights).Error
	return highlights, err
}
