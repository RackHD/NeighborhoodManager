package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

// NewSkusPostParams creates a new SkusPostParams object
// with the default values initialized.
func NewSkusPostParams() SkusPostParams {
	var ()
	return SkusPostParams{}
}

// SkusPostParams contains all the bound params for the skus post operation
// typically these are obtained from a http.Request
//
// swagger:parameters skusPost
type SkusPostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The properties used to define the new SKU
	  Required: true
	  In: body
	*/
	Body *models.Skus20SkusUpsert
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *SkusPostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Skus20SkusUpsert
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
