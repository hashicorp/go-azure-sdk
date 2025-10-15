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

type MetricsConfigurationsUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ClusterMetricsConfiguration
}

type MetricsConfigurationsUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultMetricsConfigurationsUpdateOperationOptions() MetricsConfigurationsUpdateOperationOptions {
	return MetricsConfigurationsUpdateOperationOptions{}
}

func (o MetricsConfigurationsUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o MetricsConfigurationsUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o MetricsConfigurationsUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MetricsConfigurationsUpdate ...
func (c NetworkcloudsClient) MetricsConfigurationsUpdate(ctx context.Context, id MetricsConfigurationId, input ClusterMetricsConfigurationPatchParameters, options MetricsConfigurationsUpdateOperationOptions) (result MetricsConfigurationsUpdateOperationResponse, err error) {
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

// MetricsConfigurationsUpdateThenPoll performs MetricsConfigurationsUpdate then polls until it's completed
func (c NetworkcloudsClient) MetricsConfigurationsUpdateThenPoll(ctx context.Context, id MetricsConfigurationId, input ClusterMetricsConfigurationPatchParameters, options MetricsConfigurationsUpdateOperationOptions) error {
	result, err := c.MetricsConfigurationsUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing MetricsConfigurationsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after MetricsConfigurationsUpdate: %+v", err)
	}

	return nil
}
