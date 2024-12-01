-- +goose Up
-- +goose StatementBegin
INSERT INTO role_permissions (role_id, permission_id) VALUES
('e64f26ea-f6c5-4767-9750-d859619563db', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- 'users:read' for 'Admin' role
('e64f26ea-f6c5-4767-9750-d859619563db', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- 'users:update' for 'Admin' role
('e64f26ea-f6c5-4767-9750-d859619563db', 'b763bc91-0885-4016-bd79-ded4f736a598'), -- 'users:create' for 'Admin' role
('e64f26ea-f6c5-4767-9750-d859619563db', 'f0c6d4b8-e2b3-459a-a04a-9b28107010bb'), -- 'users:delete' for 'Admin' role
('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- 'users:read' for 'Employee' role
('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- 'users:update' for 'Employee' role
('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', '44c57f6f-7a79-42c1-bae5-447f4e2be6cd'), -- 'profile:read' for 'Employee' role
('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', '7b5a31b1-6790-47b1-9264-d3c9a9e358fd'); -- 'profile:update' for 'Employee' role

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM role_permissions WHERE (role_id, permission_id) IN (
    ('e64f26ea-f6c5-4767-9750-d859619563db', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- 'users:read' for 'Admin' role
    ('e64f26ea-f6c5-4767-9750-d859619563db', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- 'users:update' for 'Admin' role
    ('e64f26ea-f6c5-4767-9750-d859619563db', 'b763bc91-0885-4016-bd79-ded4f736a598'), -- 'users:create' for 'Admin' role
    ('e64f26ea-f6c5-4767-9750-d859619563db', 'f0c6d4b8-e2b3-459a-a04a-9b28107010bb'), -- 'users:delete' for 'Admin' role
    ('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', 'f7e548cd-2e52-47f8-8e6a-1f9b2782bc44'), -- 'users:read' for 'Employee' role
    ('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', 'c2e69755-d017-4f4b-941a-cfe58e8bbd43'), -- 'users:update' for 'Employee' role
    ('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', '44c57f6f-7a79-42c1-bae5-447f4e2be6cd'), -- 'profile:read' for 'Employee' role
    ('f54e4ae4-6a9e-4f43-bcfb-cfe2f03a43f9', '7b5a31b1-6790-47b1-9264-d3c9a9e358fd')  -- 'profile:update' for 'Employee' role
);
-- +goose StatementEnd
