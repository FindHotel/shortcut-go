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

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/files/{file-public-id}] getFile", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

/*
GetFileOK describes a response with status code 200, with default header values.

Resource
*/
type GetFileOK struct {
	Payload *models.UploadedFile
}

// IsSuccess returns true when this get file o k response has a 2xx status code
func (o *GetFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file o k response has a 3xx status code
func (o *GetFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file o k response has a 4xx status code
func (o *GetFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file o k response has a 5xx status code
func (o *GetFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file o k response a status code equal to that given
func (o *GetFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file o k response
func (o *GetFileOK) Code() int {
	return 200
}

func (o *GetFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileOK %s", 200, payload)
}

func (o *GetFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileOK %s", 200, payload)
}

func (o *GetFileOK) GetPayload() *models.UploadedFile {
	return o.Payload
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadedFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileBadRequest creates a GetFileBadRequest with default headers values
func NewGetFileBadRequest() *GetFileBadRequest {
	return &GetFileBadRequest{}
}

/*
GetFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetFileBadRequest struct {
}

// IsSuccess returns true when this get file bad request response has a 2xx status code
func (o *GetFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file bad request response has a 3xx status code
func (o *GetFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file bad request response has a 4xx status code
func (o *GetFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file bad request response has a 5xx status code
func (o *GetFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get file bad request response a status code equal to that given
func (o *GetFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get file bad request response
func (o *GetFileBadRequest) Code() int {
	return 400
}

func (o *GetFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileBadRequest", 400)
}

func (o *GetFileBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileBadRequest", 400)
}

func (o *GetFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileNotFound creates a GetFileNotFound with default headers values
func NewGetFileNotFound() *GetFileNotFound {
	return &GetFileNotFound{}
}

/*
GetFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetFileNotFound struct {
}

// IsSuccess returns true when this get file not found response has a 2xx status code
func (o *GetFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file not found response has a 3xx status code
func (o *GetFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file not found response has a 4xx status code
func (o *GetFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file not found response has a 5xx status code
func (o *GetFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get file not found response a status code equal to that given
func (o *GetFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get file not found response
func (o *GetFileNotFound) Code() int {
	return 404
}

func (o *GetFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileNotFound", 404)
}

func (o *GetFileNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileNotFound", 404)
}

func (o *GetFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileUnprocessableEntity creates a GetFileUnprocessableEntity with default headers values
func NewGetFileUnprocessableEntity() *GetFileUnprocessableEntity {
	return &GetFileUnprocessableEntity{}
}

/*
GetFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetFileUnprocessableEntity struct {
}

// IsSuccess returns true when this get file unprocessable entity response has a 2xx status code
func (o *GetFileUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file unprocessable entity response has a 3xx status code
func (o *GetFileUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file unprocessable entity response has a 4xx status code
func (o *GetFileUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file unprocessable entity response has a 5xx status code
func (o *GetFileUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get file unprocessable entity response a status code equal to that given
func (o *GetFileUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get file unprocessable entity response
func (o *GetFileUnprocessableEntity) Code() int {
	return 422
}

func (o *GetFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileUnprocessableEntity", 422)
}

func (o *GetFileUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/files/{file-public-id}][%d] getFileUnprocessableEntity", 422)
}

func (o *GetFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
