package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/igorlev91/GlobantBookStore/source/objects"
)

var GetBookByIdMethod = func(w http.ResponseWriter, r *http.Request) {

}

var CreateBookMethod = func(w http.ResponseWriter, r *http.Request) {
	book := &objects.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO

	w.WriteHeader(http.StatusCreated)
	id := []byte(strconv.FormatInt(int64(book.Id), 10))
	w.Write([]byte(id))
}
