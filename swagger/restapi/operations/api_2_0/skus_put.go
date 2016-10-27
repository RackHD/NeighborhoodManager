package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SkusPutHandlerFunc turns a function with the right signature into a skus put handler
type SkusPutHandlerFunc func(SkusPutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SkusPutHandlerFunc) Handle(params SkusPutParams) middleware.Responder {
	return fn(params)
}

// SkusPutHandler interface for that can handle valid skus put params
type SkusPutHandler interface {
	Handle(SkusPutParams) middleware.Responder
}

// NewSkusPut creates a new http.Handler for the skus put operation
func NewSkusPut(ctx *middleware.Context, handler SkusPutHandler) *SkusPut {
	return &SkusPut{Context: ctx, Handler: handler}
}

/*SkusPut swagger:route PUT /skus /api/2.0 skusPut

Put a SKU

Create or modify a SKU.

*/
type SkusPut struct {
	Context *middleware.Context
	Handler SkusPutHandler
}

func (o *SkusPut) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewSkusPutParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SkusPutCreatedBody skus put created body
// swagger:model SkusPutCreatedBody
type SkusPutCreatedBody interface{}
