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

type ListOperationResponse struct {
	HttpResponse *http.Response
	Model        *Response
}

type ListOperationOptions struct {
	Aggregation         *string
	AutoAdjustTimegrain *bool
	Filter              *string
	Interval            *string
	Metricnames         *string
	Metricnamespace     *string
	Orderby             *string
	ResultType          *ResultType
	Rollupby            *string
	Timespan            *string
	Top                 *int64
	ValidateDimensions  *bool
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListOperationOptions) toQueryString() map[string]interface{} {
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

// List ...
func (c MetricsClient) List(ctx context.Context, id commonids.ScopeId, options ListOperationOptions) (result ListOperationResponse, err error) {
	req, err := c.preparerForList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "List", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "List", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "List", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForList prepares the List request.
func (c MetricsClient) preparerForList(ctx context.Context, id commonids.ScopeId, options ListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Insights/metrics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForList handles the response to the List request. The method always
// closes the http.Response Body.
func (c MetricsClient) responderForList(resp *http.Response) (result ListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
