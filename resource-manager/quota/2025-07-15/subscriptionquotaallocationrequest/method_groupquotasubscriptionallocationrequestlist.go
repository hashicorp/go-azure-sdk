package subscriptionquotaallocationrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaSubscriptionAllocationRequestListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]QuotaAllocationRequestStatus
}

type GroupQuotaSubscriptionAllocationRequestListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []QuotaAllocationRequestStatus
}

type GroupQuotaSubscriptionAllocationRequestListOperationOptions struct {
	Filter *string
}

func DefaultGroupQuotaSubscriptionAllocationRequestListOperationOptions() GroupQuotaSubscriptionAllocationRequestListOperationOptions {
	return GroupQuotaSubscriptionAllocationRequestListOperationOptions{}
}

func (o GroupQuotaSubscriptionAllocationRequestListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GroupQuotaSubscriptionAllocationRequestListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GroupQuotaSubscriptionAllocationRequestListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type GroupQuotaSubscriptionAllocationRequestListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotaSubscriptionAllocationRequestListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotaSubscriptionAllocationRequestList ...
func (c SubscriptionQuotaAllocationRequestClient) GroupQuotaSubscriptionAllocationRequestList(ctx context.Context, id GroupQuotaResourceProviderId, options GroupQuotaSubscriptionAllocationRequestListOperationOptions) (result GroupQuotaSubscriptionAllocationRequestListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GroupQuotaSubscriptionAllocationRequestListCustomPager{},
		Path:          fmt.Sprintf("%s/quotaAllocationRequests", id.ID()),
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
		Values *[]QuotaAllocationRequestStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotaSubscriptionAllocationRequestListComplete retrieves all the results into a single object
func (c SubscriptionQuotaAllocationRequestClient) GroupQuotaSubscriptionAllocationRequestListComplete(ctx context.Context, id GroupQuotaResourceProviderId, options GroupQuotaSubscriptionAllocationRequestListOperationOptions) (GroupQuotaSubscriptionAllocationRequestListCompleteResult, error) {
	return c.GroupQuotaSubscriptionAllocationRequestListCompleteMatchingPredicate(ctx, id, options, QuotaAllocationRequestStatusOperationPredicate{})
}

// GroupQuotaSubscriptionAllocationRequestListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionQuotaAllocationRequestClient) GroupQuotaSubscriptionAllocationRequestListCompleteMatchingPredicate(ctx context.Context, id GroupQuotaResourceProviderId, options GroupQuotaSubscriptionAllocationRequestListOperationOptions, predicate QuotaAllocationRequestStatusOperationPredicate) (result GroupQuotaSubscriptionAllocationRequestListCompleteResult, err error) {
	items := make([]QuotaAllocationRequestStatus, 0)

	resp, err := c.GroupQuotaSubscriptionAllocationRequestList(ctx, id, options)
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

	result = GroupQuotaSubscriptionAllocationRequestListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
