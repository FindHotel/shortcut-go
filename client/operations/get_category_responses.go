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

// GetCategoryReader is a Reader for the GetCategory structure.
type GetCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCategoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCategoryOK creates a GetCategoryOK with default headers values
func NewGetCategoryOK() *GetCategoryOK {
	return &GetCategoryOK{}
}

/*
GetCategoryOK describes a response with status code 200, with default header values.

Resource
*/
type GetCategoryOK struct {
	Payload *models.Category
}

// IsSuccess returns true when this get category o k response has a 2xx status code
func (o *GetCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get category o k response has a 3xx status code
func (o *GetCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get category o k response has a 4xx status code
func (o *GetCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get category o k response has a 5xx status code
func (o *GetCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get category o k response a status code equal to that given
func (o *GetCategoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCategoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryOK  %+v", 200, o.Payload)
}

func (o *GetCategoryOK) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryOK  %+v", 200, o.Payload)
}

func (o *GetCategoryOK) GetPayload() *models.Category {
	return o.Payload
}

func (o *GetCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Category)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoryBadRequest creates a GetCategoryBadRequest with default headers values
func NewGetCategoryBadRequest() *GetCategoryBadRequest {
	return &GetCategoryBadRequest{}
}

/*
GetCategoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetCategoryBadRequest struct {
}

// IsSuccess returns true when this get category bad request response has a 2xx status code
func (o *GetCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get category bad request response has a 3xx status code
func (o *GetCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get category bad request response has a 4xx status code
func (o *GetCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get category bad request response has a 5xx status code
func (o *GetCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get category bad request response a status code equal to that given
func (o *GetCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCategoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryBadRequest ", 400)
}

func (o *GetCategoryBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryBadRequest ", 400)
}

func (o *GetCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCategoryNotFound creates a GetCategoryNotFound with default headers values
func NewGetCategoryNotFound() *GetCategoryNotFound {
	return &GetCategoryNotFound{}
}

/*
GetCategoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetCategoryNotFound struct {
}

// IsSuccess returns true when this get category not found response has a 2xx status code
func (o *GetCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get category not found response has a 3xx status code
func (o *GetCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get category not found response has a 4xx status code
func (o *GetCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get category not found response has a 5xx status code
func (o *GetCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get category not found response a status code equal to that given
func (o *GetCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCategoryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryNotFound ", 404)
}

func (o *GetCategoryNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryNotFound ", 404)
}

func (o *GetCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCategoryUnprocessableEntity creates a GetCategoryUnprocessableEntity with default headers values
func NewGetCategoryUnprocessableEntity() *GetCategoryUnprocessableEntity {
	return &GetCategoryUnprocessableEntity{}
}

/*
GetCategoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetCategoryUnprocessableEntity struct {
}

// IsSuccess returns true when this get category unprocessable entity response has a 2xx status code
func (o *GetCategoryUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get category unprocessable entity response has a 3xx status code
func (o *GetCategoryUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get category unprocessable entity response has a 4xx status code
func (o *GetCategoryUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get category unprocessable entity response has a 5xx status code
func (o *GetCategoryUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get category unprocessable entity response a status code equal to that given
func (o *GetCategoryUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetCategoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryUnprocessableEntity ", 422)
}

func (o *GetCategoryUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}][%d] getCategoryUnprocessableEntity ", 422)
}

func (o *GetCategoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
