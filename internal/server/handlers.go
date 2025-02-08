package server

import (
	"fmt"
	"my_website/internal/articles"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (s *Server) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (s *Server) BlogPage(c echo.Context) error {
	repo := articles.Repository(s.db)
	articles, err := repo.GetAllPublic()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "blog.html", articles)
}

func (s *Server) ContactPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func (s *Server) ProjectsPage(c echo.Context) error {
	return c.Render(http.StatusOK, "projects.html", nil)
}

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

func (s *Server) AdminLoginPage(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "admin/login.html", nil)
	}
	if c.FormValue("password") != os.Getenv("ADMIN_PASSWORD") {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	if c.FormValue("username") != os.Getenv("ADMIN_USERNAME") {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	s.CreateSession(c, c.FormValue("username"))
	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (s *Server) AdminLogoutPage(c echo.Context) error {
	s.DestroySession(c)
	return c.Render(http.StatusOK, "admin/logout.html", nil)
}

func (s *Server) AdminHomePage(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.Redirect(http.StatusSeeOther, "/admin/login")
	}
	repo := articles.Repository(s.db)
	articles, err := repo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "admin/home.html", articles)
}

func (s *Server) AdminArticleEditPage(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.Redirect(http.StatusSeeOther, "/admin/login")
	}
	id := c.QueryParam("id")
	repo := articles.Repository(s.db)
	article, err := repo.GetOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "admin/edit_article.html", article)
}

func (s *Server) ArticleDeleteHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	id := c.QueryParam("id")
	repo := articles.Repository(s.db)
	if err := repo.DeleteOne(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted",
	})
}

func (s *Server) ArticleUpdateHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	id := c.QueryParam("id")
	var article articles.Article
	if err = c.Bind(&article); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	repo := articles.Repository(s.db)
	if err := repo.UpdateOne(id, &article); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "updated",
	})
}

func (s *Server) GetImageHandler(c echo.Context) error {
	return c.File("media/me1final.jpg")
}

func (s *Server) HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World",
	})
}

func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
