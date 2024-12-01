-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,                     -- UUID as CHAR(36) for MySQL
    username VARCHAR(255) UNIQUE NOT NULL,       -- Unique username
    email VARCHAR(255) UNIQUE NOT NULL,          -- Unique email
    password TEXT NOT NULL,                      -- Encrypted password
    role_id CHAR(36) NOT NULL,                   -- Role ID (UUID)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Created at timestamp
    updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP  -- Updated at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
