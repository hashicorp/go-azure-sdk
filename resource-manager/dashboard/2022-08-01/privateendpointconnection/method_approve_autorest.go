package privateendpointconnection

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

type ApproveOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Approve ...
func (c PrivateEndpointConnectionClient) Approve(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (result ApproveOperationResponse, err error) {
	req, err := c.preparerForApprove(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "Approve", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApprove(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "Approve", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApproveThenPoll performs Approve then polls until it's completed
func (c PrivateEndpointConnectionClient) ApproveThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) error {
	result, err := c.Approve(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Approve: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Approve: %+v", err)
	}

	return nil
}

// preparerForApprove prepares the Approve request.
func (c PrivateEndpointConnectionClient) preparerForApprove(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
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

// senderForApprove sends the Approve request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointConnectionClient) senderForApprove(ctx context.Context, req *http.Request) (future ApproveOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
