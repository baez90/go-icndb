// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthOKCode is the HTTP code returned for type GetHealthOK
const GetHealthOKCode int = 200

/*GetHealthOK Service is healthy

swagger:response getHealthOK
*/
type GetHealthOK struct {
}

// NewGetHealthOK creates GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {

	return &GetHealthOK{}
}

// WriteResponse to the client
func (o *GetHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetHealthInternalServerErrorCode is the HTTP code returned for type GetHealthInternalServerError
const GetHealthInternalServerErrorCode int = 500

/*GetHealthInternalServerError Service is not healty

swagger:response getHealthInternalServerError
*/
type GetHealthInternalServerError struct {
}

// NewGetHealthInternalServerError creates GetHealthInternalServerError with default headers values
func NewGetHealthInternalServerError() *GetHealthInternalServerError {

	return &GetHealthInternalServerError{}
}

// WriteResponse to the client
func (o *GetHealthInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
