// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Application application
//
// swagger:model application
type Application struct {

	// m n n
	// Required: true
	MNN *string `json:"MNN"`

	// s m n n
	// Required: true
	SMNN *string `json:"SMNN"`

	// UUID
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"UUID"`

	// consumer unit
	// Required: true
	ConsumerUnit *string `json:"consumerUnit"`

	// division o ID
	// Required: true
	DivisionOID *string `json:"divisionOID"`

	// dosage
	// Required: true
	Dosage *string `json:"dosage"`

	// form
	// Required: true
	Form *string `json:"form"`

	// item unit
	// Required: true
	ItemUnit *string `json:"itemUnit"`

	// medical organization o ID
	// Required: true
	MedicalOrganizationOID *string `json:"medicalOrganizationOID"`

	// preferential direction
	// Required: true
	PreferentialDirection *string `json:"preferentialDirection"`

	// preferential direction code
	// Required: true
	PreferentialDirectionCode *string `json:"preferentialDirectionCode"`

	// year
	// Required: true
	Year *int64 `json:"year"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Application) UnmarshalJSON(data []byte) error {
	var props struct {

		// m n n
		// Required: true
		MNN *string `json:"MNN"`

		// s m n n
		// Required: true
		SMNN *string `json:"SMNN"`

		// UUID
		// Required: true
		// Format: uuid
		UUID *strfmt.UUID `json:"UUID"`

		// consumer unit
		// Required: true
		ConsumerUnit *string `json:"consumerUnit"`

		// division o ID
		// Required: true
		DivisionOID *string `json:"divisionOID"`

		// dosage
		// Required: true
		Dosage *string `json:"dosage"`

		// form
		// Required: true
		Form *string `json:"form"`

		// item unit
		// Required: true
		ItemUnit *string `json:"itemUnit"`

		// medical organization o ID
		// Required: true
		MedicalOrganizationOID *string `json:"medicalOrganizationOID"`

		// preferential direction
		// Required: true
		PreferentialDirection *string `json:"preferentialDirection"`

		// preferential direction code
		// Required: true
		PreferentialDirectionCode *string `json:"preferentialDirectionCode"`

		// year
		// Required: true
		Year *int64 `json:"year"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.MNN = props.MNN
	m.SMNN = props.SMNN
	m.UUID = props.UUID
	m.ConsumerUnit = props.ConsumerUnit
	m.DivisionOID = props.DivisionOID
	m.Dosage = props.Dosage
	m.Form = props.Form
	m.ItemUnit = props.ItemUnit
	m.MedicalOrganizationOID = props.MedicalOrganizationOID
	m.PreferentialDirection = props.PreferentialDirection
	m.PreferentialDirectionCode = props.PreferentialDirectionCode
	m.Year = props.Year
	return nil
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMNN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMNN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumerUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivisionOID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDosage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedicalOrganizationOID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferentialDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferentialDirectionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateMNN(formats strfmt.Registry) error {

	if err := validate.Required("MNN", "body", m.MNN); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateSMNN(formats strfmt.Registry) error {

	if err := validate.Required("SMNN", "body", m.SMNN); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("UUID", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("UUID", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateConsumerUnit(formats strfmt.Registry) error {

	if err := validate.Required("consumerUnit", "body", m.ConsumerUnit); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateDivisionOID(formats strfmt.Registry) error {

	if err := validate.Required("divisionOID", "body", m.DivisionOID); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateDosage(formats strfmt.Registry) error {

	if err := validate.Required("dosage", "body", m.Dosage); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateForm(formats strfmt.Registry) error {

	if err := validate.Required("form", "body", m.Form); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateItemUnit(formats strfmt.Registry) error {

	if err := validate.Required("itemUnit", "body", m.ItemUnit); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateMedicalOrganizationOID(formats strfmt.Registry) error {

	if err := validate.Required("medicalOrganizationOID", "body", m.MedicalOrganizationOID); err != nil {
		return err
	}

	return nil
}

func (m *Application) validatePreferentialDirection(formats strfmt.Registry) error {

	if err := validate.Required("preferentialDirection", "body", m.PreferentialDirection); err != nil {
		return err
	}

	return nil
}

func (m *Application) validatePreferentialDirectionCode(formats strfmt.Registry) error {

	if err := validate.Required("preferentialDirectionCode", "body", m.PreferentialDirectionCode); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateYear(formats strfmt.Registry) error {

	if err := validate.Required("year", "body", m.Year); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application based on context it is used
func (m *Application) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
