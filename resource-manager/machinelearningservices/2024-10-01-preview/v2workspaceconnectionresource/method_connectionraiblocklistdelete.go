package v2workspaceconnectionresource

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

type ConnectionRaiBlocklistDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type ConnectionRaiBlocklistDeleteOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiBlocklistDeleteOperationOptions() ConnectionRaiBlocklistDeleteOperationOptions {
	return ConnectionRaiBlocklistDeleteOperationOptions{}
}

func (o ConnectionRaiBlocklistDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiBlocklistDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiBlocklistDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

// ConnectionRaiBlocklistDelete ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistDelete(ctx context.Context, id RaiBlocklistId, options ConnectionRaiBlocklistDeleteOperationOptions) (result ConnectionRaiBlocklistDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

// ConnectionRaiBlocklistDeleteThenPoll performs ConnectionRaiBlocklistDelete then polls until it's completed
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistDeleteThenPoll(ctx context.Context, id RaiBlocklistId, options ConnectionRaiBlocklistDeleteOperationOptions) error {
	result, err := c.ConnectionRaiBlocklistDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ConnectionRaiBlocklistDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConnectionRaiBlocklistDelete: %+v", err)
	}

	return nil
}
