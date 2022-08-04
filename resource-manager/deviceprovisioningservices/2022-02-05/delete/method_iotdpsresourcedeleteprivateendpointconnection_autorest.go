package delete

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

type IotDpsResourceDeletePrivateEndpointConnectionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// IotDpsResourceDeletePrivateEndpointConnection ...
func (c DELETEClient) IotDpsResourceDeletePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (result IotDpsResourceDeletePrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceDeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "IotDpsResourceDeletePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForIotDpsResourceDeletePrivateEndpointConnection(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "IotDpsResourceDeletePrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// IotDpsResourceDeletePrivateEndpointConnectionThenPoll performs IotDpsResourceDeletePrivateEndpointConnection then polls until it's completed
func (c DELETEClient) IotDpsResourceDeletePrivateEndpointConnectionThenPoll(ctx context.Context, id PrivateEndpointConnectionId) error {
	result, err := c.IotDpsResourceDeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		return fmt.Errorf("performing IotDpsResourceDeletePrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after IotDpsResourceDeletePrivateEndpointConnection: %+v", err)
	}

	return nil
}

// preparerForIotDpsResourceDeletePrivateEndpointConnection prepares the IotDpsResourceDeletePrivateEndpointConnection request.
func (c DELETEClient) preparerForIotDpsResourceDeletePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// senderForIotDpsResourceDeletePrivateEndpointConnection sends the IotDpsResourceDeletePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (c DELETEClient) senderForIotDpsResourceDeletePrivateEndpointConnection(ctx context.Context, req *http.Request) (future IotDpsResourceDeletePrivateEndpointConnectionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
