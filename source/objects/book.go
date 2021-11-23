package objects

import (
	"net/http"
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
