package subscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSubscriptionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CompanySubscription
}

type ListSubscriptionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CompanySubscription
}

type ListSubscriptionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSubscriptionsOperationOptions() ListSubscriptionsOperationOptions {
	return ListSubscriptionsOperationOptions{}
}

func (o ListSubscriptionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSubscriptionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSubscriptionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSubscriptionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSubscriptionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSubscriptions - List subscriptions. Get the list of commercial subscriptions that an organization acquired.
func (c SubscriptionClient) ListSubscriptions(ctx context.Context, options ListSubscriptionsOperationOptions) (result ListSubscriptionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSubscriptionsCustomPager{},
		Path:          "/directory/subscriptions",
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
		Values *[]stable.CompanySubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSubscriptionsComplete retrieves all the results into a single object
func (c SubscriptionClient) ListSubscriptionsComplete(ctx context.Context, options ListSubscriptionsOperationOptions) (ListSubscriptionsCompleteResult, error) {
	return c.ListSubscriptionsCompleteMatchingPredicate(ctx, options, CompanySubscriptionOperationPredicate{})
}

// ListSubscriptionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionClient) ListSubscriptionsCompleteMatchingPredicate(ctx context.Context, options ListSubscriptionsOperationOptions, predicate CompanySubscriptionOperationPredicate) (result ListSubscriptionsCompleteResult, err error) {
	items := make([]stable.CompanySubscription, 0)

	resp, err := c.ListSubscriptions(ctx, options)
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

	result = ListSubscriptionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
