CREATE TABLE users (
    id serial primary key,
    first_name varchar(255) NOT NULL,
    middle_name varchar(255),
    last_name varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);