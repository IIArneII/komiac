package dal

import (
	"fmt"
	"komiac/internal/app/entities"

	uuid "github.com/satori/go.uuid"
)

var notMatshError = "Does not match: expected '%s', actual '%s'"

func СompareApplication(expected *entities.Application, actual *entities.Application) error {
	if expected.UUID.String() != actual.UUID.String() {
		return fmt.Errorf(notMatshError, expected.UUID.String(), actual.UUID.String())
	}
	if expected.MedicalOrganizationOID != actual.MedicalOrganizationOID {
		return fmt.Errorf(notMatshError, expected.MedicalOrganizationOID, expected.MedicalOrganizationOID)
	}
	if expected.DivisionOID != actual.DivisionOID {
		return fmt.Errorf(notMatshError, expected.DivisionOID, actual.DivisionOID)
	}
	if expected.Year != actual.Year {
		return fmt.Errorf(notMatshError, expected.Year, actual.Year)
	}
	if expected.SMNN != actual.SMNN {
		return fmt.Errorf(notMatshError, expected.SMNN, actual.SMNN)
	}
	if expected.MNN != actual.MNN {
		return fmt.Errorf(notMatshError, expected.MNN, actual.MNN)
	}
	if expected.Form != actual.Form {
		return fmt.Errorf(notMatshError, expected.Form, actual.Form)
	}
	if expected.Dosage != actual.Dosage {
		return fmt.Errorf(notMatshError, expected.Dosage, actual.Dosage)
	}
	if expected.ConsumerUnit != actual.ConsumerUnit {
		return fmt.Errorf(notMatshError, expected.ConsumerUnit, actual.ConsumerUnit)
	}
	if expected.ItemUnit != actual.ItemUnit {
		return fmt.Errorf(notMatshError, expected.ItemUnit, actual.ItemUnit)
	}
	if expected.PrivilegeProgramCode != actual.PrivilegeProgramCode {
		return fmt.Errorf(notMatshError, expected.PrivilegeProgramCode, actual.PrivilegeProgramCode)
	}
	if expected.PrivilegeProgram != actual.PrivilegeProgram {
		return fmt.Errorf(notMatshError, expected.PrivilegeProgram, actual.PrivilegeProgram)
	}
	return nil
}

func CreateUniqueApplication() *entities.Application {
	return &entities.Application{
		UUID:                   uuid.NewV4(),
		MedicalOrganizationOID: "1.2.643.5.1.13.13.12.2.42.4044",
		DivisionOID:            "1.2.643.5.1.13.13.12.2.42.4041.0.150024",
		Year:                   2023,
		SMNN:                   "21.20.10.116-000008-1-00180-0000000000000",
		MNN:                    "МЕСАЛАЗИН",
		Form:                   "ТАБЛЕТКИ С ПРОЛОНГИРОВАННЫМ ВЫСВОБОЖДЕНИЕМ",
		Dosage:                 "500 мг",
		ConsumerUnit:           "Штука",
		ItemUnit:               "мг",
		PrivilegeProgramCode:   "ОНЛП",
		PrivilegeProgram:       "ЛЛО в соответствии с ФЗ от 17.07.1999 №178-ФЗ \"О государственной социальной помощи\"",
	}
}
