package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"my_website/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HomePage)
	e.GET("/contact", s.ContactPage)
	e.GET("/projects", s.ProjectsPage)
	e.GET("/blog", s.BlogPage)
	e.GET("/blog/:article_id", s.ArticleView)
	e.GET("/hello", s.HelloHandler)
	e.GET("/health", s.HealthHandler)

	e.GET("/admin", s.AdminHomePage)
	e.Any("/admin/login", s.AdminLoginPage)
	e.Any("/admin/logout", s.AdminLogoutPage)
	e.GET("/admin/edit-article", s.AdminArticleEditPage)

	e.DELETE("/articles/delete", s.ArticleDeleteHandler)
	e.PUT("/articles/update", s.ArticleUpdateHandler)

	e.GET("/img", s.GetImageHandler)

	return e
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
