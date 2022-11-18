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

type ScriptExecutionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ScriptExecutionsCreateOrUpdate ...
func (c ScriptsClient) ScriptExecutionsCreateOrUpdate(ctx context.Context, id ScriptExecutionId, input ScriptExecution) (result ScriptExecutionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForScriptExecutionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForScriptExecutionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ScriptExecutionsCreateOrUpdateThenPoll performs ScriptExecutionsCreateOrUpdate then polls until it's completed
func (c ScriptsClient) ScriptExecutionsCreateOrUpdateThenPoll(ctx context.Context, id ScriptExecutionId, input ScriptExecution) error {
	result, err := c.ScriptExecutionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ScriptExecutionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ScriptExecutionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForScriptExecutionsCreateOrUpdate prepares the ScriptExecutionsCreateOrUpdate request.
func (c ScriptsClient) preparerForScriptExecutionsCreateOrUpdate(ctx context.Context, id ScriptExecutionId, input ScriptExecution) (*http.Request, error) {
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

// senderForScriptExecutionsCreateOrUpdate sends the ScriptExecutionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ScriptsClient) senderForScriptExecutionsCreateOrUpdate(ctx context.Context, req *http.Request) (future ScriptExecutionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
