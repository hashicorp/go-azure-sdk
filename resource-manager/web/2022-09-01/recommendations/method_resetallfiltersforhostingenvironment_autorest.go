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

type ResetAllFiltersForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
}

type ResetAllFiltersForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultResetAllFiltersForHostingEnvironmentOperationOptions() ResetAllFiltersForHostingEnvironmentOperationOptions {
	return ResetAllFiltersForHostingEnvironmentOperationOptions{}
}

func (o ResetAllFiltersForHostingEnvironmentOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ResetAllFiltersForHostingEnvironmentOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EnvironmentName != nil {
		out["environmentName"] = *o.EnvironmentName
	}

	return out
}

// ResetAllFiltersForHostingEnvironment ...
func (c RecommendationsClient) ResetAllFiltersForHostingEnvironment(ctx context.Context, id HostingEnvironmentId, options ResetAllFiltersForHostingEnvironmentOperationOptions) (result ResetAllFiltersForHostingEnvironmentOperationResponse, err error) {
	req, err := c.preparerForResetAllFiltersForHostingEnvironment(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForHostingEnvironment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForHostingEnvironment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetAllFiltersForHostingEnvironment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ResetAllFiltersForHostingEnvironment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetAllFiltersForHostingEnvironment prepares the ResetAllFiltersForHostingEnvironment request.
func (c RecommendationsClient) preparerForResetAllFiltersForHostingEnvironment(ctx context.Context, id HostingEnvironmentId, options ResetAllFiltersForHostingEnvironmentOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/recommendations/reset", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetAllFiltersForHostingEnvironment handles the response to the ResetAllFiltersForHostingEnvironment request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForResetAllFiltersForHostingEnvironment(resp *http.Response) (result ResetAllFiltersForHostingEnvironmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
