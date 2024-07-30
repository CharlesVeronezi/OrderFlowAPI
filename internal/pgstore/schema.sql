
CREATE TABLE IF NOT EXISTS address
(
    ad_addressid uuid PRIMARY KEY NOT NULL,
    ad_street    VARCHAR(256)     NOT NULL,
    ad_city      VARCHAR(256)     NOT NULL,
    ad_state     VARCHAR(256)     NOT NULL,
    ad_zip       VARCHAR(256)     NOT NULL,
    ad_country   VARCHAR(256)     NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    us_userid    uuid PRIMARY KEY NOT NULL,
    us_firstname VARCHAR(256)     NOT NULL,
    us_lastname  VARCHAR(256),
    us_email     VARCHAR(256),
    us_vbactive  BOOLEAN          NOT NULL
);

CREATE TABLE IF NOT EXISTS products
(
    pr_productid   uuid PRIMARY KEY NOT NULL,
    pr_description VARCHAR(512)     NOT NULL,
    pr_stock       INTEGER          NOT NULL,
    pr_price       DOUBLE PRECISION   NOT NULL,
    pr_vbactive    BOOLEAN          NOT NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    or_orderid         uuid PRIMARY KEY NOT NULL,
    or_userid          uuid REFERENCES users (us_userid) ON DELETE CASCADE,
    or_totalamount     NUMERIC(10, 2)   NOT NULL,
    or_orderstatus     VARCHAR(256)     NOT NULL,
    or_paymentmethod   VARCHAR(256)     NOT NULL,
    or_shippingaddress uuid REFERENCES address (ad_addressid) ON DELETE CASCADE,
    or_createdat       TIMESTAMP        NOT NULL,
    or_updatedat       TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders_products
(
    orp_orderid    uuid REFERENCES orders (or_orderid) ON DELETE CASCADE,
    orp_productid  uuid REFERENCES products (pr_productid) ON DELETE CASCADE,
    orp_quantidad  INTEGER        NOT NULL,
    orp_totalprice DOUBLE PRECISION NOT NULL
);
