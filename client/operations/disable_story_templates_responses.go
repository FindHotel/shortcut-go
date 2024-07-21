// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DisableStoryTemplatesReader is a Reader for the DisableStoryTemplates structure.
type DisableStoryTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableStoryTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisableStoryTemplatesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisableStoryTemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableStoryTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDisableStoryTemplatesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v3/entity-templates/disable] disableStoryTemplates", response, response.Code())
	}
}

// NewDisableStoryTemplatesNoContent creates a DisableStoryTemplatesNoContent with default headers values
func NewDisableStoryTemplatesNoContent() *DisableStoryTemplatesNoContent {
	return &DisableStoryTemplatesNoContent{}
}

/*
DisableStoryTemplatesNoContent describes a response with status code 204, with default header values.

No Content
*/
type DisableStoryTemplatesNoContent struct {
}

// IsSuccess returns true when this disable story templates no content response has a 2xx status code
func (o *DisableStoryTemplatesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable story templates no content response has a 3xx status code
func (o *DisableStoryTemplatesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable story templates no content response has a 4xx status code
func (o *DisableStoryTemplatesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable story templates no content response has a 5xx status code
func (o *DisableStoryTemplatesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this disable story templates no content response a status code equal to that given
func (o *DisableStoryTemplatesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the disable story templates no content response
func (o *DisableStoryTemplatesNoContent) Code() int {
	return 204
}

func (o *DisableStoryTemplatesNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesNoContent", 204)
}

func (o *DisableStoryTemplatesNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesNoContent", 204)
}

func (o *DisableStoryTemplatesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableStoryTemplatesBadRequest creates a DisableStoryTemplatesBadRequest with default headers values
func NewDisableStoryTemplatesBadRequest() *DisableStoryTemplatesBadRequest {
	return &DisableStoryTemplatesBadRequest{}
}

/*
DisableStoryTemplatesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DisableStoryTemplatesBadRequest struct {
}

// IsSuccess returns true when this disable story templates bad request response has a 2xx status code
func (o *DisableStoryTemplatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable story templates bad request response has a 3xx status code
func (o *DisableStoryTemplatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable story templates bad request response has a 4xx status code
func (o *DisableStoryTemplatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable story templates bad request response has a 5xx status code
func (o *DisableStoryTemplatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this disable story templates bad request response a status code equal to that given
func (o *DisableStoryTemplatesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the disable story templates bad request response
func (o *DisableStoryTemplatesBadRequest) Code() int {
	return 400
}

func (o *DisableStoryTemplatesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesBadRequest", 400)
}

func (o *DisableStoryTemplatesBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesBadRequest", 400)
}

func (o *DisableStoryTemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableStoryTemplatesNotFound creates a DisableStoryTemplatesNotFound with default headers values
func NewDisableStoryTemplatesNotFound() *DisableStoryTemplatesNotFound {
	return &DisableStoryTemplatesNotFound{}
}

/*
DisableStoryTemplatesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DisableStoryTemplatesNotFound struct {
}

// IsSuccess returns true when this disable story templates not found response has a 2xx status code
func (o *DisableStoryTemplatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable story templates not found response has a 3xx status code
func (o *DisableStoryTemplatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable story templates not found response has a 4xx status code
func (o *DisableStoryTemplatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable story templates not found response has a 5xx status code
func (o *DisableStoryTemplatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this disable story templates not found response a status code equal to that given
func (o *DisableStoryTemplatesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the disable story templates not found response
func (o *DisableStoryTemplatesNotFound) Code() int {
	return 404
}

func (o *DisableStoryTemplatesNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesNotFound", 404)
}

func (o *DisableStoryTemplatesNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesNotFound", 404)
}

func (o *DisableStoryTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableStoryTemplatesUnprocessableEntity creates a DisableStoryTemplatesUnprocessableEntity with default headers values
func NewDisableStoryTemplatesUnprocessableEntity() *DisableStoryTemplatesUnprocessableEntity {
	return &DisableStoryTemplatesUnprocessableEntity{}
}

/*
DisableStoryTemplatesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DisableStoryTemplatesUnprocessableEntity struct {
}

// IsSuccess returns true when this disable story templates unprocessable entity response has a 2xx status code
func (o *DisableStoryTemplatesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable story templates unprocessable entity response has a 3xx status code
func (o *DisableStoryTemplatesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable story templates unprocessable entity response has a 4xx status code
func (o *DisableStoryTemplatesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable story templates unprocessable entity response has a 5xx status code
func (o *DisableStoryTemplatesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this disable story templates unprocessable entity response a status code equal to that given
func (o *DisableStoryTemplatesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the disable story templates unprocessable entity response
func (o *DisableStoryTemplatesUnprocessableEntity) Code() int {
	return 422
}

func (o *DisableStoryTemplatesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesUnprocessableEntity", 422)
}

func (o *DisableStoryTemplatesUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] disableStoryTemplatesUnprocessableEntity", 422)
}

func (o *DisableStoryTemplatesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
