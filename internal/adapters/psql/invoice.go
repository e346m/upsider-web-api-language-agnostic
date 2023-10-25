package psql

import (
	"context"
	"errors"
	"time"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	ericDecimal "github.com/ericlagergren/decimal"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"golang.org/x/sync/errgroup"
)

func (p *PSQL) GetInvoices(ctx context.Context, from, to *time.Time, orgID string, status domains.InvoiceStatus) ([]*domains.Invoice, error) {
	ex := p.getExecutor(ctx)

	orgid, err := p.id.StringToBinary(orgID)
	if err != nil {
		return nil, err
	}

	where := dbmodel.InvoiceWhere
	condition := []qm.QueryMod{
		where.Status.EQ(int16(status)),
		where.OrganizationID.EQ(orgid),
		qm.Load(dbmodel.InvoiceRels.Client),
	}
	if from != nil && !from.IsZero() {
		condition = append(condition, where.DueDate.GTE(*from))
	}

	if to != nil && !to.IsZero() {
		condition = append(condition, where.DueDate.LTE(*to))
	}

	rows, err := dbmodel.Invoices(
		condition...,
	).All(ctx, ex)
	if err != nil {
		return nil, err
	}

	doms := make([]*domains.Invoice, len(rows))

	eg, _ := errgroup.WithContext(ctx)
	for idx, row := range rows {
		idx := idx
		row := row
		eg.Go(func() error {
			invoice, err := p.mapInvoice(row)
			if err != nil {
				return err
			}
			if row.R.Client == nil {
				return errors.New("internal error")
			}
			client, err := p.mapClient(row.R.Client)
			if err != nil {
				return err
			}

			invoice.Client = client

			doms[idx] = invoice
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return doms, nil
}

func (p *PSQL) mapInvoice(row *dbmodel.Invoice) (*domains.Invoice, error) {
	id, err := p.id.BinaryToString(row.ID)
	if err != nil {
		return nil, err
	}

	amountBilled, err := revertDecimal(row.AmountBilled)
	if err != nil {
		return nil, err
	}

	totalAmount, err := revertDecimal(row.TotalAmount)
	if err != nil {
		return nil, err
	}

	commission, err := revertDecimal(row.Commission)
	if err != nil {
		return nil, err
	}

	commissionRate, err := revertDecimal(row.CommissionRate)
	if err != nil {
		return nil, err
	}

	consumptionTax, err := revertNullDecimal(row.ConsumptionTax)
	if err != nil {
		return nil, err
	}

	consumptionRate, err := revertNullDecimal(row.ConsumptionTaxRate)
	if err != nil {
		return nil, err
	}

	return &domains.Invoice{
		ID:              id,
		AmountBilled:    amountBilled,
		TotalAmount:     totalAmount,
		Commission:      commission,
		CommissionRate:  commissionRate,
		ConsumptionTax:  consumptionTax,
		ConsumptionRate: consumptionRate,
		IssueDate:       row.IssueDate,
		DueDate:         row.DueDate,
		Status:          domains.InvoiceStatus(row.Status),
	}, nil
}

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
		TotalAmount:        convertDecimal(dom.TotalAmount),
		AmountBilled:       convertDecimal(dom.AmountBilled),
		CommissionRate:     convertDecimal(dom.CommissionRate),
		ConsumptionTax:     convertNullDecimal(dom.ConsumptionTax),
		ConsumptionTaxRate: convertNullDecimal(dom.ConsumptionRate),
		IssueDate:          dom.IssueDate,
		DueDate:            dom.DueDate, Status: int16(dom.Status),
	}

	err = db.Insert(ctx, ex, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func revertDecimal(d types.Decimal) (decimal.Decimal, error) {
	intPart, ok := d.Int64()
	if !ok {
		return decimal.Zero, errors.New("cannot get int")
	}
	exp := -d.Scale()
	return decimal.New(intPart, int32(exp)), nil
}

func revertNullDecimal(d types.NullDecimal) (decimal.Decimal, error) {
	intPart, ok := d.Int64()
	if !ok {
		return decimal.Zero, errors.New("cannot get int")
	}
	exp := -d.Scale()
	return decimal.New(intPart, int32(exp)), nil
}

func convertDecimal(d decimal.Decimal) types.Decimal {
	// ericlagergrenとshopspringのdecimalではscaleの指定が逆転しているため
	//  shopDecimal.New(1234, -3) // 1.234
	//  ericDecimal.New(1234, 3) // 1.234
	big := ericDecimal.New(d.CoefficientInt64(), -1*int(d.Exponent()))
	return types.NewDecimal(big)
}

func convertNullDecimal(d decimal.Decimal) types.NullDecimal {
	// ericlagergrenとshopspringのdecimalではscaleの指定が逆転しているため
	//  shopDecimal.New(1234, -3) // 1.234
	//  ericDecimal.New(1234, 3) // 1.234
	big := ericDecimal.New(d.CoefficientInt64(), -1*int(d.Exponent()))
	return types.NewNullDecimal(big)
}
