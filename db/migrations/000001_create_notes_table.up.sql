CREATE TABLE IF NOT EXISTS notes(
    user_id serial PRIMARY KEY,
    title VARCHAR (100) UNIQUE NOT NULL,
    password TEXT,
    is_public BOOLEAN
);
