package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSubscriptionLevelAlertStateToDismissOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateSubscriptionLevelAlertStateToDismiss ...
func (c AlertsClient) UpdateSubscriptionLevelAlertStateToDismiss(ctx context.Context, id AlertId) (result UpdateSubscriptionLevelAlertStateToDismissOperationResponse, err error) {
	req, err := c.preparerForUpdateSubscriptionLevelAlertStateToDismiss(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToDismiss", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToDismiss", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSubscriptionLevelAlertStateToDismiss(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToDismiss", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSubscriptionLevelAlertStateToDismiss prepares the UpdateSubscriptionLevelAlertStateToDismiss request.
func (c AlertsClient) preparerForUpdateSubscriptionLevelAlertStateToDismiss(ctx context.Context, id AlertId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dismiss", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSubscriptionLevelAlertStateToDismiss handles the response to the UpdateSubscriptionLevelAlertStateToDismiss request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateSubscriptionLevelAlertStateToDismiss(resp *http.Response) (result UpdateSubscriptionLevelAlertStateToDismissOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
