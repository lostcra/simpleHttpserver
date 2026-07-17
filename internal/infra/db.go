package infra

import (
	"os"

	"github.com/lostcra/simpleHttpserver/internal/secret"
)

type PGConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
}

func loadPGConfig() {}

func getDBPassword() (string, error) {
	path := os.Getenv("POSTGRES_PASSWORD_FILE")

	fr, err := secret.New(path)
	if err != nil {
		return "", err
	}

	pwd, err := fr.ReadFile()
	if err != nil {
		return "", err
	}

	return pwd, nil
}
