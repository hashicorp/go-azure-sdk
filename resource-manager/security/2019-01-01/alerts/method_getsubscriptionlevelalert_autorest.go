package alerts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSubscriptionLevelAlertOperationResponse struct {
	HttpResponse *http.Response
	Model        *Alert
}

// GetSubscriptionLevelAlert ...
func (c AlertsClient) GetSubscriptionLevelAlert(ctx context.Context, id AlertId) (result GetSubscriptionLevelAlertOperationResponse, err error) {
	req, err := c.preparerForGetSubscriptionLevelAlert(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetSubscriptionLevelAlert", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetSubscriptionLevelAlert", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSubscriptionLevelAlert(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetSubscriptionLevelAlert", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSubscriptionLevelAlert prepares the GetSubscriptionLevelAlert request.
func (c AlertsClient) preparerForGetSubscriptionLevelAlert(ctx context.Context, id AlertId) (*http.Request, error) {
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

// responderForGetSubscriptionLevelAlert handles the response to the GetSubscriptionLevelAlert request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForGetSubscriptionLevelAlert(resp *http.Response) (result GetSubscriptionLevelAlertOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
