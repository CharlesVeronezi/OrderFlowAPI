package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/api"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/store/pgstore/pgstore"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	slog.Error("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	handler := api.NewHandler(pgstore.New(pool), conn)

	go func() {
		if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
