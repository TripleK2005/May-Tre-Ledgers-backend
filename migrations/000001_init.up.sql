CREATE TABLE IF NOT EXISTS schema_migrations_test (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    hello_world varchar(50)
);
