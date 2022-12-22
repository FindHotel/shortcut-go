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

// UpdateFileReader is a Reader for the UpdateFile structure.
type UpdateFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFileOK creates a UpdateFileOK with default headers values
func NewUpdateFileOK() *UpdateFileOK {
	return &UpdateFileOK{}
}

/*
UpdateFileOK describes a response with status code 200, with default header values.

Resource
*/
type UpdateFileOK struct {
	Payload *models.UploadedFile
}

// IsSuccess returns true when this update file o k response has a 2xx status code
func (o *UpdateFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update file o k response has a 3xx status code
func (o *UpdateFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update file o k response has a 4xx status code
func (o *UpdateFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update file o k response has a 5xx status code
func (o *UpdateFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update file o k response a status code equal to that given
func (o *UpdateFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateFileOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileOK  %+v", 200, o.Payload)
}

func (o *UpdateFileOK) String() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileOK  %+v", 200, o.Payload)
}

func (o *UpdateFileOK) GetPayload() *models.UploadedFile {
	return o.Payload
}

func (o *UpdateFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadedFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFileBadRequest creates a UpdateFileBadRequest with default headers values
func NewUpdateFileBadRequest() *UpdateFileBadRequest {
	return &UpdateFileBadRequest{}
}

/*
UpdateFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UpdateFileBadRequest struct {
}

// IsSuccess returns true when this update file bad request response has a 2xx status code
func (o *UpdateFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update file bad request response has a 3xx status code
func (o *UpdateFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update file bad request response has a 4xx status code
func (o *UpdateFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update file bad request response has a 5xx status code
func (o *UpdateFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update file bad request response a status code equal to that given
func (o *UpdateFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileBadRequest ", 400)
}

func (o *UpdateFileBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileBadRequest ", 400)
}

func (o *UpdateFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFileNotFound creates a UpdateFileNotFound with default headers values
func NewUpdateFileNotFound() *UpdateFileNotFound {
	return &UpdateFileNotFound{}
}

/*
UpdateFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UpdateFileNotFound struct {
}

// IsSuccess returns true when this update file not found response has a 2xx status code
func (o *UpdateFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update file not found response has a 3xx status code
func (o *UpdateFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update file not found response has a 4xx status code
func (o *UpdateFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update file not found response has a 5xx status code
func (o *UpdateFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update file not found response a status code equal to that given
func (o *UpdateFileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileNotFound ", 404)
}

func (o *UpdateFileNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileNotFound ", 404)
}

func (o *UpdateFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFileUnprocessableEntity creates a UpdateFileUnprocessableEntity with default headers values
func NewUpdateFileUnprocessableEntity() *UpdateFileUnprocessableEntity {
	return &UpdateFileUnprocessableEntity{}
}

/*
UpdateFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UpdateFileUnprocessableEntity struct {
}

// IsSuccess returns true when this update file unprocessable entity response has a 2xx status code
func (o *UpdateFileUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update file unprocessable entity response has a 3xx status code
func (o *UpdateFileUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update file unprocessable entity response has a 4xx status code
func (o *UpdateFileUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update file unprocessable entity response has a 5xx status code
func (o *UpdateFileUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update file unprocessable entity response a status code equal to that given
func (o *UpdateFileUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UpdateFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileUnprocessableEntity ", 422)
}

func (o *UpdateFileUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] updateFileUnprocessableEntity ", 422)
}

func (o *UpdateFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
