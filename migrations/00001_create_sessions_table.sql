-- noinspection SqlNoDataSourceInspectionForFile

-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions(
    id BIGSERIAL PRIMARY KEY,
    token_hash text UNIQUE NOT NULL,
    key text UNIQUE NOT NULL,
    value text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
