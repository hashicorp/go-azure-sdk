package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventsListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EventSummary
}

type EventsListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EventSummary
}

type EventsListByBillingAccountOperationOptions struct {
	Filter *string
}

func DefaultEventsListByBillingAccountOperationOptions() EventsListByBillingAccountOperationOptions {
	return EventsListByBillingAccountOperationOptions{}
}

func (o EventsListByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EventsListByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o EventsListByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type EventsListByBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EventsListByBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EventsListByBillingAccount ...
func (c OpenapisClient) EventsListByBillingAccount(ctx context.Context, id BillingAccountId, options EventsListByBillingAccountOperationOptions) (result EventsListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &EventsListByBillingAccountCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/events", id.ID()),
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
		Values *[]EventSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// EventsListByBillingAccountComplete retrieves all the results into a single object
func (c OpenapisClient) EventsListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options EventsListByBillingAccountOperationOptions) (EventsListByBillingAccountCompleteResult, error) {
	return c.EventsListByBillingAccountCompleteMatchingPredicate(ctx, id, options, EventSummaryOperationPredicate{})
}

// EventsListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) EventsListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options EventsListByBillingAccountOperationOptions, predicate EventSummaryOperationPredicate) (result EventsListByBillingAccountCompleteResult, err error) {
	items := make([]EventSummary, 0)

	resp, err := c.EventsListByBillingAccount(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = EventsListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
