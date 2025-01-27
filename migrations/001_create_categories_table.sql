CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT NOW(),
    modified_by VARCHAR(255)
);
