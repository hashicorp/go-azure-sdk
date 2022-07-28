package operationalinsights

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueriesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// QueriesDelete ...
func (c OperationalInsightsClient) QueriesDelete(ctx context.Context, id QueriesId) (result QueriesDeleteOperationResponse, err error) {
	req, err := c.preparerForQueriesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationalInsightsClient", "QueriesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationalInsightsClient", "QueriesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForQueriesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.OperationalInsightsClient", "QueriesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForQueriesDelete prepares the QueriesDelete request.
func (c OperationalInsightsClient) preparerForQueriesDelete(ctx context.Context, id QueriesId) (*http.Request, error) {
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

// responderForQueriesDelete handles the response to the QueriesDelete request. The method always
// closes the http.Response Body.
func (c OperationalInsightsClient) responderForQueriesDelete(resp *http.Response) (result QueriesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
