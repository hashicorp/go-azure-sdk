package actions

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

type ServerEndpointsrecallActionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerEndpointsrecallAction ...
func (c ActionsClient) ServerEndpointsrecallAction(ctx context.Context, id ServerEndpointId, input RecallActionParameters) (result ServerEndpointsrecallActionOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsrecallAction(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "ServerEndpointsrecallAction", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerEndpointsrecallAction(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "ServerEndpointsrecallAction", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerEndpointsrecallActionThenPoll performs ServerEndpointsrecallAction then polls until it's completed
func (c ActionsClient) ServerEndpointsrecallActionThenPoll(ctx context.Context, id ServerEndpointId, input RecallActionParameters) error {
	result, err := c.ServerEndpointsrecallAction(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServerEndpointsrecallAction: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerEndpointsrecallAction: %+v", err)
	}

	return nil
}

// preparerForServerEndpointsrecallAction prepares the ServerEndpointsrecallAction request.
func (c ActionsClient) preparerForServerEndpointsrecallAction(ctx context.Context, id ServerEndpointId, input RecallActionParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recallAction", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServerEndpointsrecallAction sends the ServerEndpointsrecallAction request. The method will close the
// http.Response Body if it receives an error.
func (c ActionsClient) senderForServerEndpointsrecallAction(ctx context.Context, req *http.Request) (future ServerEndpointsrecallActionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
