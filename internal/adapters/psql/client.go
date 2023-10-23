package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
