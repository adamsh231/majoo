-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchants"
(
    "id"         char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name"       varchar(128) NOT NULL,
    "phone"      varchar(128) NOT NULL,
    "alamat"     text         NOT NULL,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp    NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "merchants";