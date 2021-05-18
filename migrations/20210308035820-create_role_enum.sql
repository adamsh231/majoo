-- +migrate Up

-- +migrate StatementBegin
DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM ('seller', 'supplier', 'customer');
    END IF;
END
$$;
-- +migrate StatementEnd

-- +migrate Down
DROP TYPE role;