package di

import (
	"database/sql"

	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/adapters/psql"
	"github.com/e346m/upsider-wala/internal/ports/http"
	"github.com/e346m/upsider-wala/internal/usecases"
	"github.com/e346m/upsider-wala/utils"
	"go.opentelemetry.io/otel/trace"
)

func New(db *sql.DB, cfg *config.Config, tp trace.Tracer, identifier *utils.Identifier) *http.Handler {
	psqlClient := psql.NewPSQLClient(db, identifier)
	usecase := usecases.NewUsecase(psqlClient, tp)
	handler := http.NewHandler(usecase)
	return handler
}
