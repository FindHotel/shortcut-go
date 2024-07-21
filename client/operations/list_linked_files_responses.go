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

// ListLinkedFilesReader is a Reader for the ListLinkedFiles structure.
type ListLinkedFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLinkedFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLinkedFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListLinkedFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListLinkedFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListLinkedFilesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/linked-files] listLinkedFiles", response, response.Code())
	}
}

// NewListLinkedFilesOK creates a ListLinkedFilesOK with default headers values
func NewListLinkedFilesOK() *ListLinkedFilesOK {
	return &ListLinkedFilesOK{}
}

/*
ListLinkedFilesOK describes a response with status code 200, with default header values.

Resource
*/
type ListLinkedFilesOK struct {
	Payload []*models.LinkedFile
}

// IsSuccess returns true when this list linked files o k response has a 2xx status code
func (o *ListLinkedFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list linked files o k response has a 3xx status code
func (o *ListLinkedFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list linked files o k response has a 4xx status code
func (o *ListLinkedFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list linked files o k response has a 5xx status code
func (o *ListLinkedFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list linked files o k response a status code equal to that given
func (o *ListLinkedFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list linked files o k response
func (o *ListLinkedFilesOK) Code() int {
	return 200
}

func (o *ListLinkedFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesOK %s", 200, payload)
}

func (o *ListLinkedFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesOK %s", 200, payload)
}

func (o *ListLinkedFilesOK) GetPayload() []*models.LinkedFile {
	return o.Payload
}

func (o *ListLinkedFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLinkedFilesBadRequest creates a ListLinkedFilesBadRequest with default headers values
func NewListLinkedFilesBadRequest() *ListLinkedFilesBadRequest {
	return &ListLinkedFilesBadRequest{}
}

/*
ListLinkedFilesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListLinkedFilesBadRequest struct {
}

// IsSuccess returns true when this list linked files bad request response has a 2xx status code
func (o *ListLinkedFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list linked files bad request response has a 3xx status code
func (o *ListLinkedFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list linked files bad request response has a 4xx status code
func (o *ListLinkedFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list linked files bad request response has a 5xx status code
func (o *ListLinkedFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list linked files bad request response a status code equal to that given
func (o *ListLinkedFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list linked files bad request response
func (o *ListLinkedFilesBadRequest) Code() int {
	return 400
}

func (o *ListLinkedFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesBadRequest", 400)
}

func (o *ListLinkedFilesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesBadRequest", 400)
}

func (o *ListLinkedFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLinkedFilesNotFound creates a ListLinkedFilesNotFound with default headers values
func NewListLinkedFilesNotFound() *ListLinkedFilesNotFound {
	return &ListLinkedFilesNotFound{}
}

/*
ListLinkedFilesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListLinkedFilesNotFound struct {
}

// IsSuccess returns true when this list linked files not found response has a 2xx status code
func (o *ListLinkedFilesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list linked files not found response has a 3xx status code
func (o *ListLinkedFilesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list linked files not found response has a 4xx status code
func (o *ListLinkedFilesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list linked files not found response has a 5xx status code
func (o *ListLinkedFilesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list linked files not found response a status code equal to that given
func (o *ListLinkedFilesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list linked files not found response
func (o *ListLinkedFilesNotFound) Code() int {
	return 404
}

func (o *ListLinkedFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesNotFound", 404)
}

func (o *ListLinkedFilesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesNotFound", 404)
}

func (o *ListLinkedFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLinkedFilesUnprocessableEntity creates a ListLinkedFilesUnprocessableEntity with default headers values
func NewListLinkedFilesUnprocessableEntity() *ListLinkedFilesUnprocessableEntity {
	return &ListLinkedFilesUnprocessableEntity{}
}

/*
ListLinkedFilesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListLinkedFilesUnprocessableEntity struct {
}

// IsSuccess returns true when this list linked files unprocessable entity response has a 2xx status code
func (o *ListLinkedFilesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list linked files unprocessable entity response has a 3xx status code
func (o *ListLinkedFilesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list linked files unprocessable entity response has a 4xx status code
func (o *ListLinkedFilesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list linked files unprocessable entity response has a 5xx status code
func (o *ListLinkedFilesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list linked files unprocessable entity response a status code equal to that given
func (o *ListLinkedFilesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list linked files unprocessable entity response
func (o *ListLinkedFilesUnprocessableEntity) Code() int {
	return 422
}

func (o *ListLinkedFilesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesUnprocessableEntity", 422)
}

func (o *ListLinkedFilesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/linked-files][%d] listLinkedFilesUnprocessableEntity", 422)
}

func (o *ListLinkedFilesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
