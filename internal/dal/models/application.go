package dal

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Application struct {
	UUID                   uuid.UUID `db:"uuid"`
	MedicalOrganizationOID string    `db:"medical_organization_oid"`
	DivisionOID            string    `db:"division_oid"`
	Year                   int       `db:"year"`
	SMNN                   string    `db:"smnn"`
	MNN                    string    `db:"mnn"`
	Form                   string    `db:"form"`
	Dosage                 string    `db:"dosage"`
	ConsumerUnit           string    `db:"consumer_unit"`
	ItemUnit               string    `db:"item_unit"`
	PrivilegeProgramCode   string    `db:"privilege_program_code"`
	PrivilegeProgram       string    `db:"privilege_program"`
	Created_at             time.Time `db:"created_at"`
	Deleted_at             time.Time `db:"deleted_at"`
}
