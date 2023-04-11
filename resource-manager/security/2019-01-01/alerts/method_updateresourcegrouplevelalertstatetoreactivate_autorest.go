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

type UpdateResourceGroupLevelAlertStateToReactivateOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateResourceGroupLevelAlertStateToReactivate ...
func (c AlertsClient) UpdateResourceGroupLevelAlertStateToReactivate(ctx context.Context, id LocationAlertId) (result UpdateResourceGroupLevelAlertStateToReactivateOperationResponse, err error) {
	req, err := c.preparerForUpdateResourceGroupLevelAlertStateToReactivate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToReactivate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToReactivate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateResourceGroupLevelAlertStateToReactivate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelAlertStateToReactivate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateResourceGroupLevelAlertStateToReactivate prepares the UpdateResourceGroupLevelAlertStateToReactivate request.
func (c AlertsClient) preparerForUpdateResourceGroupLevelAlertStateToReactivate(ctx context.Context, id LocationAlertId) (*http.Request, error) {
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

// responderForUpdateResourceGroupLevelAlertStateToReactivate handles the response to the UpdateResourceGroupLevelAlertStateToReactivate request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateResourceGroupLevelAlertStateToReactivate(resp *http.Response) (result UpdateResourceGroupLevelAlertStateToReactivateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
