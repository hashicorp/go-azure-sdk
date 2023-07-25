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

type DisableRecommendationForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
}

type DisableRecommendationForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultDisableRecommendationForHostingEnvironmentOperationOptions() DisableRecommendationForHostingEnvironmentOperationOptions {
	return DisableRecommendationForHostingEnvironmentOperationOptions{}
}

func (o DisableRecommendationForHostingEnvironmentOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DisableRecommendationForHostingEnvironmentOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EnvironmentName != nil {
		out["environmentName"] = *o.EnvironmentName
	}

	return out
}

// DisableRecommendationForHostingEnvironment ...
func (c RecommendationsClient) DisableRecommendationForHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options DisableRecommendationForHostingEnvironmentOperationOptions) (result DisableRecommendationForHostingEnvironmentOperationResponse, err error) {
	req, err := c.preparerForDisableRecommendationForHostingEnvironment(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForHostingEnvironment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForHostingEnvironment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisableRecommendationForHostingEnvironment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForHostingEnvironment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisableRecommendationForHostingEnvironment prepares the DisableRecommendationForHostingEnvironment request.
func (c RecommendationsClient) preparerForDisableRecommendationForHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options DisableRecommendationForHostingEnvironmentOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisableRecommendationForHostingEnvironment handles the response to the DisableRecommendationForHostingEnvironment request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForDisableRecommendationForHostingEnvironment(resp *http.Response) (result DisableRecommendationForHostingEnvironmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
