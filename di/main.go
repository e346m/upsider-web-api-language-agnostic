package di

import (
	"database/sql"

	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/adapters/psql"
	"github.com/e346m/upsider-wala/internal/ports/http"
	"github.com/e346m/upsider-wala/internal/usecases"
)

func NewCollectorHTTP(db *sql.DB, cfg *config.Config) *http.Handler {
	service := newCollectorSet(db, cfg)
	handler := http.NewHandler(service)
	return handler
}

func newCollectorSet(db *sql.DB, cfg *config.Config) *usecases.Usecase {
	psqlClient := psql.NewPSQLClient(db)
	return usecases.NewUsecase(psqlClient)
}
