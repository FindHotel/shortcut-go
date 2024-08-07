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

// UpdateTaskReader is a Reader for the UpdateTask structure.
type UpdateTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}] updateTask", response, response.Code())
	}
}

// NewUpdateTaskOK creates a UpdateTaskOK with default headers values
func NewUpdateTaskOK() *UpdateTaskOK {
	return &UpdateTaskOK{}
}

/*
UpdateTaskOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this update task o k response has a 2xx status code
func (o *UpdateTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update task o k response has a 3xx status code
func (o *UpdateTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task o k response has a 4xx status code
func (o *UpdateTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update task o k response has a 5xx status code
func (o *UpdateTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update task o k response a status code equal to that given
func (o *UpdateTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update task o k response
func (o *UpdateTaskOK) Code() int {
	return 200
}

func (o *UpdateTaskOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskOK %s", 200, payload)
}

func (o *UpdateTaskOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskOK %s", 200, payload)
}

func (o *UpdateTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTaskBadRequest creates a UpdateTaskBadRequest with default headers values
func NewUpdateTaskBadRequest() *UpdateTaskBadRequest {
	return &UpdateTaskBadRequest{}
}

/*
UpdateTaskBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateTaskBadRequest struct {
}

// IsSuccess returns true when this update task bad request response has a 2xx status code
func (o *UpdateTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task bad request response has a 3xx status code
func (o *UpdateTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task bad request response has a 4xx status code
func (o *UpdateTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update task bad request response has a 5xx status code
func (o *UpdateTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update task bad request response a status code equal to that given
func (o *UpdateTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update task bad request response
func (o *UpdateTaskBadRequest) Code() int {
	return 400
}

func (o *UpdateTaskBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskBadRequest", 400)
}

func (o *UpdateTaskBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskBadRequest", 400)
}

func (o *UpdateTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTaskNotFound creates a UpdateTaskNotFound with default headers values
func NewUpdateTaskNotFound() *UpdateTaskNotFound {
	return &UpdateTaskNotFound{}
}

/*
UpdateTaskNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateTaskNotFound struct {
}

// IsSuccess returns true when this update task not found response has a 2xx status code
func (o *UpdateTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task not found response has a 3xx status code
func (o *UpdateTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task not found response has a 4xx status code
func (o *UpdateTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update task not found response has a 5xx status code
func (o *UpdateTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update task not found response a status code equal to that given
func (o *UpdateTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update task not found response
func (o *UpdateTaskNotFound) Code() int {
	return 404
}

func (o *UpdateTaskNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskNotFound", 404)
}

func (o *UpdateTaskNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskNotFound", 404)
}

func (o *UpdateTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTaskUnprocessableEntity creates a UpdateTaskUnprocessableEntity with default headers values
func NewUpdateTaskUnprocessableEntity() *UpdateTaskUnprocessableEntity {
	return &UpdateTaskUnprocessableEntity{}
}

/*
UpdateTaskUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateTaskUnprocessableEntity struct {
}

// IsSuccess returns true when this update task unprocessable entity response has a 2xx status code
func (o *UpdateTaskUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task unprocessable entity response has a 3xx status code
func (o *UpdateTaskUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task unprocessable entity response has a 4xx status code
func (o *UpdateTaskUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update task unprocessable entity response has a 5xx status code
func (o *UpdateTaskUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update task unprocessable entity response a status code equal to that given
func (o *UpdateTaskUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update task unprocessable entity response
func (o *UpdateTaskUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskUnprocessableEntity", 422)
}

func (o *UpdateTaskUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/tasks/{task-public-id}][%d] updateTaskUnprocessableEntity", 422)
}

func (o *UpdateTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
