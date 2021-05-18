-- +migrate Up
CREATE TABLE IF NOT EXISTS "credentials"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "email"      varchar(100) NOT NULL,
    "password"   varchar(128) NOT NULL,
    "role"       role         NOT NULL,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp    NOT NULL,
    "deleted_at" timestamp
    );

-- +migrate Down
DROP TABLE IF EXISTS "credentials";