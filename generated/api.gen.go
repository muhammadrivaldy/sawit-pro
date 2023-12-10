// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// ResponseErrorBadRequest defines model for responseErrorBadRequest.
type ResponseErrorBadRequest struct {
	Code    *int                    `json:"code,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message *string                 `json:"message,omitempty"`
}

// ResponseErrorConflict defines model for responseErrorConflict.
type ResponseErrorConflict struct {
	Code    *int                    `json:"code,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message *string                 `json:"message,omitempty"`
}

// ResponseErrorSystem defines model for responseErrorSystem.
type ResponseErrorSystem struct {
	Code    *int                    `json:"code,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message *string                 `json:"message,omitempty"`
}

// ResponseRegistration defines model for responseRegistration.
type ResponseRegistration struct {
	Code *int `json:"code,omitempty"`
	Data *struct {
		FullName    string `json:"full_name"`
		Id          int    `json:"id"`
		PhoneNumber string `json:"phone_number"`
	} `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
}

// PostUsersFormdataBody defines parameters for PostUsers.
type PostUsersFormdataBody struct {
	FullName    string `form:"full_name" json:"full_name"`
	Password    string `form:"password" json:"password"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}

// PostUsersFormdataRequestBody defines body for PostUsers for application/x-www-form-urlencoded ContentType.
type PostUsersFormdataRequestBody PostUsersFormdataBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /users)
	PostUsers(ctx echo.Context) error

	// (POST /users/login)
	PostUsersLogin(ctx echo.Context) error

	// (GET /users/{id})
	GetUsersId(ctx echo.Context, id int) error

	// (PUT /users/{id})
	PutUsersId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// PostUsersLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersLogin(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersLogin(ctx)
	return err
}

// GetUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersId(ctx, id)
	return err
}

// PutUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutUsersId(ctx, id)
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

	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.POST(baseURL+"/users/login", wrapper.PostUsersLogin)
	router.GET(baseURL+"/users/:id", wrapper.GetUsersId)
	router.PUT(baseURL+"/users/:id", wrapper.PutUsersId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yWYW/yNhDHv0rk7d3SJwGeZwXeja6aKm1SVdZXCFXGPhJXie3ZFyhC+e6THWiTkMJg",
	"TNqL5124+C53v/P9jy1hKtdKgkRLxlti4K8CLE4UF1A3PEEiLBqKQklnZkoiSHSPVOtMMP8mertZr9c3",
	"S2Xym8JkIJniwN0hy1LIqXvSRmkwuAu/LLLsRdIc3A94o7nOgIzJk1jRjG+C6WRKQoIb7YwWjZAJKUOi",
	"qbVrZXjTKd887uyjUdzplioJL7LIF2Carj/93B/e3g77g2FvOOyPDp3L0JMQxpUza0YKa1XUcpu/B1GL",
	"V2BIymYUNAVUFquVtHvc1Y97Y5SZUP5U4T+C/NVWHfmMsGtBo9ivcfyemZAICRgHh1P03q2cQ5KDtTRp",
	"NegxA2oh0AUGTBkDDIMVzQroJNeBgYNlRujqOpEJ5cHuopEaEQ/hTsllJtiVEYz+PYJfKdKA7bO7qO5m",
	"iHbl041FyK9Z97drtH6qcsBUyCRYpyCDtVEyuaz8ZwlvGhgCD8AVXCfwDwXnbAT94wiuoU6iqUu9/qDr",
	"i6fEKI5Hvf6gF5/UIsHbClSPfKhCn7TVFoyBtZd1crpzdm8ssMII3ExdSyqOC2oF+6XA9L1VLpy3fnwv",
	"RdQuuwVQA+bwtDe3j7sPCrlU7mgmGEjry6p6Rv54+NNFRIG+xGcLxgZTMCvBHKgVGFvl3/sSf4ndUaVB",
	"Ui3ImAy8yek5pr6IqHDu/pKoSpGbDO4l10pIDJbKBKZ+fX3Y6vmBO/FUFn0uJKwt3I0L+aOBJRmTH6KP",
	"tRw1dnLUtZDbW6Qfx58H252LOietDMnXc5zbe8r7j870v6sp4Lezv76TSX9PaWLdSFRw585UdS3KVCLk",
	"VXv3u4/YDf4/FatL55GMZ41JnM3L+TFkW8FLl00Cp4AlgIHzOYD1G1SsHrgfJENzQD9Dsy1x7fDDRcL9",
	"uHola/5BCmuwjiuqK+Z/34s27ZDo4hTdQnOK0A34sfgO+CjgagDMag+lMNlud4yjKFOMZqkThHJe/h0A",
	"AP//tptqcgoNAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
