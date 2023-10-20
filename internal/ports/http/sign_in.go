package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignIn(c echo.Context) error {
	var req SignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res := &SignInResponse{
		Token: "test",
	}

	return c.JSON(http.StatusOK, res)
}
