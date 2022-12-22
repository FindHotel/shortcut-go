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

// CreateEntityTemplateReader is a Reader for the CreateEntityTemplate structure.
type CreateEntityTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEntityTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEntityTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEntityTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEntityTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateEntityTemplateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEntityTemplateCreated creates a CreateEntityTemplateCreated with default headers values
func NewCreateEntityTemplateCreated() *CreateEntityTemplateCreated {
	return &CreateEntityTemplateCreated{}
}

/*
CreateEntityTemplateCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateEntityTemplateCreated struct {
	Payload *models.EntityTemplate
}

// IsSuccess returns true when this create entity template created response has a 2xx status code
func (o *CreateEntityTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create entity template created response has a 3xx status code
func (o *CreateEntityTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create entity template created response has a 4xx status code
func (o *CreateEntityTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create entity template created response has a 5xx status code
func (o *CreateEntityTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create entity template created response a status code equal to that given
func (o *CreateEntityTemplateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateEntityTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateEntityTemplateCreated) String() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateEntityTemplateCreated) GetPayload() *models.EntityTemplate {
	return o.Payload
}

func (o *CreateEntityTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEntityTemplateBadRequest creates a CreateEntityTemplateBadRequest with default headers values
func NewCreateEntityTemplateBadRequest() *CreateEntityTemplateBadRequest {
	return &CreateEntityTemplateBadRequest{}
}

/*
CreateEntityTemplateBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateEntityTemplateBadRequest struct {
}

// IsSuccess returns true when this create entity template bad request response has a 2xx status code
func (o *CreateEntityTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create entity template bad request response has a 3xx status code
func (o *CreateEntityTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create entity template bad request response has a 4xx status code
func (o *CreateEntityTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create entity template bad request response has a 5xx status code
func (o *CreateEntityTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create entity template bad request response a status code equal to that given
func (o *CreateEntityTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateEntityTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateBadRequest ", 400)
}

func (o *CreateEntityTemplateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateBadRequest ", 400)
}

func (o *CreateEntityTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEntityTemplateNotFound creates a CreateEntityTemplateNotFound with default headers values
func NewCreateEntityTemplateNotFound() *CreateEntityTemplateNotFound {
	return &CreateEntityTemplateNotFound{}
}

/*
CreateEntityTemplateNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateEntityTemplateNotFound struct {
}

// IsSuccess returns true when this create entity template not found response has a 2xx status code
func (o *CreateEntityTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create entity template not found response has a 3xx status code
func (o *CreateEntityTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create entity template not found response has a 4xx status code
func (o *CreateEntityTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create entity template not found response has a 5xx status code
func (o *CreateEntityTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create entity template not found response a status code equal to that given
func (o *CreateEntityTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateEntityTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateNotFound ", 404)
}

func (o *CreateEntityTemplateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateNotFound ", 404)
}

func (o *CreateEntityTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEntityTemplateUnprocessableEntity creates a CreateEntityTemplateUnprocessableEntity with default headers values
func NewCreateEntityTemplateUnprocessableEntity() *CreateEntityTemplateUnprocessableEntity {
	return &CreateEntityTemplateUnprocessableEntity{}
}

/*
CreateEntityTemplateUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateEntityTemplateUnprocessableEntity struct {
}

// IsSuccess returns true when this create entity template unprocessable entity response has a 2xx status code
func (o *CreateEntityTemplateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create entity template unprocessable entity response has a 3xx status code
func (o *CreateEntityTemplateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create entity template unprocessable entity response has a 4xx status code
func (o *CreateEntityTemplateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create entity template unprocessable entity response has a 5xx status code
func (o *CreateEntityTemplateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create entity template unprocessable entity response a status code equal to that given
func (o *CreateEntityTemplateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *CreateEntityTemplateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateUnprocessableEntity ", 422)
}

func (o *CreateEntityTemplateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/entity-templates][%d] createEntityTemplateUnprocessableEntity ", 422)
}

func (o *CreateEntityTemplateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}