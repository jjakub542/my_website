package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

/*
func (s *Server) BlogPage(c echo.Context) error {
	repo := articles.Repository(s.db)
	articles, err := repo.GetAllPublic()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "blog.html", articles)
}*/

func (h *Handler) ContactPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func (h *Handler) ProjectsPage(c echo.Context) error {
	return c.Render(http.StatusOK, "projects.html", nil)
}

/*
func (s *Server) ArticleView(c echo.Context) error {
	repo := articles.Repository(s.db)
	id := c.Param("article_id")
	fmt.Println(id)
	article, err := repo.GetOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "article.html", article)
}
*/
