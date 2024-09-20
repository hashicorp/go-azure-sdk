package driveitemsubscription

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

type ListDriveItemSubscriptionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Subscription
}

type ListDriveItemSubscriptionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Subscription
}

type ListDriveItemSubscriptionsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDriveItemSubscriptionsOperationOptions() ListDriveItemSubscriptionsOperationOptions {
	return ListDriveItemSubscriptionsOperationOptions{}
}

func (o ListDriveItemSubscriptionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemSubscriptionsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListDriveItemSubscriptionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemSubscriptionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemSubscriptionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemSubscriptions - Get subscriptions from groups. The set of subscriptions on the item. Only supported on
// the root of a drive.
func (c DriveItemSubscriptionClient) ListDriveItemSubscriptions(ctx context.Context, id stable.GroupIdDriveIdItemId, options ListDriveItemSubscriptionsOperationOptions) (result ListDriveItemSubscriptionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveItemSubscriptionsCustomPager{},
		Path:          fmt.Sprintf("%s/subscriptions", id.ID()),
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
		Values *[]stable.Subscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveItemSubscriptionsComplete retrieves all the results into a single object
func (c DriveItemSubscriptionClient) ListDriveItemSubscriptionsComplete(ctx context.Context, id stable.GroupIdDriveIdItemId, options ListDriveItemSubscriptionsOperationOptions) (ListDriveItemSubscriptionsCompleteResult, error) {
	return c.ListDriveItemSubscriptionsCompleteMatchingPredicate(ctx, id, options, SubscriptionOperationPredicate{})
}

// ListDriveItemSubscriptionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemSubscriptionClient) ListDriveItemSubscriptionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdDriveIdItemId, options ListDriveItemSubscriptionsOperationOptions, predicate SubscriptionOperationPredicate) (result ListDriveItemSubscriptionsCompleteResult, err error) {
	items := make([]stable.Subscription, 0)

	resp, err := c.ListDriveItemSubscriptions(ctx, id, options)
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

	result = ListDriveItemSubscriptionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
