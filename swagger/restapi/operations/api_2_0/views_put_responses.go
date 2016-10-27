package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*ViewsPutCreated Successfully created the specified template

swagger:response viewsPutCreated
*/
type ViewsPutCreated struct {

	// In: body
	Payload ViewsPutCreatedBody `json:"body,omitempty"`
}

// NewViewsPutCreated creates ViewsPutCreated with default headers values
func NewViewsPutCreated() *ViewsPutCreated {
	return &ViewsPutCreated{}
}

// WithPayload adds the payload to the views put created response
func (o *ViewsPutCreated) WithPayload(payload ViewsPutCreatedBody) *ViewsPutCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the views put created response
func (o *ViewsPutCreated) SetPayload(payload ViewsPutCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ViewsPutCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
