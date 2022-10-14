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

type CreatePublicIPOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreatePublicIP ...
func (c WorkloadNetworksClient) CreatePublicIP(ctx context.Context, id PublicIPId, input WorkloadNetworkPublicIP) (result CreatePublicIPOperationResponse, err error) {
	req, err := c.preparerForCreatePublicIP(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreatePublicIP", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreatePublicIP(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreatePublicIP", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreatePublicIPThenPoll performs CreatePublicIP then polls until it's completed
func (c WorkloadNetworksClient) CreatePublicIPThenPoll(ctx context.Context, id PublicIPId, input WorkloadNetworkPublicIP) error {
	result, err := c.CreatePublicIP(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreatePublicIP: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreatePublicIP: %+v", err)
	}

	return nil
}

// preparerForCreatePublicIP prepares the CreatePublicIP request.
func (c WorkloadNetworksClient) preparerForCreatePublicIP(ctx context.Context, id PublicIPId, input WorkloadNetworkPublicIP) (*http.Request, error) {
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

// senderForCreatePublicIP sends the CreatePublicIP request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreatePublicIP(ctx context.Context, req *http.Request) (future CreatePublicIPOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
