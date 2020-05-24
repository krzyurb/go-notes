CREATE TABLE IF NOT EXISTS notes(
    id serial PRIMARY KEY,
    title VARCHAR (100) UNIQUE NOT NULL,
    content TEXT,
    is_public BOOLEAN,
    created_at DATE,
    deleted_at DATE,
    updated_at DATE
);
