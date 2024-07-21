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

// GetMilestoneReader is a Reader for the GetMilestone structure.
type GetMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMilestoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetMilestoneUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/milestones/{milestone-public-id}] getMilestone", response, response.Code())
	}
}

// NewGetMilestoneOK creates a GetMilestoneOK with default headers values
func NewGetMilestoneOK() *GetMilestoneOK {
	return &GetMilestoneOK{}
}

/*
GetMilestoneOK describes a response with status code 200, with default header values.

Resource
*/
type GetMilestoneOK struct {
	Payload *models.Milestone
}

// IsSuccess returns true when this get milestone o k response has a 2xx status code
func (o *GetMilestoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get milestone o k response has a 3xx status code
func (o *GetMilestoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get milestone o k response has a 4xx status code
func (o *GetMilestoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get milestone o k response has a 5xx status code
func (o *GetMilestoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get milestone o k response a status code equal to that given
func (o *GetMilestoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get milestone o k response
func (o *GetMilestoneOK) Code() int {
	return 200
}

func (o *GetMilestoneOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneOK %s", 200, payload)
}

func (o *GetMilestoneOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneOK %s", 200, payload)
}

func (o *GetMilestoneOK) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *GetMilestoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMilestoneBadRequest creates a GetMilestoneBadRequest with default headers values
func NewGetMilestoneBadRequest() *GetMilestoneBadRequest {
	return &GetMilestoneBadRequest{}
}

/*
GetMilestoneBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetMilestoneBadRequest struct {
}

// IsSuccess returns true when this get milestone bad request response has a 2xx status code
func (o *GetMilestoneBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get milestone bad request response has a 3xx status code
func (o *GetMilestoneBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get milestone bad request response has a 4xx status code
func (o *GetMilestoneBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get milestone bad request response has a 5xx status code
func (o *GetMilestoneBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get milestone bad request response a status code equal to that given
func (o *GetMilestoneBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get milestone bad request response
func (o *GetMilestoneBadRequest) Code() int {
	return 400
}

func (o *GetMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneBadRequest", 400)
}

func (o *GetMilestoneBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneBadRequest", 400)
}

func (o *GetMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMilestoneNotFound creates a GetMilestoneNotFound with default headers values
func NewGetMilestoneNotFound() *GetMilestoneNotFound {
	return &GetMilestoneNotFound{}
}

/*
GetMilestoneNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetMilestoneNotFound struct {
}

// IsSuccess returns true when this get milestone not found response has a 2xx status code
func (o *GetMilestoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get milestone not found response has a 3xx status code
func (o *GetMilestoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get milestone not found response has a 4xx status code
func (o *GetMilestoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get milestone not found response has a 5xx status code
func (o *GetMilestoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get milestone not found response a status code equal to that given
func (o *GetMilestoneNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get milestone not found response
func (o *GetMilestoneNotFound) Code() int {
	return 404
}

func (o *GetMilestoneNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneNotFound", 404)
}

func (o *GetMilestoneNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneNotFound", 404)
}

func (o *GetMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMilestoneUnprocessableEntity creates a GetMilestoneUnprocessableEntity with default headers values
func NewGetMilestoneUnprocessableEntity() *GetMilestoneUnprocessableEntity {
	return &GetMilestoneUnprocessableEntity{}
}

/*
GetMilestoneUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetMilestoneUnprocessableEntity struct {
}

// IsSuccess returns true when this get milestone unprocessable entity response has a 2xx status code
func (o *GetMilestoneUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get milestone unprocessable entity response has a 3xx status code
func (o *GetMilestoneUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get milestone unprocessable entity response has a 4xx status code
func (o *GetMilestoneUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get milestone unprocessable entity response has a 5xx status code
func (o *GetMilestoneUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get milestone unprocessable entity response a status code equal to that given
func (o *GetMilestoneUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get milestone unprocessable entity response
func (o *GetMilestoneUnprocessableEntity) Code() int {
	return 422
}

func (o *GetMilestoneUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneUnprocessableEntity", 422)
}

func (o *GetMilestoneUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getMilestoneUnprocessableEntity", 422)
}

func (o *GetMilestoneUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
