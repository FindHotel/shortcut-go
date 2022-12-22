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

// CreateLinkedFileReader is a Reader for the CreateLinkedFile structure.
type CreateLinkedFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLinkedFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLinkedFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLinkedFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLinkedFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateLinkedFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLinkedFileCreated creates a CreateLinkedFileCreated with default headers values
func NewCreateLinkedFileCreated() *CreateLinkedFileCreated {
	return &CreateLinkedFileCreated{}
}

/*
CreateLinkedFileCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateLinkedFileCreated struct {
	Payload *models.LinkedFile
}

// IsSuccess returns true when this create linked file created response has a 2xx status code
func (o *CreateLinkedFileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create linked file created response has a 3xx status code
func (o *CreateLinkedFileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create linked file created response has a 4xx status code
func (o *CreateLinkedFileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create linked file created response has a 5xx status code
func (o *CreateLinkedFileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create linked file created response a status code equal to that given
func (o *CreateLinkedFileCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateLinkedFileCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileCreated  %+v", 201, o.Payload)
}

func (o *CreateLinkedFileCreated) String() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileCreated  %+v", 201, o.Payload)
}

func (o *CreateLinkedFileCreated) GetPayload() *models.LinkedFile {
	return o.Payload
}

func (o *CreateLinkedFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LinkedFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLinkedFileBadRequest creates a CreateLinkedFileBadRequest with default headers values
func NewCreateLinkedFileBadRequest() *CreateLinkedFileBadRequest {
	return &CreateLinkedFileBadRequest{}
}

/*
CreateLinkedFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateLinkedFileBadRequest struct {
}

// IsSuccess returns true when this create linked file bad request response has a 2xx status code
func (o *CreateLinkedFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create linked file bad request response has a 3xx status code
func (o *CreateLinkedFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create linked file bad request response has a 4xx status code
func (o *CreateLinkedFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create linked file bad request response has a 5xx status code
func (o *CreateLinkedFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create linked file bad request response a status code equal to that given
func (o *CreateLinkedFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateLinkedFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileBadRequest ", 400)
}

func (o *CreateLinkedFileBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileBadRequest ", 400)
}

func (o *CreateLinkedFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLinkedFileNotFound creates a CreateLinkedFileNotFound with default headers values
func NewCreateLinkedFileNotFound() *CreateLinkedFileNotFound {
	return &CreateLinkedFileNotFound{}
}

/*
CreateLinkedFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateLinkedFileNotFound struct {
}

// IsSuccess returns true when this create linked file not found response has a 2xx status code
func (o *CreateLinkedFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create linked file not found response has a 3xx status code
func (o *CreateLinkedFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create linked file not found response has a 4xx status code
func (o *CreateLinkedFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create linked file not found response has a 5xx status code
func (o *CreateLinkedFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create linked file not found response a status code equal to that given
func (o *CreateLinkedFileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateLinkedFileNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileNotFound ", 404)
}

func (o *CreateLinkedFileNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileNotFound ", 404)
}

func (o *CreateLinkedFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLinkedFileUnprocessableEntity creates a CreateLinkedFileUnprocessableEntity with default headers values
func NewCreateLinkedFileUnprocessableEntity() *CreateLinkedFileUnprocessableEntity {
	return &CreateLinkedFileUnprocessableEntity{}
}

/*
CreateLinkedFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateLinkedFileUnprocessableEntity struct {
}

// IsSuccess returns true when this create linked file unprocessable entity response has a 2xx status code
func (o *CreateLinkedFileUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create linked file unprocessable entity response has a 3xx status code
func (o *CreateLinkedFileUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create linked file unprocessable entity response has a 4xx status code
func (o *CreateLinkedFileUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create linked file unprocessable entity response has a 5xx status code
func (o *CreateLinkedFileUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create linked file unprocessable entity response a status code equal to that given
func (o *CreateLinkedFileUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *CreateLinkedFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileUnprocessableEntity ", 422)
}

func (o *CreateLinkedFileUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/linked-files][%d] createLinkedFileUnprocessableEntity ", 422)
}

func (o *CreateLinkedFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
