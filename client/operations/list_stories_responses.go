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

// ListStoriesReader is a Reader for the ListStories structure.
type ListStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListStoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/projects/{project-public-id}/stories] listStories", response, response.Code())
	}
}

// NewListStoriesOK creates a ListStoriesOK with default headers values
func NewListStoriesOK() *ListStoriesOK {
	return &ListStoriesOK{}
}

/*
ListStoriesOK describes a response with status code 200, with default header values.

Resource
*/
type ListStoriesOK struct {
	Payload []*models.StorySlim
}

// IsSuccess returns true when this list stories o k response has a 2xx status code
func (o *ListStoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list stories o k response has a 3xx status code
func (o *ListStoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list stories o k response has a 4xx status code
func (o *ListStoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list stories o k response has a 5xx status code
func (o *ListStoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list stories o k response a status code equal to that given
func (o *ListStoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list stories o k response
func (o *ListStoriesOK) Code() int {
	return 200
}

func (o *ListStoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesOK %s", 200, payload)
}

func (o *ListStoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesOK %s", 200, payload)
}

func (o *ListStoriesOK) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *ListStoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListStoriesBadRequest creates a ListStoriesBadRequest with default headers values
func NewListStoriesBadRequest() *ListStoriesBadRequest {
	return &ListStoriesBadRequest{}
}

/*
ListStoriesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListStoriesBadRequest struct {
}

// IsSuccess returns true when this list stories bad request response has a 2xx status code
func (o *ListStoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list stories bad request response has a 3xx status code
func (o *ListStoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list stories bad request response has a 4xx status code
func (o *ListStoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list stories bad request response has a 5xx status code
func (o *ListStoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list stories bad request response a status code equal to that given
func (o *ListStoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list stories bad request response
func (o *ListStoriesBadRequest) Code() int {
	return 400
}

func (o *ListStoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesBadRequest", 400)
}

func (o *ListStoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesBadRequest", 400)
}

func (o *ListStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStoriesNotFound creates a ListStoriesNotFound with default headers values
func NewListStoriesNotFound() *ListStoriesNotFound {
	return &ListStoriesNotFound{}
}

/*
ListStoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListStoriesNotFound struct {
}

// IsSuccess returns true when this list stories not found response has a 2xx status code
func (o *ListStoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list stories not found response has a 3xx status code
func (o *ListStoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list stories not found response has a 4xx status code
func (o *ListStoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list stories not found response has a 5xx status code
func (o *ListStoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list stories not found response a status code equal to that given
func (o *ListStoriesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list stories not found response
func (o *ListStoriesNotFound) Code() int {
	return 404
}

func (o *ListStoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesNotFound", 404)
}

func (o *ListStoriesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesNotFound", 404)
}

func (o *ListStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListStoriesUnprocessableEntity creates a ListStoriesUnprocessableEntity with default headers values
func NewListStoriesUnprocessableEntity() *ListStoriesUnprocessableEntity {
	return &ListStoriesUnprocessableEntity{}
}

/*
ListStoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListStoriesUnprocessableEntity struct {
}

// IsSuccess returns true when this list stories unprocessable entity response has a 2xx status code
func (o *ListStoriesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list stories unprocessable entity response has a 3xx status code
func (o *ListStoriesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list stories unprocessable entity response has a 4xx status code
func (o *ListStoriesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list stories unprocessable entity response has a 5xx status code
func (o *ListStoriesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list stories unprocessable entity response a status code equal to that given
func (o *ListStoriesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list stories unprocessable entity response
func (o *ListStoriesUnprocessableEntity) Code() int {
	return 422
}

func (o *ListStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesUnprocessableEntity", 422)
}

func (o *ListStoriesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/projects/{project-public-id}/stories][%d] listStoriesUnprocessableEntity", 422)
}

func (o *ListStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
