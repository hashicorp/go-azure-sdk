package groupquotasubscriptionrequeststatuses

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaSubscriptionRequestsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GroupQuotaSubscriptionRequestStatus
}

type GroupQuotaSubscriptionRequestsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GroupQuotaSubscriptionRequestStatus
}

type GroupQuotaSubscriptionRequestsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GroupQuotaSubscriptionRequestsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GroupQuotaSubscriptionRequestsList ...
func (c GroupQuotaSubscriptionRequestStatusesClient) GroupQuotaSubscriptionRequestsList(ctx context.Context, id GroupQuotaId) (result GroupQuotaSubscriptionRequestsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GroupQuotaSubscriptionRequestsListCustomPager{},
		Path:       fmt.Sprintf("%s/subscriptionRequests", id.ID()),
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
		Values *[]GroupQuotaSubscriptionRequestStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GroupQuotaSubscriptionRequestsListComplete retrieves all the results into a single object
func (c GroupQuotaSubscriptionRequestStatusesClient) GroupQuotaSubscriptionRequestsListComplete(ctx context.Context, id GroupQuotaId) (GroupQuotaSubscriptionRequestsListCompleteResult, error) {
	return c.GroupQuotaSubscriptionRequestsListCompleteMatchingPredicate(ctx, id, GroupQuotaSubscriptionRequestStatusOperationPredicate{})
}

// GroupQuotaSubscriptionRequestsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupQuotaSubscriptionRequestStatusesClient) GroupQuotaSubscriptionRequestsListCompleteMatchingPredicate(ctx context.Context, id GroupQuotaId, predicate GroupQuotaSubscriptionRequestStatusOperationPredicate) (result GroupQuotaSubscriptionRequestsListCompleteResult, err error) {
	items := make([]GroupQuotaSubscriptionRequestStatus, 0)

	resp, err := c.GroupQuotaSubscriptionRequestsList(ctx, id)
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

	result = GroupQuotaSubscriptionRequestsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
