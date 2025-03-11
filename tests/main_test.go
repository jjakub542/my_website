package tests

import (
	"context"
	"fmt"
	"log"
	"my_website/internal/database"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbname   = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	TestDB   *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	fmt.Println(dbname, port, password, username)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)
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
