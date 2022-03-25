CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    title text NOT NULL,
    body text NOT NULL,
    tags varchar(100)[],
    created_at timestamp DEFAULT current_timestamp
);