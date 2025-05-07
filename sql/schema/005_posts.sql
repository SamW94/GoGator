-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL, 
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feedid
    FOREIGN KEY (feed_id)
    REFERENCES public.feeds(id)
);
-- +goose Down
DROP TABLE posts;