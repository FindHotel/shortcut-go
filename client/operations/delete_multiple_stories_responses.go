// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMultipleStoriesReader is a Reader for the DeleteMultipleStories structure.
type DeleteMultipleStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMultipleStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMultipleStoriesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMultipleStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMultipleStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteMultipleStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v3/stories/bulk] deleteMultipleStories", response, response.Code())
	}
}

// NewDeleteMultipleStoriesNoContent creates a DeleteMultipleStoriesNoContent with default headers values
func NewDeleteMultipleStoriesNoContent() *DeleteMultipleStoriesNoContent {
	return &DeleteMultipleStoriesNoContent{}
}

/*
DeleteMultipleStoriesNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteMultipleStoriesNoContent struct {
}

// IsSuccess returns true when this delete multiple stories no content response has a 2xx status code
func (o *DeleteMultipleStoriesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete multiple stories no content response has a 3xx status code
func (o *DeleteMultipleStoriesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete multiple stories no content response has a 4xx status code
func (o *DeleteMultipleStoriesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete multiple stories no content response has a 5xx status code
func (o *DeleteMultipleStoriesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete multiple stories no content response a status code equal to that given
func (o *DeleteMultipleStoriesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete multiple stories no content response
func (o *DeleteMultipleStoriesNoContent) Code() int {
	return 204
}

func (o *DeleteMultipleStoriesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesNoContent", 204)
}

func (o *DeleteMultipleStoriesNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesNoContent", 204)
}

func (o *DeleteMultipleStoriesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMultipleStoriesBadRequest creates a DeleteMultipleStoriesBadRequest with default headers values
func NewDeleteMultipleStoriesBadRequest() *DeleteMultipleStoriesBadRequest {
	return &DeleteMultipleStoriesBadRequest{}
}

/*
DeleteMultipleStoriesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteMultipleStoriesBadRequest struct {
}

// IsSuccess returns true when this delete multiple stories bad request response has a 2xx status code
func (o *DeleteMultipleStoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete multiple stories bad request response has a 3xx status code
func (o *DeleteMultipleStoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete multiple stories bad request response has a 4xx status code
func (o *DeleteMultipleStoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete multiple stories bad request response has a 5xx status code
func (o *DeleteMultipleStoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete multiple stories bad request response a status code equal to that given
func (o *DeleteMultipleStoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete multiple stories bad request response
func (o *DeleteMultipleStoriesBadRequest) Code() int {
	return 400
}

func (o *DeleteMultipleStoriesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesBadRequest", 400)
}

func (o *DeleteMultipleStoriesBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesBadRequest", 400)
}

func (o *DeleteMultipleStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMultipleStoriesNotFound creates a DeleteMultipleStoriesNotFound with default headers values
func NewDeleteMultipleStoriesNotFound() *DeleteMultipleStoriesNotFound {
	return &DeleteMultipleStoriesNotFound{}
}

/*
DeleteMultipleStoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteMultipleStoriesNotFound struct {
}

// IsSuccess returns true when this delete multiple stories not found response has a 2xx status code
func (o *DeleteMultipleStoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete multiple stories not found response has a 3xx status code
func (o *DeleteMultipleStoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete multiple stories not found response has a 4xx status code
func (o *DeleteMultipleStoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete multiple stories not found response has a 5xx status code
func (o *DeleteMultipleStoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete multiple stories not found response a status code equal to that given
func (o *DeleteMultipleStoriesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete multiple stories not found response
func (o *DeleteMultipleStoriesNotFound) Code() int {
	return 404
}

func (o *DeleteMultipleStoriesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesNotFound", 404)
}

func (o *DeleteMultipleStoriesNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesNotFound", 404)
}

func (o *DeleteMultipleStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMultipleStoriesUnprocessableEntity creates a DeleteMultipleStoriesUnprocessableEntity with default headers values
func NewDeleteMultipleStoriesUnprocessableEntity() *DeleteMultipleStoriesUnprocessableEntity {
	return &DeleteMultipleStoriesUnprocessableEntity{}
}

/*
DeleteMultipleStoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteMultipleStoriesUnprocessableEntity struct {
}

// IsSuccess returns true when this delete multiple stories unprocessable entity response has a 2xx status code
func (o *DeleteMultipleStoriesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete multiple stories unprocessable entity response has a 3xx status code
func (o *DeleteMultipleStoriesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete multiple stories unprocessable entity response has a 4xx status code
func (o *DeleteMultipleStoriesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete multiple stories unprocessable entity response has a 5xx status code
func (o *DeleteMultipleStoriesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete multiple stories unprocessable entity response a status code equal to that given
func (o *DeleteMultipleStoriesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete multiple stories unprocessable entity response
func (o *DeleteMultipleStoriesUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteMultipleStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesUnprocessableEntity", 422)
}

func (o *DeleteMultipleStoriesUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/bulk][%d] deleteMultipleStoriesUnprocessableEntity", 422)
}

func (o *DeleteMultipleStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
