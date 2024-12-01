-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, username, email, password, role_id) VALUES
('f47f1f0b-b56f-4f59-80b7-267cb2fc95b5', 'admin', 'admin@yopmail.com', '$2a$10$l9pMXaqHl2mS/BSBLdhq0e5HAas0FHaH4POGUymLkv0bhh6tTTJ.S', 'f0bce51b-8c3d-4f76-bf1d-38308c4699a7'),  -- Admin
('89b7c504-1050-4f71-b1de-b1d5f8bcb67b', 'employee', 'employee@yopmail.com', '$2a$10$SqILe1XO6TJ.kQZZhFpZWucGspNNyQyo5rOiy4A8rjdbJX5ZLrAFK', 'a123bfe6-d39f-42b2-b1e2-c7f7f4fc184d'); -- Employee
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id IN (
    'f47f1f0b-b56f-4f59-80b7-267cb2fc95b5', 
    '89b7c504-1050-4f71-b1de-b1d5f8bcb67b'  
);
-- +goose StatementEnd
