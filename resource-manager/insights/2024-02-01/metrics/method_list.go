package metrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
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

func (o ListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Aggregation != nil {
		out.Append("aggregation", fmt.Sprintf("%v", *o.Aggregation))
	}
	if o.AutoAdjustTimegrain != nil {
		out.Append("AutoAdjustTimegrain", fmt.Sprintf("%v", *o.AutoAdjustTimegrain))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Interval != nil {
		out.Append("interval", fmt.Sprintf("%v", *o.Interval))
	}
	if o.Metricnames != nil {
		out.Append("metricnames", fmt.Sprintf("%v", *o.Metricnames))
	}
	if o.Metricnamespace != nil {
		out.Append("metricnamespace", fmt.Sprintf("%v", *o.Metricnamespace))
	}
	if o.Orderby != nil {
		out.Append("orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.ResultType != nil {
		out.Append("resultType", fmt.Sprintf("%v", *o.ResultType))
	}
	if o.Rollupby != nil {
		out.Append("rollupby", fmt.Sprintf("%v", *o.Rollupby))
	}
	if o.Timespan != nil {
		out.Append("timespan", fmt.Sprintf("%v", *o.Timespan))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	if o.ValidateDimensions != nil {
		out.Append("ValidateDimensions", fmt.Sprintf("%v", *o.ValidateDimensions))
	}
	return &out
}

// List ...
func (c MetricsClient) List(ctx context.Context, id commonids.ScopeId, options ListOperationOptions) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Insights/metrics", id.ID()),
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

	var model Response
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
