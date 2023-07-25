package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWorkflowOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkflowEnvelope
}

// GetWorkflow ...
func (c WebAppsClient) GetWorkflow(ctx context.Context, id WorkflowId) (result GetWorkflowOperationResponse, err error) {
	req, err := c.preparerForGetWorkflow(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWorkflow", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWorkflow", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetWorkflow(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWorkflow", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetWorkflow prepares the GetWorkflow request.
func (c WebAppsClient) preparerForGetWorkflow(ctx context.Context, id WorkflowId) (*http.Request, error) {
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

// responderForGetWorkflow handles the response to the GetWorkflow request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetWorkflow(resp *http.Response) (result GetWorkflowOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
