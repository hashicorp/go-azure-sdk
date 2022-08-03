package metrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveMetricGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PredictiveResponse
}

type PredictiveMetricGetOperationOptions struct {
	Aggregation     *string
	Interval        *string
	MetricName      *string
	MetricNamespace *string
	Timespan        *string
}

func DefaultPredictiveMetricGetOperationOptions() PredictiveMetricGetOperationOptions {
	return PredictiveMetricGetOperationOptions{}
}

func (o PredictiveMetricGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o PredictiveMetricGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Aggregation != nil {
		out["aggregation"] = *o.Aggregation
	}

	if o.Interval != nil {
		out["interval"] = *o.Interval
	}

	if o.MetricName != nil {
		out["metricName"] = *o.MetricName
	}

	if o.MetricNamespace != nil {
		out["metricNamespace"] = *o.MetricNamespace
	}

	if o.Timespan != nil {
		out["timespan"] = *o.Timespan
	}

	return out
}

// PredictiveMetricGet ...
func (c MetricsClient) PredictiveMetricGet(ctx context.Context, id AutoScaleSettingId, options PredictiveMetricGetOperationOptions) (result PredictiveMetricGetOperationResponse, err error) {
	req, err := c.preparerForPredictiveMetricGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "PredictiveMetricGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "PredictiveMetricGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPredictiveMetricGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metrics.MetricsClient", "PredictiveMetricGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPredictiveMetricGet prepares the PredictiveMetricGet request.
func (c MetricsClient) preparerForPredictiveMetricGet(ctx context.Context, id AutoScaleSettingId, options PredictiveMetricGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/predictiveMetrics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPredictiveMetricGet handles the response to the PredictiveMetricGet request. The method always
// closes the http.Response Body.
func (c MetricsClient) responderForPredictiveMetricGet(resp *http.Response) (result PredictiveMetricGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
