package generaterecommendations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationsGetGenerateStatusOperationResponse struct {
	HttpResponse *http.Response
}

// RecommendationsGetGenerateStatus ...
func (c GenerateRecommendationsClient) RecommendationsGetGenerateStatus(ctx context.Context, id GenerateRecommendationId) (result RecommendationsGetGenerateStatusOperationResponse, err error) {
	req, err := c.preparerForRecommendationsGetGenerateStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGetGenerateStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGetGenerateStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecommendationsGetGenerateStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGetGenerateStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecommendationsGetGenerateStatus prepares the RecommendationsGetGenerateStatus request.
func (c GenerateRecommendationsClient) preparerForRecommendationsGetGenerateStatus(ctx context.Context, id GenerateRecommendationId) (*http.Request, error) {
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

// responderForRecommendationsGetGenerateStatus handles the response to the RecommendationsGetGenerateStatus request. The method always
// closes the http.Response Body.
func (c GenerateRecommendationsClient) responderForRecommendationsGetGenerateStatus(resp *http.Response) (result RecommendationsGetGenerateStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
