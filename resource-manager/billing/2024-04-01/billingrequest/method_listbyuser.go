package billingrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRequest
}

type ListByUserCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRequest
}

type ListByUserOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *string
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListByUserOperationOptions() ListByUserOperationOptions {
	return ListByUserOperationOptions{}
}

func (o ListByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByUserCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByUserCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByUser ...
func (c BillingRequestClient) ListByUser(ctx context.Context, options ListByUserOperationOptions) (result ListByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByUserCustomPager{},
		Path:          "/providers/Microsoft.Billing/billingRequests",
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
		Values *[]BillingRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByUserComplete retrieves all the results into a single object
func (c BillingRequestClient) ListByUserComplete(ctx context.Context, options ListByUserOperationOptions) (ListByUserCompleteResult, error) {
	return c.ListByUserCompleteMatchingPredicate(ctx, options, BillingRequestOperationPredicate{})
}

// ListByUserCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRequestClient) ListByUserCompleteMatchingPredicate(ctx context.Context, options ListByUserOperationOptions, predicate BillingRequestOperationPredicate) (result ListByUserCompleteResult, err error) {
	items := make([]BillingRequest, 0)

	resp, err := c.ListByUser(ctx, options)
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

	result = ListByUserCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
