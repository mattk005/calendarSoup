package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	dat, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling JSON	%s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("responding with 5XX error: %s", msg)
	}
	type returnError struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, returnError{
		Error: msg,
	})
}
