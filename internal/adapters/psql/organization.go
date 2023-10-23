package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (p *PSQL) SaveOrganization(ctx context.Context, dom *domains.Organization) error {
	ex := p.getExecutor(ctx)

	id, err := p.id.StringToBinary(dom.ID)
	if err != nil {
		return err
	}

	db := dbmodel.Organization{
		ID:             id,
		Name:           dom.Name,
		Representative: dom.Name,
		PhoneNumber:    dom.PhoneNumber,
		Address:        dom.Address,
	}

	err = db.Insert(ctx, ex, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (p *PSQL) GetOrganizationByID(ctx context.Context, orgID string) (*domains.Organization, error) {
	ex := p.getExecutor(ctx)

	orgid, err := p.id.StringToBinary(orgID)
	if err != nil {
		return nil, err
	}

	row, err := dbmodel.Organizations(
		qm.Where("id=?", orgid),
	).One(ctx, ex)
	if err != nil {
		return nil, err
	}

	dom := &domains.Organization{
		ID:             orgID,
		Name:           row.Name,
		Representative: row.Representative,
		PhoneNumber:    row.PhoneNumber,
		Address:        row.Address,
	}

	return dom, nil
}
