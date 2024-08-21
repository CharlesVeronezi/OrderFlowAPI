CREATE TABLE IF NOT EXISTS orders
(
    "or_orderid"         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "or_userid"          uuid             NOT NULL,
    "or_totalamount"     DOUBLE PRECISION NOT NULL,
    "or_orderstatus"     VARCHAR(256)     NOT NULL,
    "or_paymentmethod"   VARCHAR(256)     NOT NULL,
    "or_shippingaddress" uuid             NOT NULL,
    "or_createdat"       TIMESTAMP        NOT NULL,
    "or_updatedat"       TIMESTAMP,

    FOREIGN KEY (or_userid) REFERENCES users (us_userid),
    FOREIGN KEY (or_shippingaddress) REFERENCES address (ad_addressid)
)

---- create above / drop below ----

DROP TABLE IF EXISTS orders;