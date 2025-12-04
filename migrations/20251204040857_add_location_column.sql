-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN location JSONB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS location;
-- +goose StatementEnd
