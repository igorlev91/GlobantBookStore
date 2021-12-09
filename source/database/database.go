package database

import (
	"errors"

	"github.com/igorlev91/GlobantBookStore/source/objects"
	"gorm.io/gorm"
)

func GetBookByID(id uint, db *gorm.DB) (objects.Book, bool, error) {
	book := objects.Book{}

	err := db.Preload(objects.GENRE).First(&book, objects.Book{BookID: id}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, nil
	}

	return book, true, nil
}

func GetAllBooks(urlParams map[string]string, db *gorm.DB) ([]objects.Book, error) {
	books := []objects.Book{}

	if err := db.Debug().Preload(objects.GENRE).Order("name").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func DeleteBook(id uint, db *gorm.DB) error {
	var book objects.Book

	if err := CheckDatabase(db); err != nil {
		return err
	}

	if err := db.Where("id=?", id).First(&book).Error; err != nil {
		return err
	}

	if err := db.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil

}

func UpdateBook(db *gorm.DB, book *objects.Book) error {
	if err := CheckDatabase(db); err != nil {
		return err
	}

	return nil

}

func CheckDatabase(db *gorm.DB) error {
	if db == nil {
		return errors.New("Session isnt exists")
	}
	return nil
}
