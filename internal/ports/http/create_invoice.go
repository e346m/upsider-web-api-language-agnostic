package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateInvoice(c echo.Context) error {
	var req CreateInvoiceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	dom, err := h.usecase.CreateInvoice(
		c.Request().Context(),
		req.AmountBilled,
		req.DueDate,
		"01HDDPWNWNH3BECY9074BJ2T1G", // req.ClientId,
		"01ARZ3NDEKTSV4RRFFQ69G5FAV",
	)
	if err != nil {
		return errorResponse(err, c)
	}

	res := &CreateInvoiceResponse{
		AmountBilled: dom.AmountBilled.CoefficientInt64(),
		TotalAmount:  dom.TotalAmount.CoefficientInt64(),
		ClientId:     dom.Client.ID,
		DueDate:      dom.DueDate,
		IssueDate:    dom.IssueDate,
		Status:       dom.Status.String(),
	}

	return c.JSON(http.StatusOK, res)
}
