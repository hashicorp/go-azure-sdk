package policyevents

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

type ListQueryResultsForResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyEvent
}

type ListQueryResultsForResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyEvent
}

type ListQueryResultsForResourceOperationOptions struct {
	Apply   *string
	Expand  *string
	Filter  *string
	From    *string
	Orderby *string
	Select  *string
	To      *string
	Top     *int64
}

func DefaultListQueryResultsForResourceOperationOptions() ListQueryResultsForResourceOperationOptions {
	return ListQueryResultsForResourceOperationOptions{}
}

func (o ListQueryResultsForResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListQueryResultsForResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListQueryResultsForResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Apply != nil {
		out.Append("$apply", fmt.Sprintf("%v", *o.Apply))
	}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.From != nil {
		out.Append("$from", fmt.Sprintf("%v", *o.From))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.To != nil {
		out.Append("$to", fmt.Sprintf("%v", *o.To))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListQueryResultsForResourceCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListQueryResultsForResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListQueryResultsForResource ...
func (c PolicyEventsClient) ListQueryResultsForResource(ctx context.Context, id commonids.ScopeId, options ListQueryResultsForResourceOperationOptions) (result ListQueryResultsForResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListQueryResultsForResourceCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/policyEvents/default/queryResults", id.ID()),
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
		Values *[]PolicyEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListQueryResultsForResourceComplete retrieves all the results into a single object
func (c PolicyEventsClient) ListQueryResultsForResourceComplete(ctx context.Context, id commonids.ScopeId, options ListQueryResultsForResourceOperationOptions) (ListQueryResultsForResourceCompleteResult, error) {
	return c.ListQueryResultsForResourceCompleteMatchingPredicate(ctx, id, options, PolicyEventOperationPredicate{})
}

// ListQueryResultsForResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyEventsClient) ListQueryResultsForResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ListQueryResultsForResourceOperationOptions, predicate PolicyEventOperationPredicate) (result ListQueryResultsForResourceCompleteResult, err error) {
	items := make([]PolicyEvent, 0)

	resp, err := c.ListQueryResultsForResource(ctx, id, options)
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

	result = ListQueryResultsForResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
