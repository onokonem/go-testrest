package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/onokonem/go-testrest/proto/models"
)

// GetAccountByIDOKCode is the HTTP code returned for type GetAccountByIDOK
const GetAccountByIDOKCode int = 200

/*GetAccountByIDOK successful operation

swagger:response getAccountByIdOK
*/
type GetAccountByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewGetAccountByIDOK creates GetAccountByIDOK with default headers values
func NewGetAccountByIDOK() *GetAccountByIDOK {
	return &GetAccountByIDOK{}
}

// WithPayload adds the payload to the get account by Id o k response
func (o *GetAccountByIDOK) WithPayload(payload *models.Account) *GetAccountByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account by Id o k response
func (o *GetAccountByIDOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountByIDBadRequestCode is the HTTP code returned for type GetAccountByIDBadRequest
const GetAccountByIDBadRequestCode int = 400

/*GetAccountByIDBadRequest Invalid ID supplied

swagger:response getAccountByIdBadRequest
*/
type GetAccountByIDBadRequest struct {
}

// NewGetAccountByIDBadRequest creates GetAccountByIDBadRequest with default headers values
func NewGetAccountByIDBadRequest() *GetAccountByIDBadRequest {
	return &GetAccountByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// GetAccountByIDNotFoundCode is the HTTP code returned for type GetAccountByIDNotFound
const GetAccountByIDNotFoundCode int = 404

/*GetAccountByIDNotFound Pet not found

swagger:response getAccountByIdNotFound
*/
type GetAccountByIDNotFound struct {
}

// NewGetAccountByIDNotFound creates GetAccountByIDNotFound with default headers values
func NewGetAccountByIDNotFound() *GetAccountByIDNotFound {
	return &GetAccountByIDNotFound{}
}

// WriteResponse to the client
func (o *GetAccountByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
