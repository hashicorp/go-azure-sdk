package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetAllFiltersOperationResponse struct {
	HttpResponse *http.Response
}

// ResetAllFilters ...
func (c RecommendationsClient) ResetAllFilters(ctx context.Context, id commonids.SubscriptionId) (result ResetAllFiltersOperationResponse, err error) {
	req, err := c.preparerForResetAllFilters(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFilters", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFilters", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetAllFilters(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFilters", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetAllFilters prepares the ResetAllFilters request.
func (c RecommendationsClient) preparerForResetAllFilters(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/recommendations/reset", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetAllFilters handles the response to the ResetAllFilters request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForResetAllFilters(resp *http.Response) (result ResetAllFiltersOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
