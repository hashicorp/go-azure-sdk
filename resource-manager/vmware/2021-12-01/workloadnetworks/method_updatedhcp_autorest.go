package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDhcpOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateDhcp ...
func (c WorkloadNetworksClient) UpdateDhcp(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) (result UpdateDhcpOperationResponse, err error) {
	req, err := c.preparerForUpdateDhcp(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDhcp", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateDhcp(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDhcp", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateDhcpThenPoll performs UpdateDhcp then polls until it's completed
func (c WorkloadNetworksClient) UpdateDhcpThenPoll(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) error {
	result, err := c.UpdateDhcp(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateDhcp: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateDhcp: %+v", err)
	}

	return nil
}

// preparerForUpdateDhcp prepares the UpdateDhcp request.
func (c WorkloadNetworksClient) preparerForUpdateDhcp(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateDhcp sends the UpdateDhcp request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdateDhcp(ctx context.Context, req *http.Request) (future UpdateDhcpOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
