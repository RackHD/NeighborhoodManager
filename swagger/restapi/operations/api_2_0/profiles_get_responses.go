package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*ProfilesGetOK Successfully retrieved a list of profiles for specified mac / ip

swagger:response profilesGetOK
*/
type ProfilesGetOK struct {

	// In: body
	Payload ProfilesGetOKBody `json:"body,omitempty"`
}

// NewProfilesGetOK creates ProfilesGetOK with default headers values
func NewProfilesGetOK() *ProfilesGetOK {
	return &ProfilesGetOK{}
}

// WithPayload adds the payload to the profiles get o k response
func (o *ProfilesGetOK) WithPayload(payload ProfilesGetOKBody) *ProfilesGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profiles get o k response
func (o *ProfilesGetOK) SetPayload(payload ProfilesGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfilesGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ProfilesGetDefault Unexpected error

swagger:response profilesGetDefault
*/
type ProfilesGetDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewProfilesGetDefault creates ProfilesGetDefault with default headers values
func NewProfilesGetDefault(code int) *ProfilesGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ProfilesGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the profiles get default response
func (o *ProfilesGetDefault) WithStatusCode(code int) *ProfilesGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the profiles get default response
func (o *ProfilesGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the profiles get default response
func (o *ProfilesGetDefault) WithPayload(payload *models.Error) *ProfilesGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profiles get default response
func (o *ProfilesGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfilesGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
