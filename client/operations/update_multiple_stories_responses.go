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

// UpdateMultipleStoriesReader is a Reader for the UpdateMultipleStories structure.
type UpdateMultipleStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMultipleStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMultipleStoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMultipleStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMultipleStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMultipleStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateMultipleStoriesOK creates a UpdateMultipleStoriesOK with default headers values
func NewUpdateMultipleStoriesOK() *UpdateMultipleStoriesOK {
	return &UpdateMultipleStoriesOK{}
}

/*
UpdateMultipleStoriesOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateMultipleStoriesOK struct {
	Payload []*models.StorySlim
}

// IsSuccess returns true when this update multiple stories o k response has a 2xx status code
func (o *UpdateMultipleStoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update multiple stories o k response has a 3xx status code
func (o *UpdateMultipleStoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update multiple stories o k response has a 4xx status code
func (o *UpdateMultipleStoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update multiple stories o k response has a 5xx status code
func (o *UpdateMultipleStoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update multiple stories o k response a status code equal to that given
func (o *UpdateMultipleStoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateMultipleStoriesOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesOK  %+v", 200, o.Payload)
}

func (o *UpdateMultipleStoriesOK) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesOK  %+v", 200, o.Payload)
}

func (o *UpdateMultipleStoriesOK) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *UpdateMultipleStoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMultipleStoriesBadRequest creates a UpdateMultipleStoriesBadRequest with default headers values
func NewUpdateMultipleStoriesBadRequest() *UpdateMultipleStoriesBadRequest {
	return &UpdateMultipleStoriesBadRequest{}
}

/*
UpdateMultipleStoriesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateMultipleStoriesBadRequest struct {
}

// IsSuccess returns true when this update multiple stories bad request response has a 2xx status code
func (o *UpdateMultipleStoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update multiple stories bad request response has a 3xx status code
func (o *UpdateMultipleStoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update multiple stories bad request response has a 4xx status code
func (o *UpdateMultipleStoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update multiple stories bad request response has a 5xx status code
func (o *UpdateMultipleStoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update multiple stories bad request response a status code equal to that given
func (o *UpdateMultipleStoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateMultipleStoriesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesBadRequest ", 400)
}

func (o *UpdateMultipleStoriesBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesBadRequest ", 400)
}

func (o *UpdateMultipleStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMultipleStoriesNotFound creates a UpdateMultipleStoriesNotFound with default headers values
func NewUpdateMultipleStoriesNotFound() *UpdateMultipleStoriesNotFound {
	return &UpdateMultipleStoriesNotFound{}
}

/*
UpdateMultipleStoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateMultipleStoriesNotFound struct {
}

// IsSuccess returns true when this update multiple stories not found response has a 2xx status code
func (o *UpdateMultipleStoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update multiple stories not found response has a 3xx status code
func (o *UpdateMultipleStoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update multiple stories not found response has a 4xx status code
func (o *UpdateMultipleStoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update multiple stories not found response has a 5xx status code
func (o *UpdateMultipleStoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update multiple stories not found response a status code equal to that given
func (o *UpdateMultipleStoriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateMultipleStoriesNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesNotFound ", 404)
}

func (o *UpdateMultipleStoriesNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesNotFound ", 404)
}

func (o *UpdateMultipleStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMultipleStoriesUnprocessableEntity creates a UpdateMultipleStoriesUnprocessableEntity with default headers values
func NewUpdateMultipleStoriesUnprocessableEntity() *UpdateMultipleStoriesUnprocessableEntity {
	return &UpdateMultipleStoriesUnprocessableEntity{}
}

/*
UpdateMultipleStoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateMultipleStoriesUnprocessableEntity struct {
}

// IsSuccess returns true when this update multiple stories unprocessable entity response has a 2xx status code
func (o *UpdateMultipleStoriesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update multiple stories unprocessable entity response has a 3xx status code
func (o *UpdateMultipleStoriesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update multiple stories unprocessable entity response has a 4xx status code
func (o *UpdateMultipleStoriesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update multiple stories unprocessable entity response has a 5xx status code
func (o *UpdateMultipleStoriesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update multiple stories unprocessable entity response a status code equal to that given
func (o *UpdateMultipleStoriesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UpdateMultipleStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesUnprocessableEntity ", 422)
}

func (o *UpdateMultipleStoriesUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] updateMultipleStoriesUnprocessableEntity ", 422)
}

func (o *UpdateMultipleStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}