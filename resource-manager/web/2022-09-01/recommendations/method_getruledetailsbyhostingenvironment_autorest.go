package recommendations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRuleDetailsByHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecommendationRule
}

type GetRuleDetailsByHostingEnvironmentOperationOptions struct {
	RecommendationId *string
	UpdateSeen       *bool
}

func DefaultGetRuleDetailsByHostingEnvironmentOperationOptions() GetRuleDetailsByHostingEnvironmentOperationOptions {
	return GetRuleDetailsByHostingEnvironmentOperationOptions{}
}

func (o GetRuleDetailsByHostingEnvironmentOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetRuleDetailsByHostingEnvironmentOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RecommendationId != nil {
		out["recommendationId"] = *o.RecommendationId
	}

	if o.UpdateSeen != nil {
		out["updateSeen"] = *o.UpdateSeen
	}

	return out
}

// GetRuleDetailsByHostingEnvironment ...
func (c RecommendationsClient) GetRuleDetailsByHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options GetRuleDetailsByHostingEnvironmentOperationOptions) (result GetRuleDetailsByHostingEnvironmentOperationResponse, err error) {
	req, err := c.preparerForGetRuleDetailsByHostingEnvironment(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByHostingEnvironment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByHostingEnvironment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRuleDetailsByHostingEnvironment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByHostingEnvironment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRuleDetailsByHostingEnvironment prepares the GetRuleDetailsByHostingEnvironment request.
func (c RecommendationsClient) preparerForGetRuleDetailsByHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options GetRuleDetailsByHostingEnvironmentOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetRuleDetailsByHostingEnvironment handles the response to the GetRuleDetailsByHostingEnvironment request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForGetRuleDetailsByHostingEnvironment(resp *http.Response) (result GetRuleDetailsByHostingEnvironmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
