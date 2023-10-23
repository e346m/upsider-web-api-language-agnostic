package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (p *PSQL) SaveClient(ctx context.Context, dom *domains.Client) error {
	ex := p.getExecutor(ctx)

	id, err := p.id.StringToBinary(dom.ID)
	if err != nil {
		return err
	}

	orgID, err := p.id.StringToBinary(dom.Organization.ID)
	if err != nil {
		return err
	}

	db := dbmodel.Client{
		ID:             id,
		Name:           dom.Name,
		Representative: dom.Name,
		PhoneNumber:    dom.PhoneNumber,
		Address:        dom.Address,
		OrganizationID: orgID,
	}

	err = db.Insert(ctx, ex, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (p *PSQL) GetClientByIDWithOrg(ctx context.Context, clientID, orgID string) (*domains.Client, error) {
	ex := p.getExecutor(ctx)

	id, err := p.id.StringToBinary(clientID)
	if err != nil {
		return nil, err
	}

	orgid, err := p.id.StringToBinary(orgID)
	if err != nil {
		return nil, err
	}

	row, err := dbmodel.Clients(
		qm.Where("id=?", id),
		qm.And("organization_id=?", orgid),
	).One(ctx, ex)
	if err != nil {
		return nil, err
	}

	dom := &domains.Client{
		ID:             clientID,
		Name:           row.Name,
		Representative: row.Representative,
		PhoneNumber:    row.PhoneNumber,
		Address:        row.Address,
		Organization:   &domains.Organization{ID: orgID},
	}

	return dom, nil
}
