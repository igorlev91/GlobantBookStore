package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/igorlev91/GlobantBookStore/source/objects"
)

var GetBookByIdMethod = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.ParseInt(vars["id"], 10, 32)
	book, _ := objects.GetBook(int32(book_id))
	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(book)
}

var CreateBookMethod = func(w http.ResponseWriter, r *http.Request) {
	book := &objects.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//todo Insert
	w.WriteHeader(http.StatusCreated)
	id := []byte(strconv.FormatInt(int64(book.Id), 10))
	w.Write([]byte(id))
}
