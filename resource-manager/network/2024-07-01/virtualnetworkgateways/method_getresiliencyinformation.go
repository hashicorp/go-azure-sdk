package virtualnetworkgateways

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

type GetResiliencyInformationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *GatewayResiliencyInformation
}

type GetResiliencyInformationOperationOptions struct {
	AttemptRefresh *bool
}

func DefaultGetResiliencyInformationOperationOptions() GetResiliencyInformationOperationOptions {
	return GetResiliencyInformationOperationOptions{}
}

func (o GetResiliencyInformationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetResiliencyInformationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetResiliencyInformationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.AttemptRefresh != nil {
		out.Append("attemptRefresh", fmt.Sprintf("%v", *o.AttemptRefresh))
	}
	return &out
}

// GetResiliencyInformation ...
func (c VirtualNetworkGatewaysClient) GetResiliencyInformation(ctx context.Context, id VirtualNetworkGatewayId, options GetResiliencyInformationOperationOptions) (result GetResiliencyInformationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/getResiliencyInformation", id.ID()),
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

// GetResiliencyInformationThenPoll performs GetResiliencyInformation then polls until it's completed
func (c VirtualNetworkGatewaysClient) GetResiliencyInformationThenPoll(ctx context.Context, id VirtualNetworkGatewayId, options GetResiliencyInformationOperationOptions) error {
	result, err := c.GetResiliencyInformation(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing GetResiliencyInformation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GetResiliencyInformation: %+v", err)
	}

	return nil
}
