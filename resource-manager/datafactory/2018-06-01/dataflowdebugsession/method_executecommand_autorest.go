package dataflowdebugsession

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

type ExecuteCommandOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ExecuteCommand ...
func (c DataFlowDebugSessionClient) ExecuteCommand(ctx context.Context, id FactoryId, input DataFlowDebugCommandRequest) (result ExecuteCommandOperationResponse, err error) {
	req, err := c.preparerForExecuteCommand(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "ExecuteCommand", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForExecuteCommand(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "ExecuteCommand", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ExecuteCommandThenPoll performs ExecuteCommand then polls until it's completed
func (c DataFlowDebugSessionClient) ExecuteCommandThenPoll(ctx context.Context, id FactoryId, input DataFlowDebugCommandRequest) error {
	result, err := c.ExecuteCommand(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ExecuteCommand: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ExecuteCommand: %+v", err)
	}

	return nil
}

// preparerForExecuteCommand prepares the ExecuteCommand request.
func (c DataFlowDebugSessionClient) preparerForExecuteCommand(ctx context.Context, id FactoryId, input DataFlowDebugCommandRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/executeDataFlowDebugCommand", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForExecuteCommand sends the ExecuteCommand request. The method will close the
// http.Response Body if it receives an error.
func (c DataFlowDebugSessionClient) senderForExecuteCommand(ctx context.Context, req *http.Request) (future ExecuteCommandOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
