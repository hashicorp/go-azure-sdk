package sharesubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerSourceDataSetsListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ConsumerSourceDataSet
}

type ConsumerSourceDataSetsListByShareSubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ConsumerSourceDataSet
}

type ConsumerSourceDataSetsListByShareSubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConsumerSourceDataSetsListByShareSubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConsumerSourceDataSetsListByShareSubscription ...
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (result ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ConsumerSourceDataSetsListByShareSubscriptionCustomPager{},
		Path:       fmt.Sprintf("%s/consumerSourceDataSets", id.ID()),
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
		Values *[]ConsumerSourceDataSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConsumerSourceDataSetsListByShareSubscriptionComplete retrieves all the results into a single object
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId) (ConsumerSourceDataSetsListByShareSubscriptionCompleteResult, error) {
	return c.ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate(ctx, id, ConsumerSourceDataSetOperationPredicate{})
}

// ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate ConsumerSourceDataSetOperationPredicate) (result ConsumerSourceDataSetsListByShareSubscriptionCompleteResult, err error) {
	items := make([]ConsumerSourceDataSet, 0)

	resp, err := c.ConsumerSourceDataSetsListByShareSubscription(ctx, id)
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

	result = ConsumerSourceDataSetsListByShareSubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
