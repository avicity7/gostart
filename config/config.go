package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// To evaluate performance using single connection vs Pool in the future.
var Dbpool *pgx.Conn
var err error

func Connect() {
	Dbpool, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		println(err)
		println("unable to connect")
	}
}
