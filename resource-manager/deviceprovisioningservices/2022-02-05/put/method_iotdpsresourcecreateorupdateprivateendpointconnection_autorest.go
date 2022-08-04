package put

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

type IotDpsResourceCreateOrUpdatePrivateEndpointConnectionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// IotDpsResourceCreateOrUpdatePrivateEndpointConnection ...
func (c PUTClient) IotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (result IotDpsResourceCreateOrUpdatePrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "IotDpsResourceCreateOrUpdatePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForIotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "put.PUTClient", "IotDpsResourceCreateOrUpdatePrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// IotDpsResourceCreateOrUpdatePrivateEndpointConnectionThenPoll performs IotDpsResourceCreateOrUpdatePrivateEndpointConnection then polls until it's completed
func (c PUTClient) IotDpsResourceCreateOrUpdatePrivateEndpointConnectionThenPoll(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) error {
	result, err := c.IotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing IotDpsResourceCreateOrUpdatePrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after IotDpsResourceCreateOrUpdatePrivateEndpointConnection: %+v", err)
	}

	return nil
}

// preparerForIotDpsResourceCreateOrUpdatePrivateEndpointConnection prepares the IotDpsResourceCreateOrUpdatePrivateEndpointConnection request.
func (c PUTClient) preparerForIotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
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

// senderForIotDpsResourceCreateOrUpdatePrivateEndpointConnection sends the IotDpsResourceCreateOrUpdatePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (c PUTClient) senderForIotDpsResourceCreateOrUpdatePrivateEndpointConnection(ctx context.Context, req *http.Request) (future IotDpsResourceCreateOrUpdatePrivateEndpointConnectionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
