package handlers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("to signup: /signup \nto login: /login \n "))
}
