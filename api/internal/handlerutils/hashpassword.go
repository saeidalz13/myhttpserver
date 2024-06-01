package handlerutils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) ([]byte, error) {
	bytesHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytesHashed, nil
}

func ComparePasswords(hashedPassword []byte, password string) error {
    return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be a minimum of 8 characters")
	}

	uppercaseRegex, err := regexp.Compile(`[A-Z]`)
	if err != nil {
		return err
	}
	digitRegex, err := regexp.Compile(`\d`)
	if err != nil {
		return err
	}

	if !uppercaseRegex.MatchString(password) || !digitRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter and one digit")
	}

	return nil
}
