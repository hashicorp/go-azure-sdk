package workflowresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Workflow
}

// WorkflowsGet ...
func (c WorkflowResourceClient) WorkflowsGet(ctx context.Context, id WorkflowId) (result WorkflowsGetOperationResponse, err error) {
	req, err := c.preparerForWorkflowsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkflowsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkflowsGet prepares the WorkflowsGet request.
func (c WorkflowResourceClient) preparerForWorkflowsGet(ctx context.Context, id WorkflowId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkflowsGet handles the response to the WorkflowsGet request. The method always
// closes the http.Response Body.
func (c WorkflowResourceClient) responderForWorkflowsGet(resp *http.Response) (result WorkflowsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
