package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
)

func dsnString() (string, error) {
	conf, err := LoadEnvConfig()
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.UserName, conf.Password, conf.Host, conf.Port, conf.DBName)

	return dsn, nil
}

func ConnectToDB() (*pgx.Conn, error) {
	dsn, err := dsnString()
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
