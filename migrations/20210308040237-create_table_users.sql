-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         char(36)     PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name"       varchar(128) NOT NULL,
    "email"      varchar(128) NOT NULL,
    "password"   varchar(128) NOT NULL,
    "role"       role,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp    NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "users";