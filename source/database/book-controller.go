package database

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/handers"
	"github.com/igorlev91/GlobantBookStore/source/objects"
)

func (db_book *Database) CreateBookMethod(w http.ResponseWriter, r *http.Request) {
	book := objects.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db_book.Connetion.Create(&book).Error; err != nil {
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handers.RespondJSON(true, w, http.StatusOK, book.BookID)
}

func (db_book *Database) GetBookByIdMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sId := vars["id"]

	book_id, err := strconv.ParseUint(sId, 10, 32)
	if err != nil {
		panic(err)
	}

	book, ok, err := objects.GetBookByID(uint(book_id), db_book.Connetion)
	if err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		handers.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	handers.RespondJSON(true, w, http.StatusOK, book)
}

func (db_book *Database) GetBooksByFilterMethod(w http.ResponseWriter, r *http.Request) {
	books, err := objects.GetAllBooks(db_book.Connetion)
	if err != nil {
		fmt.Println(err)
		handers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handers.RespondJSON(true, w, http.StatusOK, books)
}
