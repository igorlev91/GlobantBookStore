package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/handers"
	"github.com/igorlev91/GlobantBookStore/source/objects"
)

type BookORM struct {
	DB *gorm.DB
}

func (db_book *BookORM) CreateBookMethod(w http.ResponseWriter, r *http.Request) {
	book := objects.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db_book.DB.Create(&book).Error; err != nil {
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handers.RespondJSON(true, w, http.StatusOK, book.Id)
}

func (db_book *BookORM) GetBookByIdMethod(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	book, exists, err := objects.GetBookByID(id, db_book.DB)
	if err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		handers.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	handers.RespondJSON(true, w, http.StatusOK, book)
}

func (db_book *BookORM) GetBooksByFilterMethod(w http.ResponseWriter, r *http.Request) {
	books, err := objects.GetAllBooks(db_book.DB)
	if err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handers.RespondJSON(true, w, http.StatusOK, books)
}

func getId(r *http.Request) uint {
	vars := mux.Vars(r)
	sId := vars["id"]
	return handers.StringToUint(sId)
}
