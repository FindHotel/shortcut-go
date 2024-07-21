// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/FindHotel/shortcut-go/models"
)

// UpdateKeyResultReader is a Reader for the UpdateKeyResult structure.
type UpdateKeyResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKeyResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKeyResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateKeyResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateKeyResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateKeyResultUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/key-results/{key-result-public-id}] updateKeyResult", response, response.Code())
	}
}

// NewUpdateKeyResultOK creates a UpdateKeyResultOK with default headers values
func NewUpdateKeyResultOK() *UpdateKeyResultOK {
	return &UpdateKeyResultOK{}
}

/*
UpdateKeyResultOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateKeyResultOK struct {
	Payload *models.KeyResult
}

// IsSuccess returns true when this update key result o k response has a 2xx status code
func (o *UpdateKeyResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update key result o k response has a 3xx status code
func (o *UpdateKeyResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update key result o k response has a 4xx status code
func (o *UpdateKeyResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update key result o k response has a 5xx status code
func (o *UpdateKeyResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update key result o k response a status code equal to that given
func (o *UpdateKeyResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update key result o k response
func (o *UpdateKeyResultOK) Code() int {
	return 200
}

func (o *UpdateKeyResultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultOK %s", 200, payload)
}

func (o *UpdateKeyResultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultOK %s", 200, payload)
}

func (o *UpdateKeyResultOK) GetPayload() *models.KeyResult {
	return o.Payload
}

func (o *UpdateKeyResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKeyResultBadRequest creates a UpdateKeyResultBadRequest with default headers values
func NewUpdateKeyResultBadRequest() *UpdateKeyResultBadRequest {
	return &UpdateKeyResultBadRequest{}
}

/*
UpdateKeyResultBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateKeyResultBadRequest struct {
}

// IsSuccess returns true when this update key result bad request response has a 2xx status code
func (o *UpdateKeyResultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update key result bad request response has a 3xx status code
func (o *UpdateKeyResultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update key result bad request response has a 4xx status code
func (o *UpdateKeyResultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update key result bad request response has a 5xx status code
func (o *UpdateKeyResultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update key result bad request response a status code equal to that given
func (o *UpdateKeyResultBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update key result bad request response
func (o *UpdateKeyResultBadRequest) Code() int {
	return 400
}

func (o *UpdateKeyResultBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultBadRequest", 400)
}

func (o *UpdateKeyResultBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultBadRequest", 400)
}

func (o *UpdateKeyResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKeyResultNotFound creates a UpdateKeyResultNotFound with default headers values
func NewUpdateKeyResultNotFound() *UpdateKeyResultNotFound {
	return &UpdateKeyResultNotFound{}
}

/*
UpdateKeyResultNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateKeyResultNotFound struct {
}

// IsSuccess returns true when this update key result not found response has a 2xx status code
func (o *UpdateKeyResultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update key result not found response has a 3xx status code
func (o *UpdateKeyResultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update key result not found response has a 4xx status code
func (o *UpdateKeyResultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update key result not found response has a 5xx status code
func (o *UpdateKeyResultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update key result not found response a status code equal to that given
func (o *UpdateKeyResultNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update key result not found response
func (o *UpdateKeyResultNotFound) Code() int {
	return 404
}

func (o *UpdateKeyResultNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultNotFound", 404)
}

func (o *UpdateKeyResultNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultNotFound", 404)
}

func (o *UpdateKeyResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKeyResultUnprocessableEntity creates a UpdateKeyResultUnprocessableEntity with default headers values
func NewUpdateKeyResultUnprocessableEntity() *UpdateKeyResultUnprocessableEntity {
	return &UpdateKeyResultUnprocessableEntity{}
}

/*
UpdateKeyResultUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateKeyResultUnprocessableEntity struct {
}

// IsSuccess returns true when this update key result unprocessable entity response has a 2xx status code
func (o *UpdateKeyResultUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update key result unprocessable entity response has a 3xx status code
func (o *UpdateKeyResultUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update key result unprocessable entity response has a 4xx status code
func (o *UpdateKeyResultUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update key result unprocessable entity response has a 5xx status code
func (o *UpdateKeyResultUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update key result unprocessable entity response a status code equal to that given
func (o *UpdateKeyResultUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update key result unprocessable entity response
func (o *UpdateKeyResultUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateKeyResultUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultUnprocessableEntity", 422)
}

func (o *UpdateKeyResultUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/key-results/{key-result-public-id}][%d] updateKeyResultUnprocessableEntity", 422)
}

func (o *UpdateKeyResultUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
