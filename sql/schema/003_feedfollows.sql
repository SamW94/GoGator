-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_userid
    FOREIGN KEY (user_id)
    REFERENCES public.users(id)
    ON DELETE CASCADE,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feedid
    FOREIGN KEY (feed_id)
    REFERENCES public.feeds(id)
    ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);
-- +goose Down
DROP TABLE feed_follows;