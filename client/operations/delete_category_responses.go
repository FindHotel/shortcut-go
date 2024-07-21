// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCategoryReader is a Reader for the DeleteCategory structure.
type DeleteCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCategoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteCategoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v3/categories/{category-public-id}] deleteCategory", response, response.Code())
	}
}

// NewDeleteCategoryNoContent creates a DeleteCategoryNoContent with default headers values
func NewDeleteCategoryNoContent() *DeleteCategoryNoContent {
	return &DeleteCategoryNoContent{}
}

/*
DeleteCategoryNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteCategoryNoContent struct {
}

// IsSuccess returns true when this delete category no content response has a 2xx status code
func (o *DeleteCategoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete category no content response has a 3xx status code
func (o *DeleteCategoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category no content response has a 4xx status code
func (o *DeleteCategoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete category no content response has a 5xx status code
func (o *DeleteCategoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category no content response a status code equal to that given
func (o *DeleteCategoryNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete category no content response
func (o *DeleteCategoryNoContent) Code() int {
	return 204
}

func (o *DeleteCategoryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNoContent", 204)
}

func (o *DeleteCategoryNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNoContent", 204)
}

func (o *DeleteCategoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryBadRequest creates a DeleteCategoryBadRequest with default headers values
func NewDeleteCategoryBadRequest() *DeleteCategoryBadRequest {
	return &DeleteCategoryBadRequest{}
}

/*
DeleteCategoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteCategoryBadRequest struct {
}

// IsSuccess returns true when this delete category bad request response has a 2xx status code
func (o *DeleteCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete category bad request response has a 3xx status code
func (o *DeleteCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category bad request response has a 4xx status code
func (o *DeleteCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete category bad request response has a 5xx status code
func (o *DeleteCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category bad request response a status code equal to that given
func (o *DeleteCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete category bad request response
func (o *DeleteCategoryBadRequest) Code() int {
	return 400
}

func (o *DeleteCategoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryBadRequest", 400)
}

func (o *DeleteCategoryBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryBadRequest", 400)
}

func (o *DeleteCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryNotFound creates a DeleteCategoryNotFound with default headers values
func NewDeleteCategoryNotFound() *DeleteCategoryNotFound {
	return &DeleteCategoryNotFound{}
}

/*
DeleteCategoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteCategoryNotFound struct {
}

// IsSuccess returns true when this delete category not found response has a 2xx status code
func (o *DeleteCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete category not found response has a 3xx status code
func (o *DeleteCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category not found response has a 4xx status code
func (o *DeleteCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete category not found response has a 5xx status code
func (o *DeleteCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category not found response a status code equal to that given
func (o *DeleteCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete category not found response
func (o *DeleteCategoryNotFound) Code() int {
	return 404
}

func (o *DeleteCategoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNotFound", 404)
}

func (o *DeleteCategoryNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryNotFound", 404)
}

func (o *DeleteCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryUnprocessableEntity creates a DeleteCategoryUnprocessableEntity with default headers values
func NewDeleteCategoryUnprocessableEntity() *DeleteCategoryUnprocessableEntity {
	return &DeleteCategoryUnprocessableEntity{}
}

/*
DeleteCategoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteCategoryUnprocessableEntity struct {
}

// IsSuccess returns true when this delete category unprocessable entity response has a 2xx status code
func (o *DeleteCategoryUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete category unprocessable entity response has a 3xx status code
func (o *DeleteCategoryUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category unprocessable entity response has a 4xx status code
func (o *DeleteCategoryUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete category unprocessable entity response has a 5xx status code
func (o *DeleteCategoryUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category unprocessable entity response a status code equal to that given
func (o *DeleteCategoryUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete category unprocessable entity response
func (o *DeleteCategoryUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteCategoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryUnprocessableEntity", 422)
}

func (o *DeleteCategoryUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/categories/{category-public-id}][%d] deleteCategoryUnprocessableEntity", 422)
}

func (o *DeleteCategoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
