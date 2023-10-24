package usecases

import (
	"context"
	"time"

	"github.com/e346m/upsider-wala/internal/domains"
)

func (u *Usecase) GetInvoices(ctx context.Context, to, from time.Time, orgID string) ([]*domains.Invoice, error) {
	ctx, span := u.tracer.Start(ctx, "GetInvoices")
	defer span.End()

	doms, err := u.repo.GetInvoices(ctx, to, from, orgID, domains.Waiting)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return doms, nil
}
