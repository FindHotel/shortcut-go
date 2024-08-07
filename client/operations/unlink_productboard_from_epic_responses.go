// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UnlinkProductboardFromEpicReader is a Reader for the UnlinkProductboardFromEpic structure.
type UnlinkProductboardFromEpicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlinkProductboardFromEpicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnlinkProductboardFromEpicNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnlinkProductboardFromEpicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnlinkProductboardFromEpicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUnlinkProductboardFromEpicUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/epics/{epic-public-id}/unlink-productboard] unlinkProductboardFromEpic", response, response.Code())
	}
}

// NewUnlinkProductboardFromEpicNoContent creates a UnlinkProductboardFromEpicNoContent with default headers values
func NewUnlinkProductboardFromEpicNoContent() *UnlinkProductboardFromEpicNoContent {
	return &UnlinkProductboardFromEpicNoContent{}
}

/*
UnlinkProductboardFromEpicNoContent describes a response with status code 204, with default header values.

No Content
*/
type UnlinkProductboardFromEpicNoContent struct {
}

// IsSuccess returns true when this unlink productboard from epic no content response has a 2xx status code
func (o *UnlinkProductboardFromEpicNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unlink productboard from epic no content response has a 3xx status code
func (o *UnlinkProductboardFromEpicNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink productboard from epic no content response has a 4xx status code
func (o *UnlinkProductboardFromEpicNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this unlink productboard from epic no content response has a 5xx status code
func (o *UnlinkProductboardFromEpicNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink productboard from epic no content response a status code equal to that given
func (o *UnlinkProductboardFromEpicNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the unlink productboard from epic no content response
func (o *UnlinkProductboardFromEpicNoContent) Code() int {
	return 204
}

func (o *UnlinkProductboardFromEpicNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicNoContent", 204)
}

func (o *UnlinkProductboardFromEpicNoContent) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicNoContent", 204)
}

func (o *UnlinkProductboardFromEpicNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkProductboardFromEpicBadRequest creates a UnlinkProductboardFromEpicBadRequest with default headers values
func NewUnlinkProductboardFromEpicBadRequest() *UnlinkProductboardFromEpicBadRequest {
	return &UnlinkProductboardFromEpicBadRequest{}
}

/*
UnlinkProductboardFromEpicBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UnlinkProductboardFromEpicBadRequest struct {
}

// IsSuccess returns true when this unlink productboard from epic bad request response has a 2xx status code
func (o *UnlinkProductboardFromEpicBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink productboard from epic bad request response has a 3xx status code
func (o *UnlinkProductboardFromEpicBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink productboard from epic bad request response has a 4xx status code
func (o *UnlinkProductboardFromEpicBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink productboard from epic bad request response has a 5xx status code
func (o *UnlinkProductboardFromEpicBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink productboard from epic bad request response a status code equal to that given
func (o *UnlinkProductboardFromEpicBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the unlink productboard from epic bad request response
func (o *UnlinkProductboardFromEpicBadRequest) Code() int {
	return 400
}

func (o *UnlinkProductboardFromEpicBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicBadRequest", 400)
}

func (o *UnlinkProductboardFromEpicBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicBadRequest", 400)
}

func (o *UnlinkProductboardFromEpicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkProductboardFromEpicNotFound creates a UnlinkProductboardFromEpicNotFound with default headers values
func NewUnlinkProductboardFromEpicNotFound() *UnlinkProductboardFromEpicNotFound {
	return &UnlinkProductboardFromEpicNotFound{}
}

/*
UnlinkProductboardFromEpicNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UnlinkProductboardFromEpicNotFound struct {
}

// IsSuccess returns true when this unlink productboard from epic not found response has a 2xx status code
func (o *UnlinkProductboardFromEpicNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink productboard from epic not found response has a 3xx status code
func (o *UnlinkProductboardFromEpicNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink productboard from epic not found response has a 4xx status code
func (o *UnlinkProductboardFromEpicNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink productboard from epic not found response has a 5xx status code
func (o *UnlinkProductboardFromEpicNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink productboard from epic not found response a status code equal to that given
func (o *UnlinkProductboardFromEpicNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the unlink productboard from epic not found response
func (o *UnlinkProductboardFromEpicNotFound) Code() int {
	return 404
}

func (o *UnlinkProductboardFromEpicNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicNotFound", 404)
}

func (o *UnlinkProductboardFromEpicNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicNotFound", 404)
}

func (o *UnlinkProductboardFromEpicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkProductboardFromEpicUnprocessableEntity creates a UnlinkProductboardFromEpicUnprocessableEntity with default headers values
func NewUnlinkProductboardFromEpicUnprocessableEntity() *UnlinkProductboardFromEpicUnprocessableEntity {
	return &UnlinkProductboardFromEpicUnprocessableEntity{}
}

/*
UnlinkProductboardFromEpicUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UnlinkProductboardFromEpicUnprocessableEntity struct {
}

// IsSuccess returns true when this unlink productboard from epic unprocessable entity response has a 2xx status code
func (o *UnlinkProductboardFromEpicUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink productboard from epic unprocessable entity response has a 3xx status code
func (o *UnlinkProductboardFromEpicUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink productboard from epic unprocessable entity response has a 4xx status code
func (o *UnlinkProductboardFromEpicUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink productboard from epic unprocessable entity response has a 5xx status code
func (o *UnlinkProductboardFromEpicUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink productboard from epic unprocessable entity response a status code equal to that given
func (o *UnlinkProductboardFromEpicUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the unlink productboard from epic unprocessable entity response
func (o *UnlinkProductboardFromEpicUnprocessableEntity) Code() int {
	return 422
}

func (o *UnlinkProductboardFromEpicUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicUnprocessableEntity", 422)
}

func (o *UnlinkProductboardFromEpicUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/unlink-productboard][%d] unlinkProductboardFromEpicUnprocessableEntity", 422)
}

func (o *UnlinkProductboardFromEpicUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
