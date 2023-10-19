// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClusterNodesHandlerFunc turns a function with the right signature into a get cluster nodes handler
type GetClusterNodesHandlerFunc func(GetClusterNodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterNodesHandlerFunc) Handle(params GetClusterNodesParams) middleware.Responder {
	return fn(params)
}

// GetClusterNodesHandler interface for that can handle valid get cluster nodes params
type GetClusterNodesHandler interface {
	Handle(GetClusterNodesParams) middleware.Responder
}

// NewGetClusterNodes creates a new http.Handler for the get cluster nodes operation
func NewGetClusterNodes(ctx *middleware.Context, handler GetClusterNodesHandler) *GetClusterNodes {
	return &GetClusterNodes{Context: ctx, Handler: handler}
}

/*
	GetClusterNodes swagger:route GET /cluster/nodes daemon getClusterNodes

Get nodes information stored in the cilium-agent
*/
type GetClusterNodes struct {
	Context *middleware.Context
	Handler GetClusterNodesHandler
}

func (o *GetClusterNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClusterNodesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
