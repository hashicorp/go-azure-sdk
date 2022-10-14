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

type DeleteDhcpOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteDhcp ...
func (c WorkloadNetworksClient) DeleteDhcp(ctx context.Context, id DhcpConfigurationId) (result DeleteDhcpOperationResponse, err error) {
	req, err := c.preparerForDeleteDhcp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteDhcp", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteDhcp(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteDhcp", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteDhcpThenPoll performs DeleteDhcp then polls until it's completed
func (c WorkloadNetworksClient) DeleteDhcpThenPoll(ctx context.Context, id DhcpConfigurationId) error {
	result, err := c.DeleteDhcp(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteDhcp: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteDhcp: %+v", err)
	}

	return nil
}

// preparerForDeleteDhcp prepares the DeleteDhcp request.
func (c WorkloadNetworksClient) preparerForDeleteDhcp(ctx context.Context, id DhcpConfigurationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteDhcp sends the DeleteDhcp request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeleteDhcp(ctx context.Context, req *http.Request) (future DeleteDhcpOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
