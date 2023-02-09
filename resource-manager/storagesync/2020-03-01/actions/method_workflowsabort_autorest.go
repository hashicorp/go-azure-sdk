package actions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowsAbortOperationResponse struct {
	HttpResponse *http.Response
}

// WorkflowsAbort ...
func (c ActionsClient) WorkflowsAbort(ctx context.Context, id WorkflowId) (result WorkflowsAbortOperationResponse, err error) {
	req, err := c.preparerForWorkflowsAbort(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "WorkflowsAbort", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "WorkflowsAbort", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkflowsAbort(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "WorkflowsAbort", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkflowsAbort prepares the WorkflowsAbort request.
func (c ActionsClient) preparerForWorkflowsAbort(ctx context.Context, id WorkflowId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/abort", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkflowsAbort handles the response to the WorkflowsAbort request. The method always
// closes the http.Response Body.
func (c ActionsClient) responderForWorkflowsAbort(resp *http.Response) (result WorkflowsAbortOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
