package metadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationMetadataGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MetadataEntity
}

// RecommendationMetadataGet ...
func (c MetadataClient) RecommendationMetadataGet(ctx context.Context, id MetadataId) (result RecommendationMetadataGetOperationResponse, err error) {
	req, err := c.preparerForRecommendationMetadataGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecommendationMetadataGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecommendationMetadataGet prepares the RecommendationMetadataGet request.
func (c MetadataClient) preparerForRecommendationMetadataGet(ctx context.Context, id MetadataId) (*http.Request, error) {
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

// responderForRecommendationMetadataGet handles the response to the RecommendationMetadataGet request. The method always
// closes the http.Response Body.
func (c MetadataClient) responderForRecommendationMetadataGet(resp *http.Response) (result RecommendationMetadataGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
