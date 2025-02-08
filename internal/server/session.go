package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const name = "session_token"

var sessions = map[string]string{}

func sessionCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(72 * time.Hour),
	}
}

func destroyCookie() *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	}
}

func (s *Server) CreateSession(c echo.Context, userid string) {
	token := uuid.NewString()
	c.SetCookie(sessionCookie(token))
	sessions[token] = userid
}

func (s *Server) DestroySession(c echo.Context) error {
	cookie, err := c.Cookie(name)
	if err != nil {
		return err
	}
	c.SetCookie(destroyCookie())
	delete(sessions, cookie.Value)
	return nil
}

func (s *Server) GetSession(c echo.Context) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	userid := sessions[cookie.Value]
	if userid == "" {
		c.SetCookie(destroyCookie())
		return "", err
	}
	return userid, nil
}
