package server

import (
	"my_website/internal/handlers"
	"my_website/internal/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) Router() http.Handler {
	e := echo.New()
	e.Use(session.Middleware(s.store))
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handlers.Handler{Repository: s.repository}

	e.GET("/", h.HomePage)
	e.GET("/contact", h.ContactPage)
	e.GET("/projects", h.ProjectsPage)
	e.GET("/blog", h.BlogPage)
	e.GET("/blog/:article_id", h.ArticleView)

	adminGroup := e.Group("/admin")
	adminGroup.GET("/articles", session.AdminAuth(h.AdminHomePage))
	adminGroup.POST("/articles/create", session.AdminAuth(h.ArticleCreate))
	adminGroup.POST("/articles/:article_id/delete", session.AdminAuth(h.ArticleDelete))
	adminGroup.POST("/articles/:article_id/update", session.AdminAuth(h.ArticleUpdate))
	adminGroup.POST("/articles/:article_id/attach-image", session.AdminAuth(h.ArticleAttachImage))
	adminGroup.POST("/articles/delete-image", session.AdminAuth(h.ArticleDeleteImage))
	adminGroup.GET("/articles/:article_id/edit", session.AdminAuth(h.ArticleEditPage))
	adminGroup.Any("/login", h.LoginPage)
	adminGroup.Any("/logout", h.LogoutPage)

	return e
}
