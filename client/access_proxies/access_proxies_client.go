// Code generated by go-swagger; DO NOT EDIT.

package access_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new access proxies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access proxies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateProxy creates proxy
*/
func (a *Client) CreateProxy(params *CreateProxyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProxyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProxyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProxy",
		Method:             "POST",
		PathPattern:        "/access_proxies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProxyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProxy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteProxy deletes proxy
*/
func (a *Client) DeleteProxy(params *DeleteProxyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProxyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProxyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProxy",
		Method:             "DELETE",
		PathPattern:        "/access_proxies/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProxyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProxy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EditProxy edits proxy
*/
func (a *Client) EditProxy(params *EditProxyParams, authInfo runtime.ClientAuthInfoWriter) (*EditProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditProxyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editProxy",
		Method:             "PATCH",
		PathPattern:        "/access_proxies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EditProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editProxy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProxy retrieves information about a proxy
*/
func (a *Client) GetProxy(params *GetProxyParams, authInfo runtime.ClientAuthInfoWriter) (*GetProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProxyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProxy",
		Method:             "GET",
		PathPattern:        "/access_proxies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProxyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProxy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListProxies lists proxies
*/
func (a *Client) ListProxies(params *ListProxiesParams, authInfo runtime.ClientAuthInfoWriter) (*ListProxiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProxiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProxies",
		Method:             "GET",
		PathPattern:        "/access_proxies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProxiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProxiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProxies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}