package dal

import (
	"fmt"
	"komiac/internal/app/entities"

	uuid "github.com/satori/go.uuid"
)

var notMatshError = "Does not match: expected '%s', actual '%s'"

func –°ompareApplication(expected *entities.Application, actual *entities.Application) error {
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
		MNN:                    "–ú–ē–°–ź–õ–ź–ó–ė–Ě",
		Form:                   "–Ę–ź–Ď–õ–ē–Ę–ö–ė –° –ü–†–ě–õ–ě–Ě–ď–ė–†–ě–í–ź–Ě–Ě–ę–ú –í–ę–°–í–ě–Ď–ě–Ė–Ē–ē–Ě–ė–ē–ú",
		Dosage:                 "500 –ľ–≥",
		ConsumerUnit:           "–®—ā—É–ļ–į",
		ItemUnit:               "–ľ–≥",
		PrivilegeProgramCode:   "–ě–Ě–õ–ü",
		PrivilegeProgram:       "–õ–õ–ě –≤ —Ā–ĺ–ĺ—ā–≤–Ķ—ā—Ā—ā–≤–ł–ł —Ā –§–ó –ĺ—ā 17.07.1999 ‚ĄĖ178-–§–ó \"–ě –≥–ĺ—Ā—É–ī–į—Ä—Ā—ā–≤–Ķ–Ĺ–Ĺ–ĺ–Ļ —Ā–ĺ—Ü–ł–į–Ľ—Ć–Ĺ–ĺ–Ļ –Ņ–ĺ–ľ–ĺ—Č–ł\"",
	}
}
