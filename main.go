package main

import (
	"net/http"

	"github.com/MelihEmreGuler/login-register-app-with-golang/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
