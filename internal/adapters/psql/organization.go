package psql

import (
	"context"

	dbmodel "github.com/e346m/upsider-wala/db/schema"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
