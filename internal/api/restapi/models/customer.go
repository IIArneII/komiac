// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Customer customer
//
// swagger:model customer
type Customer struct {

	// o ID
	OID string `json:"OID,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Customer) UnmarshalJSON(data []byte) error {
	var props struct {

		// o ID
		OID string `json:"OID,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.OID = props.OID
	return nil
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer based on context it is used
func (m *Customer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
