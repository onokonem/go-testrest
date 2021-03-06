package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/onokonem/go-testrest/proto/models"
)

// AddAccountOKCode is the HTTP code returned for type AddAccountOK
const AddAccountOKCode int = 200

/*AddAccountOK successful operation

swagger:response addAccountOK
*/
type AddAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewAddAccountOK creates AddAccountOK with default headers values
func NewAddAccountOK() *AddAccountOK {
	return &AddAccountOK{}
}

// WithPayload adds the payload to the add account o k response
func (o *AddAccountOK) WithPayload(payload *models.Account) *AddAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add account o k response
func (o *AddAccountOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAccountMethodNotAllowedCode is the HTTP code returned for type AddAccountMethodNotAllowed
const AddAccountMethodNotAllowedCode int = 405

/*AddAccountMethodNotAllowed Invalid input

swagger:response addAccountMethodNotAllowed
*/
type AddAccountMethodNotAllowed struct {
}

// NewAddAccountMethodNotAllowed creates AddAccountMethodNotAllowed with default headers values
func NewAddAccountMethodNotAllowed() *AddAccountMethodNotAllowed {
	return &AddAccountMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddAccountMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
}

// AddAccountInternalServerErrorCode is the HTTP code returned for type AddAccountInternalServerError
const AddAccountInternalServerErrorCode int = 500

/*AddAccountInternalServerError Operation error

swagger:response addAccountInternalServerError
*/
type AddAccountInternalServerError struct {
}

// NewAddAccountInternalServerError creates AddAccountInternalServerError with default headers values
func NewAddAccountInternalServerError() *AddAccountInternalServerError {
	return &AddAccountInternalServerError{}
}

// WriteResponse to the client
func (o *AddAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}
