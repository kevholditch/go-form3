// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PayportAssociationAttributes payport association attributes
// swagger:model PayportAssociationAttributes

type PayportAssociationAttributes struct {

	// customer sending fps institution
	CustomerSendingFpsInstitution string `json:"customer_sending_fps_institution,omitempty"`

	// participant id
	ParticipantID string `json:"participant_id,omitempty"`

	// sponsor account number
	SponsorAccountNumber string `json:"sponsor_account_number,omitempty"`

	// sponsor bank id
	SponsorBankID string `json:"sponsor_bank_id,omitempty"`
}

/* polymorph PayportAssociationAttributes customer_sending_fps_institution false */

/* polymorph PayportAssociationAttributes participant_id false */

/* polymorph PayportAssociationAttributes sponsor_account_number false */

/* polymorph PayportAssociationAttributes sponsor_bank_id false */

// Validate validates this payport association attributes
func (m *PayportAssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PayportAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayportAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res PayportAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
