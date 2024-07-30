CREATE TABLE IF NOT EXISTS address (
    "ad_addressid" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "ad_street"    VARCHAR(256)     NOT NULL,
    "ad_city"      VARCHAR(256)     NOT NULL,
    "ad_state"     VARCHAR(256)     NOT NULL,
    "ad_zip"       VARCHAR(256)     NOT NULL,
    "ad_country"   VARCHAR(256)     NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS address;
