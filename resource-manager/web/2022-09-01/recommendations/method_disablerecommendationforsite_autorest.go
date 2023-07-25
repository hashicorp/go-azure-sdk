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

type DisableRecommendationForSiteOperationResponse struct {
	HttpResponse *http.Response
}

// DisableRecommendationForSite ...
func (c RecommendationsClient) DisableRecommendationForSite(ctx context.Context, id SiteRecommendationId) (result DisableRecommendationForSiteOperationResponse, err error) {
	req, err := c.preparerForDisableRecommendationForSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisableRecommendationForSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisableRecommendationForSite prepares the DisableRecommendationForSite request.
func (c RecommendationsClient) preparerForDisableRecommendationForSite(ctx context.Context, id SiteRecommendationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisableRecommendationForSite handles the response to the DisableRecommendationForSite request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForDisableRecommendationForSite(resp *http.Response) (result DisableRecommendationForSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
