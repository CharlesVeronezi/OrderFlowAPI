CREATE TABLE IF NOT EXISTS orders_products
(
    "orp_orderid"    uuid REFERENCES orders (or_orderid) ON DELETE CASCADE,
    "orp_productid"  uuid REFERENCES products (pr_productid) ON DELETE CASCADE,
    "orp_quantidad"  INTEGER        NOT NULL,
    "orp_totalprice" DOUBLE PRECISION NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS orders_products;
