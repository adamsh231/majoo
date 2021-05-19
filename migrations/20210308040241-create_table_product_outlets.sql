-- +migrate Up
CREATE TABLE IF NOT EXISTS "product_outlets"
(
    "id"          char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "product_id"  char(36)     NOT NULL,
    "outlet_id"   char(36)     NOT NULL,
    "price"       numeric      NOT NULL,
    "stock"       int          NOT NULL,
    "created_at"  timestamp    NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "product_outlets";