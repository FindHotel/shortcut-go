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

// CreateStoryFromTemplateReader is a Reader for the CreateStoryFromTemplate structure.
type CreateStoryFromTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStoryFromTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStoryFromTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStoryFromTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateStoryFromTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateStoryFromTemplateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/stories/from-template] createStoryFromTemplate", response, response.Code())
	}
}

// NewCreateStoryFromTemplateCreated creates a CreateStoryFromTemplateCreated with default headers values
func NewCreateStoryFromTemplateCreated() *CreateStoryFromTemplateCreated {
	return &CreateStoryFromTemplateCreated{}
}

/*
CreateStoryFromTemplateCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateStoryFromTemplateCreated struct {
	Payload *models.Story
}

// IsSuccess returns true when this create story from template created response has a 2xx status code
func (o *CreateStoryFromTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create story from template created response has a 3xx status code
func (o *CreateStoryFromTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create story from template created response has a 4xx status code
func (o *CreateStoryFromTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create story from template created response has a 5xx status code
func (o *CreateStoryFromTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create story from template created response a status code equal to that given
func (o *CreateStoryFromTemplateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create story from template created response
func (o *CreateStoryFromTemplateCreated) Code() int {
	return 201
}

func (o *CreateStoryFromTemplateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateCreated %s", 201, payload)
}

func (o *CreateStoryFromTemplateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateCreated %s", 201, payload)
}

func (o *CreateStoryFromTemplateCreated) GetPayload() *models.Story {
	return o.Payload
}

func (o *CreateStoryFromTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Story)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStoryFromTemplateBadRequest creates a CreateStoryFromTemplateBadRequest with default headers values
func NewCreateStoryFromTemplateBadRequest() *CreateStoryFromTemplateBadRequest {
	return &CreateStoryFromTemplateBadRequest{}
}

/*
CreateStoryFromTemplateBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateStoryFromTemplateBadRequest struct {
}

// IsSuccess returns true when this create story from template bad request response has a 2xx status code
func (o *CreateStoryFromTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create story from template bad request response has a 3xx status code
func (o *CreateStoryFromTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create story from template bad request response has a 4xx status code
func (o *CreateStoryFromTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create story from template bad request response has a 5xx status code
func (o *CreateStoryFromTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create story from template bad request response a status code equal to that given
func (o *CreateStoryFromTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create story from template bad request response
func (o *CreateStoryFromTemplateBadRequest) Code() int {
	return 400
}

func (o *CreateStoryFromTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateBadRequest", 400)
}

func (o *CreateStoryFromTemplateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateBadRequest", 400)
}

func (o *CreateStoryFromTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStoryFromTemplateNotFound creates a CreateStoryFromTemplateNotFound with default headers values
func NewCreateStoryFromTemplateNotFound() *CreateStoryFromTemplateNotFound {
	return &CreateStoryFromTemplateNotFound{}
}

/*
CreateStoryFromTemplateNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateStoryFromTemplateNotFound struct {
}

// IsSuccess returns true when this create story from template not found response has a 2xx status code
func (o *CreateStoryFromTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create story from template not found response has a 3xx status code
func (o *CreateStoryFromTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create story from template not found response has a 4xx status code
func (o *CreateStoryFromTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create story from template not found response has a 5xx status code
func (o *CreateStoryFromTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create story from template not found response a status code equal to that given
func (o *CreateStoryFromTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create story from template not found response
func (o *CreateStoryFromTemplateNotFound) Code() int {
	return 404
}

func (o *CreateStoryFromTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateNotFound", 404)
}

func (o *CreateStoryFromTemplateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateNotFound", 404)
}

func (o *CreateStoryFromTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStoryFromTemplateUnprocessableEntity creates a CreateStoryFromTemplateUnprocessableEntity with default headers values
func NewCreateStoryFromTemplateUnprocessableEntity() *CreateStoryFromTemplateUnprocessableEntity {
	return &CreateStoryFromTemplateUnprocessableEntity{}
}

/*
CreateStoryFromTemplateUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateStoryFromTemplateUnprocessableEntity struct {
}

// IsSuccess returns true when this create story from template unprocessable entity response has a 2xx status code
func (o *CreateStoryFromTemplateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create story from template unprocessable entity response has a 3xx status code
func (o *CreateStoryFromTemplateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create story from template unprocessable entity response has a 4xx status code
func (o *CreateStoryFromTemplateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create story from template unprocessable entity response has a 5xx status code
func (o *CreateStoryFromTemplateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create story from template unprocessable entity response a status code equal to that given
func (o *CreateStoryFromTemplateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create story from template unprocessable entity response
func (o *CreateStoryFromTemplateUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateStoryFromTemplateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateUnprocessableEntity", 422)
}

func (o *CreateStoryFromTemplateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/from-template][%d] createStoryFromTemplateUnprocessableEntity", 422)
}

func (o *CreateStoryFromTemplateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
