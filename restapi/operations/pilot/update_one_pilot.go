package pilot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/dominictracey/rugby-scores/models"
)

// UpdateOnePilotHandlerFunc turns a function with the right signature into a update one pilot handler
type UpdateOnePilotHandlerFunc func(UpdateOnePilotParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateOnePilotHandlerFunc) Handle(params UpdateOnePilotParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateOnePilotHandler interface for that can handle valid update one pilot params
type UpdateOnePilotHandler interface {
	Handle(UpdateOnePilotParams, *models.Principal) middleware.Responder
}

// NewUpdateOnePilot creates a new http.Handler for the update one pilot operation
func NewUpdateOnePilot(ctx *middleware.Context, handler UpdateOnePilotHandler) *UpdateOnePilot {
	return &UpdateOnePilot{Context: ctx, Handler: handler}
}

/*UpdateOnePilot swagger:route PUT /pilot/{id} pilot updateOnePilot

UpdateOnePilot update one pilot API

*/
type UpdateOnePilot struct {
	Context *middleware.Context
	Handler UpdateOnePilotHandler
}

func (o *UpdateOnePilot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUpdateOnePilotParams()

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
