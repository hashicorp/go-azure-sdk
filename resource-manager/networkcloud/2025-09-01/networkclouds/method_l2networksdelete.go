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

type L2NetworksDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type L2NetworksDeleteOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultL2NetworksDeleteOperationOptions() L2NetworksDeleteOperationOptions {
	return L2NetworksDeleteOperationOptions{}
}

func (o L2NetworksDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o L2NetworksDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o L2NetworksDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// L2NetworksDelete ...
func (c NetworkcloudsClient) L2NetworksDelete(ctx context.Context, id L2NetworkId, options L2NetworksDeleteOperationOptions) (result L2NetworksDeleteOperationResponse, err error) {
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

// L2NetworksDeleteThenPoll performs L2NetworksDelete then polls until it's completed
func (c NetworkcloudsClient) L2NetworksDeleteThenPoll(ctx context.Context, id L2NetworkId, options L2NetworksDeleteOperationOptions) error {
	result, err := c.L2NetworksDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing L2NetworksDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after L2NetworksDelete: %+v", err)
	}

	return nil
}
