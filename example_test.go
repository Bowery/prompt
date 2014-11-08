package prompt_test

import (
	"flag"
	"net/mail"
	"os"

	"github.com/Bowery/prompt"
	"golang.org/x/crypto/bcrypt"
)

func ExamplePassword() ([]byte, error) {
	clear, err := prompt.Password("Password")
	if err != nil {
		return nil, err
	}

	return bcrypt.GenerateFromPassword(clear, bcrypt.DefaultCost)
}

func ExamplePrompt() (string, error) {
	email, err := prompt.Prompt("Email")
	if err != nil {
		return "", err
	}

	_, err = mail.ParseAddress(email)
	return email, err
}
