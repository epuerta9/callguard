-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ALTER COLUMN metadata TYPE TEXT USING metadata::TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
ALTER COLUMN metadata TYPE JSONB USING metadata::JSONB;
-- +goose StatementEnd 