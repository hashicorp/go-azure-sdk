package iotsecuritysolutionsanalytics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregatedAlertsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IoTSecurityAggregatedAlert
}

type AggregatedAlertsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IoTSecurityAggregatedAlert
}

type AggregatedAlertsListOperationOptions struct {
	Top *int64
}

func DefaultAggregatedAlertsListOperationOptions() AggregatedAlertsListOperationOptions {
	return AggregatedAlertsListOperationOptions{}
}

func (o AggregatedAlertsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AggregatedAlertsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AggregatedAlertsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// AggregatedAlertsList ...
func (c IoTSecuritySolutionsAnalyticsClient) AggregatedAlertsList(ctx context.Context, id IotSecuritySolutionId, options AggregatedAlertsListOperationOptions) (result AggregatedAlertsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/analyticsModels/default/aggregatedAlerts", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]IoTSecurityAggregatedAlert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AggregatedAlertsListComplete retrieves all the results into a single object
func (c IoTSecuritySolutionsAnalyticsClient) AggregatedAlertsListComplete(ctx context.Context, id IotSecuritySolutionId, options AggregatedAlertsListOperationOptions) (AggregatedAlertsListCompleteResult, error) {
	return c.AggregatedAlertsListCompleteMatchingPredicate(ctx, id, options, IoTSecurityAggregatedAlertOperationPredicate{})
}

// AggregatedAlertsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IoTSecuritySolutionsAnalyticsClient) AggregatedAlertsListCompleteMatchingPredicate(ctx context.Context, id IotSecuritySolutionId, options AggregatedAlertsListOperationOptions, predicate IoTSecurityAggregatedAlertOperationPredicate) (result AggregatedAlertsListCompleteResult, err error) {
	items := make([]IoTSecurityAggregatedAlert, 0)

	resp, err := c.AggregatedAlertsList(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = AggregatedAlertsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
