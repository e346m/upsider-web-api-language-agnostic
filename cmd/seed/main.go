package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/adapters/psql"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/e346m/upsider-wala/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.LoadConfig()
	conn := initDB(cfg)
	identifier := utils.NewIdentifier()

	sql := psql.NewPSQLClient(conn, identifier)
	defer func() {
		conn.Close()
	}()

	// import adapter to lessen a burden
	// organization
	ctx := context.TODO()
	org := &domains.Organization{
		ID:             "01ARZ3NDEKTSV4RRFFQ69G5FAV",
		Name:           gofakeit.Company(),
		Representative: gofakeit.Name(),
		PhoneNumber:    gofakeit.Phone(),
		Address:        gofakeit.StreetName(),
	}
	err := sql.SaveOrganization(ctx, org)
	// Not long lived seed system
	if err != nil {
		fmt.Println("Already seeded")
		os.Exit(0)
	}
	// members
	m := &domains.Member{
		ID:           "01ARZ4NDEKTSV4RRFFQ69G5FAV",
		FullName:     gofakeit.Name(),
		Email:        "test@example.com",
		Organization: org,
	}
	domains.SetGeneratePassword("password", m)
	sql.SaveMember(ctx, m)
	// clients
	client := &domains.Client{
		ID:             "01HDDPWNWNH3BECY9074BJ2T1G",
		Name:           gofakeit.Company(),
		Representative: gofakeit.Name(),
		PhoneNumber:    gofakeit.Phone(),
		Address:        gofakeit.StreetName(),
		Organization:   org,
	}
	sql.SaveClient(ctx, client)

	// invoice
	numOfRows := int(1e6)
	var ids [1e6]string

	for i := 0; i < numOfRows; i++ {
		ids[i] = sql.GenID(ctx)
	}

	fmt.Println("generate invoice ids")

	now := time.Now()
	eg, _ := errgroup.WithContext(ctx)
	_, err = sql.DoInTx(ctx, func(ctx context.Context) (any, error) {
		for i := 0; i < numOfRows; i++ {
			i := i
			eg.Go(func() error {
				invoice := domains.NewInvoice()
				invoice.ID = ids[i]
				invoice.Client = client
				invoice.Organization = org

				var dueDate time.Time
				if i%2 == 0 {
					dueDate = gofakeit.FutureDate()
				} else {
					dueDate = gofakeit.Date()
				}
				invoice.DueDate = dueDate

				if dueDate.Before(now) {
					// paid:2 | failed: 3
					invoice.Status = domains.InvoiceStatus(rand.Intn(1) + 2)
				} else if dueDate.After(now) {
					// waiting
					invoice.Status = domains.InvoiceStatus(0)
				} else {
					// waiting:0 | ongoing: 1
					invoice.Status = domains.InvoiceStatus(rand.Intn(1))
				}
				invoice.SetIntAmountBilled(rand.Int63n(214748367))

				invoice.Calc()

				err = sql.SaveInvoice(ctx, invoice)
				if err != nil {
					return err
				}

				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			fmt.Println(err)
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func initDB(cfg *config.Config) *sql.DB {
	conn, err := sql.Open(cfg.DBType(), cfg.DBUrl())
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic(err)
	}

	return conn
}
