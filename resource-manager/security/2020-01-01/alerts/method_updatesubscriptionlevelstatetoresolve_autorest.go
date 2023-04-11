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

type UpdateSubscriptionLevelStateToResolveOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateSubscriptionLevelStateToResolve ...
func (c AlertsClient) UpdateSubscriptionLevelStateToResolve(ctx context.Context, id AlertId) (result UpdateSubscriptionLevelStateToResolveOperationResponse, err error) {
	req, err := c.preparerForUpdateSubscriptionLevelStateToResolve(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelStateToResolve", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelStateToResolve", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSubscriptionLevelStateToResolve(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateSubscriptionLevelStateToResolve", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSubscriptionLevelStateToResolve prepares the UpdateSubscriptionLevelStateToResolve request.
func (c AlertsClient) preparerForUpdateSubscriptionLevelStateToResolve(ctx context.Context, id AlertId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resolve", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSubscriptionLevelStateToResolve handles the response to the UpdateSubscriptionLevelStateToResolve request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateSubscriptionLevelStateToResolve(resp *http.Response) (result UpdateSubscriptionLevelStateToResolveOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
