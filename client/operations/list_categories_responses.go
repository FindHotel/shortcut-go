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

// ListCategoriesReader is a Reader for the ListCategories structure.
type ListCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCategoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListCategoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCategoriesOK creates a ListCategoriesOK with default headers values
func NewListCategoriesOK() *ListCategoriesOK {
	return &ListCategoriesOK{}
}

/*
ListCategoriesOK describes a response with status code 200, with default header values.

Resource
*/
type ListCategoriesOK struct {
	Payload []*models.Category
}

// IsSuccess returns true when this list categories o k response has a 2xx status code
func (o *ListCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list categories o k response has a 3xx status code
func (o *ListCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories o k response has a 4xx status code
func (o *ListCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list categories o k response has a 5xx status code
func (o *ListCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories o k response a status code equal to that given
func (o *ListCategoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesOK) String() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesOK) GetPayload() []*models.Category {
	return o.Payload
}

func (o *ListCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCategoriesBadRequest creates a ListCategoriesBadRequest with default headers values
func NewListCategoriesBadRequest() *ListCategoriesBadRequest {
	return &ListCategoriesBadRequest{}
}

/*
ListCategoriesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListCategoriesBadRequest struct {
}

// IsSuccess returns true when this list categories bad request response has a 2xx status code
func (o *ListCategoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list categories bad request response has a 3xx status code
func (o *ListCategoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories bad request response has a 4xx status code
func (o *ListCategoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list categories bad request response has a 5xx status code
func (o *ListCategoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories bad request response a status code equal to that given
func (o *ListCategoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListCategoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesBadRequest ", 400)
}

func (o *ListCategoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesBadRequest ", 400)
}

func (o *ListCategoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoriesNotFound creates a ListCategoriesNotFound with default headers values
func NewListCategoriesNotFound() *ListCategoriesNotFound {
	return &ListCategoriesNotFound{}
}

/*
ListCategoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListCategoriesNotFound struct {
}

// IsSuccess returns true when this list categories not found response has a 2xx status code
func (o *ListCategoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list categories not found response has a 3xx status code
func (o *ListCategoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories not found response has a 4xx status code
func (o *ListCategoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list categories not found response has a 5xx status code
func (o *ListCategoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories not found response a status code equal to that given
func (o *ListCategoriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesNotFound ", 404)
}

func (o *ListCategoriesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesNotFound ", 404)
}

func (o *ListCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoriesUnprocessableEntity creates a ListCategoriesUnprocessableEntity with default headers values
func NewListCategoriesUnprocessableEntity() *ListCategoriesUnprocessableEntity {
	return &ListCategoriesUnprocessableEntity{}
}

/*
ListCategoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListCategoriesUnprocessableEntity struct {
}

// IsSuccess returns true when this list categories unprocessable entity response has a 2xx status code
func (o *ListCategoriesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list categories unprocessable entity response has a 3xx status code
func (o *ListCategoriesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories unprocessable entity response has a 4xx status code
func (o *ListCategoriesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list categories unprocessable entity response has a 5xx status code
func (o *ListCategoriesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories unprocessable entity response a status code equal to that given
func (o *ListCategoriesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *ListCategoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesUnprocessableEntity ", 422)
}

func (o *ListCategoriesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/categories][%d] listCategoriesUnprocessableEntity ", 422)
}

func (o *ListCategoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
