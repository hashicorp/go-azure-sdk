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

type ConnectionRaiBlocklistCreateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RaiBlocklistPropertiesBasicResource
}

type ConnectionRaiBlocklistCreateOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiBlocklistCreateOperationOptions() ConnectionRaiBlocklistCreateOperationOptions {
	return ConnectionRaiBlocklistCreateOperationOptions{}
}

func (o ConnectionRaiBlocklistCreateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiBlocklistCreateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiBlocklistCreateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

// ConnectionRaiBlocklistCreate ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistCreate(ctx context.Context, id RaiBlocklistId, input RaiBlocklistPropertiesBasicResource, options ConnectionRaiBlocklistCreateOperationOptions) (result ConnectionRaiBlocklistCreateOperationResponse, err error) {
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

// ConnectionRaiBlocklistCreateThenPoll performs ConnectionRaiBlocklistCreate then polls until it's completed
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistCreateThenPoll(ctx context.Context, id RaiBlocklistId, input RaiBlocklistPropertiesBasicResource, options ConnectionRaiBlocklistCreateOperationOptions) error {
	result, err := c.ConnectionRaiBlocklistCreate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing ConnectionRaiBlocklistCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConnectionRaiBlocklistCreate: %+v", err)
	}

	return nil
}
