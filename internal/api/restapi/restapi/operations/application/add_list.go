// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddListHandlerFunc turns a function with the right signature into a add list handler
type AddListHandlerFunc func(AddListParams) AddListResponder

// Handle executing the request and returning a response
func (fn AddListHandlerFunc) Handle(params AddListParams) AddListResponder {
	return fn(params)
}

// AddListHandler interface for that can handle valid add list params
type AddListHandler interface {
	Handle(AddListParams) AddListResponder
}

// NewAddList creates a new http.Handler for the add list operation
func NewAddList(ctx *middleware.Context, handler AddListHandler) *AddList {
	return &AddList{Context: ctx, Handler: handler}
}

/*
	AddList swagger:route POST /application Application addList

Add new and update existing applications
*/
type AddList struct {
	Context *middleware.Context
	Handler AddListHandler
}

func (o *AddList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
