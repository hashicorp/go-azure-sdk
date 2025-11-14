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

type EventsListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EventSummary
}

type EventsListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EventSummary
}

type EventsListByBillingProfileOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultEventsListByBillingProfileOperationOptions() EventsListByBillingProfileOperationOptions {
	return EventsListByBillingProfileOperationOptions{}
}

func (o EventsListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EventsListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o EventsListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.StartDate != nil {
		out.Append("startDate", fmt.Sprintf("%v", *o.StartDate))
	}
	return &out
}

type EventsListByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EventsListByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EventsListByBillingProfile ...
func (c OpenapisClient) EventsListByBillingProfile(ctx context.Context, id BillingProfileId, options EventsListByBillingProfileOperationOptions) (result EventsListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &EventsListByBillingProfileCustomPager{},
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

// EventsListByBillingProfileComplete retrieves all the results into a single object
func (c OpenapisClient) EventsListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options EventsListByBillingProfileOperationOptions) (EventsListByBillingProfileCompleteResult, error) {
	return c.EventsListByBillingProfileCompleteMatchingPredicate(ctx, id, options, EventSummaryOperationPredicate{})
}

// EventsListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) EventsListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options EventsListByBillingProfileOperationOptions, predicate EventSummaryOperationPredicate) (result EventsListByBillingProfileCompleteResult, err error) {
	items := make([]EventSummary, 0)

	resp, err := c.EventsListByBillingProfile(ctx, id, options)
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

	result = EventsListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
