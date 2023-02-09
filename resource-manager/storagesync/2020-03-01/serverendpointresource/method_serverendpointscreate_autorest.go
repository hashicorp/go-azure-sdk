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

type ServerEndpointsCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerEndpointsCreate ...
func (c ServerEndpointResourceClient) ServerEndpointsCreate(ctx context.Context, id ServerEndpointId, input ServerEndpointCreateParameters) (result ServerEndpointsCreateOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerEndpointsCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerEndpointsCreateThenPoll performs ServerEndpointsCreate then polls until it's completed
func (c ServerEndpointResourceClient) ServerEndpointsCreateThenPoll(ctx context.Context, id ServerEndpointId, input ServerEndpointCreateParameters) error {
	result, err := c.ServerEndpointsCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServerEndpointsCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerEndpointsCreate: %+v", err)
	}

	return nil
}

// preparerForServerEndpointsCreate prepares the ServerEndpointsCreate request.
func (c ServerEndpointResourceClient) preparerForServerEndpointsCreate(ctx context.Context, id ServerEndpointId, input ServerEndpointCreateParameters) (*http.Request, error) {
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

// senderForServerEndpointsCreate sends the ServerEndpointsCreate request. The method will close the
// http.Response Body if it receives an error.
func (c ServerEndpointResourceClient) senderForServerEndpointsCreate(ctx context.Context, req *http.Request) (future ServerEndpointsCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
