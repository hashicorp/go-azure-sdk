package metrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveMetricGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
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

func (o PredictiveMetricGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PredictiveMetricGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o PredictiveMetricGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Aggregation != nil {
		out.Append("aggregation", fmt.Sprintf("%v", *o.Aggregation))
	}
	if o.Interval != nil {
		out.Append("interval", fmt.Sprintf("%v", *o.Interval))
	}
	if o.MetricName != nil {
		out.Append("metricName", fmt.Sprintf("%v", *o.MetricName))
	}
	if o.MetricNamespace != nil {
		out.Append("metricNamespace", fmt.Sprintf("%v", *o.MetricNamespace))
	}
	if o.Timespan != nil {
		out.Append("timespan", fmt.Sprintf("%v", *o.Timespan))
	}
	return &out
}

// PredictiveMetricGet ...
func (c MetricsClient) PredictiveMetricGet(ctx context.Context, id AutoScaleSettingId, options PredictiveMetricGetOperationOptions) (result PredictiveMetricGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/predictiveMetrics", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model PredictiveResponse
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
