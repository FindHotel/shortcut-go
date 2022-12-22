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

// GetLabelReader is a Reader for the GetLabel structure.
type GetLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetLabelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelOK creates a GetLabelOK with default headers values
func NewGetLabelOK() *GetLabelOK {
	return &GetLabelOK{}
}

/*
GetLabelOK describes a response with status code 200, with default header values.

Resource
*/
type GetLabelOK struct {
	Payload *models.Label
}

// IsSuccess returns true when this get label o k response has a 2xx status code
func (o *GetLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get label o k response has a 3xx status code
func (o *GetLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get label o k response has a 4xx status code
func (o *GetLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get label o k response has a 5xx status code
func (o *GetLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get label o k response a status code equal to that given
func (o *GetLabelOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLabelOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelOK  %+v", 200, o.Payload)
}

func (o *GetLabelOK) String() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelOK  %+v", 200, o.Payload)
}

func (o *GetLabelOK) GetPayload() *models.Label {
	return o.Payload
}

func (o *GetLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelBadRequest creates a GetLabelBadRequest with default headers values
func NewGetLabelBadRequest() *GetLabelBadRequest {
	return &GetLabelBadRequest{}
}

/*
GetLabelBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type GetLabelBadRequest struct {
}

// IsSuccess returns true when this get label bad request response has a 2xx status code
func (o *GetLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get label bad request response has a 3xx status code
func (o *GetLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get label bad request response has a 4xx status code
func (o *GetLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get label bad request response has a 5xx status code
func (o *GetLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get label bad request response a status code equal to that given
func (o *GetLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLabelBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelBadRequest ", 400)
}

func (o *GetLabelBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelBadRequest ", 400)
}

func (o *GetLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelNotFound creates a GetLabelNotFound with default headers values
func NewGetLabelNotFound() *GetLabelNotFound {
	return &GetLabelNotFound{}
}

/*
GetLabelNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type GetLabelNotFound struct {
}

// IsSuccess returns true when this get label not found response has a 2xx status code
func (o *GetLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get label not found response has a 3xx status code
func (o *GetLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get label not found response has a 4xx status code
func (o *GetLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get label not found response has a 5xx status code
func (o *GetLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get label not found response a status code equal to that given
func (o *GetLabelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLabelNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelNotFound ", 404)
}

func (o *GetLabelNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelNotFound ", 404)
}

func (o *GetLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelUnprocessableEntity creates a GetLabelUnprocessableEntity with default headers values
func NewGetLabelUnprocessableEntity() *GetLabelUnprocessableEntity {
	return &GetLabelUnprocessableEntity{}
}

/*
GetLabelUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type GetLabelUnprocessableEntity struct {
}

// IsSuccess returns true when this get label unprocessable entity response has a 2xx status code
func (o *GetLabelUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get label unprocessable entity response has a 3xx status code
func (o *GetLabelUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get label unprocessable entity response has a 4xx status code
func (o *GetLabelUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get label unprocessable entity response has a 5xx status code
func (o *GetLabelUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get label unprocessable entity response a status code equal to that given
func (o *GetLabelUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetLabelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelUnprocessableEntity ", 422)
}

func (o *GetLabelUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}][%d] getLabelUnprocessableEntity ", 422)
}

func (o *GetLabelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
