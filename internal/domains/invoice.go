package domains

import (
	"time"

	"github.com/shopspring/decimal"
)

var (
	commissionRate  = decimal.NewFromFloat(0.04)
	consumptionRate = decimal.NewFromFloat(0.10)
)

const (
	Waiting InvoiceStatus = iota
	Ongoing
	Paid
	Failed
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
	InvoiceStatus uint8
	InvoiceOption struct {
		CommissionRate  *float64
		ConsumptionRate *float64
		Status          *uint8
	}
)

func NewInvoice(opts ...InvoiceOption) *Invoice {
	i := &Invoice{
		CommissionRate:  commissionRate,
		ConsumptionRate: consumptionRate,
		Status:          Waiting,
		IssueDate:       time.Now(),
	}

	for _, op := range opts {
		if op.CommissionRate != nil {
			i.CommissionRate = decimal.NewFromFloat(*op.CommissionRate)
		}
		if op.ConsumptionRate != nil {
			i.ConsumptionRate = decimal.NewFromFloat(*op.ConsumptionRate)
		}

		if op.Status != nil {
			// should be handle error case but for now trust value saved in database
			i.Status = InvoiceStatus(*op.Status)
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

func (i *Invoice) SetDueDate(dueDate time.Time) error {
	now := time.Now()
	if dueDate.Before(now) {
		return &DomainError{
			Kind:    Validation,
			Message: "due date must be future date",
		}
	}

	i.DueDate = dueDate
	return nil
}

func (i *Invoice) Calc() error {
	if i.AmountBilled.IsZero() || i.AmountBilled.IsNegative() {
		return &DomainError{
			Kind:    Validation,
			Message: "amount billed must be positive",
		}
	}

	if i.CommissionRate.IsZero() || i.CommissionRate.IsNegative() {
		return &DomainError{
			Kind:    Validation,
			Message: "commission rate must be positive",
		}
	}

	if i.ConsumptionRate.IsZero() || i.ConsumptionRate.IsNegative() {
		return &DomainError{
			Kind:    Validation,
			Message: "consumption rate must be positive",
		}
	}

	i.Commission = i.AmountBilled.Mul(i.CommissionRate)
	i.ConsumptionTax = i.Commission.Mul(i.ConsumptionRate)
	i.TotalAmount = i.AmountBilled.Add(i.Commission).Add(i.ConsumptionTax)

	return nil
}

func (i *Invoice) RoundUpAmountBilled() int64 {
	return i.AmountBilled.RoundUp(0).IntPart()
}

func (i *Invoice) RoundUpTotalAmount() int64 {
	return i.TotalAmount.RoundUp(0).IntPart()
}

func (is InvoiceStatus) String() string {
	switch is {
	case Waiting:
		return "未処理"
	case Ongoing:
		return "処理中"
	case Paid:
		return "支払い済み"
	case Failed:
		return "エラー"
	default:
		return ""
	}
}
