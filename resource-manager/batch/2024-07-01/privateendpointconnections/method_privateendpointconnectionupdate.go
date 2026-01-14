package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PrivateEndpointConnection
}

type PrivateEndpointConnectionUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultPrivateEndpointConnectionUpdateOperationOptions() PrivateEndpointConnectionUpdateOperationOptions {
	return PrivateEndpointConnectionUpdateOperationOptions{}
}

func (o PrivateEndpointConnectionUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o PrivateEndpointConnectionUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PrivateEndpointConnectionUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PrivateEndpointConnectionUpdate ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection, options PrivateEndpointConnectionUpdateOperationOptions) (result PrivateEndpointConnectionUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// PrivateEndpointConnectionUpdateThenPoll performs PrivateEndpointConnectionUpdate then polls until it's completed
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionUpdateThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection, options PrivateEndpointConnectionUpdateOperationOptions) error {
	result, err := c.PrivateEndpointConnectionUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionUpdate: %+v", err)
	}

	return nil
}
