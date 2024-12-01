-- +goose Up
-- +goose StatementBegin
INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
-- Role: Admin
('f0bce51b-8c3d-4f76-bf1d-38308c4699a7', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- users:read
('f0bce51b-8c3d-4f76-bf1d-38308c4699a7', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- users:update
('f0bce51b-8c3d-4f76-bf1d-38308c4699a7', 'b763bc91-0885-4016-bd79-ded4f736a598'), -- users:create
('f0bce51b-8c3d-4f76-bf1d-38308c4699a7', 'f0c6d4b8-e2b3-459a-a04a-9b28107010bb'), -- users:delete

-- Role: Employee
('a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- users:read
('a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- users:update
('a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d', '44c57f6f-7a79-42c1-bae5-447f4e2be6cd'), -- profile:read
('a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d', '7b5a31b1-6790-47b1-9264-d3c9a9e358fd'), -- profile:update

-- Role: Customer
('8ec8bb2c-d50a-4d9b-9534-dc3be0e11b4f', '44c57f6f-7a79-42c1-bae5-447f4e2be6cd'), -- profile:read
('8ec8bb2c-d50a-4d9b-9534-dc3be0e11b4f', '7b5a31b1-6790-47b1-9264-d3c9a9e358fd');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `role_permissions`
WHERE `role_id` IN (
    'f0bce51b-8c3d-4f76-bf1d-38308c4699a7',
    'a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d',
    '8ec8bb2c-d50a-4d9b-9534-dc3be0e11b4f'
) AND `permission_id` IN (
    'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44',
    'c2e69755-d017-4f4b-941a-cfe58e8bbd43',
    'b763bc91-0885-4016-bd79-ded4f736a598',
    'f0c6d4b8-e2b3-459a-a04a-9b28107010bb',
    '44c57f6f-7a79-42c1-bae5-447f4e2be6cd',
    '7b5a31b1-6790-47b1-9264-d3c9a9e358fd'
);
-- +goose StatementEnd
