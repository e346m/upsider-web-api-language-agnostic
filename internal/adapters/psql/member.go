package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (p *PSQL) GetMemberByEmail(ctx context.Context, email string) (*domains.Member, error) {
	ex := p.getExecutor(ctx)
	row, err := dbmodel.Members(
		qm.Where("email = ?", email),
		qm.Load(dbmodel.MemberRels.Organization),
	).One(ctx, ex)
	if err != nil {
		return &domains.Member{}, nil
	}

	id, err := p.id.BinaryToString(row.ID)
	if err != nil {
		return &domains.Member{}, nil
	}

	dom := &domains.Member{
		ID:       id,
		FullName: row.FullName,
		Email:    row.Email,
		Password: row.Password,
	}

	orgRow := row.R.Organization
	if orgRow == nil {
		return &domains.Member{}, nil
	}

	orgId, err := p.id.BinaryToString(row.R.Organization.ID)
	if err != nil {
		return &domains.Member{}, nil
	}

	org := &domains.Organization{
		ID:              orgId,
		Name:            orgRow.Name,
		Rrepresentative: orgRow.Representative,
		PhoneNumber:     orgRow.PhoneNumber,
		Address:         orgRow.Address,
	}

	dom.Organization = org

	return dom, nil
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
