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

// UpdateIterationReader is a Reader for the UpdateIteration structure.
type UpdateIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIterationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateIterationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateIterationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/iterations/{iteration-public-id}] updateIteration", response, response.Code())
	}
}

// NewUpdateIterationOK creates a UpdateIterationOK with default headers values
func NewUpdateIterationOK() *UpdateIterationOK {
	return &UpdateIterationOK{}
}

/*
UpdateIterationOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateIterationOK struct {
	Payload *models.Iteration
}

// IsSuccess returns true when this update iteration o k response has a 2xx status code
func (o *UpdateIterationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update iteration o k response has a 3xx status code
func (o *UpdateIterationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update iteration o k response has a 4xx status code
func (o *UpdateIterationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update iteration o k response has a 5xx status code
func (o *UpdateIterationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update iteration o k response a status code equal to that given
func (o *UpdateIterationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update iteration o k response
func (o *UpdateIterationOK) Code() int {
	return 200
}

func (o *UpdateIterationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationOK %s", 200, payload)
}

func (o *UpdateIterationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationOK %s", 200, payload)
}

func (o *UpdateIterationOK) GetPayload() *models.Iteration {
	return o.Payload
}

func (o *UpdateIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Iteration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIterationBadRequest creates a UpdateIterationBadRequest with default headers values
func NewUpdateIterationBadRequest() *UpdateIterationBadRequest {
	return &UpdateIterationBadRequest{}
}

/*
UpdateIterationBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateIterationBadRequest struct {
}

// IsSuccess returns true when this update iteration bad request response has a 2xx status code
func (o *UpdateIterationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update iteration bad request response has a 3xx status code
func (o *UpdateIterationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update iteration bad request response has a 4xx status code
func (o *UpdateIterationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update iteration bad request response has a 5xx status code
func (o *UpdateIterationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update iteration bad request response a status code equal to that given
func (o *UpdateIterationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update iteration bad request response
func (o *UpdateIterationBadRequest) Code() int {
	return 400
}

func (o *UpdateIterationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationBadRequest", 400)
}

func (o *UpdateIterationBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationBadRequest", 400)
}

func (o *UpdateIterationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIterationNotFound creates a UpdateIterationNotFound with default headers values
func NewUpdateIterationNotFound() *UpdateIterationNotFound {
	return &UpdateIterationNotFound{}
}

/*
UpdateIterationNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateIterationNotFound struct {
}

// IsSuccess returns true when this update iteration not found response has a 2xx status code
func (o *UpdateIterationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update iteration not found response has a 3xx status code
func (o *UpdateIterationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update iteration not found response has a 4xx status code
func (o *UpdateIterationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update iteration not found response has a 5xx status code
func (o *UpdateIterationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update iteration not found response a status code equal to that given
func (o *UpdateIterationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update iteration not found response
func (o *UpdateIterationNotFound) Code() int {
	return 404
}

func (o *UpdateIterationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationNotFound", 404)
}

func (o *UpdateIterationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationNotFound", 404)
}

func (o *UpdateIterationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIterationUnprocessableEntity creates a UpdateIterationUnprocessableEntity with default headers values
func NewUpdateIterationUnprocessableEntity() *UpdateIterationUnprocessableEntity {
	return &UpdateIterationUnprocessableEntity{}
}

/*
UpdateIterationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateIterationUnprocessableEntity struct {
}

// IsSuccess returns true when this update iteration unprocessable entity response has a 2xx status code
func (o *UpdateIterationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update iteration unprocessable entity response has a 3xx status code
func (o *UpdateIterationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update iteration unprocessable entity response has a 4xx status code
func (o *UpdateIterationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update iteration unprocessable entity response has a 5xx status code
func (o *UpdateIterationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update iteration unprocessable entity response a status code equal to that given
func (o *UpdateIterationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update iteration unprocessable entity response
func (o *UpdateIterationUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateIterationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationUnprocessableEntity", 422)
}

func (o *UpdateIterationUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/{iteration-public-id}][%d] updateIterationUnprocessableEntity", 422)
}

func (o *UpdateIterationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
