-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by TIMESTAMP
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR,
    image_url VARCHAR,
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR,
    category_id INTEGER REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by TIMESTAMP
);

-- +migrate Down
-- DROP TABLE IF EXISTS books;
-- DROP TABLE IF EXISTS categories;
-- DROP TABLE IF EXISTS users;
