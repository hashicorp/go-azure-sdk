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

type ConnectionRaiPolicyDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type ConnectionRaiPolicyDeleteOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiPolicyDeleteOperationOptions() ConnectionRaiPolicyDeleteOperationOptions {
	return ConnectionRaiPolicyDeleteOperationOptions{}
}

func (o ConnectionRaiPolicyDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiPolicyDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiPolicyDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

// ConnectionRaiPolicyDelete ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiPolicyDelete(ctx context.Context, id ConnectionRaiPolicyId, options ConnectionRaiPolicyDeleteOperationOptions) (result ConnectionRaiPolicyDeleteOperationResponse, err error) {
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

// ConnectionRaiPolicyDeleteThenPoll performs ConnectionRaiPolicyDelete then polls until it's completed
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiPolicyDeleteThenPoll(ctx context.Context, id ConnectionRaiPolicyId, options ConnectionRaiPolicyDeleteOperationOptions) error {
	result, err := c.ConnectionRaiPolicyDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ConnectionRaiPolicyDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConnectionRaiPolicyDelete: %+v", err)
	}

	return nil
}
