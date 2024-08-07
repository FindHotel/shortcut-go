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

// UnlinkCommentThreadFromSlackReader is a Reader for the UnlinkCommentThreadFromSlack structure.
type UnlinkCommentThreadFromSlackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlinkCommentThreadFromSlackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUnlinkCommentThreadFromSlackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnlinkCommentThreadFromSlackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnlinkCommentThreadFromSlackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUnlinkCommentThreadFromSlackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack] unlinkCommentThreadFromSlack", response, response.Code())
	}
}

// NewUnlinkCommentThreadFromSlackCreated creates a UnlinkCommentThreadFromSlackCreated with default headers values
func NewUnlinkCommentThreadFromSlackCreated() *UnlinkCommentThreadFromSlackCreated {
	return &UnlinkCommentThreadFromSlackCreated{}
}

/*
UnlinkCommentThreadFromSlackCreated describes a response with status code 201, with default header values.

Resource
*/
type UnlinkCommentThreadFromSlackCreated struct {
	Payload *models.StoryComment
}

// IsSuccess returns true when this unlink comment thread from slack created response has a 2xx status code
func (o *UnlinkCommentThreadFromSlackCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unlink comment thread from slack created response has a 3xx status code
func (o *UnlinkCommentThreadFromSlackCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink comment thread from slack created response has a 4xx status code
func (o *UnlinkCommentThreadFromSlackCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this unlink comment thread from slack created response has a 5xx status code
func (o *UnlinkCommentThreadFromSlackCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink comment thread from slack created response a status code equal to that given
func (o *UnlinkCommentThreadFromSlackCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the unlink comment thread from slack created response
func (o *UnlinkCommentThreadFromSlackCreated) Code() int {
	return 201
}

func (o *UnlinkCommentThreadFromSlackCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackCreated %s", 201, payload)
}

func (o *UnlinkCommentThreadFromSlackCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackCreated %s", 201, payload)
}

func (o *UnlinkCommentThreadFromSlackCreated) GetPayload() *models.StoryComment {
	return o.Payload
}

func (o *UnlinkCommentThreadFromSlackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoryComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnlinkCommentThreadFromSlackBadRequest creates a UnlinkCommentThreadFromSlackBadRequest with default headers values
func NewUnlinkCommentThreadFromSlackBadRequest() *UnlinkCommentThreadFromSlackBadRequest {
	return &UnlinkCommentThreadFromSlackBadRequest{}
}

/*
UnlinkCommentThreadFromSlackBadRequest describes a response with status code 400, with default header values.

Schema mismatch
*/
type UnlinkCommentThreadFromSlackBadRequest struct {
}

// IsSuccess returns true when this unlink comment thread from slack bad request response has a 2xx status code
func (o *UnlinkCommentThreadFromSlackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink comment thread from slack bad request response has a 3xx status code
func (o *UnlinkCommentThreadFromSlackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink comment thread from slack bad request response has a 4xx status code
func (o *UnlinkCommentThreadFromSlackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink comment thread from slack bad request response has a 5xx status code
func (o *UnlinkCommentThreadFromSlackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink comment thread from slack bad request response a status code equal to that given
func (o *UnlinkCommentThreadFromSlackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the unlink comment thread from slack bad request response
func (o *UnlinkCommentThreadFromSlackBadRequest) Code() int {
	return 400
}

func (o *UnlinkCommentThreadFromSlackBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackBadRequest", 400)
}

func (o *UnlinkCommentThreadFromSlackBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackBadRequest", 400)
}

func (o *UnlinkCommentThreadFromSlackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkCommentThreadFromSlackNotFound creates a UnlinkCommentThreadFromSlackNotFound with default headers values
func NewUnlinkCommentThreadFromSlackNotFound() *UnlinkCommentThreadFromSlackNotFound {
	return &UnlinkCommentThreadFromSlackNotFound{}
}

/*
UnlinkCommentThreadFromSlackNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type UnlinkCommentThreadFromSlackNotFound struct {
}

// IsSuccess returns true when this unlink comment thread from slack not found response has a 2xx status code
func (o *UnlinkCommentThreadFromSlackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink comment thread from slack not found response has a 3xx status code
func (o *UnlinkCommentThreadFromSlackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink comment thread from slack not found response has a 4xx status code
func (o *UnlinkCommentThreadFromSlackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink comment thread from slack not found response has a 5xx status code
func (o *UnlinkCommentThreadFromSlackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink comment thread from slack not found response a status code equal to that given
func (o *UnlinkCommentThreadFromSlackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the unlink comment thread from slack not found response
func (o *UnlinkCommentThreadFromSlackNotFound) Code() int {
	return 404
}

func (o *UnlinkCommentThreadFromSlackNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackNotFound", 404)
}

func (o *UnlinkCommentThreadFromSlackNotFound) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackNotFound", 404)
}

func (o *UnlinkCommentThreadFromSlackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkCommentThreadFromSlackUnprocessableEntity creates a UnlinkCommentThreadFromSlackUnprocessableEntity with default headers values
func NewUnlinkCommentThreadFromSlackUnprocessableEntity() *UnlinkCommentThreadFromSlackUnprocessableEntity {
	return &UnlinkCommentThreadFromSlackUnprocessableEntity{}
}

/*
UnlinkCommentThreadFromSlackUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type UnlinkCommentThreadFromSlackUnprocessableEntity struct {
}

// IsSuccess returns true when this unlink comment thread from slack unprocessable entity response has a 2xx status code
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unlink comment thread from slack unprocessable entity response has a 3xx status code
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unlink comment thread from slack unprocessable entity response has a 4xx status code
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this unlink comment thread from slack unprocessable entity response has a 5xx status code
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this unlink comment thread from slack unprocessable entity response a status code equal to that given
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the unlink comment thread from slack unprocessable entity response
func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) Code() int {
	return 422
}

func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackUnprocessableEntity", 422)
}

func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack][%d] unlinkCommentThreadFromSlackUnprocessableEntity", 422)
}

func (o *UnlinkCommentThreadFromSlackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
