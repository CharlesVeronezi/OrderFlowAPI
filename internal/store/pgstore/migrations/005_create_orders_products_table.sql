CREATE TABLE IF NOT EXISTS orders_products
(
    "orp_orderid"    uuid             NOT NULL,
    "orp_productid"  uuid             NOT NULL,
    "orp_quantidad"  INTEGER          NOT NULL,
    "orp_totalprice" DOUBLE PRECISION NOT NULL,

    FOREIGN KEY (orp_orderid) REFERENCES orders (or_orderid),
    FOREIGN KEY (orp_productid) REFERENCES products (pr_productid)
)

---- create above / drop below ----

DROP TABLE IF EXISTS orders_products;