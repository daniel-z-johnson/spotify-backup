-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
    id BIGSERIAL PRIMARY KEY,
    token_hash      text,
    access_token    text,
    token_type      text,
    scope           text,
    expires_in      timestamptz,
    refresh_token   text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
-- +goose StatementEnd
