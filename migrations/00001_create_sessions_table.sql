-- noinspection SqlNoDataSourceInspectionForFile

-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions(
    id BIGSERIAL PRIMARY KEY,
    token_hash text UNIQUE NOT NULL,
    key text NOT NULL,
    value text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
