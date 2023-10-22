package domains

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
)

func TestNewInvoice(t *testing.T) {
	t.Run("Get Default Invoice",
		func(t *testing.T) {
			wanted := &Invoice{
				CommissionRate:  commissionRate,
				ConsumptionRate: consumptionRate,
				Status:          waiting,
				IssueDate:       time.Now(),
			}

			got := NewInvoice()
			cmpOpt := cmpopts.IgnoreFields(Invoice{}, "IssueDate")

			if diff := cmp.Diff(wanted, got, cmpOpt); diff != "" {
				t.Errorf("NewInvoice() mismatch (-wanted +got):\n%s", diff)
			}
		},
	)
	t.Run(
		"Override invoice with option",
		func(t *testing.T) {
			cmmir := 1.12
			conr := 80.0
			status := uint8(paid)
			opt := InvoiceOption{
				CommissionRate:  &cmmir,
				ConsumptionRate: &conr,
				Status:          &status,
			}
			wanted := &Invoice{
				CommissionRate:  decimal.NewFromFloat(cmmir),
				ConsumptionRate: decimal.NewFromFloat(conr),
				Status:          paid,
				IssueDate:       time.Now(),
			}

			got := NewInvoice(opt)

			cmpOpt := cmpopts.IgnoreFields(Invoice{}, "IssueDate")
			if diff := cmp.Diff(wanted, got, cmpOpt); diff != "" {
				t.Errorf("NewInvoice() mismatch (-wanted +got):\n%s", diff)
			}
		},
	)
	t.Run(
		"Override some of invoice field with option",
		func(t *testing.T) {
			cmmir := 1.12
			status := uint8(paid)
			opt := InvoiceOption{
				CommissionRate: &cmmir,
				Status:         &status,
			}
			wanted := &Invoice{
				CommissionRate:  decimal.NewFromFloat(cmmir),
				ConsumptionRate: decimal.NewFromFloat(0.10),
				Status:          paid,
				IssueDate:       time.Now(),
			}

			got := NewInvoice(opt)

			cmpOpt := cmpopts.IgnoreFields(Invoice{}, "IssueDate")
			if diff := cmp.Diff(wanted, got, cmpOpt); diff != "" {
				t.Errorf("NewInvoice() mismatch (-wanted +got):\n%s", diff)
			}
		},
	)
}

// 時間があれば、fuzzing testを追加してもいいかも
func TestCalc(t *testing.T) {
	t.Run("",
		func(t *testing.T) {
			wantedCommission := "40"
			wantedConsumptionTax := "4"
			wantedTotalAmount := "1044"
			i := NewInvoice()
			i.SetIntAmountBilled(1000)
			err := i.Calc()
			cmp.Diff(i.Commission.String(), wantedCommission)
			if diff := cmp.Diff(i.Commission.String(), wantedCommission); diff != "" {
				t.Errorf("Commision mismatch (-wanted +got):\n%s", diff)
			}

			if diff := cmp.Diff(i.ConsumptionTax.String(), wantedConsumptionTax); diff != "" {
				t.Errorf("Commision mismatch (-wanted +got):\n%s", diff)
			}

			if diff := cmp.Diff(i.TotalAmount.String(), wantedTotalAmount); diff != "" {
				t.Errorf("Commision mismatch (-wanted +got):\n%s", diff)
			}

			if err != nil {
				t.Fatalf("test should pass")
			}
		},
	)
	t.Run(
		"When amount billed is missing",
		func(t *testing.T) {
			t.Parallel()
			i := &Invoice{}
			err := i.Calc()
			if err == nil {
				t.Fatalf("it must be return err")
			}
		},
	)
	t.Run(
		"When commission rates missing",
		func(t *testing.T) {
			t.Parallel()
			i := &Invoice{}
			i.SetIntAmountBilled(1000)
			err := i.Calc()
			if err == nil {
				t.Fatalf("it must be return err")
			}
		},
	)
	t.Run(
		"When consumption rates missing",
		func(t *testing.T) {
			t.Parallel()
			i := &Invoice{
				CommissionRate: decimal.NewFromFloat(0.04),
			}
			i.SetIntAmountBilled(1000)
			err := i.Calc()
			if err == nil {
				t.Fatalf("it must be return err")
			}
		},
	)
}

func TestSetDueDate(t *testing.T) {
	t.Run("When due date is past",
		func(t *testing.T) {
			i := NewInvoice()
			err := i.SetDueDate(time.Now().AddDate(0, 0, -1))

			if err == nil {
				t.Fatalf("test should fail")
			}
		},
	)
	t.Run(
		"When due date is future",
		func(t *testing.T) {
			i := NewInvoice()
			err := i.SetDueDate(time.Now().AddDate(0, 0, 1))
			if err != nil {
				t.Fatalf("test should pass")
			}
		},
	)
}

func TestSetIntAmountBilled(t *testing.T) {
	t.Run("When SetIntAmountBilled() is called",
		func(t *testing.T) {
			var wanted int64 = 1000
			i := NewInvoice().SetIntAmountBilled(wanted)
			if !i.AmountBilled.Equal(decimal.NewFromInt(wanted)) {
				t.Fatalf("passed value must be set")
			}
		},
	)
}

func TestSetAmountBilled(t *testing.T) {
	t.Run("When SetAmountBilled() is called",
		func(t *testing.T) {
			wanted := decimal.NewFromInt(1000)
			i := NewInvoice().SetAmountBilled(wanted)
			if !i.AmountBilled.Equal(wanted) {
				t.Fatalf("passed value must be set")
			}
		},
	)
}
