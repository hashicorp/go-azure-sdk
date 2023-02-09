package serverendpointresource

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

type ServerEndpointsUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerEndpointsUpdate ...
func (c ServerEndpointResourceClient) ServerEndpointsUpdate(ctx context.Context, id ServerEndpointId, input ServerEndpointUpdateParameters) (result ServerEndpointsUpdateOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerEndpointsUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerEndpointsUpdateThenPoll performs ServerEndpointsUpdate then polls until it's completed
func (c ServerEndpointResourceClient) ServerEndpointsUpdateThenPoll(ctx context.Context, id ServerEndpointId, input ServerEndpointUpdateParameters) error {
	result, err := c.ServerEndpointsUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServerEndpointsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerEndpointsUpdate: %+v", err)
	}

	return nil
}

// preparerForServerEndpointsUpdate prepares the ServerEndpointsUpdate request.
func (c ServerEndpointResourceClient) preparerForServerEndpointsUpdate(ctx context.Context, id ServerEndpointId, input ServerEndpointUpdateParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServerEndpointsUpdate sends the ServerEndpointsUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ServerEndpointResourceClient) senderForServerEndpointsUpdate(ctx context.Context, req *http.Request) (future ServerEndpointsUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
