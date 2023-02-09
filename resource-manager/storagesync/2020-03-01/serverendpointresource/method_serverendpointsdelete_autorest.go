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

type ServerEndpointsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerEndpointsDelete ...
func (c ServerEndpointResourceClient) ServerEndpointsDelete(ctx context.Context, id ServerEndpointId) (result ServerEndpointsDeleteOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerEndpointsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerEndpointsDeleteThenPoll performs ServerEndpointsDelete then polls until it's completed
func (c ServerEndpointResourceClient) ServerEndpointsDeleteThenPoll(ctx context.Context, id ServerEndpointId) error {
	result, err := c.ServerEndpointsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServerEndpointsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerEndpointsDelete: %+v", err)
	}

	return nil
}

// preparerForServerEndpointsDelete prepares the ServerEndpointsDelete request.
func (c ServerEndpointResourceClient) preparerForServerEndpointsDelete(ctx context.Context, id ServerEndpointId) (*http.Request, error) {
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

// senderForServerEndpointsDelete sends the ServerEndpointsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ServerEndpointResourceClient) senderForServerEndpointsDelete(ctx context.Context, req *http.Request) (future ServerEndpointsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
