package workflowresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowsListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkflowArray
}

// WorkflowsListByStorageSyncService ...
func (c WorkflowResourceClient) WorkflowsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result WorkflowsListByStorageSyncServiceOperationResponse, err error) {
	req, err := c.preparerForWorkflowsListByStorageSyncService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsListByStorageSyncService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkflowsListByStorageSyncService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowresource.WorkflowResourceClient", "WorkflowsListByStorageSyncService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkflowsListByStorageSyncService prepares the WorkflowsListByStorageSyncService request.
func (c WorkflowResourceClient) preparerForWorkflowsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workflows", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkflowsListByStorageSyncService handles the response to the WorkflowsListByStorageSyncService request. The method always
// closes the http.Response Body.
func (c WorkflowResourceClient) responderForWorkflowsListByStorageSyncService(resp *http.Response) (result WorkflowsListByStorageSyncServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
