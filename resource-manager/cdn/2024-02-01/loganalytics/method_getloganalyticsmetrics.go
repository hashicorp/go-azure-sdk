package loganalytics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLogAnalyticsMetricsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MetricsResponse
}

type GetLogAnalyticsMetricsOperationOptions struct {
	Continents       *[]string
	CountryOrRegions *[]string
	CustomDomains    *[]string
	DateTimeBegin    *string
	DateTimeEnd      *string
	Granularity      *LogMetricsGranularity
	GroupBy          *[]string
	Metrics          *[]string
	Protocols        *[]string
}

func DefaultGetLogAnalyticsMetricsOperationOptions() GetLogAnalyticsMetricsOperationOptions {
	return GetLogAnalyticsMetricsOperationOptions{}
}

func (o GetLogAnalyticsMetricsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLogAnalyticsMetricsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetLogAnalyticsMetricsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Continents != nil {
		out.Append("continents", fmt.Sprintf("%v", *o.Continents))
	}
	if o.CountryOrRegions != nil {
		out.Append("countryOrRegions", fmt.Sprintf("%v", *o.CountryOrRegions))
	}
	if o.CustomDomains != nil {
		out.Append("customDomains", fmt.Sprintf("%v", *o.CustomDomains))
	}
	if o.DateTimeBegin != nil {
		out.Append("dateTimeBegin", fmt.Sprintf("%v", *o.DateTimeBegin))
	}
	if o.DateTimeEnd != nil {
		out.Append("dateTimeEnd", fmt.Sprintf("%v", *o.DateTimeEnd))
	}
	if o.Granularity != nil {
		out.Append("granularity", fmt.Sprintf("%v", *o.Granularity))
	}
	if o.GroupBy != nil {
		out.Append("groupBy", fmt.Sprintf("%v", *o.GroupBy))
	}
	if o.Metrics != nil {
		out.Append("metrics", fmt.Sprintf("%v", *o.Metrics))
	}
	if o.Protocols != nil {
		out.Append("protocols", fmt.Sprintf("%v", *o.Protocols))
	}
	return &out
}

// GetLogAnalyticsMetrics ...
func (c LogAnalyticsClient) GetLogAnalyticsMetrics(ctx context.Context, id ProfileId, options GetLogAnalyticsMetricsOperationOptions) (result GetLogAnalyticsMetricsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/getLogAnalyticsMetrics", id.ID()),
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

	var model MetricsResponse
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
