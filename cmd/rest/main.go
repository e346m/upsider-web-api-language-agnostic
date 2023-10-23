package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/e346m/upsider-wala/config"
	"github.com/e346m/upsider-wala/di"
	"github.com/e346m/upsider-wala/internal/adapters/auth"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"

	_ "github.com/jackc/pgx/v4/stdlib"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	cfg := config.LoadConfig()
	conn := initDB(cfg)
	tp := initTracer(cfg)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down trace provider: %v", err)
		}

		conn.Close()
	}()

	handler := di.New(conn, cfg, tp.Tracer("upsider-wala"))

	e := initEcho()

	api := e.Group("/api")
	{
		api.GET("/health", healthCheck)
		api.POST("/session", handler.SignIn)

		private := api.Group("")
		private.Use(echoJwtConfig(cfg.SecretKey()))
		{
			private.POST("/invoices", handler.CreateInvoice)
		}
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func initTracer(c *config.Config) *sdktrace.TracerProvider {
	var exporter sdktrace.SpanExporter
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
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

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(echoRecover())
	e.Use(echoSecureHeaderConfig())
	e.Use(otelecho.Middleware("wala"))
	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

func echoJwtConfig(key string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.CustomClaims)
		},
		SigningKey: []byte(key),
	})
}

func echoRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisableStackAll:   false,
		DisablePrintStack: false,
	})
}

func echoSecureHeaderConfig() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	})
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
