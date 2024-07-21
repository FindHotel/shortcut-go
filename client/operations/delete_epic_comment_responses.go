// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEpicCommentReader is a Reader for the DeleteEpicComment structure.
type DeleteEpicCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEpicCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEpicCommentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEpicCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEpicCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteEpicCommentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}] deleteEpicComment", response, response.Code())
	}
}

// NewDeleteEpicCommentNoContent creates a DeleteEpicCommentNoContent with default headers values
func NewDeleteEpicCommentNoContent() *DeleteEpicCommentNoContent {
	return &DeleteEpicCommentNoContent{}
}

/*
DeleteEpicCommentNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteEpicCommentNoContent struct {
}

// IsSuccess returns true when this delete epic comment no content response has a 2xx status code
func (o *DeleteEpicCommentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete epic comment no content response has a 3xx status code
func (o *DeleteEpicCommentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete epic comment no content response has a 4xx status code
func (o *DeleteEpicCommentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete epic comment no content response has a 5xx status code
func (o *DeleteEpicCommentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete epic comment no content response a status code equal to that given
func (o *DeleteEpicCommentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete epic comment no content response
func (o *DeleteEpicCommentNoContent) Code() int {
	return 204
}

func (o *DeleteEpicCommentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNoContent", 204)
}

func (o *DeleteEpicCommentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNoContent", 204)
}

func (o *DeleteEpicCommentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentBadRequest creates a DeleteEpicCommentBadRequest with default headers values
func NewDeleteEpicCommentBadRequest() *DeleteEpicCommentBadRequest {
	return &DeleteEpicCommentBadRequest{}
}

/*
DeleteEpicCommentBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type DeleteEpicCommentBadRequest struct {
}

// IsSuccess returns true when this delete epic comment bad request response has a 2xx status code
func (o *DeleteEpicCommentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete epic comment bad request response has a 3xx status code
func (o *DeleteEpicCommentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete epic comment bad request response has a 4xx status code
func (o *DeleteEpicCommentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete epic comment bad request response has a 5xx status code
func (o *DeleteEpicCommentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete epic comment bad request response a status code equal to that given
func (o *DeleteEpicCommentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete epic comment bad request response
func (o *DeleteEpicCommentBadRequest) Code() int {
	return 400
}

func (o *DeleteEpicCommentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentBadRequest", 400)
}

func (o *DeleteEpicCommentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentBadRequest", 400)
}

func (o *DeleteEpicCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentNotFound creates a DeleteEpicCommentNotFound with default headers values
func NewDeleteEpicCommentNotFound() *DeleteEpicCommentNotFound {
	return &DeleteEpicCommentNotFound{}
}

/*
DeleteEpicCommentNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type DeleteEpicCommentNotFound struct {
}

// IsSuccess returns true when this delete epic comment not found response has a 2xx status code
func (o *DeleteEpicCommentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete epic comment not found response has a 3xx status code
func (o *DeleteEpicCommentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete epic comment not found response has a 4xx status code
func (o *DeleteEpicCommentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete epic comment not found response has a 5xx status code
func (o *DeleteEpicCommentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete epic comment not found response a status code equal to that given
func (o *DeleteEpicCommentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete epic comment not found response
func (o *DeleteEpicCommentNotFound) Code() int {
	return 404
}

func (o *DeleteEpicCommentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNotFound", 404)
}

func (o *DeleteEpicCommentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentNotFound", 404)
}

func (o *DeleteEpicCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEpicCommentUnprocessableEntity creates a DeleteEpicCommentUnprocessableEntity with default headers values
func NewDeleteEpicCommentUnprocessableEntity() *DeleteEpicCommentUnprocessableEntity {
	return &DeleteEpicCommentUnprocessableEntity{}
}

/*
DeleteEpicCommentUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type DeleteEpicCommentUnprocessableEntity struct {
}

// IsSuccess returns true when this delete epic comment unprocessable entity response has a 2xx status code
func (o *DeleteEpicCommentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete epic comment unprocessable entity response has a 3xx status code
func (o *DeleteEpicCommentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete epic comment unprocessable entity response has a 4xx status code
func (o *DeleteEpicCommentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete epic comment unprocessable entity response has a 5xx status code
func (o *DeleteEpicCommentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete epic comment unprocessable entity response a status code equal to that given
func (o *DeleteEpicCommentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete epic comment unprocessable entity response
func (o *DeleteEpicCommentUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteEpicCommentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentUnprocessableEntity", 422)
}

func (o *DeleteEpicCommentUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] deleteEpicCommentUnprocessableEntity", 422)
}

func (o *DeleteEpicCommentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
