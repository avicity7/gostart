package services

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// interface{} is used here to dictate that the value in the key:value of the returned map can be of ANY type.
func GetTestService() []map[string]interface{} {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		println(err)
		println("unable to connect")
	}
	results := make([]map[string]interface{}, 0)
	rows, err := dbpool.Query(context.Background(), "SELECT * FROM test_data")
	if err != nil {
		println(err)
		println("unable to query")
	} else {
		for rows.Next() {
			var id string
			var data float32
			var timestamp pgtype.Timestamp
			rows.Scan(&id, &data, &timestamp)
			result := map[string]interface{}{"id": id, "data": data, "timestamp": timestamp}
			results = append(results, result)
		}
	}

	return results
}
