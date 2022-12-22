// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/FindHotel/shortcut-go/models"
)

// GetCustomFieldReader is a Reader for the GetCustomField structure.
type GetCustomFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomFieldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomFieldBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomFieldNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCustomFieldUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomFieldOK creates a GetCustomFieldOK with default headers values
func NewGetCustomFieldOK() *GetCustomFieldOK {
	return &GetCustomFieldOK{}
}

/*
GetCustomFieldOK describes a response with status code 200, with default header values.

Resource
*/
type GetCustomFieldOK struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this get custom field o k response has a 2xx status code
func (o *GetCustomFieldOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get custom field o k response has a 3xx status code
func (o *GetCustomFieldOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field o k response has a 4xx status code
func (o *GetCustomFieldOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom field o k response has a 5xx status code
func (o *GetCustomFieldOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field o k response a status code equal to that given
func (o *GetCustomFieldOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomFieldOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldOK  %+v", 200, o.Payload)
}

func (o *GetCustomFieldOK) String() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldOK  %+v", 200, o.Payload)
}

func (o *GetCustomFieldOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *GetCustomFieldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomFieldBadRequest creates a GetCustomFieldBadRequest with default headers values
func NewGetCustomFieldBadRequest() *GetCustomFieldBadRequest {
	return &GetCustomFieldBadRequest{}
}

/*
GetCustomFieldBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetCustomFieldBadRequest struct {
}

// IsSuccess returns true when this get custom field bad request response has a 2xx status code
func (o *GetCustomFieldBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom field bad request response has a 3xx status code
func (o *GetCustomFieldBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field bad request response has a 4xx status code
func (o *GetCustomFieldBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom field bad request response has a 5xx status code
func (o *GetCustomFieldBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field bad request response a status code equal to that given
func (o *GetCustomFieldBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCustomFieldBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldBadRequest ", 400)
}

func (o *GetCustomFieldBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldBadRequest ", 400)
}

func (o *GetCustomFieldBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomFieldNotFound creates a GetCustomFieldNotFound with default headers values
func NewGetCustomFieldNotFound() *GetCustomFieldNotFound {
	return &GetCustomFieldNotFound{}
}

/*
GetCustomFieldNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetCustomFieldNotFound struct {
}

// IsSuccess returns true when this get custom field not found response has a 2xx status code
func (o *GetCustomFieldNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom field not found response has a 3xx status code
func (o *GetCustomFieldNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field not found response has a 4xx status code
func (o *GetCustomFieldNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom field not found response has a 5xx status code
func (o *GetCustomFieldNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field not found response a status code equal to that given
func (o *GetCustomFieldNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCustomFieldNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldNotFound ", 404)
}

func (o *GetCustomFieldNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldNotFound ", 404)
}

func (o *GetCustomFieldNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomFieldUnprocessableEntity creates a GetCustomFieldUnprocessableEntity with default headers values
func NewGetCustomFieldUnprocessableEntity() *GetCustomFieldUnprocessableEntity {
	return &GetCustomFieldUnprocessableEntity{}
}

/*
GetCustomFieldUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetCustomFieldUnprocessableEntity struct {
}

// IsSuccess returns true when this get custom field unprocessable entity response has a 2xx status code
func (o *GetCustomFieldUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom field unprocessable entity response has a 3xx status code
func (o *GetCustomFieldUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field unprocessable entity response has a 4xx status code
func (o *GetCustomFieldUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom field unprocessable entity response has a 5xx status code
func (o *GetCustomFieldUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field unprocessable entity response a status code equal to that given
func (o *GetCustomFieldUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetCustomFieldUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldUnprocessableEntity ", 422)
}

func (o *GetCustomFieldUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/custom-fields/{custom-field-public-id}][%d] getCustomFieldUnprocessableEntity ", 422)
}

func (o *GetCustomFieldUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
