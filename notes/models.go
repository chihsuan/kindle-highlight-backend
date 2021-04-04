package notes

import (
	"kindle-notes/books"
	"kindle-notes/common"
)

type NoteModel struct {
	ID      uint `gorm:"primary_key"`
	Content string
	BookID  uint
	UserID  uint
	Book    books.BookModel
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&NoteModel{})
}

func FindAll(condition interface{}) ([]NoteModel, error) {
	db := common.GetDB()
	var models []NoteModel
	err := db.Where(condition).Find(&models).Error
	return models, err
}

func FindOneOrCreate(condition interface{}, book NoteModel) (NoteModel, error) {
	bookInDB, err := FindOne(condition)
	if err != nil {
		bookInDB, err := SaveOne(book)
		return bookInDB, err
	}
	return bookInDB, err
}

func FindOne(condition interface{}) (NoteModel, error) {
	db := common.GetDB()
	var model NoteModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(book NoteModel) (NoteModel, error) {
	db := common.GetDB()
	err := db.Save(&book).Error
	return book, err
}
