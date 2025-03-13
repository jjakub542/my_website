package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (h *Handler) BlogPage(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit := 8
	offset := (page - 1) * limit
	articles, err := h.Repository.Article.GetAllPublicBetween(limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	count, err := h.Repository.Article.GetCount()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "blog.html", map[string]interface{}{
		"articles": articles,
		"count":    count,
	})
}

func (h *Handler) ContactPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func (h *Handler) ProjectsPage(c echo.Context) error {
	return c.Render(http.StatusOK, "projects.html", nil)
}

func (h *Handler) ArticleView(c echo.Context) error {
	id := c.Param("article_id")
	article, err := h.Repository.Article.GetOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "article.html", article)
}
