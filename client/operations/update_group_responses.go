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

// UpdateGroupReader is a Reader for the UpdateGroup structure.
type UpdateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/groups/{group-public-id}] updateGroup", response, response.Code())
	}
}

// NewUpdateGroupOK creates a UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {
	return &UpdateGroupOK{}
}

/*
UpdateGroupOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateGroupOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this update group o k response has a 2xx status code
func (o *UpdateGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update group o k response has a 3xx status code
func (o *UpdateGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group o k response has a 4xx status code
func (o *UpdateGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update group o k response has a 5xx status code
func (o *UpdateGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update group o k response a status code equal to that given
func (o *UpdateGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update group o k response
func (o *UpdateGroupOK) Code() int {
	return 200
}

func (o *UpdateGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupOK %s", 200, payload)
}

func (o *UpdateGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupOK %s", 200, payload)
}

func (o *UpdateGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UpdateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupBadRequest creates a UpdateGroupBadRequest with default headers values
func NewUpdateGroupBadRequest() *UpdateGroupBadRequest {
	return &UpdateGroupBadRequest{}
}

/*
UpdateGroupBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateGroupBadRequest struct {
}

// IsSuccess returns true when this update group bad request response has a 2xx status code
func (o *UpdateGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group bad request response has a 3xx status code
func (o *UpdateGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group bad request response has a 4xx status code
func (o *UpdateGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group bad request response has a 5xx status code
func (o *UpdateGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update group bad request response a status code equal to that given
func (o *UpdateGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update group bad request response
func (o *UpdateGroupBadRequest) Code() int {
	return 400
}

func (o *UpdateGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupBadRequest", 400)
}

func (o *UpdateGroupBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupBadRequest", 400)
}

func (o *UpdateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupForbidden creates a UpdateGroupForbidden with default headers values
func NewUpdateGroupForbidden() *UpdateGroupForbidden {
	return &UpdateGroupForbidden{}
}

/*
UpdateGroupForbidden describes a response with status code 403, with default header values.

UpdateGroupForbidden update group forbidden
*/
type UpdateGroupForbidden struct {
	Payload *models.UnusableEntitlementError
}

// IsSuccess returns true when this update group forbidden response has a 2xx status code
func (o *UpdateGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group forbidden response has a 3xx status code
func (o *UpdateGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group forbidden response has a 4xx status code
func (o *UpdateGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group forbidden response has a 5xx status code
func (o *UpdateGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update group forbidden response a status code equal to that given
func (o *UpdateGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update group forbidden response
func (o *UpdateGroupForbidden) Code() int {
	return 403
}

func (o *UpdateGroupForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupForbidden %s", 403, payload)
}

func (o *UpdateGroupForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupForbidden %s", 403, payload)
}

func (o *UpdateGroupForbidden) GetPayload() *models.UnusableEntitlementError {
	return o.Payload
}

func (o *UpdateGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnusableEntitlementError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupNotFound creates a UpdateGroupNotFound with default headers values
func NewUpdateGroupNotFound() *UpdateGroupNotFound {
	return &UpdateGroupNotFound{}
}

/*
UpdateGroupNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateGroupNotFound struct {
}

// IsSuccess returns true when this update group not found response has a 2xx status code
func (o *UpdateGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group not found response has a 3xx status code
func (o *UpdateGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group not found response has a 4xx status code
func (o *UpdateGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group not found response has a 5xx status code
func (o *UpdateGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update group not found response a status code equal to that given
func (o *UpdateGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update group not found response
func (o *UpdateGroupNotFound) Code() int {
	return 404
}

func (o *UpdateGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupNotFound", 404)
}

func (o *UpdateGroupNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupNotFound", 404)
}

func (o *UpdateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupUnprocessableEntity creates a UpdateGroupUnprocessableEntity with default headers values
func NewUpdateGroupUnprocessableEntity() *UpdateGroupUnprocessableEntity {
	return &UpdateGroupUnprocessableEntity{}
}

/*
UpdateGroupUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateGroupUnprocessableEntity struct {
}

// IsSuccess returns true when this update group unprocessable entity response has a 2xx status code
func (o *UpdateGroupUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update group unprocessable entity response has a 3xx status code
func (o *UpdateGroupUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group unprocessable entity response has a 4xx status code
func (o *UpdateGroupUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update group unprocessable entity response has a 5xx status code
func (o *UpdateGroupUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update group unprocessable entity response a status code equal to that given
func (o *UpdateGroupUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update group unprocessable entity response
func (o *UpdateGroupUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupUnprocessableEntity", 422)
}

func (o *UpdateGroupUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/groups/{group-public-id}][%d] updateGroupUnprocessableEntity", 422)
}

func (o *UpdateGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
