package psql

import (
	"context"

	"github.com/e346m/upsider-wala/internal/domains"
)

func (p *PSQL) GetMemberByEmail(ctx context.Context, email string) (*domains.Member, error) {
	return &domains.Member{}, nil
}

func (p *PSQL) SaveMember(ctx context.Context, dom *domains.Member) error {
	return nil
}
