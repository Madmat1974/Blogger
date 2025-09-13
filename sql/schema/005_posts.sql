-- +goose Up
CREATE TABLE posts (
id              UUID PRIMARY KEY,
feed_id         UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
title           TEXT NULL,
url             TEXT NOT NULL UNIQUE,
description     TEXT NULL,
published_at    TIMESTAMPTZ NULL,
created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS posts;