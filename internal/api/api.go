package api

import (
	"encoding/json"
	"github.com/CharlesVeronezi/OrderFlowAPI/internal/store/pgstore/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgtype"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"net/http"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
	c *amqp.Connection
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries, conn *amqp.Connection) http.Handler {
	a := apiHandler{
		q: q,
		c: conn,
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/address", func(r chi.Router) {
			r.Post("/", a.handleCreateAddress)
		})
		r.Route("/products", func(r chi.Router) {
			r.Post("/", a.handleCreateProduct)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/", a.handleCreateUser)
		})
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", a.handleCreateOrder)
			r.Get("/:orderID", a.handleGetOrder)
			r.Put("/:orderID/conclude", a.handleConcludeOrder)
		})
	})

	a.r = r

	return a
}

func (h apiHandler) handleCreateAddress(w http.ResponseWriter, r *http.Request) {
	var body Address
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	addressID, err := h.q.InsertAddress(r.Context(),
		pgstore.InsertAddressParams{
			AdStreet:  body.AdStreet,
			AdCity:    body.AdCity,
			AdState:   body.AdState,
			AdZip:     body.AdZip,
			AdCountry: body.AdCountry,
		},
	)

	if err != nil {
		slog.Error("Failed to insert address", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	sendJSON(w, DefaultResponseID{ID: addressID.String(), Message: "Address entered successfully"})

}

func (h apiHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var body Product
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	productID, err := h.q.InsertProducts(r.Context(),
		pgstore.InsertProductsParams{
			PrDescription: body.PrDescription,
			PrStock:       int32(body.PrStock),
			PrPrice:       body.PrPrice,
			PrVbactive:    body.PrVbActive,
		},
	)

	if err != nil {
		slog.Error("Failed to insert product", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	sendJSON(w, DefaultResponseID{ID: productID.String(), Message: "Product entered successfully"})
}

func (h apiHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	userID, err := h.q.InsertUsers(r.Context(),
		pgstore.InsertUsersParams{
			UsFirstname: body.UsFirstname,
			UsLastname:  pgtype.Text{String: body.UsLastname, Valid: true},
			UsEmail:     pgtype.Text{String: body.UsEmail, Valid: true},
			UsVbactive:  body.UsVbActive,
		},
	)

	if err != nil {
		slog.Error("Failed to insert product", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	sendJSON(w, DefaultResponseID{ID: userID.String(), Message: "User entered successfully"})
}

func (h apiHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var body Order

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	orderBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Failed to marshal order data", http.StatusInternalServerError)
		return
	}

	err = publishOrderToRabbitMQ(orderBytes, h.c)
	if err != nil {
		slog.Error("Failed to publish order to RabbitMQ:", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	sendJSON(w, "Order created successfully")

}

func (h apiHandler) handleGetOrder(w http.ResponseWriter, r *http.Request)      {}
func (h apiHandler) handleConcludeOrder(w http.ResponseWriter, r *http.Request) {}
