package workflowtriggerhistories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResubmitOperationResponse struct {
	HttpResponse *http.Response
}

// Resubmit ...
func (c WorkflowTriggerHistoriesClient) Resubmit(ctx context.Context, id HistoryId) (result ResubmitOperationResponse, err error) {
	req, err := c.preparerForResubmit(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowtriggerhistories.WorkflowTriggerHistoriesClient", "Resubmit", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowtriggerhistories.WorkflowTriggerHistoriesClient", "Resubmit", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResubmit(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowtriggerhistories.WorkflowTriggerHistoriesClient", "Resubmit", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResubmit prepares the Resubmit request.
func (c WorkflowTriggerHistoriesClient) preparerForResubmit(ctx context.Context, id HistoryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resubmit", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResubmit handles the response to the Resubmit request. The method always
// closes the http.Response Body.
func (c WorkflowTriggerHistoriesClient) responderForResubmit(resp *http.Response) (result ResubmitOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
