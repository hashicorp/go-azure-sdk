package generaterecommendations

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

type RecommendationsGenerateOperationResponse struct {
	HttpResponse *http.Response
}

// RecommendationsGenerate ...
func (c GenerateRecommendationsClient) RecommendationsGenerate(ctx context.Context, id commonids.SubscriptionId) (result RecommendationsGenerateOperationResponse, err error) {
	req, err := c.preparerForRecommendationsGenerate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGenerate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGenerate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecommendationsGenerate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "generaterecommendations.GenerateRecommendationsClient", "RecommendationsGenerate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecommendationsGenerate prepares the RecommendationsGenerate request.
func (c GenerateRecommendationsClient) preparerForRecommendationsGenerate(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Advisor/generateRecommendations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRecommendationsGenerate handles the response to the RecommendationsGenerate request. The method always
// closes the http.Response Body.
func (c GenerateRecommendationsClient) responderForRecommendationsGenerate(resp *http.Response) (result RecommendationsGenerateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
