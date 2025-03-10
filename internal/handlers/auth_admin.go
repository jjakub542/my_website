package handlers

import (
	"fmt"
	"my_website/internal/domain"
	"my_website/internal/session"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) LoginPage(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	if c.Request().Method == http.MethodGet {
		role, ok := store.Get(sessionID, "role")
		if ok || role == "admin" {
			return c.Redirect(http.StatusSeeOther, "/admin/articles")
		}
		return c.Render(http.StatusOK, "admin/login.html", nil)
	}

	requestUser := domain.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	requestUser.CreatePasswordHash()

	user, err := h.Repository.User.GetOneByEmail(c.FormValue("email"))

	fmt.Println(user.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User with this email does not exist")
	}

	if user.PasswordHash != requestUser.PasswordHash {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
	}

	if !user.IsSuperuser {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store.Set(sessionID, "authenticated", true)
	store.Set(sessionID, "role", "admin")

	return c.Redirect(http.StatusSeeOther, "/admin/articles")
}

func (h *Handler) LogoutPage(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	store.Delete(sessionID)
	return c.Render(http.StatusOK, "admin/logout.html", nil)
}
