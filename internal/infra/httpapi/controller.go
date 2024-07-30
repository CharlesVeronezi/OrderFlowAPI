package httpapi

import (
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/domain/orders"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func PostOrders(c *fiber.Ctx) error {
	var dto orders.Order
	if err := c.BodyParser(&dto); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	// Obter o pool de conexões do contexto
	pool := c.Locals("pool").(*pgxpool.Pool)

	// Iniciar uma transação
	tx, err := pool.Begin(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to begin transaction"})
	}
	defer tx.Rollback(c.Context())

	// Inserir os dados do pedido no banco de dados
	query := `
		INSERT INTO orders (
			or_userid, or_totalamount, or_orderstatus, 
			or_paymentmethod, or_shippingaddress, 
			or_createdat, or_updatedat
		) VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING or_orderid`
	err = tx.QueryRow(c.Context(), query,
		dto.OrUserid, dto.OrTotalamount, dto.OrOrderstatus,
		dto.OrPaymentmethod, dto.OrShippingaddress.AdAddresid,
		time.Now(), time.Now()).Scan(&dto.OrOrderid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create order"})
	}

	// Inserir os produtos do pedido no banco de dados
	for _, product := range dto.OrProducts {
		query := `
			INSERT INTO order_products (
				or_orderid, pr_productid, pr_description, pr_price, pr_vbactive
			) VALUES ($1, $2, $3, $4, $5)`
		_, err := tx.Exec(c.Context(), query, dto.OrOrderid, product.PrProductID, product.PrDescription, product.PrPrice, product.PrVbActive)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create order products"})
		}
	}

	// Commit a transação
	err = tx.Commit(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to commit transaction"})
	}

	// Retornar a resposta com o pedido criado
	return c.Status(fiber.StatusCreated).JSON(dto)
}

func GetOrdersOrderID(c *fiber.Ctx) error {
	return nil
}

func GetOrdersUserID(c *fiber.Ctx) error {
	return nil
}
