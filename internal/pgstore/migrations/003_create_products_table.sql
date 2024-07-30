CREATE TABLE IF NOT EXISTS products
(
    "pr_productid"   uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "pr_description" VARCHAR(512)     NOT NULL,
    "pr_stock"       INTEGER          NOT NULL,
    "pr_price"       DOUBLE PRECISION   NOT NULL,
    "pr_vbactive"    BOOLEAN          NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS products;