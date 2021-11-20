package objects

import (
	"errors"
	"net/http"

	"github.com/igorlev91/GlobantBookStore/database"
	"github.com/upper/db/v4"
)

type Book struct {
	Id     int32   `json,db:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Genre  uint    `json:"genre"`
	Amount uint    `json:"amount"`
}

type Genre struct {
	Id   uint   `json,db:"id,omitempty"`
	Name string `json:"name"`
}

type Endpoints interface {
	GetBookByIdMethod(idParam string) func(w http.ResponseWriter, r *http.Request)

	CreateBookMethod() func(w http.ResponseWriter, r *http.Request)

	GetBooksByfiltering() func(w http.ResponseWriter, r *http.Request)

	UpdateBookMethod(idParam string) func(w http.ResponseWriter, r *http.Request)

	DeleteBookMethod(idParam string) func(w http.ResponseWriter, r *http.Request)
}

func (b *Book) Validate() error {

	books := database.GetSession().Collection("book")

	name_duplications, _ := books.Find(db.Cond{"name": b.Name}).Count()
	if name_duplications != 0 {
		return errors.New("Name is not unique")
	}

	switch {
	case b.Price < 0:
		return errors.New("Bad price")
	case b.Amount < 0:
		return errors.New("Bad amount")
	default:
		return nil
	}
}

func GetBook(book_id int32) (*Book, error) {
	books := database.GetSession().Collection("book")
	result := &Book{}

	err := books.Find(db.Cond{"id": book_id}).One(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
