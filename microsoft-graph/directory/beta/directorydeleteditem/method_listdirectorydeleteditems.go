package directorydeleteditem

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

type ListDirectoryDeletedItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListDirectoryDeletedItemsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListDirectoryDeletedItems ...
func (c DirectoryDeletedItemClient) ListDirectoryDeletedItems(ctx context.Context) (result ListDirectoryDeletedItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/deletedItems",
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryDeletedItemsComplete retrieves all the results into a single object
func (c DirectoryDeletedItemClient) ListDirectoryDeletedItemsComplete(ctx context.Context) (ListDirectoryDeletedItemsCompleteResult, error) {
	return c.ListDirectoryDeletedItemsCompleteMatchingPredicate(ctx, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListDirectoryDeletedItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryDeletedItemClient) ListDirectoryDeletedItemsCompleteMatchingPredicate(ctx context.Context, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListDirectoryDeletedItemsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListDirectoryDeletedItems(ctx)
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

	result = ListDirectoryDeletedItemsCompleteResult{
		Items: items,
	}
	return
}
