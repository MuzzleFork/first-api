// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/muzzlefork/first-api/models"
)

// GetConfigOKCode is the HTTP code returned for type GetConfigOK
const GetConfigOKCode int = 200

/*GetConfigOK Successfully returns the current configuration

swagger:response getConfigOK
*/
type GetConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.Config `json:"body,omitempty"`
}

// NewGetConfigOK creates GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {

	return &GetConfigOK{}
}

// WithPayload adds the payload to the get config o k response
func (o *GetConfigOK) WithPayload(payload *models.Config) *GetConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config o k response
func (o *GetConfigOK) SetPayload(payload *models.Config) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigUnauthorizedCode is the HTTP code returned for type GetConfigUnauthorized
const GetConfigUnauthorizedCode int = 401

/*GetConfigUnauthorized Unauthorized

swagger:response getConfigUnauthorized
*/
type GetConfigUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.AuthError `json:"body,omitempty"`
}

// NewGetConfigUnauthorized creates GetConfigUnauthorized with default headers values
func NewGetConfigUnauthorized() *GetConfigUnauthorized {

	return &GetConfigUnauthorized{}
}

// WithPayload adds the payload to the get config unauthorized response
func (o *GetConfigUnauthorized) WithPayload(payload *models.AuthError) *GetConfigUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config unauthorized response
func (o *GetConfigUnauthorized) SetPayload(payload *models.AuthError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigForbiddenCode is the HTTP code returned for type GetConfigForbidden
const GetConfigForbiddenCode int = 403

/*GetConfigForbidden Forbidden

swagger:response getConfigForbidden
*/
type GetConfigForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.AuthError `json:"body,omitempty"`
}

// NewGetConfigForbidden creates GetConfigForbidden with default headers values
func NewGetConfigForbidden() *GetConfigForbidden {

	return &GetConfigForbidden{}
}

// WithPayload adds the payload to the get config forbidden response
func (o *GetConfigForbidden) WithPayload(payload *models.AuthError) *GetConfigForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config forbidden response
func (o *GetConfigForbidden) SetPayload(payload *models.AuthError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigInternalServerErrorCode is the HTTP code returned for type GetConfigInternalServerError
const GetConfigInternalServerErrorCode int = 500

/*GetConfigInternalServerError Server Error

swagger:response getConfigInternalServerError
*/
type GetConfigInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigInternalServerError creates GetConfigInternalServerError with default headers values
func NewGetConfigInternalServerError() *GetConfigInternalServerError {

	return &GetConfigInternalServerError{}
}

// WithPayload adds the payload to the get config internal server error response
func (o *GetConfigInternalServerError) WithPayload(payload *models.Error) *GetConfigInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config internal server error response
func (o *GetConfigInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
