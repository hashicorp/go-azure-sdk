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

type GetDirectoryDeletedItemsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObject
}

type GetDirectoryDeletedItemsByIdsCompleteResult struct {
	Items []models.DirectoryObject
}

// GetDirectoryDeletedItemsByIds ...
func (c DirectoryDeletedItemClient) GetDirectoryDeletedItemsByIds(ctx context.Context) (result GetDirectoryDeletedItemsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/deletedItems/getByIds",
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
		Values *[]models.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDirectoryDeletedItemsByIdsComplete retrieves all the results into a single object
func (c DirectoryDeletedItemClient) GetDirectoryDeletedItemsByIdsComplete(ctx context.Context) (GetDirectoryDeletedItemsByIdsCompleteResult, error) {
	return c.GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate(ctx, models.DirectoryObjectOperationPredicate{})
}

// GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryDeletedItemClient) GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate(ctx context.Context, predicate models.DirectoryObjectOperationPredicate) (result GetDirectoryDeletedItemsByIdsCompleteResult, err error) {
	items := make([]models.DirectoryObject, 0)

	resp, err := c.GetDirectoryDeletedItemsByIds(ctx)
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

	result = GetDirectoryDeletedItemsByIdsCompleteResult{
		Items: items,
	}
	return
}
