package forecast

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

type UsageOperationResponse struct {
	HttpResponse *http.Response
	Model        *ForecastResult
}

type UsageOperationOptions struct {
	Filter *string
}

func DefaultUsageOperationOptions() UsageOperationOptions {
	return UsageOperationOptions{}
}

func (o UsageOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o UsageOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// Usage ...
func (c ForecastClient) Usage(ctx context.Context, id commonids.ScopeId, input ForecastDefinition, options UsageOperationOptions) (result UsageOperationResponse, err error) {
	req, err := c.preparerForUsage(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "Usage", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "Usage", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUsage(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "forecast.ForecastClient", "Usage", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUsage prepares the Usage request.
func (c ForecastClient) preparerForUsage(ctx context.Context, id commonids.ScopeId, input ForecastDefinition, options UsageOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/forecast", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUsage handles the response to the Usage request. The method always
// closes the http.Response Body.
func (c ForecastClient) responderForUsage(resp *http.Response) (result UsageOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
