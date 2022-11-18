package listallhybridrunbookworkergroupinautomationaccount

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerGroupDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// HybridRunbookWorkerGroupDelete ...
func (c ListAllHybridRunbookWorkerGroupInAutomationAccountClient) HybridRunbookWorkerGroupDelete(ctx context.Context, id HybridRunbookWorkerGroupId) (result HybridRunbookWorkerGroupDeleteOperationResponse, err error) {
	req, err := c.preparerForHybridRunbookWorkerGroupDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listallhybridrunbookworkergroupinautomationaccount.ListAllHybridRunbookWorkerGroupInAutomationAccountClient", "HybridRunbookWorkerGroupDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "listallhybridrunbookworkergroupinautomationaccount.ListAllHybridRunbookWorkerGroupInAutomationAccountClient", "HybridRunbookWorkerGroupDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHybridRunbookWorkerGroupDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "listallhybridrunbookworkergroupinautomationaccount.ListAllHybridRunbookWorkerGroupInAutomationAccountClient", "HybridRunbookWorkerGroupDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHybridRunbookWorkerGroupDelete prepares the HybridRunbookWorkerGroupDelete request.
func (c ListAllHybridRunbookWorkerGroupInAutomationAccountClient) preparerForHybridRunbookWorkerGroupDelete(ctx context.Context, id HybridRunbookWorkerGroupId) (*http.Request, error) {
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

// responderForHybridRunbookWorkerGroupDelete handles the response to the HybridRunbookWorkerGroupDelete request. The method always
// closes the http.Response Body.
func (c ListAllHybridRunbookWorkerGroupInAutomationAccountClient) responderForHybridRunbookWorkerGroupDelete(resp *http.Response) (result HybridRunbookWorkerGroupDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
