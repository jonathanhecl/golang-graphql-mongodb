package main

import (
	"errors"
	"regexp"
	"strings"
)

// isRequired
func isRequired(list map[string]string) error {
	var err error = nil

	// Email Format
	email := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	for k, v := range list {
		if len(v) == 0 {
			err = errors.New(strings.Title(k) + " required.")
			return err
		}
		if strings.ToLower(k) == "email" {
			if !email.MatchString(v) {
				err = errors.New(strings.Title(k) + " invalid.")
				return err
			}
		}
	}
	return err
}
