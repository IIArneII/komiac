package entities

import uuid "github.com/satori/go.uuid"

type Application struct {
	UUID                   uuid.UUID
	MedicalOrganizationOID string
	DivisionOID            string
	Year                   int
	SMNN                   string
	MNN                    string
	Form                   string
	Dosage                 string
	ConsumerUnit           string
	ItemUnit               string
	PrivilegeProgramCode   string
	PrivilegeProgram       string
}

type ApplicationFilter struct {
	DivisionOID string
	Year        int64
	MNN         string
}
