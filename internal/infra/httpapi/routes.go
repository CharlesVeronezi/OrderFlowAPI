package httpapi

import (
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/pgstore"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Handler(app *fiber.App, pool *pgxpool.Pool, logger *zap.Logger) {
	queries := pgstore.New(pool)
	controller := NewController(queries, logger, pool)

	// Middleware para adicionar o contexto ao Fiber
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("queries", queries)
		c.Locals("logger", logger)
		c.Locals("pool", pool)
		return c.Next()
	})

	app.Post("/address", controller.PostAddress)
	app.Post("/products", controller.PostProducts)
	app.Post("/users", controller.PostUsers)
	app.Post("/orders", controller.PostOrders)
}
