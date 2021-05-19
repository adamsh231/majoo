-- +migrate Up
CREATE TABLE IF NOT EXISTS "product_images"
(
    "id"          char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "product_id"  char(36)     NOT NULL,
    "path"        varchar(128) NOT NULL,
    "alt"         varchar(128) NOT NULL,
    "created_at"  timestamp    NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "product_images";