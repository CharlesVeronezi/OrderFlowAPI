package httpapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Handler(app *fiber.App, pool *pgxpool.Pool) {
	app.Post("/orders", PostOrders)
	app.Get("/orders/:orderId", GetOrdersOrderID)
	app.Get("/orders/:userId", GetOrdersUserID)
}
