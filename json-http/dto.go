package jsonhttp

import (
	"errors"
	"net/mail"
)

// User type
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	User    *User  `json:"user"`
}

func validate(in User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
