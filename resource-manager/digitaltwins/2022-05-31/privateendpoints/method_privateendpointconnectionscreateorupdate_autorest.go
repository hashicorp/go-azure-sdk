package privateendpoints

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

type PrivateEndpointConnectionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdate ...
func (c PrivateEndpointsClient) PrivateEndpointConnectionsCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (result PrivateEndpointConnectionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoints.PrivateEndpointsClient", "PrivateEndpointConnectionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPrivateEndpointConnectionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoints.PrivateEndpointsClient", "PrivateEndpointConnectionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PrivateEndpointConnectionsCreateOrUpdateThenPoll performs PrivateEndpointConnectionsCreateOrUpdate then polls until it's completed
func (c PrivateEndpointsClient) PrivateEndpointConnectionsCreateOrUpdateThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) error {
	result, err := c.PrivateEndpointConnectionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForPrivateEndpointConnectionsCreateOrUpdate prepares the PrivateEndpointConnectionsCreateOrUpdate request.
func (c PrivateEndpointsClient) preparerForPrivateEndpointConnectionsCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
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

// senderForPrivateEndpointConnectionsCreateOrUpdate sends the PrivateEndpointConnectionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointsClient) senderForPrivateEndpointConnectionsCreateOrUpdate(ctx context.Context, req *http.Request) (future PrivateEndpointConnectionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
