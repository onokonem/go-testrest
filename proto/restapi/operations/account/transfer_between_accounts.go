package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TransferBetweenAccountsHandlerFunc turns a function with the right signature into a transfer between accounts handler
type TransferBetweenAccountsHandlerFunc func(TransferBetweenAccountsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TransferBetweenAccountsHandlerFunc) Handle(params TransferBetweenAccountsParams) middleware.Responder {
	return fn(params)
}

// TransferBetweenAccountsHandler interface for that can handle valid transfer between accounts params
type TransferBetweenAccountsHandler interface {
	Handle(TransferBetweenAccountsParams) middleware.Responder
}

// NewTransferBetweenAccounts creates a new http.Handler for the transfer between accounts operation
func NewTransferBetweenAccounts(ctx *middleware.Context, handler TransferBetweenAccountsHandler) *TransferBetweenAccounts {
	return &TransferBetweenAccounts{Context: ctx, Handler: handler}
}

/*TransferBetweenAccounts swagger:route PUT /account/transfer/{accountID} account transferBetweenAccounts

Transfer some money from some account to some other account

*/
type TransferBetweenAccounts struct {
	Context *middleware.Context
	Handler TransferBetweenAccountsHandler
}

func (o *TransferBetweenAccounts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewTransferBetweenAccountsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
