package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"my_website/internal/database"
	"my_website/internal/repository"
	"my_website/internal/session"
)

type Server struct {
	port       int
	db         *pgxpool.Pool
	store      *session.Store
	repository *repository.Repository
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.Connect()
	store := session.NewStore()
	NewServer := &Server{
		port:       port,
		db:         db,
		store:      store,
		repository: repository.New(db),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
