package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbname             = os.Getenv("DB_NAME")
	dbnametest         = os.Getenv("DB_NAME_TEST")
	password           = os.Getenv("DB_PASSWORD")
	username           = os.Getenv("DB_USERNAME")
	port               = os.Getenv("DB_PORT")
	host               = os.Getenv("DB_HOST")
	postgresClient     *pgxpool.Pool
	postgresClientTest *pgxpool.Pool
)

func Connect() *pgxpool.Pool {
	fmt.Println(username)
	if postgresClient != nil {
		return postgresClient
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	postgresClient = db
	return db
}

func ConnectTest() *pgxpool.Pool {
	if postgresClientTest != nil {
		return postgresClientTest
	}
	fmt.Println(username)
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"my_website_admin",
		"123",
		"localhost",
		"5432",
		"my_website_db_test",
	)
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	postgresClientTest = db
	return db
}

func InitTables(db *pgxpool.Pool, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sql := string(content)
	_, err = db.Exec(context.Background(), sql)
	return err
}

func DropTables(db *pgxpool.Pool) error {
	var err error

	_, err = db.Exec(context.Background(), `
	DROP TABLE articles;
	`)

	_, err = db.Exec(context.Background(), `
	DROP TABLE users;
	`)

	_, err = db.Exec(context.Background(), `
	DROP TABLE images;
	`)

	return err
}
