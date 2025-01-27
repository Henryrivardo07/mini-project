CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT,
    release_year INT CHECK (release_year >= 1980 AND release_year <= 2024),
    price INT NOT NULL,
    total_page INT,
    thickness VARCHAR(50),
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT NOW(),
    modified_by VARCHAR(255)
);
