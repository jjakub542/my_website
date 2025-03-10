package tests

import (
	"context"
	"fmt"
	"log"
	"my_website/internal/database"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TestDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "offgrid_admin", "offgrid123", "localhost", "5432", "offgrid_db_test")
	TestDB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	database.InitTables(TestDB, "../internal/database/tables.sql")

	// Run tests
	code := m.Run()
	database.DropTables(TestDB)
	TestDB.Close()
	os.Exit(code)
}
