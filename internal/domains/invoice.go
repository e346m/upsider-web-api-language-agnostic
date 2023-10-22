package domains

import (
	"time"

	"github.com/shopspring/decimal"
)

var (
	commissionRate  = decimal.NewFromFloat(4)
	consumptionRate = decimal.NewFromFloat(10)
)

const (
	waiting InvoiceStatus = iota
	prcessing
	paid
	failed
)

type (
	Invoice struct {
		ID              string
		AmountBilled    decimal.Decimal
		TotalAmount     decimal.Decimal
		Commission      decimal.Decimal
		CommissionRate  decimal.Decimal
		ConsumptionTax  decimal.Decimal
		ConsumptionRate decimal.Decimal
		IssueDate       time.Time
		DueDate         time.Time
		Status          InvoiceStatus
		*Organization
		*Client
	}
	InvoiceStatus int8
	InvoiceOption struct {
		CommissionRate  *float64
		ConsumptionRate *float64
	}
)

func NewInvoice(opts ...InvoiceOption) *Invoice {
	i := &Invoice{
		CommissionRate:  commissionRate,
		ConsumptionRate: consumptionRate,
		IssueDate:       time.Now(),
	}

	for _, op := range opts {
		if op.CommissionRate != nil {
			i.CommissionRate = decimal.NewFromFloat(*op.CommissionRate)
		}
		if op.ConsumptionRate != nil {
			i.ConsumptionRate = decimal.NewFromFloat(*op.ConsumptionRate)
		}
	}

	return i
}

func (i *Invoice) SetAmountBilled(amountBilled decimal.Decimal) *Invoice {
	i.AmountBilled = amountBilled
	return i
}

func (i *Invoice) SetIntAmountBilled(amountBilled int64) *Invoice {
	i.AmountBilled = decimal.NewFromInt(amountBilled)
	return i
}

func (i *Invoice) SetDueDate(dueDate time.Time) (*Invoice, error) {
	now := time.Now()
	if dueDate.Before(now) {
		return i, &DomainError{
			Kind:    Validation,
			Message: "due date must be future date",
		}
	}

	i.DueDate = dueDate
	return i, nil
}

func (i *Invoice) Calc() {
	// check empty value
	// set calc result
}
