package users

import (
	"time"

	"kindle-notes/apps/common"
	"kindle-notes/apps/notes"
	"kindle-notes/apps/readhistory"
)

type UserModel struct {
	ID               uint   `gorm:"primary_key"`
	Email            string `gorm:"column:email;unique_index"`
	ReadingHistories []readhistory.ReadHistoryModel
	Notes            []notes.NoteModel
	CreatedAt        time.Time // Set to current time if it is zero on creating
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&UserModel{})
}

// You could input the conditions and it will return an UserModel in database with error info.
// 	userModel, err := FindOneUser(&UserModel{Username: "username0"})
func FindOne(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// You could input an UserModel which will be saved in database returning with error info
// 	if err := SaveOne(&userModel); err != nil { ... }
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
