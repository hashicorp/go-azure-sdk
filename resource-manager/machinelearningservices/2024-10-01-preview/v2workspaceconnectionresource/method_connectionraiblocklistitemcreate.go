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

type ConnectionRaiBlocklistItemCreateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RaiBlocklistItemPropertiesBasicResource
}

type ConnectionRaiBlocklistItemCreateOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiBlocklistItemCreateOperationOptions() ConnectionRaiBlocklistItemCreateOperationOptions {
	return ConnectionRaiBlocklistItemCreateOperationOptions{}
}

func (o ConnectionRaiBlocklistItemCreateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiBlocklistItemCreateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiBlocklistItemCreateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

// ConnectionRaiBlocklistItemCreate ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemCreate(ctx context.Context, id RaiBlocklistItemId, input RaiBlocklistItemPropertiesBasicResource, options ConnectionRaiBlocklistItemCreateOperationOptions) (result ConnectionRaiBlocklistItemCreateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
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

// ConnectionRaiBlocklistItemCreateThenPoll performs ConnectionRaiBlocklistItemCreate then polls until it's completed
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemCreateThenPoll(ctx context.Context, id RaiBlocklistItemId, input RaiBlocklistItemPropertiesBasicResource, options ConnectionRaiBlocklistItemCreateOperationOptions) error {
	result, err := c.ConnectionRaiBlocklistItemCreate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing ConnectionRaiBlocklistItemCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConnectionRaiBlocklistItemCreate: %+v", err)
	}

	return nil
}
