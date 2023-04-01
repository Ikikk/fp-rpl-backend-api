package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPwd string, password string) (bool, error) {
	byteHash := []byte(hashedPwd)
	plainPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		fmt.Println("hash di utilnya salah")
		return false, err
	}
	return true, nil
}

func HashAndSalt(p string) (string, error) {
	pwd := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
