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

type L3NetworksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]L3Network
}

type L3NetworksListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []L3Network
}

// L3NetworksListByResourceGroup ...
func (c NetworkcloudsClient) L3NetworksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result L3NetworksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/l3Networks", id.ID()),
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
		Values *[]L3Network `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// L3NetworksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) L3NetworksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (L3NetworksListByResourceGroupCompleteResult, error) {
	return c.L3NetworksListByResourceGroupCompleteMatchingPredicate(ctx, id, L3NetworkOperationPredicate{})
}

// L3NetworksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) L3NetworksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate L3NetworkOperationPredicate) (result L3NetworksListByResourceGroupCompleteResult, err error) {
	items := make([]L3Network, 0)

	resp, err := c.L3NetworksListByResourceGroup(ctx, id)
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

	result = L3NetworksListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
