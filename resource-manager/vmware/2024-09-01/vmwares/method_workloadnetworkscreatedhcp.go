package vmwares

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

type WorkloadNetworksCreateDhcpOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *WorkloadNetworkDhcp
}

// WorkloadNetworksCreateDhcp ...
func (c VMwaresClient) WorkloadNetworksCreateDhcp(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) (result WorkloadNetworksCreateDhcpOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

// WorkloadNetworksCreateDhcpThenPoll performs WorkloadNetworksCreateDhcp then polls until it's completed
func (c VMwaresClient) WorkloadNetworksCreateDhcpThenPoll(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) error {
	return c.WorkloadNetworksCreateDhcpCallbackThenPoll(ctx, id, input, nil)
}

// WorkloadNetworksCreateDhcpCallbackThenPoll performs WorkloadNetworksCreateDhcp, runs the optional callback function, then polls until it's completed
func (c VMwaresClient) WorkloadNetworksCreateDhcpCallbackThenPoll(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp, callback func() error) error {
	result, err := c.WorkloadNetworksCreateDhcp(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkloadNetworksCreateDhcp: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkloadNetworksCreateDhcp: %+v", err)
	}

	return nil
}
