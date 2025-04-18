package groupquotassubscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaSubscriptionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GroupQuotaSubscriptionId
}

type GroupQuotaSubscriptionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GroupQuotaSubscriptionId
}

type GroupQuotaSubscriptionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotaSubscriptionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotaSubscriptionsList ...
func (c GroupQuotasSubscriptionsClient) GroupQuotaSubscriptionsList(ctx context.Context, id GroupQuotaId) (result GroupQuotaSubscriptionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GroupQuotaSubscriptionsListCustomPager{},
		Path:       fmt.Sprintf("%s/subscriptions", id.ID()),
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
		Values *[]GroupQuotaSubscriptionId `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotaSubscriptionsListComplete retrieves all the results into a single object
func (c GroupQuotasSubscriptionsClient) GroupQuotaSubscriptionsListComplete(ctx context.Context, id GroupQuotaId) (GroupQuotaSubscriptionsListCompleteResult, error) {
	return c.GroupQuotaSubscriptionsListCompleteMatchingPredicate(ctx, id, GroupQuotaSubscriptionIdOperationPredicate{})
}

// GroupQuotaSubscriptionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupQuotasSubscriptionsClient) GroupQuotaSubscriptionsListCompleteMatchingPredicate(ctx context.Context, id GroupQuotaId, predicate GroupQuotaSubscriptionIdOperationPredicate) (result GroupQuotaSubscriptionsListCompleteResult, err error) {
	items := make([]GroupQuotaSubscriptionId, 0)

	resp, err := c.GroupQuotaSubscriptionsList(ctx, id)
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

	result = GroupQuotaSubscriptionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
