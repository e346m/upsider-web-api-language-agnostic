package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (p *PSQL) GetMemberByEmail(ctx context.Context, email string) (*domains.Member, error) {
	return &domains.Member{}, nil
}

func (p *PSQL) SaveMember(ctx context.Context, dom *domains.Member) error {
	ex := p.getExecutor(ctx)

	id, err := p.id.StringToBinary(dom.ID)
	if err != nil {
		return err
	}

	orgId, err := p.id.StringToBinary(dom.Organization.ID)
	if err != nil {
		return err
	}

	db := dbmodel.Member{
		ID:             id,
		FullName:       dom.FullName,
		Email:          dom.Email,
		Password:       dom.Password,
		OrganizationID: orgId,
	}

	err = db.Insert(ctx, ex, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
