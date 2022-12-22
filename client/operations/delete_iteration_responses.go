// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteIterationReader is a Reader for the DeleteIteration structure.
type DeleteIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIterationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIterationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIterationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteIterationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIterationNoContent creates a DeleteIterationNoContent with default headers values
func NewDeleteIterationNoContent() *DeleteIterationNoContent {
	return &DeleteIterationNoContent{}
}

/*
DeleteIterationNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteIterationNoContent struct {
}

// IsSuccess returns true when this delete iteration no content response has a 2xx status code
func (o *DeleteIterationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete iteration no content response has a 3xx status code
func (o *DeleteIterationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete iteration no content response has a 4xx status code
func (o *DeleteIterationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete iteration no content response has a 5xx status code
func (o *DeleteIterationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete iteration no content response a status code equal to that given
func (o *DeleteIterationNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteIterationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationNoContent ", 204)
}

func (o *DeleteIterationNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationNoContent ", 204)
}

func (o *DeleteIterationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIterationBadRequest creates a DeleteIterationBadRequest with default headers values
func NewDeleteIterationBadRequest() *DeleteIterationBadRequest {
	return &DeleteIterationBadRequest{}
}

/*
DeleteIterationBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteIterationBadRequest struct {
}

// IsSuccess returns true when this delete iteration bad request response has a 2xx status code
func (o *DeleteIterationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete iteration bad request response has a 3xx status code
func (o *DeleteIterationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete iteration bad request response has a 4xx status code
func (o *DeleteIterationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete iteration bad request response has a 5xx status code
func (o *DeleteIterationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete iteration bad request response a status code equal to that given
func (o *DeleteIterationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIterationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationBadRequest ", 400)
}

func (o *DeleteIterationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationBadRequest ", 400)
}

func (o *DeleteIterationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIterationNotFound creates a DeleteIterationNotFound with default headers values
func NewDeleteIterationNotFound() *DeleteIterationNotFound {
	return &DeleteIterationNotFound{}
}

/*
DeleteIterationNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteIterationNotFound struct {
}

// IsSuccess returns true when this delete iteration not found response has a 2xx status code
func (o *DeleteIterationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete iteration not found response has a 3xx status code
func (o *DeleteIterationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete iteration not found response has a 4xx status code
func (o *DeleteIterationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete iteration not found response has a 5xx status code
func (o *DeleteIterationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete iteration not found response a status code equal to that given
func (o *DeleteIterationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIterationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationNotFound ", 404)
}

func (o *DeleteIterationNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationNotFound ", 404)
}

func (o *DeleteIterationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIterationUnprocessableEntity creates a DeleteIterationUnprocessableEntity with default headers values
func NewDeleteIterationUnprocessableEntity() *DeleteIterationUnprocessableEntity {
	return &DeleteIterationUnprocessableEntity{}
}

/*
DeleteIterationUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteIterationUnprocessableEntity struct {
}

// IsSuccess returns true when this delete iteration unprocessable entity response has a 2xx status code
func (o *DeleteIterationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete iteration unprocessable entity response has a 3xx status code
func (o *DeleteIterationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete iteration unprocessable entity response has a 4xx status code
func (o *DeleteIterationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete iteration unprocessable entity response has a 5xx status code
func (o *DeleteIterationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete iteration unprocessable entity response a status code equal to that given
func (o *DeleteIterationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *DeleteIterationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationUnprocessableEntity ", 422)
}

func (o *DeleteIterationUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/iterations/{iteration-public-id}][%d] deleteIterationUnprocessableEntity ", 422)
}

func (o *DeleteIterationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
