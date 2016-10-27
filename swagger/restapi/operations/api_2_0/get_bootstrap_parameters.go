package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBootstrapParams creates a new GetBootstrapParams object
// with the default values initialized.
func NewGetBootstrapParams() GetBootstrapParams {
	var ()
	return GetBootstrapParams{}
}

// GetBootstrapParams contains all the bound params for the get bootstrap operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBootstrap
type GetBootstrapParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Query string containing the mac address
	  In: query
	*/
	MacAddress *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetBootstrapParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMacAddress, qhkMacAddress, _ := qs.GetOK("macAddress")
	if err := o.bindMacAddress(qMacAddress, qhkMacAddress, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBootstrapParams) bindMacAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MacAddress = &raw

	return nil
}
