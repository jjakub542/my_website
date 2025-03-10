CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    is_superuser BOOLEAN
);

CREATE TABLE IF NOT EXISTS articles(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT UNIQUE,
    description TEXT, 
    content TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    public BOOLEAN
);

CREATE TABLE IF NOT EXISTS images(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    filename TEXT,
    uploaded_at TIMESTAMP DEFAULT now(),
    article_id UUID,
    CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE
);