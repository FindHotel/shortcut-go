// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteFileReader is a Reader for the DeleteFile structure.
type DeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v3/files/{file-public-id}] deleteFile", response, response.Code())
	}
}

// NewDeleteFileNoContent creates a DeleteFileNoContent with default headers values
func NewDeleteFileNoContent() *DeleteFileNoContent {
	return &DeleteFileNoContent{}
}

/*
DeleteFileNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteFileNoContent struct {
}

// IsSuccess returns true when this delete file no content response has a 2xx status code
func (o *DeleteFileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete file no content response has a 3xx status code
func (o *DeleteFileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file no content response has a 4xx status code
func (o *DeleteFileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete file no content response has a 5xx status code
func (o *DeleteFileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file no content response a status code equal to that given
func (o *DeleteFileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete file no content response
func (o *DeleteFileNoContent) Code() int {
	return 204
}

func (o *DeleteFileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileNoContent", 204)
}

func (o *DeleteFileNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileNoContent", 204)
}

func (o *DeleteFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFileBadRequest creates a DeleteFileBadRequest with default headers values
func NewDeleteFileBadRequest() *DeleteFileBadRequest {
	return &DeleteFileBadRequest{}
}

/*
DeleteFileBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteFileBadRequest struct {
}

// IsSuccess returns true when this delete file bad request response has a 2xx status code
func (o *DeleteFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file bad request response has a 3xx status code
func (o *DeleteFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file bad request response has a 4xx status code
func (o *DeleteFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete file bad request response has a 5xx status code
func (o *DeleteFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file bad request response a status code equal to that given
func (o *DeleteFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete file bad request response
func (o *DeleteFileBadRequest) Code() int {
	return 400
}

func (o *DeleteFileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileBadRequest", 400)
}

func (o *DeleteFileBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileBadRequest", 400)
}

func (o *DeleteFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFileNotFound creates a DeleteFileNotFound with default headers values
func NewDeleteFileNotFound() *DeleteFileNotFound {
	return &DeleteFileNotFound{}
}

/*
DeleteFileNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteFileNotFound struct {
}

// IsSuccess returns true when this delete file not found response has a 2xx status code
func (o *DeleteFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file not found response has a 3xx status code
func (o *DeleteFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file not found response has a 4xx status code
func (o *DeleteFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete file not found response has a 5xx status code
func (o *DeleteFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file not found response a status code equal to that given
func (o *DeleteFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete file not found response
func (o *DeleteFileNotFound) Code() int {
	return 404
}

func (o *DeleteFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileNotFound", 404)
}

func (o *DeleteFileNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileNotFound", 404)
}

func (o *DeleteFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFileUnprocessableEntity creates a DeleteFileUnprocessableEntity with default headers values
func NewDeleteFileUnprocessableEntity() *DeleteFileUnprocessableEntity {
	return &DeleteFileUnprocessableEntity{}
}

/*
DeleteFileUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteFileUnprocessableEntity struct {
}

// IsSuccess returns true when this delete file unprocessable entity response has a 2xx status code
func (o *DeleteFileUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file unprocessable entity response has a 3xx status code
func (o *DeleteFileUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file unprocessable entity response has a 4xx status code
func (o *DeleteFileUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete file unprocessable entity response has a 5xx status code
func (o *DeleteFileUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file unprocessable entity response a status code equal to that given
func (o *DeleteFileUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete file unprocessable entity response
func (o *DeleteFileUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileUnprocessableEntity", 422)
}

func (o *DeleteFileUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/files/{file-public-id}][%d] deleteFileUnprocessableEntity", 422)
}

func (o *DeleteFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
