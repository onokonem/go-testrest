package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WithdrawAccountNoContentCode is the HTTP code returned for type WithdrawAccountNoContent
const WithdrawAccountNoContentCode int = 204

/*WithdrawAccountNoContent successful operation

swagger:response withdrawAccountNoContent
*/
type WithdrawAccountNoContent struct {
}

// NewWithdrawAccountNoContent creates WithdrawAccountNoContent with default headers values
func NewWithdrawAccountNoContent() *WithdrawAccountNoContent {
	return &WithdrawAccountNoContent{}
}

// WriteResponse to the client
func (o *WithdrawAccountNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// WithdrawAccountBadRequestCode is the HTTP code returned for type WithdrawAccountBadRequest
const WithdrawAccountBadRequestCode int = 400

/*WithdrawAccountBadRequest Invalid ID supplied

swagger:response withdrawAccountBadRequest
*/
type WithdrawAccountBadRequest struct {
}

// NewWithdrawAccountBadRequest creates WithdrawAccountBadRequest with default headers values
func NewWithdrawAccountBadRequest() *WithdrawAccountBadRequest {
	return &WithdrawAccountBadRequest{}
}

// WriteResponse to the client
func (o *WithdrawAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// WithdrawAccountNotFoundCode is the HTTP code returned for type WithdrawAccountNotFound
const WithdrawAccountNotFoundCode int = 404

/*WithdrawAccountNotFound Account not found

swagger:response withdrawAccountNotFound
*/
type WithdrawAccountNotFound struct {
}

// NewWithdrawAccountNotFound creates WithdrawAccountNotFound with default headers values
func NewWithdrawAccountNotFound() *WithdrawAccountNotFound {
	return &WithdrawAccountNotFound{}
}

// WriteResponse to the client
func (o *WithdrawAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
