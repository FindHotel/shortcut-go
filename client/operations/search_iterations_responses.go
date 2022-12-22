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

// SearchIterationsReader is a Reader for the SearchIterations structure.
type SearchIterationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchIterationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchIterationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchIterationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchIterationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSearchIterationsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchIterationsOK creates a SearchIterationsOK with default headers values
func NewSearchIterationsOK() *SearchIterationsOK {
	return &SearchIterationsOK{}
}

/*
SearchIterationsOK describes a response with status code 200, with default header values.

Resource
*/
type SearchIterationsOK struct {
	Payload *models.IterationSearchResults
}

// IsSuccess returns true when this search iterations o k response has a 2xx status code
func (o *SearchIterationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search iterations o k response has a 3xx status code
func (o *SearchIterationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search iterations o k response has a 4xx status code
func (o *SearchIterationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search iterations o k response has a 5xx status code
func (o *SearchIterationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search iterations o k response a status code equal to that given
func (o *SearchIterationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchIterationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsOK  %+v", 200, o.Payload)
}

func (o *SearchIterationsOK) String() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsOK  %+v", 200, o.Payload)
}

func (o *SearchIterationsOK) GetPayload() *models.IterationSearchResults {
	return o.Payload
}

func (o *SearchIterationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IterationSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIterationsBadRequest creates a SearchIterationsBadRequest with default headers values
func NewSearchIterationsBadRequest() *SearchIterationsBadRequest {
	return &SearchIterationsBadRequest{}
}

/*
SearchIterationsBadRequest describes a response with status code 400, with default header values.

**Either:** (1) Schema mismatch **or** (2) Maximum of 1000 search results exceeded
*/
type SearchIterationsBadRequest struct {
	Payload *models.MaxSearchResultsExceededError
}

// IsSuccess returns true when this search iterations bad request response has a 2xx status code
func (o *SearchIterationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search iterations bad request response has a 3xx status code
func (o *SearchIterationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search iterations bad request response has a 4xx status code
func (o *SearchIterationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search iterations bad request response has a 5xx status code
func (o *SearchIterationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search iterations bad request response a status code equal to that given
func (o *SearchIterationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchIterationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchIterationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchIterationsBadRequest) GetPayload() *models.MaxSearchResultsExceededError {
	return o.Payload
}

func (o *SearchIterationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaxSearchResultsExceededError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIterationsNotFound creates a SearchIterationsNotFound with default headers values
func NewSearchIterationsNotFound() *SearchIterationsNotFound {
	return &SearchIterationsNotFound{}
}

/*
SearchIterationsNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type SearchIterationsNotFound struct {
}

// IsSuccess returns true when this search iterations not found response has a 2xx status code
func (o *SearchIterationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search iterations not found response has a 3xx status code
func (o *SearchIterationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search iterations not found response has a 4xx status code
func (o *SearchIterationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search iterations not found response has a 5xx status code
func (o *SearchIterationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search iterations not found response a status code equal to that given
func (o *SearchIterationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchIterationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsNotFound ", 404)
}

func (o *SearchIterationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsNotFound ", 404)
}

func (o *SearchIterationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchIterationsUnprocessableEntity creates a SearchIterationsUnprocessableEntity with default headers values
func NewSearchIterationsUnprocessableEntity() *SearchIterationsUnprocessableEntity {
	return &SearchIterationsUnprocessableEntity{}
}

/*
SearchIterationsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type SearchIterationsUnprocessableEntity struct {
}

// IsSuccess returns true when this search iterations unprocessable entity response has a 2xx status code
func (o *SearchIterationsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search iterations unprocessable entity response has a 3xx status code
func (o *SearchIterationsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search iterations unprocessable entity response has a 4xx status code
func (o *SearchIterationsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this search iterations unprocessable entity response has a 5xx status code
func (o *SearchIterationsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this search iterations unprocessable entity response a status code equal to that given
func (o *SearchIterationsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *SearchIterationsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsUnprocessableEntity ", 422)
}

func (o *SearchIterationsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/search/iterations][%d] searchIterationsUnprocessableEntity ", 422)
}

func (o *SearchIterationsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}