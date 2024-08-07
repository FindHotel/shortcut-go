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

// SearchStoriesReader is a Reader for the SearchStories structure.
type SearchStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchStoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSearchStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/search/stories] searchStories", response, response.Code())
	}
}

// NewSearchStoriesOK creates a SearchStoriesOK with default headers values
func NewSearchStoriesOK() *SearchStoriesOK {
	return &SearchStoriesOK{}
}

/*
SearchStoriesOK describes a response with status code 200, with default header values.

Resource
*/
type SearchStoriesOK struct {
	Payload *models.StorySearchResults
}

// IsSuccess returns true when this search stories o k response has a 2xx status code
func (o *SearchStoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search stories o k response has a 3xx status code
func (o *SearchStoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stories o k response has a 4xx status code
func (o *SearchStoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search stories o k response has a 5xx status code
func (o *SearchStoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search stories o k response a status code equal to that given
func (o *SearchStoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search stories o k response
func (o *SearchStoriesOK) Code() int {
	return 200
}

func (o *SearchStoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesOK %s", 200, payload)
}

func (o *SearchStoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesOK %s", 200, payload)
}

func (o *SearchStoriesOK) GetPayload() *models.StorySearchResults {
	return o.Payload
}

func (o *SearchStoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorySearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStoriesBadRequest creates a SearchStoriesBadRequest with default headers values
func NewSearchStoriesBadRequest() *SearchStoriesBadRequest {
	return &SearchStoriesBadRequest{}
}

/*
SearchStoriesBadRequest describes a response with status code 400, with default header values.

**Either:** (1) Schema mismatch **or** (2) Maximum of 1000 search results exceeded
*/
type SearchStoriesBadRequest struct {
	Payload *models.MaxSearchResultsExceededError
}

// IsSuccess returns true when this search stories bad request response has a 2xx status code
func (o *SearchStoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stories bad request response has a 3xx status code
func (o *SearchStoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stories bad request response has a 4xx status code
func (o *SearchStoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stories bad request response has a 5xx status code
func (o *SearchStoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search stories bad request response a status code equal to that given
func (o *SearchStoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the search stories bad request response
func (o *SearchStoriesBadRequest) Code() int {
	return 400
}

func (o *SearchStoriesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesBadRequest %s", 400, payload)
}

func (o *SearchStoriesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesBadRequest %s", 400, payload)
}

func (o *SearchStoriesBadRequest) GetPayload() *models.MaxSearchResultsExceededError {
	return o.Payload
}

func (o *SearchStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaxSearchResultsExceededError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStoriesNotFound creates a SearchStoriesNotFound with default headers values
func NewSearchStoriesNotFound() *SearchStoriesNotFound {
	return &SearchStoriesNotFound{}
}

/*
SearchStoriesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type SearchStoriesNotFound struct {
}

// IsSuccess returns true when this search stories not found response has a 2xx status code
func (o *SearchStoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stories not found response has a 3xx status code
func (o *SearchStoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stories not found response has a 4xx status code
func (o *SearchStoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stories not found response has a 5xx status code
func (o *SearchStoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search stories not found response a status code equal to that given
func (o *SearchStoriesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search stories not found response
func (o *SearchStoriesNotFound) Code() int {
	return 404
}

func (o *SearchStoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesNotFound", 404)
}

func (o *SearchStoriesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesNotFound", 404)
}

func (o *SearchStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchStoriesUnprocessableEntity creates a SearchStoriesUnprocessableEntity with default headers values
func NewSearchStoriesUnprocessableEntity() *SearchStoriesUnprocessableEntity {
	return &SearchStoriesUnprocessableEntity{}
}

/*
SearchStoriesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type SearchStoriesUnprocessableEntity struct {
}

// IsSuccess returns true when this search stories unprocessable entity response has a 2xx status code
func (o *SearchStoriesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search stories unprocessable entity response has a 3xx status code
func (o *SearchStoriesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search stories unprocessable entity response has a 4xx status code
func (o *SearchStoriesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this search stories unprocessable entity response has a 5xx status code
func (o *SearchStoriesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this search stories unprocessable entity response a status code equal to that given
func (o *SearchStoriesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the search stories unprocessable entity response
func (o *SearchStoriesUnprocessableEntity) Code() int {
	return 422
}

func (o *SearchStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesUnprocessableEntity", 422)
}

func (o *SearchStoriesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/search/stories][%d] searchStoriesUnprocessableEntity", 422)
}

func (o *SearchStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
