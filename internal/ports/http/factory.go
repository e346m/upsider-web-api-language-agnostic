package http

import (
	"net/http"
	"time"

	"github.com/e346m/upsider-wala/internal/domains"
	usecases "github.com/e346m/upsider-wala/internal/usecases"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase usecases.Usecase
}

func NewHandler(usecase *usecases.Usecase) *Handler {
	return &Handler{
		usecase: *usecase,
	}
}

func generateCookieWithRefreshToken(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Expires = time.Now().AddDate(0, 1, 0)
	cookie.Domain = "test.com"
	cookie.Secure = true
	return cookie
}

func errorResponse(err error, c echo.Context) error {
	e, ok := err.(*domains.DomainError)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "ask admin")
	}

	if e.IsValidation() {
		return c.JSON(http.StatusBadRequest, e.Message)
	} else if e.IsNotFound() {
		return c.JSON(http.StatusNotFound, e.Message)
	} else if e.IsPolicy() {
		return c.JSON(http.StatusForbidden, e.Message)
	} else {
		return c.JSON(http.StatusInternalServerError, e.Message)
	}
}
