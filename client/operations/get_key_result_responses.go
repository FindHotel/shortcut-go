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

// GetKeyResultReader is a Reader for the GetKeyResult structure.
type GetKeyResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeyResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeyResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKeyResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKeyResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetKeyResultUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/key-results/{key-result-public-id}] getKeyResult", response, response.Code())
	}
}

// NewGetKeyResultOK creates a GetKeyResultOK with default headers values
func NewGetKeyResultOK() *GetKeyResultOK {
	return &GetKeyResultOK{}
}

/*
GetKeyResultOK describes a response with status code 200, with default header values.

Resource
*/
type GetKeyResultOK struct {
	Payload *models.KeyResult
}

// IsSuccess returns true when this get key result o k response has a 2xx status code
func (o *GetKeyResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get key result o k response has a 3xx status code
func (o *GetKeyResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get key result o k response has a 4xx status code
func (o *GetKeyResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get key result o k response has a 5xx status code
func (o *GetKeyResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get key result o k response a status code equal to that given
func (o *GetKeyResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get key result o k response
func (o *GetKeyResultOK) Code() int {
	return 200
}

func (o *GetKeyResultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultOK %s", 200, payload)
}

func (o *GetKeyResultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultOK %s", 200, payload)
}

func (o *GetKeyResultOK) GetPayload() *models.KeyResult {
	return o.Payload
}

func (o *GetKeyResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeyResultBadRequest creates a GetKeyResultBadRequest with default headers values
func NewGetKeyResultBadRequest() *GetKeyResultBadRequest {
	return &GetKeyResultBadRequest{}
}

/*
GetKeyResultBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetKeyResultBadRequest struct {
}

// IsSuccess returns true when this get key result bad request response has a 2xx status code
func (o *GetKeyResultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get key result bad request response has a 3xx status code
func (o *GetKeyResultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get key result bad request response has a 4xx status code
func (o *GetKeyResultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get key result bad request response has a 5xx status code
func (o *GetKeyResultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get key result bad request response a status code equal to that given
func (o *GetKeyResultBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get key result bad request response
func (o *GetKeyResultBadRequest) Code() int {
	return 400
}

func (o *GetKeyResultBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultBadRequest", 400)
}

func (o *GetKeyResultBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultBadRequest", 400)
}

func (o *GetKeyResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKeyResultNotFound creates a GetKeyResultNotFound with default headers values
func NewGetKeyResultNotFound() *GetKeyResultNotFound {
	return &GetKeyResultNotFound{}
}

/*
GetKeyResultNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetKeyResultNotFound struct {
}

// IsSuccess returns true when this get key result not found response has a 2xx status code
func (o *GetKeyResultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get key result not found response has a 3xx status code
func (o *GetKeyResultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get key result not found response has a 4xx status code
func (o *GetKeyResultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get key result not found response has a 5xx status code
func (o *GetKeyResultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get key result not found response a status code equal to that given
func (o *GetKeyResultNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get key result not found response
func (o *GetKeyResultNotFound) Code() int {
	return 404
}

func (o *GetKeyResultNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultNotFound", 404)
}

func (o *GetKeyResultNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultNotFound", 404)
}

func (o *GetKeyResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKeyResultUnprocessableEntity creates a GetKeyResultUnprocessableEntity with default headers values
func NewGetKeyResultUnprocessableEntity() *GetKeyResultUnprocessableEntity {
	return &GetKeyResultUnprocessableEntity{}
}

/*
GetKeyResultUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetKeyResultUnprocessableEntity struct {
}

// IsSuccess returns true when this get key result unprocessable entity response has a 2xx status code
func (o *GetKeyResultUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get key result unprocessable entity response has a 3xx status code
func (o *GetKeyResultUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get key result unprocessable entity response has a 4xx status code
func (o *GetKeyResultUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get key result unprocessable entity response has a 5xx status code
func (o *GetKeyResultUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get key result unprocessable entity response a status code equal to that given
func (o *GetKeyResultUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get key result unprocessable entity response
func (o *GetKeyResultUnprocessableEntity) Code() int {
	return 422
}

func (o *GetKeyResultUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultUnprocessableEntity", 422)
}

func (o *GetKeyResultUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/key-results/{key-result-public-id}][%d] getKeyResultUnprocessableEntity", 422)
}

func (o *GetKeyResultUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
