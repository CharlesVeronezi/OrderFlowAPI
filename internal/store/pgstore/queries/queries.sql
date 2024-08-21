-- name: GetOrder :one
SELECT "or_orderid",
       "or_userid",
       "or_totalamount",
       "or_orderstatus",
       "or_paymentmethod",
       "or_shippingaddress",
       "or_createdat",
       "or_updatedat"
FROM orders
WHERE "or_orderid" = $1;

-- name: GetOrderOrderID :many
SELECT
    ord.or_orderid,
    ord.or_totalamount,
    ord.or_orderstatus,
    ord.or_paymentmethod,
    ord.or_createdat,
    pdt.pr_description,
    pdt.pr_price,
    orp.orp_quantidad,
    orp.orp_totalprice
FROM orders ord
         JOIN orders_products orp ON orp.orp_orderid = ord.or_orderid
         JOIN products pdt ON pdt.pr_productid = orp.orp_productid
WHERE ord.or_orderid = $1;

-- name: InsertOrder :one
INSERT INTO orders (
    "or_userid",
    "or_totalamount",
    "or_orderstatus",
    "or_paymentmethod",
    "or_shippingaddress",
    "or_createdat"
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "or_orderid";

-- name: UpdateOrder :exec
UPDATE orders
SET "or_orderstatus" = $1,
    "or_updatedat" = $2
WHERE or_orderid = $3;

-- name: InsertOrderProducts :exec
INSERT INTO orders_products (
    "orp_orderid",
    "orp_productid",
    "orp_quantidad",
    "orp_totalprice"
)
VALUES ($1, $2, $3, $4);

-- name: GetOrderProducts :many
SELECT "orp_orderid",
       "orp_productid",
       "orp_quantidad",
       "orp_totalprice"
FROM orders_products
WHERE "orp_orderid" = $1;

-- name: InsertProducts :one
INSERT INTO products (
    "pr_description",
    "pr_stock",
    "pr_price",
    "pr_vbactive"
)
VALUES ($1, $2, $3, $4)
RETURNING "pr_productid";

-- name: GetProducts :one
SELECT "pr_productid",
       "pr_description",
       "pr_stock",
       "pr_price",
       "pr_vbactive"
FROM products
WHERE "pr_productid" = $1;

-- name: InsertUsers :one
INSERT INTO users (
    "us_firstname",
    "us_lastname",
    "us_email",
    "us_vbactive"
)
VALUES ($1, $2, $3, $4)
RETURNING "us_userid";

-- name: GetUsers :one
SELECT "us_userid",
       "us_firstname",
       "us_lastname",
       "us_email",
       "us_vbactive"
FROM users
WHERE "us_userid" = $1;

-- name: InsertAddress :one
INSERT INTO address (
    "ad_street",
    "ad_city",
    "ad_state",
    "ad_zip",
    "ad_country"
)
VALUES ($1, $2, $3, $4, $5)
RETURNING "ad_addressid";

-- name: GetAddress :one
SELECT "ad_addressid",
       "ad_street",
       "ad_city",
       "ad_state",
       "ad_zip",
       "ad_country"
FROM address
WHERE "ad_addressid" = $1;

