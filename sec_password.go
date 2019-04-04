package main

import (
	"golang.org/x/crypto/bcrypt"
)

// Generate Password Hash
// Example:
/*
	passwordHash, _ := generatePassword("test")
	log.Print("Password Hash: ", passwordHash)
*/
func generatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// Compare Password
// Example:
/*
	validPassword, _ := comparePassword("test", passwordHash)
	log.Print("Valid Password: ", validPassword)
*/
func comparePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
