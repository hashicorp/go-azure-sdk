package driverootsubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveRootSubscriptionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Subscription
}

type ListDriveRootSubscriptionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Subscription
}

type ListDriveRootSubscriptionsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDriveRootSubscriptionsOperationOptions() ListDriveRootSubscriptionsOperationOptions {
	return ListDriveRootSubscriptionsOperationOptions{}
}

func (o ListDriveRootSubscriptionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootSubscriptionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootSubscriptionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootSubscriptionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootSubscriptionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootSubscriptions - Get subscriptions from groups. The set of subscriptions on the item. Only supported on
// the root of a drive.
func (c DriveRootSubscriptionClient) ListDriveRootSubscriptions(ctx context.Context, id beta.GroupIdDriveId, options ListDriveRootSubscriptionsOperationOptions) (result ListDriveRootSubscriptionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootSubscriptionsCustomPager{},
		Path:          fmt.Sprintf("%s/root/subscriptions", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.Subscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveRootSubscriptionsComplete retrieves all the results into a single object
func (c DriveRootSubscriptionClient) ListDriveRootSubscriptionsComplete(ctx context.Context, id beta.GroupIdDriveId, options ListDriveRootSubscriptionsOperationOptions) (ListDriveRootSubscriptionsCompleteResult, error) {
	return c.ListDriveRootSubscriptionsCompleteMatchingPredicate(ctx, id, options, SubscriptionOperationPredicate{})
}

// ListDriveRootSubscriptionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootSubscriptionClient) ListDriveRootSubscriptionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdDriveId, options ListDriveRootSubscriptionsOperationOptions, predicate SubscriptionOperationPredicate) (result ListDriveRootSubscriptionsCompleteResult, err error) {
	items := make([]beta.Subscription, 0)

	resp, err := c.ListDriveRootSubscriptions(ctx, id, options)
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

	result = ListDriveRootSubscriptionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
