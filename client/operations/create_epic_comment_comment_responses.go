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

// CreateEpicCommentCommentReader is a Reader for the CreateEpicCommentComment structure.
type CreateEpicCommentCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEpicCommentCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEpicCommentCommentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEpicCommentCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEpicCommentCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateEpicCommentCommentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}] createEpicCommentComment", response, response.Code())
	}
}

// NewCreateEpicCommentCommentCreated creates a CreateEpicCommentCommentCreated with default headers values
func NewCreateEpicCommentCommentCreated() *CreateEpicCommentCommentCreated {
	return &CreateEpicCommentCommentCreated{}
}

/*
CreateEpicCommentCommentCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateEpicCommentCommentCreated struct {
	Payload *models.ThreadedComment
}

// IsSuccess returns true when this create epic comment comment created response has a 2xx status code
func (o *CreateEpicCommentCommentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create epic comment comment created response has a 3xx status code
func (o *CreateEpicCommentCommentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create epic comment comment created response has a 4xx status code
func (o *CreateEpicCommentCommentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create epic comment comment created response has a 5xx status code
func (o *CreateEpicCommentCommentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create epic comment comment created response a status code equal to that given
func (o *CreateEpicCommentCommentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create epic comment comment created response
func (o *CreateEpicCommentCommentCreated) Code() int {
	return 201
}

func (o *CreateEpicCommentCommentCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentCreated %s", 201, payload)
}

func (o *CreateEpicCommentCommentCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentCreated %s", 201, payload)
}

func (o *CreateEpicCommentCommentCreated) GetPayload() *models.ThreadedComment {
	return o.Payload
}

func (o *CreateEpicCommentCommentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreadedComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEpicCommentCommentBadRequest creates a CreateEpicCommentCommentBadRequest with default headers values
func NewCreateEpicCommentCommentBadRequest() *CreateEpicCommentCommentBadRequest {
	return &CreateEpicCommentCommentBadRequest{}
}

/*
CreateEpicCommentCommentBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateEpicCommentCommentBadRequest struct {
}

// IsSuccess returns true when this create epic comment comment bad request response has a 2xx status code
func (o *CreateEpicCommentCommentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create epic comment comment bad request response has a 3xx status code
func (o *CreateEpicCommentCommentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create epic comment comment bad request response has a 4xx status code
func (o *CreateEpicCommentCommentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create epic comment comment bad request response has a 5xx status code
func (o *CreateEpicCommentCommentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create epic comment comment bad request response a status code equal to that given
func (o *CreateEpicCommentCommentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create epic comment comment bad request response
func (o *CreateEpicCommentCommentBadRequest) Code() int {
	return 400
}

func (o *CreateEpicCommentCommentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentBadRequest", 400)
}

func (o *CreateEpicCommentCommentBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentBadRequest", 400)
}

func (o *CreateEpicCommentCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEpicCommentCommentNotFound creates a CreateEpicCommentCommentNotFound with default headers values
func NewCreateEpicCommentCommentNotFound() *CreateEpicCommentCommentNotFound {
	return &CreateEpicCommentCommentNotFound{}
}

/*
CreateEpicCommentCommentNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateEpicCommentCommentNotFound struct {
}

// IsSuccess returns true when this create epic comment comment not found response has a 2xx status code
func (o *CreateEpicCommentCommentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create epic comment comment not found response has a 3xx status code
func (o *CreateEpicCommentCommentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create epic comment comment not found response has a 4xx status code
func (o *CreateEpicCommentCommentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create epic comment comment not found response has a 5xx status code
func (o *CreateEpicCommentCommentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create epic comment comment not found response a status code equal to that given
func (o *CreateEpicCommentCommentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create epic comment comment not found response
func (o *CreateEpicCommentCommentNotFound) Code() int {
	return 404
}

func (o *CreateEpicCommentCommentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentNotFound", 404)
}

func (o *CreateEpicCommentCommentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentNotFound", 404)
}

func (o *CreateEpicCommentCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEpicCommentCommentUnprocessableEntity creates a CreateEpicCommentCommentUnprocessableEntity with default headers values
func NewCreateEpicCommentCommentUnprocessableEntity() *CreateEpicCommentCommentUnprocessableEntity {
	return &CreateEpicCommentCommentUnprocessableEntity{}
}

/*
CreateEpicCommentCommentUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateEpicCommentCommentUnprocessableEntity struct {
}

// IsSuccess returns true when this create epic comment comment unprocessable entity response has a 2xx status code
func (o *CreateEpicCommentCommentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create epic comment comment unprocessable entity response has a 3xx status code
func (o *CreateEpicCommentCommentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create epic comment comment unprocessable entity response has a 4xx status code
func (o *CreateEpicCommentCommentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create epic comment comment unprocessable entity response has a 5xx status code
func (o *CreateEpicCommentCommentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create epic comment comment unprocessable entity response a status code equal to that given
func (o *CreateEpicCommentCommentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create epic comment comment unprocessable entity response
func (o *CreateEpicCommentCommentUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateEpicCommentCommentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentUnprocessableEntity", 422)
}

func (o *CreateEpicCommentCommentUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] createEpicCommentCommentUnprocessableEntity", 422)
}

func (o *CreateEpicCommentCommentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
