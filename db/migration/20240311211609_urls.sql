-- migrate:up
CREATE TABLE urls (
    id bigserial PRIMARY KEY,
    long_url varchar not null,
    short_url varchar not null,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "urls" ("short_url");

-- migrate:down
DROP TABLE IF EXISTS urls;