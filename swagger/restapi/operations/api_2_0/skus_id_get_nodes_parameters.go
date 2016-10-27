package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSkusIDGetNodesParams creates a new SkusIDGetNodesParams object
// with the default values initialized.
func NewSkusIDGetNodesParams() SkusIDGetNodesParams {
	var ()
	return SkusIDGetNodesParams{}
}

// SkusIDGetNodesParams contains all the bound params for the skus Id get nodes operation
// typically these are obtained from a http.Request
//
// swagger:parameters skusIdGetNodes
type SkusIDGetNodesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The SKU identifier
	  Required: true
	  In: path
	*/
	Identifier string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *SkusIDGetNodesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rIdentifier, rhkIdentifier, _ := route.Params.GetOK("identifier")
	if err := o.bindIdentifier(rIdentifier, rhkIdentifier, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SkusIDGetNodesParams) bindIdentifier(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Identifier = raw

	return nil
}
