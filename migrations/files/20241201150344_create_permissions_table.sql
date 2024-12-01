-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
    id CHAR(36) PRIMARY KEY,                    -- UUID as CHAR(36) for MySQL
    name VARCHAR(255) UNIQUE NOT NULL,           -- Permission name (e.g., 'users:read')
    description TEXT,                           -- Description of the permission
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Created at timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- Updated at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
