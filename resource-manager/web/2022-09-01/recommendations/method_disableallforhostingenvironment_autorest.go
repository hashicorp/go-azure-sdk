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

type DisableAllForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
}

type DisableAllForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultDisableAllForHostingEnvironmentOperationOptions() DisableAllForHostingEnvironmentOperationOptions {
	return DisableAllForHostingEnvironmentOperationOptions{}
}

func (o DisableAllForHostingEnvironmentOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DisableAllForHostingEnvironmentOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EnvironmentName != nil {
		out["environmentName"] = *o.EnvironmentName
	}

	return out
}

// DisableAllForHostingEnvironment ...
func (c RecommendationsClient) DisableAllForHostingEnvironment(ctx context.Context, id HostingEnvironmentId, options DisableAllForHostingEnvironmentOperationOptions) (result DisableAllForHostingEnvironmentOperationResponse, err error) {
	req, err := c.preparerForDisableAllForHostingEnvironment(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForHostingEnvironment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForHostingEnvironment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisableAllForHostingEnvironment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForHostingEnvironment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisableAllForHostingEnvironment prepares the DisableAllForHostingEnvironment request.
func (c RecommendationsClient) preparerForDisableAllForHostingEnvironment(ctx context.Context, id HostingEnvironmentId, options DisableAllForHostingEnvironmentOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/recommendations/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisableAllForHostingEnvironment handles the response to the DisableAllForHostingEnvironment request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForDisableAllForHostingEnvironment(resp *http.Response) (result DisableAllForHostingEnvironmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
