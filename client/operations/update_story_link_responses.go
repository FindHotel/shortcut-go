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

// UpdateStoryLinkReader is a Reader for the UpdateStoryLink structure.
type UpdateStoryLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStoryLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStoryLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateStoryLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateStoryLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateStoryLinkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/story-links/{story-link-public-id}] updateStoryLink", response, response.Code())
	}
}

// NewUpdateStoryLinkOK creates a UpdateStoryLinkOK with default headers values
func NewUpdateStoryLinkOK() *UpdateStoryLinkOK {
	return &UpdateStoryLinkOK{}
}

/*
UpdateStoryLinkOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateStoryLinkOK struct {
	Payload *models.StoryLink
}

// IsSuccess returns true when this update story link o k response has a 2xx status code
func (o *UpdateStoryLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update story link o k response has a 3xx status code
func (o *UpdateStoryLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story link o k response has a 4xx status code
func (o *UpdateStoryLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update story link o k response has a 5xx status code
func (o *UpdateStoryLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update story link o k response a status code equal to that given
func (o *UpdateStoryLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update story link o k response
func (o *UpdateStoryLinkOK) Code() int {
	return 200
}

func (o *UpdateStoryLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkOK %s", 200, payload)
}

func (o *UpdateStoryLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkOK %s", 200, payload)
}

func (o *UpdateStoryLinkOK) GetPayload() *models.StoryLink {
	return o.Payload
}

func (o *UpdateStoryLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoryLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStoryLinkBadRequest creates a UpdateStoryLinkBadRequest with default headers values
func NewUpdateStoryLinkBadRequest() *UpdateStoryLinkBadRequest {
	return &UpdateStoryLinkBadRequest{}
}

/*
UpdateStoryLinkBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateStoryLinkBadRequest struct {
}

// IsSuccess returns true when this update story link bad request response has a 2xx status code
func (o *UpdateStoryLinkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story link bad request response has a 3xx status code
func (o *UpdateStoryLinkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story link bad request response has a 4xx status code
func (o *UpdateStoryLinkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story link bad request response has a 5xx status code
func (o *UpdateStoryLinkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update story link bad request response a status code equal to that given
func (o *UpdateStoryLinkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update story link bad request response
func (o *UpdateStoryLinkBadRequest) Code() int {
	return 400
}

func (o *UpdateStoryLinkBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkBadRequest", 400)
}

func (o *UpdateStoryLinkBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkBadRequest", 400)
}

func (o *UpdateStoryLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoryLinkNotFound creates a UpdateStoryLinkNotFound with default headers values
func NewUpdateStoryLinkNotFound() *UpdateStoryLinkNotFound {
	return &UpdateStoryLinkNotFound{}
}

/*
UpdateStoryLinkNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateStoryLinkNotFound struct {
}

// IsSuccess returns true when this update story link not found response has a 2xx status code
func (o *UpdateStoryLinkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story link not found response has a 3xx status code
func (o *UpdateStoryLinkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story link not found response has a 4xx status code
func (o *UpdateStoryLinkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story link not found response has a 5xx status code
func (o *UpdateStoryLinkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update story link not found response a status code equal to that given
func (o *UpdateStoryLinkNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update story link not found response
func (o *UpdateStoryLinkNotFound) Code() int {
	return 404
}

func (o *UpdateStoryLinkNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkNotFound", 404)
}

func (o *UpdateStoryLinkNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkNotFound", 404)
}

func (o *UpdateStoryLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoryLinkUnprocessableEntity creates a UpdateStoryLinkUnprocessableEntity with default headers values
func NewUpdateStoryLinkUnprocessableEntity() *UpdateStoryLinkUnprocessableEntity {
	return &UpdateStoryLinkUnprocessableEntity{}
}

/*
UpdateStoryLinkUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateStoryLinkUnprocessableEntity struct {
}

// IsSuccess returns true when this update story link unprocessable entity response has a 2xx status code
func (o *UpdateStoryLinkUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update story link unprocessable entity response has a 3xx status code
func (o *UpdateStoryLinkUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update story link unprocessable entity response has a 4xx status code
func (o *UpdateStoryLinkUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update story link unprocessable entity response has a 5xx status code
func (o *UpdateStoryLinkUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update story link unprocessable entity response a status code equal to that given
func (o *UpdateStoryLinkUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update story link unprocessable entity response
func (o *UpdateStoryLinkUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateStoryLinkUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkUnprocessableEntity", 422)
}

func (o *UpdateStoryLinkUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] updateStoryLinkUnprocessableEntity", 422)
}

func (o *UpdateStoryLinkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
