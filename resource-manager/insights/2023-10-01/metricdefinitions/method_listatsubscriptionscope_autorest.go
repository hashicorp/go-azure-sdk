package metricdefinitions

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

type ListAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionScopeMetricDefinitionCollection
}

type ListAtSubscriptionScopeOperationOptions struct {
	Metricnamespace *string
	Region          *string
}

func DefaultListAtSubscriptionScopeOperationOptions() ListAtSubscriptionScopeOperationOptions {
	return ListAtSubscriptionScopeOperationOptions{}
}

func (o ListAtSubscriptionScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtSubscriptionScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Metricnamespace != nil {
		out["metricnamespace"] = *o.Metricnamespace
	}

	if o.Region != nil {
		out["region"] = *o.Region
	}

	return out
}

// ListAtSubscriptionScope ...
func (c MetricDefinitionsClient) ListAtSubscriptionScope(ctx context.Context, id commonids.SubscriptionId, options ListAtSubscriptionScopeOperationOptions) (result ListAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForListAtSubscriptionScope(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricdefinitions.MetricDefinitionsClient", "ListAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricdefinitions.MetricDefinitionsClient", "ListAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricdefinitions.MetricDefinitionsClient", "ListAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListAtSubscriptionScope prepares the ListAtSubscriptionScope request.
func (c MetricDefinitionsClient) preparerForListAtSubscriptionScope(ctx context.Context, id commonids.SubscriptionId, options ListAtSubscriptionScopeOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Insights/metricDefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListAtSubscriptionScope handles the response to the ListAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c MetricDefinitionsClient) responderForListAtSubscriptionScope(resp *http.Response) (result ListAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
