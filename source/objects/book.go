package objects

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	Id     uint    `json,db:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Genre  uint    `json:"genre"`
	Amount uint    `json:"amount"`
}

type Genre struct {
	Id   uint   `json,db:"id,omitempty"`
	Name string `json:"name"`
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
