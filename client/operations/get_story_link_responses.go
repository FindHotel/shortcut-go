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

// GetStoryLinkReader is a Reader for the GetStoryLink structure.
type GetStoryLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoryLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoryLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStoryLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStoryLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetStoryLinkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/story-links/{story-link-public-id}] getStoryLink", response, response.Code())
	}
}

// NewGetStoryLinkOK creates a GetStoryLinkOK with default headers values
func NewGetStoryLinkOK() *GetStoryLinkOK {
	return &GetStoryLinkOK{}
}

/*
GetStoryLinkOK describes a response with status code 200, with default header values.

Resource
*/
type GetStoryLinkOK struct {
	Payload *models.StoryLink
}

// IsSuccess returns true when this get story link o k response has a 2xx status code
func (o *GetStoryLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get story link o k response has a 3xx status code
func (o *GetStoryLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story link o k response has a 4xx status code
func (o *GetStoryLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get story link o k response has a 5xx status code
func (o *GetStoryLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get story link o k response a status code equal to that given
func (o *GetStoryLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get story link o k response
func (o *GetStoryLinkOK) Code() int {
	return 200
}

func (o *GetStoryLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkOK %s", 200, payload)
}

func (o *GetStoryLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkOK %s", 200, payload)
}

func (o *GetStoryLinkOK) GetPayload() *models.StoryLink {
	return o.Payload
}

func (o *GetStoryLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoryLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoryLinkBadRequest creates a GetStoryLinkBadRequest with default headers values
func NewGetStoryLinkBadRequest() *GetStoryLinkBadRequest {
	return &GetStoryLinkBadRequest{}
}

/*
GetStoryLinkBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetStoryLinkBadRequest struct {
}

// IsSuccess returns true when this get story link bad request response has a 2xx status code
func (o *GetStoryLinkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story link bad request response has a 3xx status code
func (o *GetStoryLinkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story link bad request response has a 4xx status code
func (o *GetStoryLinkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story link bad request response has a 5xx status code
func (o *GetStoryLinkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get story link bad request response a status code equal to that given
func (o *GetStoryLinkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get story link bad request response
func (o *GetStoryLinkBadRequest) Code() int {
	return 400
}

func (o *GetStoryLinkBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkBadRequest", 400)
}

func (o *GetStoryLinkBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkBadRequest", 400)
}

func (o *GetStoryLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoryLinkNotFound creates a GetStoryLinkNotFound with default headers values
func NewGetStoryLinkNotFound() *GetStoryLinkNotFound {
	return &GetStoryLinkNotFound{}
}

/*
GetStoryLinkNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetStoryLinkNotFound struct {
}

// IsSuccess returns true when this get story link not found response has a 2xx status code
func (o *GetStoryLinkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story link not found response has a 3xx status code
func (o *GetStoryLinkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story link not found response has a 4xx status code
func (o *GetStoryLinkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story link not found response has a 5xx status code
func (o *GetStoryLinkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get story link not found response a status code equal to that given
func (o *GetStoryLinkNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get story link not found response
func (o *GetStoryLinkNotFound) Code() int {
	return 404
}

func (o *GetStoryLinkNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkNotFound", 404)
}

func (o *GetStoryLinkNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkNotFound", 404)
}

func (o *GetStoryLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoryLinkUnprocessableEntity creates a GetStoryLinkUnprocessableEntity with default headers values
func NewGetStoryLinkUnprocessableEntity() *GetStoryLinkUnprocessableEntity {
	return &GetStoryLinkUnprocessableEntity{}
}

/*
GetStoryLinkUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetStoryLinkUnprocessableEntity struct {
}

// IsSuccess returns true when this get story link unprocessable entity response has a 2xx status code
func (o *GetStoryLinkUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story link unprocessable entity response has a 3xx status code
func (o *GetStoryLinkUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story link unprocessable entity response has a 4xx status code
func (o *GetStoryLinkUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story link unprocessable entity response has a 5xx status code
func (o *GetStoryLinkUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get story link unprocessable entity response a status code equal to that given
func (o *GetStoryLinkUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get story link unprocessable entity response
func (o *GetStoryLinkUnprocessableEntity) Code() int {
	return 422
}

func (o *GetStoryLinkUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkUnprocessableEntity", 422)
}

func (o *GetStoryLinkUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/story-links/{story-link-public-id}][%d] getStoryLinkUnprocessableEntity", 422)
}

func (o *GetStoryLinkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
