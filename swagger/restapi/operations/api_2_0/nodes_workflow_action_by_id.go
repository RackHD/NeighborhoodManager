package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesWorkflowActionByIDHandlerFunc turns a function with the right signature into a nodes workflow action by Id handler
type NodesWorkflowActionByIDHandlerFunc func(NodesWorkflowActionByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesWorkflowActionByIDHandlerFunc) Handle(params NodesWorkflowActionByIDParams) middleware.Responder {
	return fn(params)
}

// NodesWorkflowActionByIDHandler interface for that can handle valid nodes workflow action by Id params
type NodesWorkflowActionByIDHandler interface {
	Handle(NodesWorkflowActionByIDParams) middleware.Responder
}

// NewNodesWorkflowActionByID creates a new http.Handler for the nodes workflow action by Id operation
func NewNodesWorkflowActionByID(ctx *middleware.Context, handler NodesWorkflowActionByIDHandler) *NodesWorkflowActionByID {
	return &NodesWorkflowActionByID{Context: ctx, Handler: handler}
}

/*NodesWorkflowActionByID swagger:route PUT /nodes/{identifier}/workflows/action /api/2.0 nodesWorkflowActionById

Perform an action on a workflow

Perform an action on a workflow associated with a node. Currently, the cancel action is supported.


*/
type NodesWorkflowActionByID struct {
	Context *middleware.Context
	Handler NodesWorkflowActionByIDHandler
}

func (o *NodesWorkflowActionByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesWorkflowActionByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesWorkflowActionByIDAcceptedBody nodes workflow action by ID accepted body
// swagger:model NodesWorkflowActionByIDAcceptedBody
type NodesWorkflowActionByIDAcceptedBody interface{}
