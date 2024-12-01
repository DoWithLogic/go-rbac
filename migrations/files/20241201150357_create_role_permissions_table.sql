-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_permissions (
    role_id CHAR(36) NOT NULL,                   -- Role ID (UUID)
    permission_id CHAR(36) NOT NULL,             -- Permission ID (UUID)
    PRIMARY KEY (role_id, permission_id)        -- Composite primary key
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
-- +goose StatementEnd
