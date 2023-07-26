package inventoryitems

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVCenterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]InventoryItem
}

type ListByVCenterCompleteResult struct {
	Items []InventoryItem
}

// ListByVCenter ...
func (c InventoryItemsClient) ListByVCenter(ctx context.Context, id VCenterId) (result ListByVCenterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inventoryItems", id.ID()),
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
		Values *[]InventoryItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByVCenterComplete retrieves all the results into a single object
func (c InventoryItemsClient) ListByVCenterComplete(ctx context.Context, id VCenterId) (ListByVCenterCompleteResult, error) {
	return c.ListByVCenterCompleteMatchingPredicate(ctx, id, InventoryItemOperationPredicate{})
}

// ListByVCenterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InventoryItemsClient) ListByVCenterCompleteMatchingPredicate(ctx context.Context, id VCenterId, predicate InventoryItemOperationPredicate) (result ListByVCenterCompleteResult, err error) {
	items := make([]InventoryItem, 0)

	resp, err := c.ListByVCenter(ctx, id)
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

	result = ListByVCenterCompleteResult{
		Items: items,
	}
	return
}
