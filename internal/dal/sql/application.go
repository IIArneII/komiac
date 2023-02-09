package sql

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ApplicationGetParams struct {
	UUID uuid.UUID `db:"uuid"`
}

type ApplicationGetListParams struct {
	DivisionOid string `db:"division_oid"`
	Year        int64  `db:"year"`
	Mnn         string `db:"mnn"`
}

type ApplicationDeleteParams struct {
	UUID      uuid.UUID `db:"uuid"`
	DeletedAt time.Time `db:"deleted_at"`
}

var (
	AplicationDeleteSql = `
	UPDATE applications
	SET deleted_at=:deleted_at
	WHERE uuid=:uuid
	`

	AplicationGetListSql = `
	SELECT * FROM applications WHERE deleted_at IS NULL and division_oid=:division_oid and year=:year and mnn=:mnn 
	`

	ApplicationGetSql = `
	SELECT * FROM applications WHERE uuid=:uuid AND deleted_at IS NULL
	`

	ApplicationCreateSql = `
	INSERT INTO applications (uuid, medical_organization_oid, division_oid, year, smnn, mnn, form, dosage, consumer_unit, item_unit, privilege_program_code, privilege_program, created_at, deleted_at)
    VALUES (:uuid, :medical_organization_oid, :division_oid, :year, :smnn, :mnn, :form, :dosage, :consumer_unit, :item_unit, :privilege_program_code, :privilege_program, :created_at, :deleted_at)
	`

	ApplicationUpdateSql = `
	UPDATE applications SET (medical_organization_oid, division_oid, year, smnn, mnn, form, dosage, consumer_unit, item_unit, privilege_program_code, privilege_program, modified_at) =
	(:medical_organization_oid, :division_oid, :year, :smnn, :mnn, :form, :dosage, :consumer_unit, :item_unit, :privilege_program_code, :privilege_program, :modified_at)
	WHERE uuid=:uuid
	`
)
