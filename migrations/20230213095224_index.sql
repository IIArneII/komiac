-- +goose Up
-- +goose StatementBegin

DROP INDEX IF EXISTS application_index;
CREATE INDEX application_filter_index ON applications (division_oid, year, LOWER(mnn));
CREATE INDEX application_uuid_index ON applications (uuid);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS application_filter_index;
CREATE INDEX application_index ON applications (division_oid, year, mnn);
DROP INDEX IF EXISTS application_uuid_index;

-- +goose StatementEnd
