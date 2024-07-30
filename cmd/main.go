package main

import (
	"context"
	"fmt"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/infra/httpapi"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Finalized API")
}

func run(ctx context.Context) error {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	logger = logger.Named("order_flow")
	defer func() { _ = logger.Sync() }()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "user=postgres password=pgpassword host=localhost port=5432 dbname=order_flow"
	}
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error("unable to connect to database", zap.Error(err))
		return err
	}
	defer pool.Close()

	app := fiber.New()
	httpapi.Handler(app, pool)

	err = app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}
