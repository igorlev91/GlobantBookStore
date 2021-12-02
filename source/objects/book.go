package objects

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	Id      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name" gorm:"unique;size:100;not null;"`
	Price   float32 `json:"price" gorm:"not null"`
	GenreID uint    `json:"genre_id" gorm:"not null"`
	Amount  uint    `json:"amount" gorm:"not null"`
	Genre   Genre   `gorm:"foreignKey:GenreID"`
}

type Genre struct {
	Id   uint   `json:"id" gorm:"gorm:primaryKey"`
	Name string `json:"name" gorm:"unique;type:varchar(100);not null"`
}

func GetBookByID(id uint, db *gorm.DB) (Book, bool, error) {
	book := Book{}

	err := db.First(&book, Book{Id: id}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return book, false, err
	}

	return book, true, nil
}

func GetAllBooks(db *gorm.DB) ([]Book, error) {
	books := []Book{}

	return books, nil
}
