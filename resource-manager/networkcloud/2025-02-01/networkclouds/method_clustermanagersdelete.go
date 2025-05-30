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

type ClusterManagersDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type ClusterManagersDeleteOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultClusterManagersDeleteOperationOptions() ClusterManagersDeleteOperationOptions {
	return ClusterManagersDeleteOperationOptions{}
}

func (o ClusterManagersDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o ClusterManagersDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ClusterManagersDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ClusterManagersDelete ...
func (c NetworkcloudsClient) ClusterManagersDelete(ctx context.Context, id ClusterManagerId, options ClusterManagersDeleteOperationOptions) (result ClusterManagersDeleteOperationResponse, err error) {
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

// ClusterManagersDeleteThenPoll performs ClusterManagersDelete then polls until it's completed
func (c NetworkcloudsClient) ClusterManagersDeleteThenPoll(ctx context.Context, id ClusterManagerId, options ClusterManagersDeleteOperationOptions) error {
	result, err := c.ClusterManagersDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ClusterManagersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ClusterManagersDelete: %+v", err)
	}

	return nil
}
