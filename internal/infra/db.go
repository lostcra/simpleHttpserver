package infra

import (
	"errors"
	"os"
	"strings"
)

func readPassword(path string) (string, error) {
	password, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(password)), nil
}

func getDBPassword() (string, error) {
	path := os.Getenv("POSTGRES_PASSWORD_FILE")
	if path == "" {
		return "", errors.New("POSTGRES_PASSWORD_FILE isn't passed")
	}
	password, err := readPassword(path)
	if err != nil {
		return "", err
	}
	return password, nil
}
