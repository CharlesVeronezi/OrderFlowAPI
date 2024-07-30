CREATE TABLE IF NOT EXISTS users
(
    "us_userid"    uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    "us_firstname" VARCHAR(256)     NOT NULL,
    "us_lastname"  VARCHAR(256),
    "us_email"     VARCHAR(256),
    "us_vbactive"  BOOLEAN          NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS users;
