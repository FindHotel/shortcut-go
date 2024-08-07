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

// CreateCategoryReader is a Reader for the CreateCategory structure.
type CreateCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCategoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/categories] createCategory", response, response.Code())
	}
}

// NewCreateCategoryCreated creates a CreateCategoryCreated with default headers values
func NewCreateCategoryCreated() *CreateCategoryCreated {
	return &CreateCategoryCreated{}
}

/*
CreateCategoryCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateCategoryCreated struct {
	Payload *models.Category
}

// IsSuccess returns true when this create category created response has a 2xx status code
func (o *CreateCategoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create category created response has a 3xx status code
func (o *CreateCategoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create category created response has a 4xx status code
func (o *CreateCategoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create category created response has a 5xx status code
func (o *CreateCategoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create category created response a status code equal to that given
func (o *CreateCategoryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create category created response
func (o *CreateCategoryCreated) Code() int {
	return 201
}

func (o *CreateCategoryCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryCreated %s", 201, payload)
}

func (o *CreateCategoryCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryCreated %s", 201, payload)
}

func (o *CreateCategoryCreated) GetPayload() *models.Category {
	return o.Payload
}

func (o *CreateCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Category)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCategoryBadRequest creates a CreateCategoryBadRequest with default headers values
func NewCreateCategoryBadRequest() *CreateCategoryBadRequest {
	return &CreateCategoryBadRequest{}
}

/*
CreateCategoryBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateCategoryBadRequest struct {
}

// IsSuccess returns true when this create category bad request response has a 2xx status code
func (o *CreateCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create category bad request response has a 3xx status code
func (o *CreateCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create category bad request response has a 4xx status code
func (o *CreateCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create category bad request response has a 5xx status code
func (o *CreateCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create category bad request response a status code equal to that given
func (o *CreateCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create category bad request response
func (o *CreateCategoryBadRequest) Code() int {
	return 400
}

func (o *CreateCategoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryBadRequest", 400)
}

func (o *CreateCategoryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryBadRequest", 400)
}

func (o *CreateCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCategoryNotFound creates a CreateCategoryNotFound with default headers values
func NewCreateCategoryNotFound() *CreateCategoryNotFound {
	return &CreateCategoryNotFound{}
}

/*
CreateCategoryNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateCategoryNotFound struct {
}

// IsSuccess returns true when this create category not found response has a 2xx status code
func (o *CreateCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create category not found response has a 3xx status code
func (o *CreateCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create category not found response has a 4xx status code
func (o *CreateCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create category not found response has a 5xx status code
func (o *CreateCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create category not found response a status code equal to that given
func (o *CreateCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create category not found response
func (o *CreateCategoryNotFound) Code() int {
	return 404
}

func (o *CreateCategoryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryNotFound", 404)
}

func (o *CreateCategoryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryNotFound", 404)
}

func (o *CreateCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCategoryUnprocessableEntity creates a CreateCategoryUnprocessableEntity with default headers values
func NewCreateCategoryUnprocessableEntity() *CreateCategoryUnprocessableEntity {
	return &CreateCategoryUnprocessableEntity{}
}

/*
CreateCategoryUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateCategoryUnprocessableEntity struct {
}

// IsSuccess returns true when this create category unprocessable entity response has a 2xx status code
func (o *CreateCategoryUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create category unprocessable entity response has a 3xx status code
func (o *CreateCategoryUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create category unprocessable entity response has a 4xx status code
func (o *CreateCategoryUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create category unprocessable entity response has a 5xx status code
func (o *CreateCategoryUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create category unprocessable entity response a status code equal to that given
func (o *CreateCategoryUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create category unprocessable entity response
func (o *CreateCategoryUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateCategoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryUnprocessableEntity", 422)
}

func (o *CreateCategoryUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/categories][%d] createCategoryUnprocessableEntity", 422)
}

func (o *CreateCategoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
