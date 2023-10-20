package usecases

import (
	"context"

	"github.com/e346m/upsider-wala/internal/domains"
)

type RepositoryKeeper interface {
	Reader
	Writer
	Transactioner
}

type Reader interface {
	GetMemberByEmail(ctx context.Context, email string) (domain *domains.Member, err error)
}

type Writer interface {
	SaveOrganization(ctx context.Context, domain *domains.Organization) error
	SaveMember(ctx context.Context, domain *domains.Member) error
}

type Transactioner interface {
	DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
