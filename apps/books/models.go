package books

import (
	"kindle-notes/apps/common"
)

type BookModel struct {
	ID     uint `gorm:"primary_key"`
	Name   string
	Author string
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&BookModel{})
}

func FindOneOrCreate(condition interface{}, book BookModel) (BookModel, error) {
	bookInDB, err := FindOne(condition)
	if err != nil {
		bookInDB, err := SaveOne(book)
		return bookInDB, err
	}
	return bookInDB, err
}

func FindOne(condition interface{}) (BookModel, error) {
	db := common.GetDB()
	var model BookModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(book BookModel) (BookModel, error) {
	db := common.GetDB()
	err := db.Save(&book).Error
	return book, err
}
