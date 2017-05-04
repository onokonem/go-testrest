package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/onokonem/go-testrest/proto/models"
)

// FindAccountByNameOKCode is the HTTP code returned for type FindAccountByNameOK
const FindAccountByNameOKCode int = 200

/*FindAccountByNameOK successful operation

swagger:response findAccountByNameOK
*/
type FindAccountByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewFindAccountByNameOK creates FindAccountByNameOK with default headers values
func NewFindAccountByNameOK() *FindAccountByNameOK {
	return &FindAccountByNameOK{}
}

// WithPayload adds the payload to the find account by name o k response
func (o *FindAccountByNameOK) WithPayload(payload *models.Account) *FindAccountByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find account by name o k response
func (o *FindAccountByNameOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindAccountByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindAccountByNameBadRequestCode is the HTTP code returned for type FindAccountByNameBadRequest
const FindAccountByNameBadRequestCode int = 400

/*FindAccountByNameBadRequest Invalid status value

swagger:response findAccountByNameBadRequest
*/
type FindAccountByNameBadRequest struct {
}

// NewFindAccountByNameBadRequest creates FindAccountByNameBadRequest with default headers values
func NewFindAccountByNameBadRequest() *FindAccountByNameBadRequest {
	return &FindAccountByNameBadRequest{}
}

// WriteResponse to the client
func (o *FindAccountByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}
