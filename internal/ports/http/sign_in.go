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

	accessToken, refreshToken, err := h.usecase.SignIn(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return errorResponse(err, c)
	}

	cookie := generateCookieWithRefreshToken(refreshToken)
	c.SetCookie(cookie)

	res := &SignInResponse{
		Token: accessToken,
	}

	return c.JSON(http.StatusOK, res)
}
