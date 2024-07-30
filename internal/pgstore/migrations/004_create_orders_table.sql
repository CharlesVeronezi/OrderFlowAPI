CREATE TABLE IF NOT EXISTS orders
(
    "or_orderid"         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "or_userid"          uuid REFERENCES users (us_userid) ON DELETE CASCADE,
    "or_totalamount"     NUMERIC(10, 2)   NOT NULL,
    "or_orderstatus"     VARCHAR(256)     NOT NULL,
    "or_paymentmethod"   VARCHAR(256)     NOT NULL,
    "or_shippingaddress" uuid REFERENCES address (ad_addressid) ON DELETE CASCADE,
    "or_createdat"       TIMESTAMP        NOT NULL,
    "or_updatedat"       TIMESTAMP
)

---- create above / drop below ----

DROP TABLE IF EXISTS orders;
