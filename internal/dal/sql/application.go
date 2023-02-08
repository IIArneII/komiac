package sql

import uuid "github.com/satori/go.uuid"

type ApplicationGetParams struct {
	UUID uuid.UUID `db:"uuid"`
}

var (
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
