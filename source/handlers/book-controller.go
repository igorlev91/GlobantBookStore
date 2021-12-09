package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/igorlev91/GlobantBookStore/source/database"
	"github.com/igorlev91/GlobantBookStore/source/objects"
	"gorm.io/gorm"
)

type ORM struct {
	DB *gorm.DB
}

func (db_book *ORM) CreateBookMethod(w http.ResponseWriter, r *http.Request) {
	book := objects.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db_book.DB.Create(&book).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(true, w, http.StatusOK, book.BookID)
}

func (db_book *ORM) GetBookByIdMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sId := vars["id"]

	book_id, err := strconv.ParseUint(sId, 10, 32)
	if err != nil {
		panic(err)
	}

	book, ok, err := database.GetBookByID(uint(book_id), db_book.DB)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(true, w, http.StatusOK, book)
}

func (a *ORM) GetBooksByFilterMethod(w http.ResponseWriter, r *http.Request) {

	urlParams := make(map[string]string)

	v := r.URL.Query()
	urlParams["name"] = v.Get("name")
	urlParams["genreId"] = v.Get("genreId")

	books, err := database.GetAllBooks(urlParams, a.DB)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	RespondJSON(true, w, http.StatusOK, books)
}

func (a *ORM) UpdateBookMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sId := vars["id"]

	resUint32, err := strconv.ParseUint(sId, 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	id := uint(resUint32)

	_, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if !exists {
		RespondError(w, http.StatusNotFound, "record not found")
		return
	}

	updatedBook := objects.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedBook); err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, "bad request") // or http.StatusInternalServerError?
		return
	}
	defer r.Body.Close()

	if err := database.UpdateBook(a.DB, &updatedBook); err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	a.GetBookByIdMethod(w, r)
}

func (a *ORM) DeleteBookMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sId := vars["id"]

	resUint32, err := strconv.ParseUint(sId, 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	id := uint(resUint32)

	_, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if !exists {
		RespondError(w, http.StatusNotFound, "record not found")
		return
	}

	err = database.DeleteBook(id, a.DB)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	RespondJSON(true, w, http.StatusOK, nil)
}
