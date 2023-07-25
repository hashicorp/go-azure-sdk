package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetAllFiltersForWebAppOperationResponse struct {
	HttpResponse *http.Response
}

// ResetAllFiltersForWebApp ...
func (c RecommendationsClient) ResetAllFiltersForWebApp(ctx context.Context, id SiteId) (result ResetAllFiltersForWebAppOperationResponse, err error) {
	req, err := c.preparerForResetAllFiltersForWebApp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForWebApp", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForWebApp", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetAllFiltersForWebApp(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForWebApp", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetAllFiltersForWebApp prepares the ResetAllFiltersForWebApp request.
func (c RecommendationsClient) preparerForResetAllFiltersForWebApp(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recommendations/reset", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetAllFiltersForWebApp handles the response to the ResetAllFiltersForWebApp request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForResetAllFiltersForWebApp(resp *http.Response) (result ResetAllFiltersForWebAppOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
