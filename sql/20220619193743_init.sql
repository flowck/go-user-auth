-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    password_reset_token UUID,
    password_reset_token_expiry_date TIMESTAMPTZ,
    create_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    update_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS roles (
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    create_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    update_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS user_roles (
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    user_Id UUID NOT NULL,
    role_Id UUID NOT NULL,
    create_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    update_at TIMESTAMPTZ  NOT NULL DEFAULT now(),

    CONSTRAINT fk_user
        FOREIGN KEY (user_Id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role
        FOREIGN KEY (role_Id)
        REFERENCES roles(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE roles;
DROP TABLE user_roles;
-- +goose StatementEnd
