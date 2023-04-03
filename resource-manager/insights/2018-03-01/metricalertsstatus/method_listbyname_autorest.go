package metricalertsstatus

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *MetricAlertStatusCollection
}

// ListByName ...
func (c MetricAlertsStatusClient) ListByName(ctx context.Context, id StatusId) (result ListByNameOperationResponse, err error) {
	req, err := c.preparerForListByName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricalertsstatus.MetricAlertsStatusClient", "ListByName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricalertsstatus.MetricAlertsStatusClient", "ListByName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricalertsstatus.MetricAlertsStatusClient", "ListByName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByName prepares the ListByName request.
func (c MetricAlertsStatusClient) preparerForListByName(ctx context.Context, id StatusId) (*http.Request, error) {
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

// responderForListByName handles the response to the ListByName request. The method always
// closes the http.Response Body.
func (c MetricAlertsStatusClient) responderForListByName(resp *http.Response) (result ListByNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
