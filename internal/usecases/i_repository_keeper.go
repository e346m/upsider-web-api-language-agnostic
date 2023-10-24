package usecases

import (
	"context"
	"time"

	"github.com/e346m/upsider-wala/internal/domains"
)

type RepositoryKeeper interface {
	Reader
	Writer
	Transactioner
	IDGenerator
}

type Reader interface {
	GetMemberByEmail(ctx context.Context, email string) (domain *domains.Member, err error)
	GetClientByIDWithOrg(ctx context.Context, clientID, orgID string) (domain *domains.Client, err error)
	GetOrganizationByID(ctx context.Context, orgID string) (domain *domains.Organization, err error)
	GetInvoices(ctx context.Context, from, to time.Time, orgID string, status domains.InvoiceStatus) ([]*domains.Invoice, error)
}

type Writer interface {
	SaveOrganization(ctx context.Context, domain *domains.Organization) error
	SaveMember(ctx context.Context, domain *domains.Member) error
	SaveInvoice(ctx context.Context, domain *domains.Invoice) error
}

type Transactioner interface {
	DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}

type IDGenerator interface {
	GenID(context.Context) string
}
