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
	Items []ConsumerSourceDataSet
}

// ConsumerSourceDataSetsListByShareSubscription ...
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (result ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Items: items,
	}
	return
}
