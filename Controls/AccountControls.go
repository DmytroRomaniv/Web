package Controls

import (
	"../Models"
	"../Models/Constants"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
)

func Authorize (writer http.ResponseWriter, request *http.Request)  {
	sendGet := func() {
		http.ServeFile(writer, request, "./View/Index.html")
	}

	switch request.Method {
	case "GET":
		sendGet()
	case "POST":
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm error: %v", err)
			return
		}

		login := request.FormValue("login")
		password := request.FormValue("password")
		if login != Constants.EmptyValue && password != Constants.EmptyValue {
			fmt.Fprintf(writer, "Login = %s\n", login)
			fmt.Fprintf(writer, "Password = %s\n", password)
		} else {
			sendGet()
		}
	default:
		http.Error(writer, "Unknown request!", http.StatusNotFound)
	}
}

func CreateAccount (writer http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET":
		http.ServeFile(writer, request, "./View/AccountCreation.html")
	case "POST":
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm error: %v", err)
			return
		}


	default:
		http.Error(writer, "Unknown request!", http.StatusNotFound)
	}
}

func getAccountDataFromForm(request http.Request) (Models.Account, []error){
	login := request.FormValue("login")
	password := request.FormValue("password")
	confirmationPassword := request.FormValue( "confirmationPassword")
	name := request.FormValue("name")
	surname := request.FormValue("surname")
	email := request.FormValue("email")

	//TODO: Validate data
	emailError := validateEmail(email)
	loginError := validateLogin(login)
	passwordError := validatePassword(password, confirmationPassword)

	if emailError == nil && loginError == nil && passwordError == nil {
		newAccount := Models.Account{
			Id: 0,
			Login: login,
			Password: password,
			Name: name,
			Surname: surname,
			Email: email,
		}

		return newAccount, nil
	} else {
		return Models.Account{}, []error{emailError, loginError, passwordError}
	}
	//TODO: Create account

	return Models.Account{}, nil
}

func validatePassword(password string, confirmationPassword string) error {
	if password == confirmationPassword {
		return nil
	} else {
		return errors.New("Passwords are not equal.")
	}
}

func validateEmail(email string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if emailRegex.MatchString(email) {
		return nil
	} else {
		return errors.New("Wrong email address.")
	}
}

func validateLogin(login string) error {
	//TODO: Get data from DB

	if login != "" {
		return nil
	} else {
		return errors.New( "User with current login already exists.")
	}
}