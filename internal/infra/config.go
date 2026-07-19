package infra

import (
	"os"

	"github.com/lostcra/simpleHttpserver/internal/secret"
)

type pgConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
}

func defaultConfig() *pgConfig {
	return &pgConfig{
		UserName: "postgres",
		Host:     "db",
		Port:     "5432",
		DBName:   "postgres",
	}
}

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

func getEnvOrDefault(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}

func LoadEnvConfig() (*pgConfig, error) {
	conf := defaultConfig()

	pwd, err := getDBPassword()
	if err != nil {
		return nil, err
	}

	conf.Host = getEnvOrDefault("DB_HOST", conf.Host)
	conf.Port = getEnvOrDefault("DB_PORT", conf.Port)
	conf.UserName = getEnvOrDefault("DB_USER", conf.UserName)
	conf.DBName = getEnvOrDefault("DB_NAME", conf.DBName)
	conf.Password = pwd

	return conf, nil
}
