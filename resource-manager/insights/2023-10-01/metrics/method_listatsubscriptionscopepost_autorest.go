package metrics

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

type ListAtSubscriptionScopePostOperationResponse struct {
	HttpResponse *http.Response
	Model        *Response
}

type ListAtSubscriptionScopePostOperationOptions struct {
	Aggregation         *string
	AutoAdjustTimegrain *bool
	Filter              *string
	Interval            *string
	Metricnames         *string
	Metricnamespace     *string
	Orderby             *string
	Region              *string
	ResultType          *MetricResultType
	Rollupby            *string
	Timespan            *string
	Top                 *int64
	ValidateDimensions  *bool
}

func DefaultListAtSubscriptionScopePostOperationOptions() ListAtSubscriptionScopePostOperationOptions {
	return ListAtSubscriptionScopePostOperationOptions{}
}

func (o ListAtSubscriptionScopePostOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtSubscriptionScopePostOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Aggregation != nil {
		out["aggregation"] = *o.Aggregation
	}

	if o.AutoAdjustTimegrain != nil {
		out["AutoAdjustTimegrain"] = *o.AutoAdjustTimegrain
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Interval != nil {
		out["interval"] = *o.Interval
	}

	if o.Metricnames != nil {
		out["metricnames"] = *o.Metricnames
	}

	if o.Metricnamespace != nil {
		out["metricnamespace"] = *o.Metricnamespace
	}

	if o.Orderby != nil {
		out["orderby"] = *o.Orderby
	}

	if o.Region != nil {
		out["region"] = *o.Region
	}

	if o.ResultType != nil {
		out["resultType"] = *o.ResultType
	}

	if o.Rollupby != nil {
		out["rollupby"] = *o.Rollupby
	}

	if o.Timespan != nil {
		out["timespan"] = *o.Timespan
	}

	if o.Top != nil {
		out["top"] = *o.Top
	}

	if o.ValidateDimensions != nil {
		out["ValidateDimensions"] = *o.ValidateDimensions
	}

	return out
}

// ListAtSubscriptionScopePost ...
func (c MetricsClient) ListAtSubscriptionScopePost(ctx context.Context, id commonids.SubscriptionId, input SubscriptionScopeMetricsRequestBodyParameters, options ListAtSubscriptionScopePostOperationOptions) (result ListAtSubscriptionScopePostOperationResponse, err error) {
	req, err := c.preparerForListAtSubscriptionScopePost(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "ListAtSubscriptionScopePost", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "ListAtSubscriptionScopePost", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListAtSubscriptionScopePost(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "ListAtSubscriptionScopePost", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListAtSubscriptionScopePost prepares the ListAtSubscriptionScopePost request.
func (c MetricsClient) preparerForListAtSubscriptionScopePost(ctx context.Context, id commonids.SubscriptionId, input SubscriptionScopeMetricsRequestBodyParameters, options ListAtSubscriptionScopePostOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Insights/metrics", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListAtSubscriptionScopePost handles the response to the ListAtSubscriptionScopePost request. The method always
// closes the http.Response Body.
func (c MetricsClient) responderForListAtSubscriptionScopePost(resp *http.Response) (result ListAtSubscriptionScopePostOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
