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

// GetGroupReader is a Reader for the GetGroup structure.
type GetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/groups/{group-public-id}] getGroup", response, response.Code())
	}
}

// NewGetGroupOK creates a GetGroupOK with default headers values
func NewGetGroupOK() *GetGroupOK {
	return &GetGroupOK{}
}

/*
GetGroupOK describes a response with status code 200, with default header values.

Resource
*/
type GetGroupOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this get group o k response has a 2xx status code
func (o *GetGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get group o k response has a 3xx status code
func (o *GetGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group o k response has a 4xx status code
func (o *GetGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group o k response has a 5xx status code
func (o *GetGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get group o k response a status code equal to that given
func (o *GetGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get group o k response
func (o *GetGroupOK) Code() int {
	return 200
}

func (o *GetGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupOK %s", 200, payload)
}

func (o *GetGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupOK %s", 200, payload)
}

func (o *GetGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBadRequest creates a GetGroupBadRequest with default headers values
func NewGetGroupBadRequest() *GetGroupBadRequest {
	return &GetGroupBadRequest{}
}

/*
GetGroupBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetGroupBadRequest struct {
}

// IsSuccess returns true when this get group bad request response has a 2xx status code
func (o *GetGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group bad request response has a 3xx status code
func (o *GetGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group bad request response has a 4xx status code
func (o *GetGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group bad request response has a 5xx status code
func (o *GetGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get group bad request response a status code equal to that given
func (o *GetGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get group bad request response
func (o *GetGroupBadRequest) Code() int {
	return 400
}

func (o *GetGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupBadRequest", 400)
}

func (o *GetGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupBadRequest", 400)
}

func (o *GetGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupNotFound creates a GetGroupNotFound with default headers values
func NewGetGroupNotFound() *GetGroupNotFound {
	return &GetGroupNotFound{}
}

/*
GetGroupNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetGroupNotFound struct {
}

// IsSuccess returns true when this get group not found response has a 2xx status code
func (o *GetGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group not found response has a 3xx status code
func (o *GetGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group not found response has a 4xx status code
func (o *GetGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group not found response has a 5xx status code
func (o *GetGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get group not found response a status code equal to that given
func (o *GetGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get group not found response
func (o *GetGroupNotFound) Code() int {
	return 404
}

func (o *GetGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupNotFound", 404)
}

func (o *GetGroupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupNotFound", 404)
}

func (o *GetGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupUnprocessableEntity creates a GetGroupUnprocessableEntity with default headers values
func NewGetGroupUnprocessableEntity() *GetGroupUnprocessableEntity {
	return &GetGroupUnprocessableEntity{}
}

/*
GetGroupUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetGroupUnprocessableEntity struct {
}

// IsSuccess returns true when this get group unprocessable entity response has a 2xx status code
func (o *GetGroupUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group unprocessable entity response has a 3xx status code
func (o *GetGroupUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group unprocessable entity response has a 4xx status code
func (o *GetGroupUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group unprocessable entity response has a 5xx status code
func (o *GetGroupUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get group unprocessable entity response a status code equal to that given
func (o *GetGroupUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get group unprocessable entity response
func (o *GetGroupUnprocessableEntity) Code() int {
	return 422
}

func (o *GetGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupUnprocessableEntity", 422)
}

func (o *GetGroupUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/groups/{group-public-id}][%d] getGroupUnprocessableEntity", 422)
}

func (o *GetGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
