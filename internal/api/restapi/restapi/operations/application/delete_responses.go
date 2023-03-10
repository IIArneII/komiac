// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"komiac/internal/api/restapi/models"
)

// DeleteOKCode is the HTTP code returned for type DeleteOK
const DeleteOKCode int = 200

/*
DeleteOK Success

swagger:response deleteOK
*/
type DeleteOK struct {
}

// NewDeleteOK creates DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {

	return &DeleteOK{}
}

// WriteResponse to the client
func (o *DeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

func (o *DeleteOK) DeleteResponder() {}

// DeleteNotFoundCode is the HTTP code returned for type DeleteNotFound
const DeleteNotFoundCode int = 404

/*
DeleteNotFound Not Found

swagger:response deleteNotFound
*/
type DeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNotFound creates DeleteNotFound with default headers values
func NewDeleteNotFound() *DeleteNotFound {

	return &DeleteNotFound{}
}

// WithPayload adds the payload to the delete not found response
func (o *DeleteNotFound) WithPayload(payload *models.Error) *DeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete not found response
func (o *DeleteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteNotFound) DeleteResponder() {}

/*
DeleteDefault Error

swagger:response deleteDefault
*/
type DeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDefault creates DeleteDefault with default headers values
func NewDeleteDefault(code int) *DeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete default response
func (o *DeleteDefault) WithStatusCode(code int) *DeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete default response
func (o *DeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete default response
func (o *DeleteDefault) WithPayload(payload *models.Error) *DeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete default response
func (o *DeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteDefault) DeleteResponder() {}

type DeleteNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteNotImplementedResponder) DeleteResponder() {}

func DeleteNotImplemented() DeleteResponder {
	return &DeleteNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Delete has not yet been implemented",
		),
	}
}

type DeleteResponder interface {
	middleware.Responder
	DeleteResponder()
}
