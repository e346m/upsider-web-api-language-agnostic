package di

import (
	"database/sql"

	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/internal/adapters/auth"
	"github.com/e346m/upsider-wala/internal/adapters/psql"
	"github.com/e346m/upsider-wala/internal/ports/http"
	"github.com/e346m/upsider-wala/internal/usecases"
	"github.com/e346m/upsider-wala/utils"
	"go.opentelemetry.io/otel/trace"
)

func New(db *sql.DB, cfg *config.Config, tp trace.Tracer) *http.Handler {
	identifier := utils.NewIdentifier()
	psqlClient := psql.NewPSQLClient(db, identifier)
	authClient := auth.NewAuthClient(cfg.SecretKey())
	usecase := usecases.NewUsecase(psqlClient, tp, authClient)
	handler := http.NewHandler(usecase, authClient)
	return handler
}
