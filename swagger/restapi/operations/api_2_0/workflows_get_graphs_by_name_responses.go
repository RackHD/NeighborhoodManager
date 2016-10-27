package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*WorkflowsGetGraphsByNameOK Successfully retrieved the workflow graph with the specified injectable name

swagger:response workflowsGetGraphsByNameOK
*/
type WorkflowsGetGraphsByNameOK struct {

	// In: body
	Payload WorkflowsGetGraphsByNameOKBody `json:"body,omitempty"`
}

// NewWorkflowsGetGraphsByNameOK creates WorkflowsGetGraphsByNameOK with default headers values
func NewWorkflowsGetGraphsByNameOK() *WorkflowsGetGraphsByNameOK {
	return &WorkflowsGetGraphsByNameOK{}
}

// WithPayload adds the payload to the workflows get graphs by name o k response
func (o *WorkflowsGetGraphsByNameOK) WithPayload(payload WorkflowsGetGraphsByNameOKBody) *WorkflowsGetGraphsByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get graphs by name o k response
func (o *WorkflowsGetGraphsByNameOK) SetPayload(payload WorkflowsGetGraphsByNameOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetGraphsByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*WorkflowsGetGraphsByNameDefault Unexpected error

swagger:response workflowsGetGraphsByNameDefault
*/
type WorkflowsGetGraphsByNameDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsGetGraphsByNameDefault creates WorkflowsGetGraphsByNameDefault with default headers values
func NewWorkflowsGetGraphsByNameDefault(code int) *WorkflowsGetGraphsByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &WorkflowsGetGraphsByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the workflows get graphs by name default response
func (o *WorkflowsGetGraphsByNameDefault) WithStatusCode(code int) *WorkflowsGetGraphsByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the workflows get graphs by name default response
func (o *WorkflowsGetGraphsByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the workflows get graphs by name default response
func (o *WorkflowsGetGraphsByNameDefault) WithPayload(payload *models.Error) *WorkflowsGetGraphsByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get graphs by name default response
func (o *WorkflowsGetGraphsByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetGraphsByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
