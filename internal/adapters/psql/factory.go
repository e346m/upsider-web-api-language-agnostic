package psql

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var txKey = struct{}{}

type PSQL struct {
	db *sql.DB
	id Identifier
}

type Identifier interface {
	StringToBinary(string) ([]byte, error)
	BinaryToString([]byte) (string, error)
}

func NewPSQLClient(db *sql.DB, id Identifier) *PSQL {
	return &PSQL{
		db: db,
		id: id,
	}
}

func (p *PSQL) getExecutor(ctx context.Context) boil.ContextExecutor {
	tx, ok := ctx.Value(&txKey).(*sql.Tx)
	if ok {
		return tx
	}
	return p.db
}

func (p *PSQL) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, &txKey, tx)

	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}
