-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN metadata JSONB DEFAULT '{}'::jsonb;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN metadata;
-- +goose StatementEnd
