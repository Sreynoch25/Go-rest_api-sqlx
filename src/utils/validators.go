package utils

import (
	"errors"
	user_model "marketing/src/models/user"
	"regexp"

)

func ValidateEnail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(pattern)

	return re.MatchString(email)
}

func ValidatePhone(phone string) bool {
	pattern := `^\+?[0-9]{10,15}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(phone)
}

func ValidateUser(user *user_model.User) error {
	if user.UserName == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required  fields")
	}

	if !ValidateEnail(user.Email) {
		return errors.New("invalid email format")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}