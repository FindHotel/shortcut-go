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

// CreateObjectiveReader is a Reader for the CreateObjective structure.
type CreateObjectiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateObjectiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateObjectiveCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateObjectiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateObjectiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateObjectiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateObjectiveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/objectives] createObjective", response, response.Code())
	}
}

// NewCreateObjectiveCreated creates a CreateObjectiveCreated with default headers values
func NewCreateObjectiveCreated() *CreateObjectiveCreated {
	return &CreateObjectiveCreated{}
}

/*
CreateObjectiveCreated describes a response with status code 201, with default header values.

Resource
*/
type CreateObjectiveCreated struct {
	Payload *models.Objective
}

// IsSuccess returns true when this create objective created response has a 2xx status code
func (o *CreateObjectiveCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create objective created response has a 3xx status code
func (o *CreateObjectiveCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create objective created response has a 4xx status code
func (o *CreateObjectiveCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create objective created response has a 5xx status code
func (o *CreateObjectiveCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create objective created response a status code equal to that given
func (o *CreateObjectiveCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create objective created response
func (o *CreateObjectiveCreated) Code() int {
	return 201
}

func (o *CreateObjectiveCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveCreated %s", 201, payload)
}

func (o *CreateObjectiveCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveCreated %s", 201, payload)
}

func (o *CreateObjectiveCreated) GetPayload() *models.Objective {
	return o.Payload
}

func (o *CreateObjectiveCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Objective)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateObjectiveBadRequest creates a CreateObjectiveBadRequest with default headers values
func NewCreateObjectiveBadRequest() *CreateObjectiveBadRequest {
	return &CreateObjectiveBadRequest{}
}

/*
CreateObjectiveBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type CreateObjectiveBadRequest struct {
}

// IsSuccess returns true when this create objective bad request response has a 2xx status code
func (o *CreateObjectiveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create objective bad request response has a 3xx status code
func (o *CreateObjectiveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create objective bad request response has a 4xx status code
func (o *CreateObjectiveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create objective bad request response has a 5xx status code
func (o *CreateObjectiveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create objective bad request response a status code equal to that given
func (o *CreateObjectiveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create objective bad request response
func (o *CreateObjectiveBadRequest) Code() int {
	return 400
}

func (o *CreateObjectiveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveBadRequest", 400)
}

func (o *CreateObjectiveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveBadRequest", 400)
}

func (o *CreateObjectiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateObjectiveForbidden creates a CreateObjectiveForbidden with default headers values
func NewCreateObjectiveForbidden() *CreateObjectiveForbidden {
	return &CreateObjectiveForbidden{}
}

/*
CreateObjectiveForbidden describes a response with status code 403, with default header values.

CreateObjectiveForbidden create objective forbidden
*/
type CreateObjectiveForbidden struct {
	Payload *models.UnusableEntitlementError
}

// IsSuccess returns true when this create objective forbidden response has a 2xx status code
func (o *CreateObjectiveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create objective forbidden response has a 3xx status code
func (o *CreateObjectiveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create objective forbidden response has a 4xx status code
func (o *CreateObjectiveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create objective forbidden response has a 5xx status code
func (o *CreateObjectiveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create objective forbidden response a status code equal to that given
func (o *CreateObjectiveForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create objective forbidden response
func (o *CreateObjectiveForbidden) Code() int {
	return 403
}

func (o *CreateObjectiveForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveForbidden %s", 403, payload)
}

func (o *CreateObjectiveForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveForbidden %s", 403, payload)
}

func (o *CreateObjectiveForbidden) GetPayload() *models.UnusableEntitlementError {
	return o.Payload
}

func (o *CreateObjectiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnusableEntitlementError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateObjectiveNotFound creates a CreateObjectiveNotFound with default headers values
func NewCreateObjectiveNotFound() *CreateObjectiveNotFound {
	return &CreateObjectiveNotFound{}
}

/*
CreateObjectiveNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type CreateObjectiveNotFound struct {
}

// IsSuccess returns true when this create objective not found response has a 2xx status code
func (o *CreateObjectiveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create objective not found response has a 3xx status code
func (o *CreateObjectiveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create objective not found response has a 4xx status code
func (o *CreateObjectiveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create objective not found response has a 5xx status code
func (o *CreateObjectiveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create objective not found response a status code equal to that given
func (o *CreateObjectiveNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create objective not found response
func (o *CreateObjectiveNotFound) Code() int {
	return 404
}

func (o *CreateObjectiveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveNotFound", 404)
}

func (o *CreateObjectiveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveNotFound", 404)
}

func (o *CreateObjectiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateObjectiveUnprocessableEntity creates a CreateObjectiveUnprocessableEntity with default headers values
func NewCreateObjectiveUnprocessableEntity() *CreateObjectiveUnprocessableEntity {
	return &CreateObjectiveUnprocessableEntity{}
}

/*
CreateObjectiveUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type CreateObjectiveUnprocessableEntity struct {
}

// IsSuccess returns true when this create objective unprocessable entity response has a 2xx status code
func (o *CreateObjectiveUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create objective unprocessable entity response has a 3xx status code
func (o *CreateObjectiveUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create objective unprocessable entity response has a 4xx status code
func (o *CreateObjectiveUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create objective unprocessable entity response has a 5xx status code
func (o *CreateObjectiveUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create objective unprocessable entity response a status code equal to that given
func (o *CreateObjectiveUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create objective unprocessable entity response
func (o *CreateObjectiveUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateObjectiveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveUnprocessableEntity", 422)
}

func (o *CreateObjectiveUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/objectives][%d] createObjectiveUnprocessableEntity", 422)
}

func (o *CreateObjectiveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
