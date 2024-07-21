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

// SearchObjectivesReader is a Reader for the SearchObjectives structure.
type SearchObjectivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchObjectivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchObjectivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchObjectivesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchObjectivesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSearchObjectivesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v3/search/objectives] searchObjectives", response, response.Code())
	}
}

// NewSearchObjectivesOK creates a SearchObjectivesOK with default headers values
func NewSearchObjectivesOK() *SearchObjectivesOK {
	return &SearchObjectivesOK{}
}

/*
SearchObjectivesOK describes a response with status code 200, with default header values.

Resource
*/
type SearchObjectivesOK struct {
	Payload *models.ObjectiveSearchResults
}

// IsSuccess returns true when this search objectives o k response has a 2xx status code
func (o *SearchObjectivesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search objectives o k response has a 3xx status code
func (o *SearchObjectivesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objectives o k response has a 4xx status code
func (o *SearchObjectivesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search objectives o k response has a 5xx status code
func (o *SearchObjectivesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search objectives o k response a status code equal to that given
func (o *SearchObjectivesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search objectives o k response
func (o *SearchObjectivesOK) Code() int {
	return 200
}

func (o *SearchObjectivesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesOK %s", 200, payload)
}

func (o *SearchObjectivesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesOK %s", 200, payload)
}

func (o *SearchObjectivesOK) GetPayload() *models.ObjectiveSearchResults {
	return o.Payload
}

func (o *SearchObjectivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectiveSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchObjectivesBadRequest creates a SearchObjectivesBadRequest with default headers values
func NewSearchObjectivesBadRequest() *SearchObjectivesBadRequest {
	return &SearchObjectivesBadRequest{}
}

/*
SearchObjectivesBadRequest describes a response with status code 400, with default header values.

**Either:** (1) Schema mismatch **or** (2) Maximum of 1000 search results exceeded
*/
type SearchObjectivesBadRequest struct {
	Payload *models.MaxSearchResultsExceededError
}

// IsSuccess returns true when this search objectives bad request response has a 2xx status code
func (o *SearchObjectivesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objectives bad request response has a 3xx status code
func (o *SearchObjectivesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objectives bad request response has a 4xx status code
func (o *SearchObjectivesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search objectives bad request response has a 5xx status code
func (o *SearchObjectivesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search objectives bad request response a status code equal to that given
func (o *SearchObjectivesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the search objectives bad request response
func (o *SearchObjectivesBadRequest) Code() int {
	return 400
}

func (o *SearchObjectivesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesBadRequest %s", 400, payload)
}

func (o *SearchObjectivesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesBadRequest %s", 400, payload)
}

func (o *SearchObjectivesBadRequest) GetPayload() *models.MaxSearchResultsExceededError {
	return o.Payload
}

func (o *SearchObjectivesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaxSearchResultsExceededError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchObjectivesNotFound creates a SearchObjectivesNotFound with default headers values
func NewSearchObjectivesNotFound() *SearchObjectivesNotFound {
	return &SearchObjectivesNotFound{}
}

/*
SearchObjectivesNotFound describes a response with status code 404, with default header values.

Resource does not exist
*/
type SearchObjectivesNotFound struct {
}

// IsSuccess returns true when this search objectives not found response has a 2xx status code
func (o *SearchObjectivesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objectives not found response has a 3xx status code
func (o *SearchObjectivesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objectives not found response has a 4xx status code
func (o *SearchObjectivesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search objectives not found response has a 5xx status code
func (o *SearchObjectivesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search objectives not found response a status code equal to that given
func (o *SearchObjectivesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search objectives not found response
func (o *SearchObjectivesNotFound) Code() int {
	return 404
}

func (o *SearchObjectivesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesNotFound", 404)
}

func (o *SearchObjectivesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesNotFound", 404)
}

func (o *SearchObjectivesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchObjectivesUnprocessableEntity creates a SearchObjectivesUnprocessableEntity with default headers values
func NewSearchObjectivesUnprocessableEntity() *SearchObjectivesUnprocessableEntity {
	return &SearchObjectivesUnprocessableEntity{}
}

/*
SearchObjectivesUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable
*/
type SearchObjectivesUnprocessableEntity struct {
}

// IsSuccess returns true when this search objectives unprocessable entity response has a 2xx status code
func (o *SearchObjectivesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objectives unprocessable entity response has a 3xx status code
func (o *SearchObjectivesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objectives unprocessable entity response has a 4xx status code
func (o *SearchObjectivesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this search objectives unprocessable entity response has a 5xx status code
func (o *SearchObjectivesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this search objectives unprocessable entity response a status code equal to that given
func (o *SearchObjectivesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the search objectives unprocessable entity response
func (o *SearchObjectivesUnprocessableEntity) Code() int {
	return 422
}

func (o *SearchObjectivesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesUnprocessableEntity", 422)
}

func (o *SearchObjectivesUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v3/search/objectives][%d] searchObjectivesUnprocessableEntity", 422)
}

func (o *SearchObjectivesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
