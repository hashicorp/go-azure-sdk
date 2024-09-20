package billingsubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByCustomerAtBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingSubscription
}

type ListByCustomerAtBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingSubscription
}

type ListByCustomerAtBillingAccountOperationOptions struct {
	Count          *bool
	Expand         *string
	Filter         *string
	IncludeDeleted *bool
	OrderBy        *string
	Search         *string
	Skip           *int64
	Top            *int64
}

func DefaultListByCustomerAtBillingAccountOperationOptions() ListByCustomerAtBillingAccountOperationOptions {
	return ListByCustomerAtBillingAccountOperationOptions{}
}

func (o ListByCustomerAtBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByCustomerAtBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByCustomerAtBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.IncludeDeleted != nil {
		out.Append("includeDeleted", fmt.Sprintf("%v", *o.IncludeDeleted))
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

type ListByCustomerAtBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByCustomerAtBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCustomerAtBillingAccount ...
func (c BillingSubscriptionClient) ListByCustomerAtBillingAccount(ctx context.Context, id CustomerId, options ListByCustomerAtBillingAccountOperationOptions) (result ListByCustomerAtBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByCustomerAtBillingAccountCustomPager{},
		Path:          fmt.Sprintf("%s/billingSubscriptions", id.ID()),
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
		Values *[]BillingSubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByCustomerAtBillingAccountComplete retrieves all the results into a single object
func (c BillingSubscriptionClient) ListByCustomerAtBillingAccountComplete(ctx context.Context, id CustomerId, options ListByCustomerAtBillingAccountOperationOptions) (ListByCustomerAtBillingAccountCompleteResult, error) {
	return c.ListByCustomerAtBillingAccountCompleteMatchingPredicate(ctx, id, options, BillingSubscriptionOperationPredicate{})
}

// ListByCustomerAtBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingSubscriptionClient) ListByCustomerAtBillingAccountCompleteMatchingPredicate(ctx context.Context, id CustomerId, options ListByCustomerAtBillingAccountOperationOptions, predicate BillingSubscriptionOperationPredicate) (result ListByCustomerAtBillingAccountCompleteResult, err error) {
	items := make([]BillingSubscription, 0)

	resp, err := c.ListByCustomerAtBillingAccount(ctx, id, options)
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

	result = ListByCustomerAtBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
