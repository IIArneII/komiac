-- +goose Up
-- +goose StatementBegin

ALTER TABLE applications
ADD COLUMN modified_at TIMESTAMP;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE applications
DROP COLUMN modified_at;

-- +goose StatementEnd
