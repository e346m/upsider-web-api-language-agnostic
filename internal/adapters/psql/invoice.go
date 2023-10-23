package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	ericDecimal "github.com/ericlagergren/decimal"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func (p *PSQL) SaveInvoice(ctx context.Context, dom *domains.Invoice) error {
	ex := p.getExecutor(ctx)

	id, err := p.id.StringToBinary(dom.ID)
	if err != nil {
		return err
	}

	orgId, err := p.id.StringToBinary(dom.Organization.ID)
	if err != nil {
		return err
	}

	clientId, err := p.id.StringToBinary(dom.Client.ID)
	if err != nil {
		return err
	}

	db := dbmodel.Invoice{
		ID:                 id,
		ClientID:           clientId,
		OrganizationID:     orgId,
		TotalAmount:        dom.TotalAmount.String(),
		AmountBilled:       dom.AmountBilled.String(),
		Commission:         dom.Commission.String(),
		CommissionRate:     converDecimals(dom.CommissionRate),
		ConsumptionTax:     null.StringFrom(dom.ConsumptionTax.String()),
		ConsumptionTaxRate: converNullDecimals(dom.ConsumptionRate),
		IssueDate:          dom.IssueDate,
		DueDate:            dom.DueDate,
	}

	err = db.Insert(ctx, ex, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func converDecimals(d decimal.Decimal) types.Decimal {
	// ericlagergrenとshopspringのdecimalではscaleの指定が逆転しているため
	//  shopDecimal.New(1234, -3) // 1.234
	//  ericDecimal.New(1234, 3) // 1.234
	big := ericDecimal.New(d.CoefficientInt64(), -1*int(d.Exponent()))
	return types.NewDecimal(big)
}

func converNullDecimals(d decimal.Decimal) types.NullDecimal {
	// ericlagergrenとshopspringのdecimalではscaleの指定が逆転しているため
	//  shopDecimal.New(1234, -3) // 1.234
	//  ericDecimal.New(1234, 3) // 1.234
	big := ericDecimal.New(d.CoefficientInt64(), -1*int(d.Exponent()))
	return types.NewNullDecimal(big)
}
