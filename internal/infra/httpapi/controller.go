package httpapi

import (
	"context"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/domain/orders"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/pgstore"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type Controller struct {
	Queries *pgstore.Queries
	Logger  *zap.Logger
	Pool    *pgxpool.Pool
}

func NewController(queries *pgstore.Queries, logger *zap.Logger, pool *pgxpool.Pool) *Controller {
	return &Controller{
		Queries: queries,
		Logger:  logger,
		Pool:    pool,
	}
}

func (c *Controller) PostAddress(ctx *fiber.Ctx) error {
	var dto orders.Address
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid JSON"})
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addressID, err := c.Queries.InsertAddress(reqCtx,
		pgstore.InsertAddressParams{
			AdStreet:  dto.AdStreet,
			AdCity:    dto.AdCity,
			AdState:   dto.AdState,
			AdZip:     dto.AdZip,
			AdCountry: dto.AdCountry,
		})

	if err != nil {
		c.Logger.Error("failed to insert address", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert address"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"ad_addressid": addressID})
}

func (c *Controller) PostProducts(ctx *fiber.Ctx) error {
	var dto orders.Product
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid JSON"})
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productID, err := c.Queries.InsertProducts(reqCtx,
		pgstore.InsertProductsParams{
			PrDescription: dto.PrDescription,
			PrStock:       int32(dto.PrStock),
			PrPrice:       dto.PrPrice,
			PrVbactive:    dto.PrVbActive,
		})

	if err != nil {
		c.Logger.Error("failed to insert product", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert product"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"pr_productid": productID})
}

func (c *Controller) PostUsers(ctx *fiber.Ctx) error {
	var dto orders.User
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid JSON"})
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, err := c.Queries.InsertUsers(reqCtx,
		pgstore.InsertUsersParams{
			UsFirstname: dto.UsFirstname,
			UsLastname:  dto.UsLastname,
			UsEmail:     dto.UsEmail,
			UsVbactive:  dto.UsVbActive,
		})

	if err != nil {
		c.Logger.Error("failed to insert user", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert user"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"us_userid": userID})
}

func (c *Controller) PostOrders(ctx *fiber.Ctx) error {
	var dto orders.Order
	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid JSON"})
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := c.Pool.Begin(reqCtx)
	if err != nil {
		c.Logger.Error("failed to begin transaction", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(reqCtx)
		}
	}()

	orderID, err := c.Queries.InsertOrder(reqCtx, pgstore.InsertOrderParams{
		OrUserid:          dto.OrUserid,
		OrTotalamount:     dto.OrTotalamount,
		OrOrderstatus:     dto.OrOrderstatus,
		OrPaymentmethod:   dto.OrPaymentmethod,
		OrShippingaddress: dto.OrShippingaddress,
		OrCreatedat:       pgtype.Timestamp{Valid: true, Time: dto.OrCreatedat},
	})
	if err != nil {
		c.Logger.Error("failed to insert order", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	for _, product := range dto.OrProducts {
		err := c.Queries.InsertOrderProducts(reqCtx, pgstore.InsertOrderProductsParams{
			OrpOrderid:    pgtype.UUID{Valid: true, Bytes: orderID},
			OrpProductid:  product.OrpProductid,
			OrpQuantidad:  int32(product.OrpQuantidad),
			OrpTotalprice: product.OrpTotalprice,
		})
		if err != nil {
			c.Logger.Error("failed to insert order product", zap.Error(err))
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
		}
	}

	// Commit transaction if all insertions successful
	if err := tx.Commit(reqCtx); err != nil {
		c.Logger.Error("failed to commit transaction", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"order_id": orderID})

}
