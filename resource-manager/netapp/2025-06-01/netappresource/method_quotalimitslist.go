package netappresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaLimitsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionQuotaItem
}

type QuotaLimitsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubscriptionQuotaItem
}

type QuotaLimitsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *QuotaLimitsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// QuotaLimitsList ...
func (c NetAppResourceClient) QuotaLimitsList(ctx context.Context, id LocationId) (result QuotaLimitsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &QuotaLimitsListCustomPager{},
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

// QuotaLimitsListComplete retrieves all the results into a single object
func (c NetAppResourceClient) QuotaLimitsListComplete(ctx context.Context, id LocationId) (QuotaLimitsListCompleteResult, error) {
	return c.QuotaLimitsListCompleteMatchingPredicate(ctx, id, SubscriptionQuotaItemOperationPredicate{})
}

// QuotaLimitsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetAppResourceClient) QuotaLimitsListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate SubscriptionQuotaItemOperationPredicate) (result QuotaLimitsListCompleteResult, err error) {
	items := make([]SubscriptionQuotaItem, 0)

	resp, err := c.QuotaLimitsList(ctx, id)
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

	result = QuotaLimitsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
