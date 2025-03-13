package networkclouds

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

type L3NetworksDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type L3NetworksDeleteOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultL3NetworksDeleteOperationOptions() L3NetworksDeleteOperationOptions {
	return L3NetworksDeleteOperationOptions{}
}

func (o L3NetworksDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o L3NetworksDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o L3NetworksDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// L3NetworksDelete ...
func (c NetworkcloudsClient) L3NetworksDelete(ctx context.Context, id L3NetworkId, options L3NetworksDeleteOperationOptions) (result L3NetworksDeleteOperationResponse, err error) {
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

// L3NetworksDeleteThenPoll performs L3NetworksDelete then polls until it's completed
func (c NetworkcloudsClient) L3NetworksDeleteThenPoll(ctx context.Context, id L3NetworkId, options L3NetworksDeleteOperationOptions) error {
	result, err := c.L3NetworksDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing L3NetworksDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after L3NetworksDelete: %+v", err)
	}

	return nil
}
