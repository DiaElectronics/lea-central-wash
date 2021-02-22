// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetProgramReader is a Reader for the SetProgram structure.
type SetProgramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetProgramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSetProgramNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSetProgramInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetProgramNoContent creates a SetProgramNoContent with default headers values
func NewSetProgramNoContent() *SetProgramNoContent {
	return &SetProgramNoContent{}
}

/*SetProgramNoContent handles this case with default header values.

OK
*/
type SetProgramNoContent struct {
}

func (o *SetProgramNoContent) Error() string {
	return fmt.Sprintf("[POST /set-program][%d] setProgramNoContent ", 204)
}

func (o *SetProgramNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetProgramInternalServerError creates a SetProgramInternalServerError with default headers values
func NewSetProgramInternalServerError() *SetProgramInternalServerError {
	return &SetProgramInternalServerError{}
}

/*SetProgramInternalServerError handles this case with default header values.

internal error
*/
type SetProgramInternalServerError struct {
}

func (o *SetProgramInternalServerError) Error() string {
	return fmt.Sprintf("[POST /set-program][%d] setProgramInternalServerError ", 500)
}

func (o *SetProgramInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}