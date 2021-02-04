// Package controllers provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package controllers

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get state changes for requesting account
	// (GET /changes)
	GetChanges(ctx echo.Context, params GetChangesParams) error
	// Get single state change for requesting account
	// (GET /changes/{id})
	GetChangesId(ctx echo.Context, id StateIDParam) error
	// Get a list of runs for each state change
	// (GET /runs)
	GetRuns(ctx echo.Context, params GetRunsParams) error
	// Generate new runs by applying a state change
	// (POST /runs)
	PostRuns(ctx echo.Context) error
	// Get a single run
	// (GET /runs/{id})
	GetRunsId(ctx echo.Context, id RunIDParam) error
	// Get configuration state for requesting account
	// (GET /states)
	GetStates(ctx echo.Context) error
	// Update configuration state for requesting account
	// (POST /states)
	PostStates(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetChanges converts echo context to params.
func (w *ServerInterfaceWrapper) GetChanges(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChangesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetChanges(ctx, params)
	return err
}

// GetChangesId converts echo context to params.
func (w *ServerInterfaceWrapper) GetChangesId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id StateIDParam

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetChangesId(ctx, id)
	return err
}

// GetRuns converts echo context to params.
func (w *ServerInterfaceWrapper) GetRuns(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRunsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Optional query parameter "sort_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_by", ctx.QueryParams(), &params.SortBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort_by: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRuns(ctx, params)
	return err
}

// PostRuns converts echo context to params.
func (w *ServerInterfaceWrapper) PostRuns(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRuns(ctx)
	return err
}

// GetRunsId converts echo context to params.
func (w *ServerInterfaceWrapper) GetRunsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id RunIDParam

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRunsId(ctx, id)
	return err
}

// GetStates converts echo context to params.
func (w *ServerInterfaceWrapper) GetStates(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStates(ctx)
	return err
}

// PostStates converts echo context to params.
func (w *ServerInterfaceWrapper) PostStates(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostStates(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/changes", wrapper.GetChanges)
	router.GET(baseURL+"/changes/:id", wrapper.GetChangesId)
	router.GET(baseURL+"/runs", wrapper.GetRuns)
	router.POST(baseURL+"/runs", wrapper.PostRuns)
	router.GET(baseURL+"/runs/:id", wrapper.GetRunsId)
	router.GET(baseURL+"/states", wrapper.GetStates)
	router.POST(baseURL+"/states", wrapper.PostStates)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xY227jNhN+FYL/f6mNnG33Rnd76MFoig3i9moRLGhpZHEhkQo5TNcI9O7FkJItRbIj",
	"Z3PonWnOjL755sAh73iqq1orUGh5csdrYUQFCMavLmQlkX5kYFMja5Ra8YT/Kb7LylVMuWoNhumcGbCu",
	"RMtQMwPojOIRlyR648BsecSVqIAnvPQGI27TAioRLOfClciTd4uIV8EwT94uaCVVWJ1HHLc16UuFsAHD",
	"mybin/PcwgS6pcpkKhAswwKYRWFQqg2rtZUkQXBpwyNjBkqB8hYIOf1LbJSAwCwgSUqEigwJZJXAtNir",
	"HvBQB1STLvZ9Wkz6dOXU8tMlxWDsl0WBwIRJC8IrM1AocwmmA1ILLPY4ZMYjbuDGSQMZT9A46GP6v4Gc",
	"J/x/8T78cdi1sQfRwbG/yhLBkM6Uu3nYnWv6Qqyh3JleaYMftodMW23w63o7sA2KuPvCUwMCIfsqiOn9",
	"IhE2Hf5BFPLrHdcWjVQbD2BFdL4+2S0M3hCm9k/SeZ+m2imf3pX4fgFqgwVPzkMK7ZYjx6JO0dv1FW10",
	"DQYleLNib/YYqu7rTUS+zXUh4qWP77wkiALJs6z7iLWu6vU3SPG+q0tVO3wKf38U1O/aYkiLu3Fwlkqi",
	"FKjN5O5Fx95o58qpp/CtVzYPKP0lK7AoqprUip5Lx5R2rs9Km7bNUEH1SDmmsWfvUanm7AxIqyDYRNzV",
	"2elkTWXElVMfA/HjGKY6g+nu4yyjzXC63jiw1OruHxmPic0BiMtPZCLXpiKHuXO+qU1l4mrHZdeOjVOK",
	"9iNuXZqCtTziuZClMzDRfEPz922fTtcZQeF7zMIYse1s9Gida6pTmbC4a5kiy/yoIMrLQbBGfox49Cbe",
	"h1PjFQv2pJ79UuX3I021z+v8zBlE41DAZ+b9nt++NHWINygrGKs0ntpcj6v7o1a53LBKKLEBwyyYW5l6",
	"CxJLGAnwiN+CsUF3cbY4Oyc8ugYlaskT/pP/K/IjiackTguhNoGnTZiQKQsFfX6Z8YT/BvixFYkGE/+X",
	"aT73InG4ETTRg4LtdN5c02Bka61sAPR2sQh9TyGEUhB1XdLALrWKv1ntD7oTZqddVnjGh0x//oOo+jl8",
	"crj1QWTsqu2rfvRyVSXMNrDDwvjX8shybboeTDeJroxJreM6vpNZM4PwZXYy5YNB9cX4fFo6pdqUMGD1",
	"KKmmPR8OkenPj1N57F1mZuRv734yQ/o/VRaenacMn2CltP46TIHxkQORFoN4kula24l4XWrbBayN9wed",
	"bZ/M28F9p/FuPyuzuwFimuB75CliApiCfwJ36y2jz299yt8jsMv8B3sJwXhEI+k9Ljx3/s1kh1Kr7Q3G",
	"K0U89pwcrf5VkHhGD+7n1HQlvV2cv9gXe2PrwwVMMudjmfcOC1DYAmSVtMQ9y42udleMJuLvpuwvFYJR",
	"omQrMLdg2C/GaDMRz9SPLi4ErM3vQ53+WMPoxfh5W0Z4N5hg3O+yStTh9U8BZP59cw2svRmO3nyaV0/J",
	"Obnx6Pj+7d0+KcTegLcYOpQzJU94gVjbJI7TUrvszEBWCDxLdRWLWsbB/Jt2+I1vabQdYq2NzlzqF3Ti",
	"BpszdXfvvf4Br7lu/g0AAP//FKMrFQEXAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

