package database

import (
	"encoding/json"
	"fmt"
	"net/http"

	//	"gorm.io/gorm"

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
	handers.RespondJSON(true, w, http.StatusOK, book.Id)
}

func (db_book *Database) GetBookByIdMethod(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	book, exists, err := objects.GetBookByID(id, db_book.Connetion)
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

func (db_book *Database) GetBooksByFilterMethod(w http.ResponseWriter, r *http.Request) {
	books, err := objects.GetAllBooks(db_book.Connetion)
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
