package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/domains"
)

type testIdentifier struct{}

func (ti *testIdentifier) NewIdString() string {
	return "id"
}

func (ti *testIdentifier) StringToBinary(string) ([]byte, error) {
	return []byte("id"), nil
}

func (ti *testIdentifier) BinaryToString([]byte) (string, error) {
	return "id", nil
}

func initDB() *sql.DB {
	cfg := config.LoadConfig("./../../../config")
	conn, err := sql.Open(cfg.DBType(), cfg.DBUrl())
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic(err)
	}

	return conn
}

func (p *PSQL) prepareRelations(ctx context.Context) (*domains.Organization, *domains.Client) {
	org := &domains.Organization{
		ID:             "id",
		Name:           gofakeit.Company(),
		Representative: gofakeit.Name(),
		PhoneNumber:    gofakeit.Phone(),
		Address:        gofakeit.StreetName(),
	}
	p.SaveOrganization(ctx, org)

	client := &domains.Client{
		ID:             "id",
		Name:           gofakeit.Company(),
		Representative: gofakeit.Name(),
		PhoneNumber:    gofakeit.Phone(),
		Address:        gofakeit.StreetName(),
		Organization:   org,
	}
	p.SaveClient(ctx, client)

	return org, client
}

func TestSaveInvoice(t *testing.T) {
	t.Run("When saveInvoice() is called",
		func(t *testing.T) {
			wanted := domains.NewInvoice()
			wanted.ID = "id"
			wanted.SetDueDate(time.Now().AddDate(0, 0, 1))
			wanted.SetIntAmountBilled(1000)
			wanted.Calc()
			fmt.Println(wanted.Commission)
			sql := NewPSQLClient(initDB(), &testIdentifier{})
			sql.DoInTx(context.Background(), func(ctx context.Context) (any, error) {
				org, client := sql.prepareRelations(ctx)
				wanted.Organization = org
				wanted.Client = client
				row, err := sql.saveInvoice(ctx, wanted)
				if err != nil {
					t.Fatalf("passed value must be set")
				}

				row.Reload(ctx, sql.db)
				got, err := sql.mapInvoice(row)
				if err != nil {
					t.Fatalf("row must be mapped to domain")
				}

				cmpOpt := cmpopts.IgnoreFields(domains.Invoice{}, "Organization", "Client")

				if diff := cmp.Diff(wanted, got, cmpOpt); diff != "" {
					t.Errorf("NewInvoice() mismatch (-wanted +got):\n%s", diff)
				}

				// for rollback
				return nil, errors.New("for rollback")
			})
		},
	)
}
