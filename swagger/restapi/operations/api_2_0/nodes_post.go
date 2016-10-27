package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesPostHandlerFunc turns a function with the right signature into a nodes post handler
type NodesPostHandlerFunc func(NodesPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesPostHandlerFunc) Handle(params NodesPostParams) middleware.Responder {
	return fn(params)
}

// NodesPostHandler interface for that can handle valid nodes post params
type NodesPostHandler interface {
	Handle(NodesPostParams) middleware.Responder
}

// NewNodesPost creates a new http.Handler for the nodes post operation
func NewNodesPost(ctx *middleware.Context, handler NodesPostHandler) *NodesPost {
	return &NodesPost{Context: ctx, Handler: handler}
}

/*NodesPost swagger:route POST /nodes /api/2.0 nodesPost

Post a node

Create and store a new node manually.

*/
type NodesPost struct {
	Context *middleware.Context
	Handler NodesPostHandler
}

func (o *NodesPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesPostCreatedBody nodes post created body
// swagger:model NodesPostCreatedBody
type NodesPostCreatedBody interface{}
