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

// ListEpicsReader is a Reader for the ListEpics structure.
type ListEpicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEpicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEpicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEpicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListEpicsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListEpicsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEpicsOK creates a ListEpicsOK with default headers values
func NewListEpicsOK() *ListEpicsOK {
	return &ListEpicsOK{}
}

/*
ListEpicsOK describes a response with status code 200, with default header values.

Resource
*/
type ListEpicsOK struct {
	Payload []*models.EpicSlim
}

// IsSuccess returns true when this list epics o k response has a 2xx status code
func (o *ListEpicsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list epics o k response has a 3xx status code
func (o *ListEpicsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list epics o k response has a 4xx status code
func (o *ListEpicsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list epics o k response has a 5xx status code
func (o *ListEpicsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list epics o k response a status code equal to that given
func (o *ListEpicsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEpicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsOK  %+v", 200, o.Payload)
}

func (o *ListEpicsOK) String() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsOK  %+v", 200, o.Payload)
}

func (o *ListEpicsOK) GetPayload() []*models.EpicSlim {
	return o.Payload
}

func (o *ListEpicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEpicsBadRequest creates a ListEpicsBadRequest with default headers values
func NewListEpicsBadRequest() *ListEpicsBadRequest {
	return &ListEpicsBadRequest{}
}

/*
ListEpicsBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListEpicsBadRequest struct {
}

// IsSuccess returns true when this list epics bad request response has a 2xx status code
func (o *ListEpicsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list epics bad request response has a 3xx status code
func (o *ListEpicsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list epics bad request response has a 4xx status code
func (o *ListEpicsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list epics bad request response has a 5xx status code
func (o *ListEpicsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list epics bad request response a status code equal to that given
func (o *ListEpicsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListEpicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsBadRequest ", 400)
}

func (o *ListEpicsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsBadRequest ", 400)
}

func (o *ListEpicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEpicsNotFound creates a ListEpicsNotFound with default headers values
func NewListEpicsNotFound() *ListEpicsNotFound {
	return &ListEpicsNotFound{}
}

/*
ListEpicsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListEpicsNotFound struct {
}

// IsSuccess returns true when this list epics not found response has a 2xx status code
func (o *ListEpicsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list epics not found response has a 3xx status code
func (o *ListEpicsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list epics not found response has a 4xx status code
func (o *ListEpicsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list epics not found response has a 5xx status code
func (o *ListEpicsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list epics not found response a status code equal to that given
func (o *ListEpicsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListEpicsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsNotFound ", 404)
}

func (o *ListEpicsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsNotFound ", 404)
}

func (o *ListEpicsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEpicsUnprocessableEntity creates a ListEpicsUnprocessableEntity with default headers values
func NewListEpicsUnprocessableEntity() *ListEpicsUnprocessableEntity {
	return &ListEpicsUnprocessableEntity{}
}

/*
ListEpicsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListEpicsUnprocessableEntity struct {
}

// IsSuccess returns true when this list epics unprocessable entity response has a 2xx status code
func (o *ListEpicsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list epics unprocessable entity response has a 3xx status code
func (o *ListEpicsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list epics unprocessable entity response has a 4xx status code
func (o *ListEpicsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list epics unprocessable entity response has a 5xx status code
func (o *ListEpicsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list epics unprocessable entity response a status code equal to that given
func (o *ListEpicsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *ListEpicsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsUnprocessableEntity ", 422)
}

func (o *ListEpicsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] listEpicsUnprocessableEntity ", 422)
}

func (o *ListEpicsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
