package routes

import (
	"net/http"

	"github.com/MelihEmreGuler/login-register-app-with-golang/handlers"
)

func Routes() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/signup", handlers.Signup)
}
