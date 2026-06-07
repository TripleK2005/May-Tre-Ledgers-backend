CREATE TABLE roles (
    id          UUID PRIMARY KEY,
    name        VARCHAR(50) UNIQUE NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles(id, name)
VALUES
    (gen_random_uuid(), 'ADMIN'),
    (gen_random_uuid(), 'MANAGER'),
    (gen_random_uuid(), 'STAFF');

