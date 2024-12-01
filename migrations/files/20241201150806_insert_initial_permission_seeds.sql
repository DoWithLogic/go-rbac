-- +goose Up
-- +goose StatementBegin
INSERT INTO permissions (id, name, description) VALUES
('f7e548cd-2e52-47f8-8e6a-1f9b2782bc44', 'users:read', 'Retrieve the list of users'),
('c2e69755-d017-4f4b-941a-cfe58e8bbd43', 'users:update', 'Update user details'),
('b763bc91-0885-4016-bd79-ded4f736a598', 'users:create', 'Create a new user'),
('f0c6d4b8-e2b3-459a-a04a-9b28107010bb', 'users:delete', 'Delete a user'),
('44c57f6f-7a79-42c1-bae5-447f4e2be6cd', 'profile:read', 'Retrieve the authenticated users profile'),
('7b5a31b1-6790-47b1-9264-d3c9a9e358fd', 'profile:update', 'Update the authenticated users profile'),
('8f3b739b-3258-48d9-bda3-12f33f639de2', 'auth:register', 'Register a new user'),
('67f130ae-7362-44b7-b7c2-5a9089b46425', 'auth:login', 'Authenticate and log in a user');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM permissions WHERE id IN (
    'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44',  -- 'users:read'
    'c2e69755-d017-4f4b-941a-cfe58e8bbd43',  -- 'users:update'
    'b763bc91-0885-4016-bd79-ded4f736a598',  -- 'users:create'
    'f0c6d4b8-e2b3-459a-a04a-9b28107010bb',  -- 'users:delete'
    '44c57f6f-7a79-42c1-bae5-447f4e2be6cd',  -- 'profile:read'
    '7b5a31b1-6790-47b1-9264-d3c9a9e358fd',  -- 'profile:update'
    '8f3b739b-3258-48d9-bda3-12f33f639de2',  -- 'auth:register'
    '67f130ae-7362-44b7-b7c2-5a9089b46425'   -- 'auth:login'
);
-- +goose StatementEnd
