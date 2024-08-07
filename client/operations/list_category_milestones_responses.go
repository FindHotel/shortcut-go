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

// ListCategoryMilestonesReader is a Reader for the ListCategoryMilestones structure.
type ListCategoryMilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoryMilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoryMilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCategoryMilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCategoryMilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListCategoryMilestonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/categories/{category-public-id}/milestones] listCategoryMilestones", response, response.Code())
	}
}

// NewListCategoryMilestonesOK creates a ListCategoryMilestonesOK with default headers values
func NewListCategoryMilestonesOK() *ListCategoryMilestonesOK {
	return &ListCategoryMilestonesOK{}
}

/*
ListCategoryMilestonesOK describes a response with status code 200, with default header values.

Resource
*/
type ListCategoryMilestonesOK struct {
	Payload []*models.Milestone
}

// IsSuccess returns true when this list category milestones o k response has a 2xx status code
func (o *ListCategoryMilestonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list category milestones o k response has a 3xx status code
func (o *ListCategoryMilestonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list category milestones o k response has a 4xx status code
func (o *ListCategoryMilestonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list category milestones o k response has a 5xx status code
func (o *ListCategoryMilestonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list category milestones o k response a status code equal to that given
func (o *ListCategoryMilestonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list category milestones o k response
func (o *ListCategoryMilestonesOK) Code() int {
	return 200
}

func (o *ListCategoryMilestonesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesOK %s", 200, payload)
}

func (o *ListCategoryMilestonesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesOK %s", 200, payload)
}

func (o *ListCategoryMilestonesOK) GetPayload() []*models.Milestone {
	return o.Payload
}

func (o *ListCategoryMilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCategoryMilestonesBadRequest creates a ListCategoryMilestonesBadRequest with default headers values
func NewListCategoryMilestonesBadRequest() *ListCategoryMilestonesBadRequest {
	return &ListCategoryMilestonesBadRequest{}
}

/*
ListCategoryMilestonesBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type ListCategoryMilestonesBadRequest struct {
}

// IsSuccess returns true when this list category milestones bad request response has a 2xx status code
func (o *ListCategoryMilestonesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list category milestones bad request response has a 3xx status code
func (o *ListCategoryMilestonesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list category milestones bad request response has a 4xx status code
func (o *ListCategoryMilestonesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list category milestones bad request response has a 5xx status code
func (o *ListCategoryMilestonesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list category milestones bad request response a status code equal to that given
func (o *ListCategoryMilestonesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list category milestones bad request response
func (o *ListCategoryMilestonesBadRequest) Code() int {
	return 400
}

func (o *ListCategoryMilestonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesBadRequest", 400)
}

func (o *ListCategoryMilestonesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesBadRequest", 400)
}

func (o *ListCategoryMilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoryMilestonesNotFound creates a ListCategoryMilestonesNotFound with default headers values
func NewListCategoryMilestonesNotFound() *ListCategoryMilestonesNotFound {
	return &ListCategoryMilestonesNotFound{}
}

/*
ListCategoryMilestonesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type ListCategoryMilestonesNotFound struct {
}

// IsSuccess returns true when this list category milestones not found response has a 2xx status code
func (o *ListCategoryMilestonesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list category milestones not found response has a 3xx status code
func (o *ListCategoryMilestonesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list category milestones not found response has a 4xx status code
func (o *ListCategoryMilestonesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list category milestones not found response has a 5xx status code
func (o *ListCategoryMilestonesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list category milestones not found response a status code equal to that given
func (o *ListCategoryMilestonesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list category milestones not found response
func (o *ListCategoryMilestonesNotFound) Code() int {
	return 404
}

func (o *ListCategoryMilestonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesNotFound", 404)
}

func (o *ListCategoryMilestonesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesNotFound", 404)
}

func (o *ListCategoryMilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCategoryMilestonesUnprocessableEntity creates a ListCategoryMilestonesUnprocessableEntity with default headers values
func NewListCategoryMilestonesUnprocessableEntity() *ListCategoryMilestonesUnprocessableEntity {
	return &ListCategoryMilestonesUnprocessableEntity{}
}

/*
ListCategoryMilestonesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type ListCategoryMilestonesUnprocessableEntity struct {
}

// IsSuccess returns true when this list category milestones unprocessable entity response has a 2xx status code
func (o *ListCategoryMilestonesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list category milestones unprocessable entity response has a 3xx status code
func (o *ListCategoryMilestonesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list category milestones unprocessable entity response has a 4xx status code
func (o *ListCategoryMilestonesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list category milestones unprocessable entity response has a 5xx status code
func (o *ListCategoryMilestonesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list category milestones unprocessable entity response a status code equal to that given
func (o *ListCategoryMilestonesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list category milestones unprocessable entity response
func (o *ListCategoryMilestonesUnprocessableEntity) Code() int {
	return 422
}

func (o *ListCategoryMilestonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesUnprocessableEntity", 422)
}

func (o *ListCategoryMilestonesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] listCategoryMilestonesUnprocessableEntity", 422)
}

func (o *ListCategoryMilestonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
