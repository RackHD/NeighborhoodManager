package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*NodesPostCreated Successfully created node

swagger:response nodesPostCreated
*/
type NodesPostCreated struct {

	// In: body
	Payload NodesPostCreatedBody `json:"body,omitempty"`
}

// NewNodesPostCreated creates NodesPostCreated with default headers values
func NewNodesPostCreated() *NodesPostCreated {
	return &NodesPostCreated{}
}

// WithPayload adds the payload to the nodes post created response
func (o *NodesPostCreated) WithPayload(payload NodesPostCreatedBody) *NodesPostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes post created response
func (o *NodesPostCreated) SetPayload(payload NodesPostCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesPostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*NodesPostDefault Unexpected error

swagger:response nodesPostDefault
*/
type NodesPostDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesPostDefault creates NodesPostDefault with default headers values
func NewNodesPostDefault(code int) *NodesPostDefault {
	if code <= 0 {
		code = 500
	}

	return &NodesPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the nodes post default response
func (o *NodesPostDefault) WithStatusCode(code int) *NodesPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the nodes post default response
func (o *NodesPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the nodes post default response
func (o *NodesPostDefault) WithPayload(payload *models.Error) *NodesPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes post default response
func (o *NodesPostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
