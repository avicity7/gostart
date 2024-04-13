package services

import (
	"context"
	"server/config"

	"github.com/jackc/pgx/v5/pgtype"
)

// interface{} is used here to dictate that the value in the key:value of the returned map can be of ANY type.
func GetTestService() []map[string]interface{} {
	results := make([]map[string]interface{}, 0)
	rows, err := config.Dbpool.Query(context.Background(), "SELECT * FROM test_data")
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
