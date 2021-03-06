// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new subscriptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSubscriptionsID deletes a subscription
*/
func (a *Client) DeleteSubscriptionsID(params *DeleteSubscriptionsIDParams) (*DeleteSubscriptionsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSubscriptionsID",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSubscriptionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSubscriptionsIDNoContent), nil

}

/*
GetSubscriptions lists all subscriptions
*/
func (a *Client) GetSubscriptions(params *GetSubscriptionsParams) (*GetSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSubscriptions",
		Method:             "GET",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSubscriptionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionsOK), nil

}

/*
GetSubscriptionsID fetches subscription
*/
func (a *Client) GetSubscriptionsID(params *GetSubscriptionsIDParams) (*GetSubscriptionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSubscriptionsID",
		Method:             "GET",
		PathPattern:        "/subscriptions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSubscriptionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionsIDOK), nil

}

/*
PatchSubscriptionsID edits subscription details
*/
func (a *Client) PatchSubscriptionsID(params *PatchSubscriptionsIDParams) (*PatchSubscriptionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchSubscriptionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchSubscriptionsID",
		Method:             "PATCH",
		PathPattern:        "/subscriptions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchSubscriptionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchSubscriptionsIDOK), nil

}

/*
PostSubscriptions creates subscription
*/
func (a *Client) PostSubscriptions(params *PostSubscriptionsParams) (*PostSubscriptionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSubscriptions",
		Method:             "POST",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSubscriptionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSubscriptionsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
