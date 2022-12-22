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

// UpdateStoryReader is a Reader for the UpdateStory structure.
type UpdateStoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateStoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateStoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateStoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateStoryOK creates a UpdateStoryOK with default headers values
func NewUpdateStoryOK() *UpdateStoryOK {
	return &UpdateStoryOK{}
}

/*
UpdateStoryOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateStoryOK struct {
	Payload *models.Story
}

// IsSuccess returns true when this update story o k response has a 2xx status code
func (o *UpdateStoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update story o k response has a 3xx status code
func (o *UpdateStoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story o k response has a 4xx status code
func (o *UpdateStoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update story o k response has a 5xx status code
func (o *UpdateStoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update story o k response a status code equal to that given
func (o *UpdateStoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateStoryOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryOK  %+v", 200, o.Payload)
}

func (o *UpdateStoryOK) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryOK  %+v", 200, o.Payload)
}

func (o *UpdateStoryOK) GetPayload() *models.Story {
	return o.Payload
}

func (o *UpdateStoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Story)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStoryBadRequest creates a UpdateStoryBadRequest with default headers values
func NewUpdateStoryBadRequest() *UpdateStoryBadRequest {
	return &UpdateStoryBadRequest{}
}

/*
UpdateStoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateStoryBadRequest struct {
}

// IsSuccess returns true when this update story bad request response has a 2xx status code
func (o *UpdateStoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story bad request response has a 3xx status code
func (o *UpdateStoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story bad request response has a 4xx status code
func (o *UpdateStoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story bad request response has a 5xx status code
func (o *UpdateStoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update story bad request response a status code equal to that given
func (o *UpdateStoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateStoryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryBadRequest ", 400)
}

func (o *UpdateStoryBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryBadRequest ", 400)
}

func (o *UpdateStoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoryNotFound creates a UpdateStoryNotFound with default headers values
func NewUpdateStoryNotFound() *UpdateStoryNotFound {
	return &UpdateStoryNotFound{}
}

/*
UpdateStoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateStoryNotFound struct {
}

// IsSuccess returns true when this update story not found response has a 2xx status code
func (o *UpdateStoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story not found response has a 3xx status code
func (o *UpdateStoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story not found response has a 4xx status code
func (o *UpdateStoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story not found response has a 5xx status code
func (o *UpdateStoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update story not found response a status code equal to that given
func (o *UpdateStoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateStoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryNotFound ", 404)
}

func (o *UpdateStoryNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryNotFound ", 404)
}

func (o *UpdateStoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoryUnprocessableEntity creates a UpdateStoryUnprocessableEntity with default headers values
func NewUpdateStoryUnprocessableEntity() *UpdateStoryUnprocessableEntity {
	return &UpdateStoryUnprocessableEntity{}
}

/*
UpdateStoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateStoryUnprocessableEntity struct {
}

// IsSuccess returns true when this update story unprocessable entity response has a 2xx status code
func (o *UpdateStoryUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story unprocessable entity response has a 3xx status code
func (o *UpdateStoryUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story unprocessable entity response has a 4xx status code
func (o *UpdateStoryUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story unprocessable entity response has a 5xx status code
func (o *UpdateStoryUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update story unprocessable entity response a status code equal to that given
func (o *UpdateStoryUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UpdateStoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryUnprocessableEntity ", 422)
}

func (o *UpdateStoryUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] updateStoryUnprocessableEntity ", 422)
}

func (o *UpdateStoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}