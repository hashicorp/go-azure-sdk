package recommendations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRuleDetailsByWebAppOperationResponse struct {
	HttpResponse *http.Response
	Model        *RecommendationRule
}

type GetRuleDetailsByWebAppOperationOptions struct {
	RecommendationId *string
	UpdateSeen       *bool
}

func DefaultGetRuleDetailsByWebAppOperationOptions() GetRuleDetailsByWebAppOperationOptions {
	return GetRuleDetailsByWebAppOperationOptions{}
}

func (o GetRuleDetailsByWebAppOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetRuleDetailsByWebAppOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RecommendationId != nil {
		out["recommendationId"] = *o.RecommendationId
	}

	if o.UpdateSeen != nil {
		out["updateSeen"] = *o.UpdateSeen
	}

	return out
}

// GetRuleDetailsByWebApp ...
func (c RecommendationsClient) GetRuleDetailsByWebApp(ctx context.Context, id SiteRecommendationId, options GetRuleDetailsByWebAppOperationOptions) (result GetRuleDetailsByWebAppOperationResponse, err error) {
	req, err := c.preparerForGetRuleDetailsByWebApp(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByWebApp", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByWebApp", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRuleDetailsByWebApp(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "GetRuleDetailsByWebApp", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRuleDetailsByWebApp prepares the GetRuleDetailsByWebApp request.
func (c RecommendationsClient) preparerForGetRuleDetailsByWebApp(ctx context.Context, id SiteRecommendationId, options GetRuleDetailsByWebAppOperationOptions) (*http.Request, error) {
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

// responderForGetRuleDetailsByWebApp handles the response to the GetRuleDetailsByWebApp request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForGetRuleDetailsByWebApp(resp *http.Response) (result GetRuleDetailsByWebAppOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
