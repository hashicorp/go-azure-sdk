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

type DeletePublicIPOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeletePublicIP ...
func (c WorkloadNetworksClient) DeletePublicIP(ctx context.Context, id PublicIPId) (result DeletePublicIPOperationResponse, err error) {
	req, err := c.preparerForDeletePublicIP(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeletePublicIP", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeletePublicIP(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeletePublicIP", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeletePublicIPThenPoll performs DeletePublicIP then polls until it's completed
func (c WorkloadNetworksClient) DeletePublicIPThenPoll(ctx context.Context, id PublicIPId) error {
	result, err := c.DeletePublicIP(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeletePublicIP: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeletePublicIP: %+v", err)
	}

	return nil
}

// preparerForDeletePublicIP prepares the DeletePublicIP request.
func (c WorkloadNetworksClient) preparerForDeletePublicIP(ctx context.Context, id PublicIPId) (*http.Request, error) {
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

// senderForDeletePublicIP sends the DeletePublicIP request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeletePublicIP(ctx context.Context, req *http.Request) (future DeletePublicIPOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
