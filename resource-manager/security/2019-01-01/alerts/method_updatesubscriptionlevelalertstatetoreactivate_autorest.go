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

type UpdateSubscriptionLevelAlertStateToReactivateOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateSubscriptionLevelAlertStateToReactivate ...
func (c AlertsClient) UpdateSubscriptionLevelAlertStateToReactivate(ctx context.Context, id AlertId) (result UpdateSubscriptionLevelAlertStateToReactivateOperationResponse, err error) {
	req, err := c.preparerForUpdateSubscriptionLevelAlertStateToReactivate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToReactivate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToReactivate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSubscriptionLevelAlertStateToReactivate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelAlertStateToReactivate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSubscriptionLevelAlertStateToReactivate prepares the UpdateSubscriptionLevelAlertStateToReactivate request.
func (c AlertsClient) preparerForUpdateSubscriptionLevelAlertStateToReactivate(ctx context.Context, id AlertId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reactivate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSubscriptionLevelAlertStateToReactivate handles the response to the UpdateSubscriptionLevelAlertStateToReactivate request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateSubscriptionLevelAlertStateToReactivate(resp *http.Response) (result UpdateSubscriptionLevelAlertStateToReactivateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
