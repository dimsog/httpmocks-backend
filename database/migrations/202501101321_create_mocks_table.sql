-- +goose Up
CREATE TABLE mocks (
    id uuid PRIMARY KEY,
    http_code smallint,
    content_type varchar(255),
    body text default null,
    created_at timestamp,
    updated_at timestamp
);

-- +goose Down
DROP TABLE mocks;