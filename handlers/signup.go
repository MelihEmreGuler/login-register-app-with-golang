package handlers

import (
	"fmt"
	"net/http"

	"github.com/MelihEmreGuler/login-register-app-with-golang/helpers"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	uName, email, psw, pwdCnfrm := "", "", "", ""
	r.ParseForm()

	uName = r.FormValue("username")
	email = r.FormValue("email")
	psw = r.FormValue("password")
	pwdCnfrm = r.FormValue("confirm")

	if CheckSignupFormValues(uName, email, psw, pwdCnfrm, w) {
		fmt.Fprintf(w, "%s's Registration successful \n", uName)
		ShowSignedUpUserData(uName, email, psw, pwdCnfrm, w)
	} else {
		fmt.Fprintf(w, "Registration failed \n")
		return
	}

}

func CheckSignupFormValues(uName, email, psw, pwdCnfrm string, w http.ResponseWriter) bool {
	uNameCheck := helpers.IsEmpty(uName)
	emailCheck := helpers.IsEmpty(email)
	pswCheck := helpers.IsEmpty(psw)
	pwdCnfrmCheck := helpers.IsEmpty(pwdCnfrm)

	if uNameCheck || emailCheck || pswCheck || pwdCnfrmCheck {
		fmt.Fprintf(w, "Please fill all the fields \n")
		return false
	}
	if psw != pwdCnfrm {
		fmt.Fprintf(w, "Passwords do not match \n")
		return false
	}
	return true
}

func ShowSignedUpUserData(uName, email, psw, pwdCnfrm string, w http.ResponseWriter) {

	fmt.Fprintf(w, "Username: %s \n", uName)
	fmt.Fprintf(w, "Email: %s \n", email)
	fmt.Fprintf(w, "Password: %s \n", psw)
	fmt.Fprintf(w, "Password Confirmation: %s \n", pwdCnfrm)

}

// for key, value := range r.Form {
// 	fmt.Println("key:", key, "value:", value)
// }
