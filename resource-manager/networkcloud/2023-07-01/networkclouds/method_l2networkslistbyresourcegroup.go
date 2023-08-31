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

type L2NetworksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]L2Network
}

type L2NetworksListByResourceGroupCompleteResult struct {
	Items []L2Network
}

// L2NetworksListByResourceGroup ...
func (c NetworkcloudsClient) L2NetworksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result L2NetworksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/l2Networks", id.ID()),
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
		Values *[]L2Network `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// L2NetworksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) L2NetworksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (L2NetworksListByResourceGroupCompleteResult, error) {
	return c.L2NetworksListByResourceGroupCompleteMatchingPredicate(ctx, id, L2NetworkOperationPredicate{})
}

// L2NetworksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) L2NetworksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate L2NetworkOperationPredicate) (result L2NetworksListByResourceGroupCompleteResult, err error) {
	items := make([]L2Network, 0)

	resp, err := c.L2NetworksListByResourceGroup(ctx, id)
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

	result = L2NetworksListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
