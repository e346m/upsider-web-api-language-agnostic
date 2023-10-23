// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package http

import (
	"time"
)

// CreateInvoiceResponse defines model for CreateInvoiceResponse.
type CreateInvoiceResponse struct {
	AmountBilled int64     `json:"amount_billed"`
	ClientId     string    `json:"client_id"`
	DueDate      time.Time `json:"due_date"`
	IssueDate    time.Time `json:"issue_date"`
	Status       string    `json:"status"`
	TotalAmount  int64     `json:"total_amount"`
}

// SignInResponse defines model for SignInResponse.
type SignInResponse struct {
	Token string `json:"token"`
}

// CreateInvoiceRequest defines model for CreateInvoiceRequest.
type CreateInvoiceRequest struct {
	AmountBilled int64     `json:"amount_billed" validate:"required"`
	ClientId     string    `json:"client_id" validate:"required"`
	DueDate      time.Time `json:"due_date" validate:"required"`
}

// SignInRequest defines model for SignInRequest.
type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CreateInvoiceJSONBody defines parameters for CreateInvoice.
type CreateInvoiceJSONBody struct {
	AmountBilled int64     `json:"amount_billed" validate:"required"`
	ClientId     string    `json:"client_id" validate:"required"`
	DueDate      time.Time `json:"due_date" validate:"required"`
}

// SignInJSONBody defines parameters for SignIn.
type SignInJSONBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CreateInvoiceJSONRequestBody defines body for CreateInvoice for application/json ContentType.
type CreateInvoiceJSONRequestBody CreateInvoiceJSONBody

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody SignInJSONBody
