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

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}] getTask", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*
GetTaskOK describes a response with status code 200, with default header values.

Resource
*/
type GetTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get task o k response has a 2xx status code
func (o *GetTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task o k response has a 3xx status code
func (o *GetTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task o k response has a 4xx status code
func (o *GetTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task o k response has a 5xx status code
func (o *GetTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task o k response a status code equal to that given
func (o *GetTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task o k response
func (o *GetTaskOK) Code() int {
	return 200
}

func (o *GetTaskOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskOK %s", 200, payload)
}

func (o *GetTaskOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskOK %s", 200, payload)
}

func (o *GetTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskBadRequest creates a GetTaskBadRequest with default headers values
func NewGetTaskBadRequest() *GetTaskBadRequest {
	return &GetTaskBadRequest{}
}

/*
GetTaskBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetTaskBadRequest struct {
}

// IsSuccess returns true when this get task bad request response has a 2xx status code
func (o *GetTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task bad request response has a 3xx status code
func (o *GetTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task bad request response has a 4xx status code
func (o *GetTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task bad request response has a 5xx status code
func (o *GetTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get task bad request response a status code equal to that given
func (o *GetTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get task bad request response
func (o *GetTaskBadRequest) Code() int {
	return 400
}

func (o *GetTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskBadRequest", 400)
}

func (o *GetTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskBadRequest", 400)
}

func (o *GetTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*
GetTaskNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetTaskNotFound struct {
}

// IsSuccess returns true when this get task not found response has a 2xx status code
func (o *GetTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task not found response has a 3xx status code
func (o *GetTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task not found response has a 4xx status code
func (o *GetTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task not found response has a 5xx status code
func (o *GetTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task not found response a status code equal to that given
func (o *GetTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get task not found response
func (o *GetTaskNotFound) Code() int {
	return 404
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskNotFound", 404)
}

func (o *GetTaskNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskNotFound", 404)
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskUnprocessableEntity creates a GetTaskUnprocessableEntity with default headers values
func NewGetTaskUnprocessableEntity() *GetTaskUnprocessableEntity {
	return &GetTaskUnprocessableEntity{}
}

/*
GetTaskUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetTaskUnprocessableEntity struct {
}

// IsSuccess returns true when this get task unprocessable entity response has a 2xx status code
func (o *GetTaskUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task unprocessable entity response has a 3xx status code
func (o *GetTaskUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task unprocessable entity response has a 4xx status code
func (o *GetTaskUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task unprocessable entity response has a 5xx status code
func (o *GetTaskUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get task unprocessable entity response a status code equal to that given
func (o *GetTaskUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get task unprocessable entity response
func (o *GetTaskUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskUnprocessableEntity", 422)
}

func (o *GetTaskUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] getTaskUnprocessableEntity", 422)
}

func (o *GetTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
