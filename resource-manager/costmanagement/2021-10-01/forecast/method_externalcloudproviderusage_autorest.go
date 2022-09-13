package forecast

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalCloudProviderUsageOperationResponse struct {
	HttpResponse *http.Response
	Model        *ForecastResult
}

type ExternalCloudProviderUsageOperationOptions struct {
	Filter *string
}

func DefaultExternalCloudProviderUsageOperationOptions() ExternalCloudProviderUsageOperationOptions {
	return ExternalCloudProviderUsageOperationOptions{}
}

func (o ExternalCloudProviderUsageOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExternalCloudProviderUsageOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ExternalCloudProviderUsage ...
func (c ForecastClient) ExternalCloudProviderUsage(ctx context.Context, id ExternalCloudProviderTypeId, input ForecastDefinition, options ExternalCloudProviderUsageOperationOptions) (result ExternalCloudProviderUsageOperationResponse, err error) {
	req, err := c.preparerForExternalCloudProviderUsage(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "ExternalCloudProviderUsage", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "ExternalCloudProviderUsage", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExternalCloudProviderUsage(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "ExternalCloudProviderUsage", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExternalCloudProviderUsage prepares the ExternalCloudProviderUsage request.
func (c ForecastClient) preparerForExternalCloudProviderUsage(ctx context.Context, id ExternalCloudProviderTypeId, input ForecastDefinition, options ExternalCloudProviderUsageOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/forecast", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForExternalCloudProviderUsage handles the response to the ExternalCloudProviderUsage request. The method always
// closes the http.Response Body.
func (c ForecastClient) responderForExternalCloudProviderUsage(resp *http.Response) (result ExternalCloudProviderUsageOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
