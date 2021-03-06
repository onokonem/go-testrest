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

// TransferBetweenAccountsURL generates an URL for the transfer between accounts operation
type TransferBetweenAccountsURL struct {
	ID int64

	Amount int64
	Target int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TransferBetweenAccountsURL) WithBasePath(bp string) *TransferBetweenAccountsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TransferBetweenAccountsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *TransferBetweenAccountsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/account/transfer/{id}"

	id := swag.FormatInt64(o.ID)
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on TransferBetweenAccountsURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/testrest"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	amount := swag.FormatInt64(o.Amount)
	if amount != "" {
		qs.Set("amount", amount)
	}

	target := swag.FormatInt64(o.Target)
	if target != "" {
		qs.Set("target", target)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *TransferBetweenAccountsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *TransferBetweenAccountsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *TransferBetweenAccountsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on TransferBetweenAccountsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on TransferBetweenAccountsURL")
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
func (o *TransferBetweenAccountsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
