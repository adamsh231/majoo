-- +migrate Up
CREATE TABLE IF NOT EXISTS "outlets"
(
    "id"          char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "merchant_id" char(36)     NOT NULL,
    "name"        varchar(128) NOT NULL,
    "phone"       varchar(128) NOT NULL,
    "address"     text         NOT NULL,
    "created_at"  timestamp    NOT NULL,
    "updated_at"  timestamp    NOT NULL,
    "deleted_at"  timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "outlets";