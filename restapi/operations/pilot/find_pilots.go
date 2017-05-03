package pilot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/dominictracey/rugby-scores/models"
)

// FindPilotsHandlerFunc turns a function with the right signature into a find pilots handler
type FindPilotsHandlerFunc func(FindPilotsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn FindPilotsHandlerFunc) Handle(params FindPilotsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// FindPilotsHandler interface for that can handle valid find pilots params
type FindPilotsHandler interface {
	Handle(FindPilotsParams, *models.Principal) middleware.Responder
}

// NewFindPilots creates a new http.Handler for the find pilots operation
func NewFindPilots(ctx *middleware.Context, handler FindPilotsHandler) *FindPilots {
	return &FindPilots{Context: ctx, Handler: handler}
}

/*FindPilots swagger:route GET /pilot pilot findPilots

FindPilots find pilots API

*/
type FindPilots struct {
	Context *middleware.Context
	Handler FindPilotsHandler
}

func (o *FindPilots) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindPilotsParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}