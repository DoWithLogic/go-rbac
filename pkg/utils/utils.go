package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func NewPointer[T comparable](value T) *T {
	return &value
}

func GetPointerValue[T comparable](ptr *T) T {
	if ptr == nil {
		var val T

		return val
	}

	return *ptr
}

func GetValueOrDefault[T comparable](ptr *T, defaultVal T) T {
	if ptr != nil {
		return *ptr
	}

	return defaultVal
}

func HashPassword(password string) (string, error) {
	// Generate a hash from the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
