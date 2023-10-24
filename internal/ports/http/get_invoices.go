package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetInvoices(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), "token", c.Get("user"))
	pal, err := h.auther.GetPrincipal(ctx)
	if err != nil {
		return errorResponse(err, c)
	}

	var req FetchInvoiceListParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	doms, err := h.usecase.GetInvoices(
		ctx,
		req.From,
		req.To,
		pal.BelongsID,
	)
	if err != nil {
		return errorResponse(err, c)
	}

	res := make(InvoiceListResponse, len(doms))

	for idx, dom := range doms {
		res[idx] = InvoiceItem{
			AmountBilled: dom.RoundUpAmountBilled(),
			TotalAmount:  dom.RoundUpTotalAmount(),
			ClientId:     dom.Client.ID,
			DueDate:      dom.DueDate,
			IssueDate:    dom.IssueDate,
		}
	}

	return c.JSON(http.StatusOK, res)
}
