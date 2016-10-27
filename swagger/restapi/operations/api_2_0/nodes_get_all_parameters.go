package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodesGetAllParams creates a new NodesGetAllParams object
// with the default values initialized.
func NewNodesGetAllParams() NodesGetAllParams {
	var ()
	return NodesGetAllParams{}
}

// NodesGetAllParams contains all the bound params for the nodes get all operation
// typically these are obtained from a http.Request
//
// swagger:parameters nodesGetAll
type NodesGetAllParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Query string specifying properties to search for
	  In: query
	*/
	DollarSkip *int64
	/*Query string specifying properties to search for
	  In: query
	*/
	DollarTop *int64
	/*Query string specifying properties to sort with
	  Pattern: [- +]{0,1}(name|id)
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NodesGetAllParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qNrDollarSkip, qhkNrDollarSkip, _ := qs.GetOK("$skip")
	if err := o.bindDollarSkip(qNrDollarSkip, qhkNrDollarSkip, route.Formats); err != nil {
		res = append(res, err)
	}

	qNrDollarTop, qhkNrDollarTop, _ := qs.GetOK("$top")
	if err := o.bindDollarTop(qNrDollarTop, qhkNrDollarTop, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NodesGetAllParams) bindDollarSkip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("$skip", "query", "int64", raw)
	}
	o.DollarSkip = &value

	return nil
}

func (o *NodesGetAllParams) bindDollarTop(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("$top", "query", "int64", raw)
	}
	o.DollarTop = &value

	return nil
}

func (o *NodesGetAllParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sort = &raw

	if err := o.validateSort(formats); err != nil {
		return err
	}

	return nil
}

func (o *NodesGetAllParams) validateSort(formats strfmt.Registry) error {

	if err := validate.Pattern("sort", "query", string(*o.Sort), `[- +]{0,1}(name|id)`); err != nil {
		return err
	}

	return nil
}
