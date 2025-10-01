package privateendpointconnectionresources

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

type PrivateEndpointConnectionPutOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionResource
}

// PrivateEndpointConnectionPut ...
func (c PrivateEndpointConnectionResourcesClient) PrivateEndpointConnectionPut(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionResource) (result PrivateEndpointConnectionPutOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionPut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient", "PrivateEndpointConnectionPut", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPrivateEndpointConnectionPut(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient", "PrivateEndpointConnectionPut", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PrivateEndpointConnectionPutThenPoll performs PrivateEndpointConnectionPut then polls until it's completed
func (c PrivateEndpointConnectionResourcesClient) PrivateEndpointConnectionPutThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionResource) error {
	result, err := c.PrivateEndpointConnectionPut(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionPut: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionPut: %+v", err)
	}

	return nil
}

// preparerForPrivateEndpointConnectionPut prepares the PrivateEndpointConnectionPut request.
func (c PrivateEndpointConnectionResourcesClient) preparerForPrivateEndpointConnectionPut(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnectionResource) (*http.Request, error) {
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

// senderForPrivateEndpointConnectionPut sends the PrivateEndpointConnectionPut request. The method will close the
// http.Response Body if it receives an error.
func (c PrivateEndpointConnectionResourcesClient) senderForPrivateEndpointConnectionPut(ctx context.Context, req *http.Request) (future PrivateEndpointConnectionPutOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
