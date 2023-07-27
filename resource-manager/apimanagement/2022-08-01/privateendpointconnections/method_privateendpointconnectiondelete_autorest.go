package privateendpointconnections

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

type PrivateEndpointConnectionDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PrivateEndpointConnectionDelete ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionDelete(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionDeleteOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPrivateEndpointConnectionDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PrivateEndpointConnectionDeleteThenPoll performs PrivateEndpointConnectionDelete then polls until it's completed
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionDeleteThenPoll(ctx context.Context, id PrivateEndpointConnectionId) error {
	result, err := c.PrivateEndpointConnectionDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionDelete: %+v", err)
	}

	return nil
}

// preparerForPrivateEndpointConnectionDelete prepares the PrivateEndpointConnectionDelete request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionDelete(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// senderForPrivateEndpointConnectionDelete sends the PrivateEndpointConnectionDelete request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointConnectionsClient) senderForPrivateEndpointConnectionDelete(ctx context.Context, req *http.Request) (future PrivateEndpointConnectionDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
