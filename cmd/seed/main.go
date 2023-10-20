package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/adapters/psql"
	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/e346m/upsider-wala/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := config.LoadConfig()
	conn := initDB(cfg)
	identifier := utils.NewIdentifier()

	client := psql.NewPSQLClient(conn, identifier)

	defer func() {
		conn.Close()
	}()

	// import adapter to lessen a burden
	// organization
	ctx := context.TODO()
	org := &domains.Organization{
		ID:              "01ARZ3NDEKTSV4RRFFQ69G5FAV",
		Name:            gofakeit.Company(),
		Rrepresentative: gofakeit.Name(),
		PhoneNumber:     gofakeit.Phone(),
		Address:         gofakeit.StreetName(),
	}
	err := client.SaveOrganization(ctx, org)
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
	m.SetGeneratePassword("password")
	client.SaveMember(ctx, m)
	// clients
	// invoice
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
