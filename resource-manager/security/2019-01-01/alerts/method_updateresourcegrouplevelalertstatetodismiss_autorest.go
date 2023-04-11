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

type UpdateResourceGroupLevelAlertStateToDismissOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateResourceGroupLevelAlertStateToDismiss ...
func (c AlertsClient) UpdateResourceGroupLevelAlertStateToDismiss(ctx context.Context, id LocationAlertId) (result UpdateResourceGroupLevelAlertStateToDismissOperationResponse, err error) {
	req, err := c.preparerForUpdateResourceGroupLevelAlertStateToDismiss(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToDismiss", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToDismiss", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateResourceGroupLevelAlertStateToDismiss(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToDismiss", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateResourceGroupLevelAlertStateToDismiss prepares the UpdateResourceGroupLevelAlertStateToDismiss request.
func (c AlertsClient) preparerForUpdateResourceGroupLevelAlertStateToDismiss(ctx context.Context, id LocationAlertId) (*http.Request, error) {
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

// responderForUpdateResourceGroupLevelAlertStateToDismiss handles the response to the UpdateResourceGroupLevelAlertStateToDismiss request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateResourceGroupLevelAlertStateToDismiss(resp *http.Response) (result UpdateResourceGroupLevelAlertStateToDismissOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
