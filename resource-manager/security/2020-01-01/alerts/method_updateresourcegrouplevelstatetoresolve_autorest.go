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

type UpdateResourceGroupLevelStateToResolveOperationResponse struct {
	HttpResponse *http.Response
}

// UpdateResourceGroupLevelStateToResolve ...
func (c AlertsClient) UpdateResourceGroupLevelStateToResolve(ctx context.Context, id LocationAlertId) (result UpdateResourceGroupLevelStateToResolveOperationResponse, err error) {
	req, err := c.preparerForUpdateResourceGroupLevelStateToResolve(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelStateToResolve", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelStateToResolve", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateResourceGroupLevelStateToResolve(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "UpdateResourceGroupLevelStateToResolve", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateResourceGroupLevelStateToResolve prepares the UpdateResourceGroupLevelStateToResolve request.
func (c AlertsClient) preparerForUpdateResourceGroupLevelStateToResolve(ctx context.Context, id LocationAlertId) (*http.Request, error) {
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

// responderForUpdateResourceGroupLevelStateToResolve handles the response to the UpdateResourceGroupLevelStateToResolve request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForUpdateResourceGroupLevelStateToResolve(resp *http.Response) (result UpdateResourceGroupLevelStateToResolveOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
