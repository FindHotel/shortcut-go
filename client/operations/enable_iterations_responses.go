// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableIterationsReader is a Reader for the EnableIterations structure.
type EnableIterationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableIterationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableIterationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnableIterationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableIterationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewEnableIterationsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/iterations/enable] enableIterations", response, response.Code())
	}
}

// NewEnableIterationsNoContent creates a EnableIterationsNoContent with default headers values
func NewEnableIterationsNoContent() *EnableIterationsNoContent {
	return &EnableIterationsNoContent{}
}

/*
EnableIterationsNoContent describes a response with status code 204, with default header values.

No Content
*/
type EnableIterationsNoContent struct {
}

// IsSuccess returns true when this enable iterations no content response has a 2xx status code
func (o *EnableIterationsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable iterations no content response has a 3xx status code
func (o *EnableIterationsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable iterations no content response has a 4xx status code
func (o *EnableIterationsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable iterations no content response has a 5xx status code
func (o *EnableIterationsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this enable iterations no content response a status code equal to that given
func (o *EnableIterationsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the enable iterations no content response
func (o *EnableIterationsNoContent) Code() int {
	return 204
}

func (o *EnableIterationsNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsNoContent", 204)
}

func (o *EnableIterationsNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsNoContent", 204)
}

func (o *EnableIterationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableIterationsBadRequest creates a EnableIterationsBadRequest with default headers values
func NewEnableIterationsBadRequest() *EnableIterationsBadRequest {
	return &EnableIterationsBadRequest{}
}

/*
EnableIterationsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type EnableIterationsBadRequest struct {
}

// IsSuccess returns true when this enable iterations bad request response has a 2xx status code
func (o *EnableIterationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable iterations bad request response has a 3xx status code
func (o *EnableIterationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable iterations bad request response has a 4xx status code
func (o *EnableIterationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable iterations bad request response has a 5xx status code
func (o *EnableIterationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enable iterations bad request response a status code equal to that given
func (o *EnableIterationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the enable iterations bad request response
func (o *EnableIterationsBadRequest) Code() int {
	return 400
}

func (o *EnableIterationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsBadRequest", 400)
}

func (o *EnableIterationsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsBadRequest", 400)
}

func (o *EnableIterationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableIterationsNotFound creates a EnableIterationsNotFound with default headers values
func NewEnableIterationsNotFound() *EnableIterationsNotFound {
	return &EnableIterationsNotFound{}
}

/*
EnableIterationsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type EnableIterationsNotFound struct {
}

// IsSuccess returns true when this enable iterations not found response has a 2xx status code
func (o *EnableIterationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable iterations not found response has a 3xx status code
func (o *EnableIterationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable iterations not found response has a 4xx status code
func (o *EnableIterationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable iterations not found response has a 5xx status code
func (o *EnableIterationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enable iterations not found response a status code equal to that given
func (o *EnableIterationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enable iterations not found response
func (o *EnableIterationsNotFound) Code() int {
	return 404
}

func (o *EnableIterationsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsNotFound", 404)
}

func (o *EnableIterationsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsNotFound", 404)
}

func (o *EnableIterationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableIterationsUnprocessableEntity creates a EnableIterationsUnprocessableEntity with default headers values
func NewEnableIterationsUnprocessableEntity() *EnableIterationsUnprocessableEntity {
	return &EnableIterationsUnprocessableEntity{}
}

/*
EnableIterationsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type EnableIterationsUnprocessableEntity struct {
}

// IsSuccess returns true when this enable iterations unprocessable entity response has a 2xx status code
func (o *EnableIterationsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable iterations unprocessable entity response has a 3xx status code
func (o *EnableIterationsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable iterations unprocessable entity response has a 4xx status code
func (o *EnableIterationsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable iterations unprocessable entity response has a 5xx status code
func (o *EnableIterationsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this enable iterations unprocessable entity response a status code equal to that given
func (o *EnableIterationsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the enable iterations unprocessable entity response
func (o *EnableIterationsUnprocessableEntity) Code() int {
	return 422
}

func (o *EnableIterationsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsUnprocessableEntity", 422)
}

func (o *EnableIterationsUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/enable][%d] enableIterationsUnprocessableEntity", 422)
}

func (o *EnableIterationsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
