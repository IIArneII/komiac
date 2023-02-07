-- +goose Up
-- +goose StatementBegin

ALTER TABLE applications
RENAME preferential_direction TO privilege_program;

ALTER TABLE applications
RENAME preferential_direction_code TO privilege_program_code;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE applications
RENAME privilege_program TO preferential_direction;

ALTER TABLE applications
RENAME privilege_program_code TO preferential_direction_code;

-- +goose StatementEnd
