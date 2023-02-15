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

type PrivateEndpointConnectionCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PrivateEndpointConnectionCreateOrUpdate ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionRequest) (result PrivateEndpointConnectionCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPrivateEndpointConnectionCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PrivateEndpointConnectionCreateOrUpdateThenPoll performs PrivateEndpointConnectionCreateOrUpdate then polls until it's completed
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionCreateOrUpdateThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionRequest) error {
	result, err := c.PrivateEndpointConnectionCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForPrivateEndpointConnectionCreateOrUpdate prepares the PrivateEndpointConnectionCreateOrUpdate request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionRequest) (*http.Request, error) {
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

// senderForPrivateEndpointConnectionCreateOrUpdate sends the PrivateEndpointConnectionCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointConnectionsClient) senderForPrivateEndpointConnectionCreateOrUpdate(ctx context.Context, req *http.Request) (future PrivateEndpointConnectionCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
