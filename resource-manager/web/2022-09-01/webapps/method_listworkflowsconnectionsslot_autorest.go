package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWorkflowsConnectionsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkflowEnvelope
}

// ListWorkflowsConnectionsSlot ...
func (c WebAppsClient) ListWorkflowsConnectionsSlot(ctx context.Context, id SlotId) (result ListWorkflowsConnectionsSlotOperationResponse, err error) {
	req, err := c.preparerForListWorkflowsConnectionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnectionsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnectionsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListWorkflowsConnectionsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnectionsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListWorkflowsConnectionsSlot prepares the ListWorkflowsConnectionsSlot request.
func (c WebAppsClient) preparerForListWorkflowsConnectionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listWorkflowsConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListWorkflowsConnectionsSlot handles the response to the ListWorkflowsConnectionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListWorkflowsConnectionsSlot(resp *http.Response) (result ListWorkflowsConnectionsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
