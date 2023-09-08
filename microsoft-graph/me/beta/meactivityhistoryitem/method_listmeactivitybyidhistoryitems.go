package meactivityhistoryitem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeActivityByIdHistoryItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ActivityHistoryItemCollectionResponse
}

type ListMeActivityByIdHistoryItemsCompleteResult struct {
	Items []models.ActivityHistoryItemCollectionResponse
}

// ListMeActivityByIdHistoryItems ...
func (c MeActivityHistoryItemClient) ListMeActivityByIdHistoryItems(ctx context.Context, id MeActivityId) (result ListMeActivityByIdHistoryItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/historyItems", id.ID()),
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
		Values *[]models.ActivityHistoryItemCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeActivityByIdHistoryItemsComplete retrieves all the results into a single object
func (c MeActivityHistoryItemClient) ListMeActivityByIdHistoryItemsComplete(ctx context.Context, id MeActivityId) (ListMeActivityByIdHistoryItemsCompleteResult, error) {
	return c.ListMeActivityByIdHistoryItemsCompleteMatchingPredicate(ctx, id, models.ActivityHistoryItemCollectionResponseOperationPredicate{})
}

// ListMeActivityByIdHistoryItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeActivityHistoryItemClient) ListMeActivityByIdHistoryItemsCompleteMatchingPredicate(ctx context.Context, id MeActivityId, predicate models.ActivityHistoryItemCollectionResponseOperationPredicate) (result ListMeActivityByIdHistoryItemsCompleteResult, err error) {
	items := make([]models.ActivityHistoryItemCollectionResponse, 0)

	resp, err := c.ListMeActivityByIdHistoryItems(ctx, id)
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

	result = ListMeActivityByIdHistoryItemsCompleteResult{
		Items: items,
	}
	return
}
