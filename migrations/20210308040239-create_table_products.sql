-- +migrate Up
CREATE TABLE IF NOT EXISTS "products"
(
    "id"          char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "merchant_id" char(36)     NOT NULL,
    "sku"         varchar(128) NOT NULL,
    "name"        varchar(128) NOT NULL,
    "slug"        varchar(128) NOT NULL,
    "description" text,
    "created_at"  timestamp    NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "products";