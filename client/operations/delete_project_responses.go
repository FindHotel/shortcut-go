// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v3/projects/{project-public-id}] deleteProject", response, response.Code())
	}
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*
DeleteProjectNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteProjectNoContent struct {
}

// IsSuccess returns true when this delete project no content response has a 2xx status code
func (o *DeleteProjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project no content response has a 3xx status code
func (o *DeleteProjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project no content response has a 4xx status code
func (o *DeleteProjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project no content response has a 5xx status code
func (o *DeleteProjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project no content response a status code equal to that given
func (o *DeleteProjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project no content response
func (o *DeleteProjectNoContent) Code() int {
	return 204
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectNoContent", 204)
}

func (o *DeleteProjectNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectNoContent", 204)
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectBadRequest creates a DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {
	return &DeleteProjectBadRequest{}
}

/*
DeleteProjectBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteProjectBadRequest struct {
}

// IsSuccess returns true when this delete project bad request response has a 2xx status code
func (o *DeleteProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project bad request response has a 3xx status code
func (o *DeleteProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project bad request response has a 4xx status code
func (o *DeleteProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project bad request response has a 5xx status code
func (o *DeleteProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project bad request response a status code equal to that given
func (o *DeleteProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete project bad request response
func (o *DeleteProjectBadRequest) Code() int {
	return 400
}

func (o *DeleteProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectBadRequest", 400)
}

func (o *DeleteProjectBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectBadRequest", 400)
}

func (o *DeleteProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*
DeleteProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteProjectNotFound struct {
}

// IsSuccess returns true when this delete project not found response has a 2xx status code
func (o *DeleteProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project not found response has a 3xx status code
func (o *DeleteProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project not found response has a 4xx status code
func (o *DeleteProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project not found response has a 5xx status code
func (o *DeleteProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project not found response a status code equal to that given
func (o *DeleteProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete project not found response
func (o *DeleteProjectNotFound) Code() int {
	return 404
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectNotFound", 404)
}

func (o *DeleteProjectNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectNotFound", 404)
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnprocessableEntity creates a DeleteProjectUnprocessableEntity with default headers values
func NewDeleteProjectUnprocessableEntity() *DeleteProjectUnprocessableEntity {
	return &DeleteProjectUnprocessableEntity{}
}

/*
DeleteProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteProjectUnprocessableEntity struct {
}

// IsSuccess returns true when this delete project unprocessable entity response has a 2xx status code
func (o *DeleteProjectUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete project unprocessable entity response has a 3xx status code
func (o *DeleteProjectUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project unprocessable entity response has a 4xx status code
func (o *DeleteProjectUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete project unprocessable entity response has a 5xx status code
func (o *DeleteProjectUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project unprocessable entity response a status code equal to that given
func (o *DeleteProjectUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete project unprocessable entity response
func (o *DeleteProjectUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectUnprocessableEntity", 422)
}

func (o *DeleteProjectUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/projects/{project-public-id}][%d] deleteProjectUnprocessableEntity", 422)
}

func (o *DeleteProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
