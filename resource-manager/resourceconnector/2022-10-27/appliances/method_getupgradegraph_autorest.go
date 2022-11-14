package appliances

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUpgradeGraphOperationResponse struct {
	HttpResponse *http.Response
	Model        *UpgradeGraph
}

// GetUpgradeGraph ...
func (c AppliancesClient) GetUpgradeGraph(ctx context.Context, id UpgradeGraphId) (result GetUpgradeGraphOperationResponse, err error) {
	req, err := c.preparerForGetUpgradeGraph(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetUpgradeGraph", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetUpgradeGraph", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetUpgradeGraph(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetUpgradeGraph", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetUpgradeGraph prepares the GetUpgradeGraph request.
func (c AppliancesClient) preparerForGetUpgradeGraph(ctx context.Context, id UpgradeGraphId) (*http.Request, error) {
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

// responderForGetUpgradeGraph handles the response to the GetUpgradeGraph request. The method always
// closes the http.Response Body.
func (c AppliancesClient) responderForGetUpgradeGraph(resp *http.Response) (result GetUpgradeGraphOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
