package model

import (
	"regexp"
	"strconv"
	"unicode"
)

func validatePassword(password string) error {
	if len(password) < 8 {
		return &fillingError{"Password must be at least 8 characters"}
	} else if len(password) > 100 {
		return &fillingError{"Password should not exceed 100 characters"}
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return &fillingError{"Password must contain at least one uppercase letter"}
	}

	if !hasLower {
		return &fillingError{"Password must contain at least one lowercase letter"}
	}

	if !hasDigit {
		return &fillingError{"Password must contain at least one digit"}
	}

	if !hasSpecial {
		return &fillingError{"Password must contain at least one special character"}
	}

	return nil
}

func validateEmail(email string) error {
	if len(email) < 1 || len(email) > 320 {
		return &fillingError{"Invalid email length"}
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if match, _ := regexp.MatchString(pattern, email); !match {
		return &fillingError{"Invalid email address"}
	}

	return nil
}

func validateName(name string) error {
	if len(name) < 4 || len(name) > 100 {
		return &fillingError{"Invalid name length"}
	}

	pattern := `^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`

	if match, _ := regexp.MatchString(pattern, name); !match {
		return &fillingError{"Invalid name"}
	}

	return nil
}

func atoi64(s string) int64 {
	numInt, _ := strconv.Atoi(s)
	return int64(numInt)
}
