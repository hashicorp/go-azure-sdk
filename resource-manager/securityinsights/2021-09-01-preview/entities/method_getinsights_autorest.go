package entities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightsOperationResponse struct {
	HttpResponse *http.Response
	Model        *EntityGetInsightsResponse
}

// GetInsights ...
func (c EntitiesClient) GetInsights(ctx context.Context, id EntitiesId, input EntityGetInsightsParameters) (result GetInsightsOperationResponse, err error) {
	req, err := c.preparerForGetInsights(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetInsights", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetInsights", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInsights(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetInsights", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInsights prepares the GetInsights request.
func (c EntitiesClient) preparerForGetInsights(ctx context.Context, id EntitiesId, input EntityGetInsightsParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getInsights", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInsights handles the response to the GetInsights request. The method always
// closes the http.Response Body.
func (c EntitiesClient) responderForGetInsights(resp *http.Response) (result GetInsightsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
