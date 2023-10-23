package usecases

import (
	"context"
	"time"

	"github.com/e346m/upsider-wala/internal/domains"
)

func (u *Usecase) CreateInvoice(ctx context.Context, amountBilled int64, dueDate time.Time, clientID, orgID string) (*domains.Invoice, error) {
	ctx, span := u.tracer.Start(ctx, "CreateInvoice")
	defer span.End()

	// Is it better to check policy before going through this usecase? It would be more simple to find client
	client, err := u.repo.GetClientByIDWithOrg(ctx, clientID, orgID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	org, err := u.repo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	dom := domains.NewInvoice()
	dom.ID = u.repo.GenID(ctx)
	dom.Client = client
	dom.Organization = org
	dom.SetIntAmountBilled(amountBilled)
	err = dom.SetDueDate(dueDate)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	err = dom.Calc()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	err = u.repo.SaveInvoice(ctx, dom)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return dom, nil
}
