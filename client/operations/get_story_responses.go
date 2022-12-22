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

// GetStoryReader is a Reader for the GetStory structure.
type GetStoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetStoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStoryOK creates a GetStoryOK with default headers values
func NewGetStoryOK() *GetStoryOK {
	return &GetStoryOK{}
}

/*
GetStoryOK describes a response with status code 200, with default header values.

Resource
*/
type GetStoryOK struct {
	Payload *models.Story
}

// IsSuccess returns true when this get story o k response has a 2xx status code
func (o *GetStoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get story o k response has a 3xx status code
func (o *GetStoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story o k response has a 4xx status code
func (o *GetStoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get story o k response has a 5xx status code
func (o *GetStoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get story o k response a status code equal to that given
func (o *GetStoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryOK  %+v", 200, o.Payload)
}

func (o *GetStoryOK) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryOK  %+v", 200, o.Payload)
}

func (o *GetStoryOK) GetPayload() *models.Story {
	return o.Payload
}

func (o *GetStoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Story)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoryBadRequest creates a GetStoryBadRequest with default headers values
func NewGetStoryBadRequest() *GetStoryBadRequest {
	return &GetStoryBadRequest{}
}

/*
GetStoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetStoryBadRequest struct {
}

// IsSuccess returns true when this get story bad request response has a 2xx status code
func (o *GetStoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story bad request response has a 3xx status code
func (o *GetStoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story bad request response has a 4xx status code
func (o *GetStoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story bad request response has a 5xx status code
func (o *GetStoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get story bad request response a status code equal to that given
func (o *GetStoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetStoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryBadRequest ", 400)
}

func (o *GetStoryBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryBadRequest ", 400)
}

func (o *GetStoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoryNotFound creates a GetStoryNotFound with default headers values
func NewGetStoryNotFound() *GetStoryNotFound {
	return &GetStoryNotFound{}
}

/*
GetStoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetStoryNotFound struct {
}

// IsSuccess returns true when this get story not found response has a 2xx status code
func (o *GetStoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story not found response has a 3xx status code
func (o *GetStoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story not found response has a 4xx status code
func (o *GetStoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story not found response has a 5xx status code
func (o *GetStoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get story not found response a status code equal to that given
func (o *GetStoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetStoryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryNotFound ", 404)
}

func (o *GetStoryNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryNotFound ", 404)
}

func (o *GetStoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStoryUnprocessableEntity creates a GetStoryUnprocessableEntity with default headers values
func NewGetStoryUnprocessableEntity() *GetStoryUnprocessableEntity {
	return &GetStoryUnprocessableEntity{}
}

/*
GetStoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetStoryUnprocessableEntity struct {
}

// IsSuccess returns true when this get story unprocessable entity response has a 2xx status code
func (o *GetStoryUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get story unprocessable entity response has a 3xx status code
func (o *GetStoryUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get story unprocessable entity response has a 4xx status code
func (o *GetStoryUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get story unprocessable entity response has a 5xx status code
func (o *GetStoryUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get story unprocessable entity response a status code equal to that given
func (o *GetStoryUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetStoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryUnprocessableEntity ", 422)
}

func (o *GetStoryUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}][%d] getStoryUnprocessableEntity ", 422)
}

func (o *GetStoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
