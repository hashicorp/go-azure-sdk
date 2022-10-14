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

type CreateDhcpOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateDhcp ...
func (c WorkloadNetworksClient) CreateDhcp(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) (result CreateDhcpOperationResponse, err error) {
	req, err := c.preparerForCreateDhcp(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateDhcp", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateDhcp(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateDhcp", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateDhcpThenPoll performs CreateDhcp then polls until it's completed
func (c WorkloadNetworksClient) CreateDhcpThenPoll(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) error {
	result, err := c.CreateDhcp(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateDhcp: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateDhcp: %+v", err)
	}

	return nil
}

// preparerForCreateDhcp prepares the CreateDhcp request.
func (c WorkloadNetworksClient) preparerForCreateDhcp(ctx context.Context, id DhcpConfigurationId, input WorkloadNetworkDhcp) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateDhcp sends the CreateDhcp request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreateDhcp(ctx context.Context, req *http.Request) (future CreateDhcpOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
