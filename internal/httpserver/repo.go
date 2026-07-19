package httpserver

import (
	"github.com/jackc/pgx/v5"
	"github.com/lostcra/simpleHttpserver/internal/infra"
)

type dbRepo struct {
	DB *pgx.Conn
}

func newDBRepo() *dbRepo {
	conn, err := infra.ConnectToDB()
	if err != nil {
		panic("faild connection to db")
	}

	return &dbRepo{
		DB: conn,
	}
}
