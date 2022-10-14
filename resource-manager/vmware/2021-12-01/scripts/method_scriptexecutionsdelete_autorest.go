package scripts

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

type ScriptExecutionsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ScriptExecutionsDelete ...
func (c ScriptsClient) ScriptExecutionsDelete(ctx context.Context, id ScriptExecutionId) (result ScriptExecutionsDeleteOperationResponse, err error) {
	req, err := c.preparerForScriptExecutionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForScriptExecutionsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ScriptExecutionsDeleteThenPoll performs ScriptExecutionsDelete then polls until it's completed
func (c ScriptsClient) ScriptExecutionsDeleteThenPoll(ctx context.Context, id ScriptExecutionId) error {
	result, err := c.ScriptExecutionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ScriptExecutionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ScriptExecutionsDelete: %+v", err)
	}

	return nil
}

// preparerForScriptExecutionsDelete prepares the ScriptExecutionsDelete request.
func (c ScriptsClient) preparerForScriptExecutionsDelete(ctx context.Context, id ScriptExecutionId) (*http.Request, error) {
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

// senderForScriptExecutionsDelete sends the ScriptExecutionsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ScriptsClient) senderForScriptExecutionsDelete(ctx context.Context, req *http.Request) (future ScriptExecutionsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
