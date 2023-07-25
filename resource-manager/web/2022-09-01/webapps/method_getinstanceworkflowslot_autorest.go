package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceWorkflowSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkflowEnvelope
}

// GetInstanceWorkflowSlot ...
func (c WebAppsClient) GetInstanceWorkflowSlot(ctx context.Context, id SlotWorkflowId) (result GetInstanceWorkflowSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceWorkflowSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceWorkflowSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceWorkflowSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceWorkflowSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceWorkflowSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceWorkflowSlot prepares the GetInstanceWorkflowSlot request.
func (c WebAppsClient) preparerForGetInstanceWorkflowSlot(ctx context.Context, id SlotWorkflowId) (*http.Request, error) {
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

// responderForGetInstanceWorkflowSlot handles the response to the GetInstanceWorkflowSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceWorkflowSlot(resp *http.Response) (result GetInstanceWorkflowSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
