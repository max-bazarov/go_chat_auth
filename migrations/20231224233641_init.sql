-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    PASSWORD varchar(255) NOT NULL,
    role varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE user;

-- +goose StatementEnd