-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (id, name) VALUES
('f0bce51b-8c3d-4f76-bf1d-38308c4699a7', 'ADMIN'),
('a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d', 'EMPLOYEE'),
('8ec8bb2c-d50a-4d9b-9534-dc3be0e11b4f', 'CUSTOMER');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM roles WHERE id IN (
    'f0bce51b-8c3d-4f76-bf1d-38308c4699a7',  -- 'ADMIN'
    'a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d',  -- 'EMPLOYEE'
    '8ec8bb2c-d50a-4d9b-9534-dc3be0e11b4f'   -- 'CUSTOMER'
);
-- +goose StatementEnd
