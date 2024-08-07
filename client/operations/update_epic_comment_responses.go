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

// UpdateEpicCommentReader is a Reader for the UpdateEpicComment structure.
type UpdateEpicCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEpicCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEpicCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEpicCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEpicCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateEpicCommentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}] updateEpicComment", response, response.Code())
	}
}

// NewUpdateEpicCommentOK creates a UpdateEpicCommentOK with default headers values
func NewUpdateEpicCommentOK() *UpdateEpicCommentOK {
	return &UpdateEpicCommentOK{}
}

/*
UpdateEpicCommentOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateEpicCommentOK struct {
	Payload *models.ThreadedComment
}

// IsSuccess returns true when this update epic comment o k response has a 2xx status code
func (o *UpdateEpicCommentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update epic comment o k response has a 3xx status code
func (o *UpdateEpicCommentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update epic comment o k response has a 4xx status code
func (o *UpdateEpicCommentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update epic comment o k response has a 5xx status code
func (o *UpdateEpicCommentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update epic comment o k response a status code equal to that given
func (o *UpdateEpicCommentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update epic comment o k response
func (o *UpdateEpicCommentOK) Code() int {
	return 200
}

func (o *UpdateEpicCommentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentOK %s", 200, payload)
}

func (o *UpdateEpicCommentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentOK %s", 200, payload)
}

func (o *UpdateEpicCommentOK) GetPayload() *models.ThreadedComment {
	return o.Payload
}

func (o *UpdateEpicCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreadedComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEpicCommentBadRequest creates a UpdateEpicCommentBadRequest with default headers values
func NewUpdateEpicCommentBadRequest() *UpdateEpicCommentBadRequest {
	return &UpdateEpicCommentBadRequest{}
}

/*
UpdateEpicCommentBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateEpicCommentBadRequest struct {
}

// IsSuccess returns true when this update epic comment bad request response has a 2xx status code
func (o *UpdateEpicCommentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update epic comment bad request response has a 3xx status code
func (o *UpdateEpicCommentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update epic comment bad request response has a 4xx status code
func (o *UpdateEpicCommentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update epic comment bad request response has a 5xx status code
func (o *UpdateEpicCommentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update epic comment bad request response a status code equal to that given
func (o *UpdateEpicCommentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update epic comment bad request response
func (o *UpdateEpicCommentBadRequest) Code() int {
	return 400
}

func (o *UpdateEpicCommentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentBadRequest", 400)
}

func (o *UpdateEpicCommentBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentBadRequest", 400)
}

func (o *UpdateEpicCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEpicCommentNotFound creates a UpdateEpicCommentNotFound with default headers values
func NewUpdateEpicCommentNotFound() *UpdateEpicCommentNotFound {
	return &UpdateEpicCommentNotFound{}
}

/*
UpdateEpicCommentNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateEpicCommentNotFound struct {
}

// IsSuccess returns true when this update epic comment not found response has a 2xx status code
func (o *UpdateEpicCommentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update epic comment not found response has a 3xx status code
func (o *UpdateEpicCommentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update epic comment not found response has a 4xx status code
func (o *UpdateEpicCommentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update epic comment not found response has a 5xx status code
func (o *UpdateEpicCommentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update epic comment not found response a status code equal to that given
func (o *UpdateEpicCommentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update epic comment not found response
func (o *UpdateEpicCommentNotFound) Code() int {
	return 404
}

func (o *UpdateEpicCommentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentNotFound", 404)
}

func (o *UpdateEpicCommentNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentNotFound", 404)
}

func (o *UpdateEpicCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEpicCommentUnprocessableEntity creates a UpdateEpicCommentUnprocessableEntity with default headers values
func NewUpdateEpicCommentUnprocessableEntity() *UpdateEpicCommentUnprocessableEntity {
	return &UpdateEpicCommentUnprocessableEntity{}
}

/*
UpdateEpicCommentUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateEpicCommentUnprocessableEntity struct {
}

// IsSuccess returns true when this update epic comment unprocessable entity response has a 2xx status code
func (o *UpdateEpicCommentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update epic comment unprocessable entity response has a 3xx status code
func (o *UpdateEpicCommentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update epic comment unprocessable entity response has a 4xx status code
func (o *UpdateEpicCommentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update epic comment unprocessable entity response has a 5xx status code
func (o *UpdateEpicCommentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update epic comment unprocessable entity response a status code equal to that given
func (o *UpdateEpicCommentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update epic comment unprocessable entity response
func (o *UpdateEpicCommentUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateEpicCommentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentUnprocessableEntity", 422)
}

func (o *UpdateEpicCommentUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] updateEpicCommentUnprocessableEntity", 422)
}

func (o *UpdateEpicCommentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
