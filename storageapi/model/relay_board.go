// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RelayBoard relay board
// swagger:model RelayBoard
type RelayBoard string

const (

	// RelayBoardLocalGPIO captures enum value "localGPIO"
	RelayBoardLocalGPIO RelayBoard = "localGPIO"

	// RelayBoardDanBoard captures enum value "danBoard"
	RelayBoardDanBoard RelayBoard = "danBoard"
)

// for schema
var relayBoardEnum []interface{}

func init() {
	var res []RelayBoard
	if err := json.Unmarshal([]byte(`["localGPIO","danBoard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relayBoardEnum = append(relayBoardEnum, v)
	}
}

func (m RelayBoard) validateRelayBoardEnum(path, location string, value RelayBoard) error {
	if err := validate.Enum(path, location, value, relayBoardEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this relay board
func (m RelayBoard) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRelayBoardEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
