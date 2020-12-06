package prompt_test

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"github.com/Bowery/prompt"
)

func ExamplePassword() {
	clear, err := prompt.Password("Password")
	if err != nil {
		// handle error
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(clear), bcrypt.DefaultCost)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(password)
	// Output:
	fmt.Println(err)
	// Output:

	return
}

func ExampleBasic() {
	email, err := prompt.Basic("Email", true)
	if err != nil {
		// handle error
		return
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(email)
	// Output:
	fmt.Println(err)
	// Output:
}
