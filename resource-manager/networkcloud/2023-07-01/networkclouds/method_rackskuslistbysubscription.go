package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RackSkusListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RackSku
}

type RackSkusListBySubscriptionCompleteResult struct {
	Items []RackSku
}

// RackSkusListBySubscription ...
func (c NetworkcloudsClient) RackSkusListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result RackSkusListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/rackSkus", id.ID()),
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
		Values *[]RackSku `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RackSkusListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) RackSkusListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (RackSkusListBySubscriptionCompleteResult, error) {
	return c.RackSkusListBySubscriptionCompleteMatchingPredicate(ctx, id, RackSkuOperationPredicate{})
}

// RackSkusListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) RackSkusListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate RackSkuOperationPredicate) (result RackSkusListBySubscriptionCompleteResult, err error) {
	items := make([]RackSku, 0)

	resp, err := c.RackSkusListBySubscription(ctx, id)
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

	result = RackSkusListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
