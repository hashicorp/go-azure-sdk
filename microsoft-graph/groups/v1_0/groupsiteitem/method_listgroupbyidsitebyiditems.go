package groupsiteitem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdSiteByIdItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.BaseItemCollectionResponse
}

type ListGroupByIdSiteByIdItemsCompleteResult struct {
	Items []models.BaseItemCollectionResponse
}

// ListGroupByIdSiteByIdItems ...
func (c GroupSiteItemClient) ListGroupByIdSiteByIdItems(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/items", id.ID()),
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
		Values *[]models.BaseItemCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdItemsComplete retrieves all the results into a single object
func (c GroupSiteItemClient) ListGroupByIdSiteByIdItemsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdItemsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdItemsCompleteMatchingPredicate(ctx, id, models.BaseItemCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteItemClient) ListGroupByIdSiteByIdItemsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.BaseItemCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdItemsCompleteResult, err error) {
	items := make([]models.BaseItemCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdItems(ctx, id)
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

	result = ListGroupByIdSiteByIdItemsCompleteResult{
		Items: items,
	}
	return
}
