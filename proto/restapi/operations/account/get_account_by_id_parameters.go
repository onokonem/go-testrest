package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountByIDParams creates a new GetAccountByIDParams object
// with the default values initialized.
func NewGetAccountByIDParams() GetAccountByIDParams {
	var ()
	return GetAccountByIDParams{}
}

// GetAccountByIDParams contains all the bound params for the get account by ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAccountByID
type GetAccountByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of account to return
	  Required: true
	  In: path
	*/
	AccountID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetAccountByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAccountID, rhkAccountID, _ := route.Params.GetOK("accountID")
	if err := o.bindAccountID(rAccountID, rhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAccountByIDParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("accountID", "path", "int64", raw)
	}
	o.AccountID = value

	return nil
}