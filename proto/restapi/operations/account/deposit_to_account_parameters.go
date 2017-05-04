package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDepositToAccountParams creates a new DepositToAccountParams object
// with the default values initialized.
func NewDepositToAccountParams() DepositToAccountParams {
	var ()
	return DepositToAccountParams{}
}

// DepositToAccountParams contains all the bound params for the deposit to account operation
// typically these are obtained from a http.Request
//
// swagger:parameters depositToAccount
type DepositToAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Account id to update
	  Required: true
	  In: path
	*/
	AccountID int64
	/*Amount to be placed
	  In: query
	*/
	Amount *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DepositToAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAccountID, rhkAccountID, _ := route.Params.GetOK("accountID")
	if err := o.bindAccountID(rAccountID, rhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	qAmount, qhkAmount, _ := qs.GetOK("amount")
	if err := o.bindAmount(qAmount, qhkAmount, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DepositToAccountParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DepositToAccountParams) bindAmount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("amount", "query", "int64", raw)
	}
	o.Amount = &value

	return nil
}