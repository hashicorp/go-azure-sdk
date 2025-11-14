package subscriptionquotaitems

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceQuotaLimitsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionQuotaItem
}

type NetAppResourceQuotaLimitsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubscriptionQuotaItem
}

type NetAppResourceQuotaLimitsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NetAppResourceQuotaLimitsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NetAppResourceQuotaLimitsList ...
func (c SubscriptionQuotaItemsClient) NetAppResourceQuotaLimitsList(ctx context.Context, id LocationId) (result NetAppResourceQuotaLimitsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NetAppResourceQuotaLimitsListCustomPager{},
		Path:       fmt.Sprintf("%s/quotaLimits", id.ID()),
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
		Values *[]SubscriptionQuotaItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NetAppResourceQuotaLimitsListComplete retrieves all the results into a single object
func (c SubscriptionQuotaItemsClient) NetAppResourceQuotaLimitsListComplete(ctx context.Context, id LocationId) (NetAppResourceQuotaLimitsListCompleteResult, error) {
	return c.NetAppResourceQuotaLimitsListCompleteMatchingPredicate(ctx, id, SubscriptionQuotaItemOperationPredicate{})
}

// NetAppResourceQuotaLimitsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionQuotaItemsClient) NetAppResourceQuotaLimitsListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate SubscriptionQuotaItemOperationPredicate) (result NetAppResourceQuotaLimitsListCompleteResult, err error) {
	items := make([]SubscriptionQuotaItem, 0)

	resp, err := c.NetAppResourceQuotaLimitsList(ctx, id)
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

	result = NetAppResourceQuotaLimitsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
