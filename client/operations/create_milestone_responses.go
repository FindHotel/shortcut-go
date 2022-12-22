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

// CreateMilestoneReader is a Reader for the CreateMilestone structure.
type CreateMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMilestoneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMilestoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateMilestoneUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMilestoneCreated creates a CreateMilestoneCreated with default headers values
func NewCreateMilestoneCreated() *CreateMilestoneCreated {
	return &CreateMilestoneCreated{}
}

/*
CreateMilestoneCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateMilestoneCreated struct {
	Payload *models.Milestone
}

// IsSuccess returns true when this create milestone created response has a 2xx status code
func (o *CreateMilestoneCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create milestone created response has a 3xx status code
func (o *CreateMilestoneCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create milestone created response has a 4xx status code
func (o *CreateMilestoneCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create milestone created response has a 5xx status code
func (o *CreateMilestoneCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create milestone created response a status code equal to that given
func (o *CreateMilestoneCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateMilestoneCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneCreated  %+v", 201, o.Payload)
}

func (o *CreateMilestoneCreated) String() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneCreated  %+v", 201, o.Payload)
}

func (o *CreateMilestoneCreated) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *CreateMilestoneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMilestoneBadRequest creates a CreateMilestoneBadRequest with default headers values
func NewCreateMilestoneBadRequest() *CreateMilestoneBadRequest {
	return &CreateMilestoneBadRequest{}
}

/*
CreateMilestoneBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateMilestoneBadRequest struct {
}

// IsSuccess returns true when this create milestone bad request response has a 2xx status code
func (o *CreateMilestoneBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create milestone bad request response has a 3xx status code
func (o *CreateMilestoneBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create milestone bad request response has a 4xx status code
func (o *CreateMilestoneBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create milestone bad request response has a 5xx status code
func (o *CreateMilestoneBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create milestone bad request response a status code equal to that given
func (o *CreateMilestoneBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneBadRequest ", 400)
}

func (o *CreateMilestoneBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneBadRequest ", 400)
}

func (o *CreateMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMilestoneForbidden creates a CreateMilestoneForbidden with default headers values
func NewCreateMilestoneForbidden() *CreateMilestoneForbidden {
	return &CreateMilestoneForbidden{}
}

/*
CreateMilestoneForbidden describes a response with status code 403, with default header values.

CreateMilestoneForbidden create milestone forbidden
*/
type CreateMilestoneForbidden struct {
	Payload *models.UnusableEntitlementError
}

// IsSuccess returns true when this create milestone forbidden response has a 2xx status code
func (o *CreateMilestoneForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create milestone forbidden response has a 3xx status code
func (o *CreateMilestoneForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create milestone forbidden response has a 4xx status code
func (o *CreateMilestoneForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create milestone forbidden response has a 5xx status code
func (o *CreateMilestoneForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create milestone forbidden response a status code equal to that given
func (o *CreateMilestoneForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateMilestoneForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *CreateMilestoneForbidden) String() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *CreateMilestoneForbidden) GetPayload() *models.UnusableEntitlementError {
	return o.Payload
}

func (o *CreateMilestoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnusableEntitlementError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMilestoneNotFound creates a CreateMilestoneNotFound with default headers values
func NewCreateMilestoneNotFound() *CreateMilestoneNotFound {
	return &CreateMilestoneNotFound{}
}

/*
CreateMilestoneNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateMilestoneNotFound struct {
}

// IsSuccess returns true when this create milestone not found response has a 2xx status code
func (o *CreateMilestoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create milestone not found response has a 3xx status code
func (o *CreateMilestoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create milestone not found response has a 4xx status code
func (o *CreateMilestoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create milestone not found response has a 5xx status code
func (o *CreateMilestoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create milestone not found response a status code equal to that given
func (o *CreateMilestoneNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateMilestoneNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneNotFound ", 404)
}

func (o *CreateMilestoneNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneNotFound ", 404)
}

func (o *CreateMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMilestoneUnprocessableEntity creates a CreateMilestoneUnprocessableEntity with default headers values
func NewCreateMilestoneUnprocessableEntity() *CreateMilestoneUnprocessableEntity {
	return &CreateMilestoneUnprocessableEntity{}
}

/*
CreateMilestoneUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateMilestoneUnprocessableEntity struct {
}

// IsSuccess returns true when this create milestone unprocessable entity response has a 2xx status code
func (o *CreateMilestoneUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create milestone unprocessable entity response has a 3xx status code
func (o *CreateMilestoneUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create milestone unprocessable entity response has a 4xx status code
func (o *CreateMilestoneUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create milestone unprocessable entity response has a 5xx status code
func (o *CreateMilestoneUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create milestone unprocessable entity response a status code equal to that given
func (o *CreateMilestoneUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *CreateMilestoneUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneUnprocessableEntity ", 422)
}

func (o *CreateMilestoneUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] createMilestoneUnprocessableEntity ", 422)
}

func (o *CreateMilestoneUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
