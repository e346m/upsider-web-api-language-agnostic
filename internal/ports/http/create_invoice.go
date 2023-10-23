package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateInvoice(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), "token", c.Get("user"))
	pal, err := h.auther.GetPrincipal(ctx)
	if err != nil {
		return errorResponse(err, c)
	}

	var req CreateInvoiceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	dom, err := h.usecase.CreateInvoice(
		ctx,
		req.AmountBilled,
		req.DueDate,
		req.ClientId,
		pal.BelongsID,
	)
	if err != nil {
		return errorResponse(err, c)
	}

	res := &CreateInvoiceResponse{
		AmountBilled: dom.RoundUpAmountBilled(),
		TotalAmount:  dom.RoundUpTotalAmount(),
		ClientId:     dom.Client.ID,
		DueDate:      dom.DueDate,
		IssueDate:    dom.IssueDate,
		Status:       dom.Status.String(),
	}

	return c.JSON(http.StatusOK, res)
}
