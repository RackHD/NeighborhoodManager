package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*CatalogsIDGetOK Successfully retrieved the catalog

swagger:response catalogsIdGetOK
*/
type CatalogsIDGetOK struct {

	// In: body
	Payload *models.VersionsResponse `json:"body,omitempty"`
}

// NewCatalogsIDGetOK creates CatalogsIDGetOK with default headers values
func NewCatalogsIDGetOK() *CatalogsIDGetOK {
	return &CatalogsIDGetOK{}
}

// WithPayload adds the payload to the catalogs Id get o k response
func (o *CatalogsIDGetOK) WithPayload(payload *models.VersionsResponse) *CatalogsIDGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalogs Id get o k response
func (o *CatalogsIDGetOK) SetPayload(payload *models.VersionsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogsIDGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CatalogsIDGetDefault Error

swagger:response catalogsIdGetDefault
*/
type CatalogsIDGetDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCatalogsIDGetDefault creates CatalogsIDGetDefault with default headers values
func NewCatalogsIDGetDefault(code int) *CatalogsIDGetDefault {
	if code <= 0 {
		code = 500
	}

	return &CatalogsIDGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the catalogs Id get default response
func (o *CatalogsIDGetDefault) WithStatusCode(code int) *CatalogsIDGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the catalogs Id get default response
func (o *CatalogsIDGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the catalogs Id get default response
func (o *CatalogsIDGetDefault) WithPayload(payload *models.ErrorResponse) *CatalogsIDGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the catalogs Id get default response
func (o *CatalogsIDGetDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CatalogsIDGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
