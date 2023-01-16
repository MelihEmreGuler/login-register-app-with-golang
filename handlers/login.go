package handlers

import (
	"fmt"
	"net/http"

	"github.com/MelihEmreGuler/login-register-app-with-golang/helpers"
)

const (
	// login types
	typeEmail    = 0
	typeUserName = 1
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uNameOrEmail := r.FormValue("data")
	psw := r.FormValue("password")

	if CheckLoginFormValues(uNameOrEmail, psw, w) {
		fmt.Fprintf(w, "Login successful \n")
		ShowLoggedInUserData(uNameOrEmail, psw, w)
	} else {
		fmt.Fprintf(w, "Login failed \n")
		return
	}

}

func CheckLoginFormValues(uNameOrEmail, psw string, w http.ResponseWriter) bool {
	uNameOrEmailCheck := helpers.IsEmpty(uNameOrEmail)
	pswCheck := helpers.IsEmpty(psw)

	if uNameOrEmailCheck || pswCheck {
		fmt.Fprintf(w, "Please fill all the fields \n")
		return false
	}

	if !isDbAndFormDataSame(uNameOrEmail, psw, w) {
		return false
	}

	return true
}

func ShowLoggedInUserData(uNameOrEmail, psw string, w http.ResponseWriter) {

	if DetectDataType(uNameOrEmail) == typeEmail {
		fmt.Fprintf(w, "Email: %s \n", uNameOrEmail)
	} else {
		fmt.Fprintf(w, "Username: %s \n", uNameOrEmail)
	}
	fmt.Fprintf(w, "Password: %s \n", psw)
}

func DetectDataType(uNameOrEmail string) int {
	if helpers.IsEmail(uNameOrEmail) {
		return 0
	}
	return 1
}

func isDbAndFormDataSame(uNameOrEmail, psw string, w http.ResponseWriter) bool {
	dbUserName := "melihguler"
	dbEmail := "gulermelihemre@gmail.com"
	dbPsw := "54321"

	if DetectDataType(uNameOrEmail) == typeEmail {
		fmt.Fprintf(w, "Email detected \n")
		Email := uNameOrEmail

		if Email != dbEmail || psw != dbPsw {
			fmt.Fprintf(w, "Email or password is wrong \n")
			return false
		}

	} else {
		fmt.Fprintf(w, "Username detected \n")
		UserName := uNameOrEmail

		if UserName != dbUserName || psw != dbPsw {
			fmt.Fprintf(w, "Username or password is wrong \n")
			return false
		}
	}
	return true
}
