package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(empty bool, w http.ResponseWriter, status int, book interface{}) {
	var response []byte
	var err error
	if empty {
		response, _ = json.Marshal(book)
	} else {
		response, err = json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	TextResponse(w, status, response)
}

func TextResponse(w http.ResponseWriter, httpStatus int, body []byte) {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	log.Println(string(body))
	w.Write(body)
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(false, w, code, map[string]string{"error": message})
}
