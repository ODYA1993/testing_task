package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("database_url")
	if databaseUrl == "" {
		databaseUrl = "host=localhost port=5432 dbname=restapi_test user=postgres password=postgres sslmode=disable"
	}
	os.Exit(m.Run())
}
