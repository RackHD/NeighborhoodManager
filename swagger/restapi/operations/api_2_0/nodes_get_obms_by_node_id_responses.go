package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*NodesGetObmsByNodeIDOK Successfully retrieved the specified OBM service

swagger:response nodesGetObmsByNodeIdOK
*/
type NodesGetObmsByNodeIDOK struct {

	// In: body
	Payload NodesGetObmsByNodeIDOKBody `json:"body,omitempty"`
}

// NewNodesGetObmsByNodeIDOK creates NodesGetObmsByNodeIDOK with default headers values
func NewNodesGetObmsByNodeIDOK() *NodesGetObmsByNodeIDOK {
	return &NodesGetObmsByNodeIDOK{}
}

// WithPayload adds the payload to the nodes get obms by node Id o k response
func (o *NodesGetObmsByNodeIDOK) WithPayload(payload NodesGetObmsByNodeIDOKBody) *NodesGetObmsByNodeIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get obms by node Id o k response
func (o *NodesGetObmsByNodeIDOK) SetPayload(payload NodesGetObmsByNodeIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetObmsByNodeIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*NodesGetObmsByNodeIDDefault Unexpected error

swagger:response nodesGetObmsByNodeIdDefault
*/
type NodesGetObmsByNodeIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesGetObmsByNodeIDDefault creates NodesGetObmsByNodeIDDefault with default headers values
func NewNodesGetObmsByNodeIDDefault(code int) *NodesGetObmsByNodeIDDefault {
	if code <= 0 {
		code = 500
	}

	return &NodesGetObmsByNodeIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the nodes get obms by node Id default response
func (o *NodesGetObmsByNodeIDDefault) WithStatusCode(code int) *NodesGetObmsByNodeIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the nodes get obms by node Id default response
func (o *NodesGetObmsByNodeIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the nodes get obms by node Id default response
func (o *NodesGetObmsByNodeIDDefault) WithPayload(payload *models.Error) *NodesGetObmsByNodeIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get obms by node Id default response
func (o *NodesGetObmsByNodeIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetObmsByNodeIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
