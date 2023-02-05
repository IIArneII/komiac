-- +goose Up
-- +goose StatementBegin

CREATE TABLE applications
(
    uuid                        uuid        PRIMARY KEY,
    medical_organization_oid    TEXT        NOT NULL,
    division_oid                TEXT        NOT NULL,
    year                        SMALLINT    NOT NULL,
    smnn                        TEXT        NOT NULL,
    mnn                         TEXT        NOT NULL,
    form                        TEXT        NOT NULL,
    dosage                      TEXT        NOT NULL,
    consumer_unit               TEXT        NOT NULL,
    item_unit                   TEXT        NOT NULL,
    preferential_direction_code TEXT        NOT NULL,
    preferential_direction      TEXT        NOT NULL,
    created_at                  TIMESTAMP   NOT NULL,
    deleted_at                  TIMESTAMP
);

CREATE INDEX application_index ON applications (division_oid, year, mnn);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS applications;

DROP INDEX IF EXISTS application_index;

-- +goose StatementEnd
