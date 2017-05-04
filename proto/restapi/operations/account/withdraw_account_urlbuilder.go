package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// WithdrawAccountURL generates an URL for the withdraw account operation
type WithdrawAccountURL struct {
	AccountID int64

	Amount *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WithdrawAccountURL) WithBasePath(bp string) *WithdrawAccountURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WithdrawAccountURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *WithdrawAccountURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/account/withdrawal/{accountID}"

	accountID := swag.FormatInt64(o.AccountID)
	if accountID != "" {
		_path = strings.Replace(_path, "{accountID}", accountID, -1)
	} else {
		return nil, errors.New("AccountID is required on WithdrawAccountURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/testrest"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var amount string
	if o.Amount != nil {
		amount = swag.FormatInt64(*o.Amount)
	}
	if amount != "" {
		qs.Set("amount", amount)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WithdrawAccountURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WithdrawAccountURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WithdrawAccountURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WithdrawAccountURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WithdrawAccountURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WithdrawAccountURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}