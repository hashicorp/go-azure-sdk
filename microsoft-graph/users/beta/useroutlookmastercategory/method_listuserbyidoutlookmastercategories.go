package useroutlookmastercategory

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

type ListUserByIdOutlookMasterCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookCategoryCollectionResponse
}

type ListUserByIdOutlookMasterCategoriesCompleteResult struct {
	Items []models.OutlookCategoryCollectionResponse
}

// ListUserByIdOutlookMasterCategories ...
func (c UserOutlookMasterCategoryClient) ListUserByIdOutlookMasterCategories(ctx context.Context, id UserId) (result ListUserByIdOutlookMasterCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/outlook/masterCategories", id.ID()),
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
		Values *[]models.OutlookCategoryCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOutlookMasterCategoriesComplete retrieves all the results into a single object
func (c UserOutlookMasterCategoryClient) ListUserByIdOutlookMasterCategoriesComplete(ctx context.Context, id UserId) (ListUserByIdOutlookMasterCategoriesCompleteResult, error) {
	return c.ListUserByIdOutlookMasterCategoriesCompleteMatchingPredicate(ctx, id, models.OutlookCategoryCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookMasterCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookMasterCategoryClient) ListUserByIdOutlookMasterCategoriesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OutlookCategoryCollectionResponseOperationPredicate) (result ListUserByIdOutlookMasterCategoriesCompleteResult, err error) {
	items := make([]models.OutlookCategoryCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookMasterCategories(ctx, id)
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

	result = ListUserByIdOutlookMasterCategoriesCompleteResult{
		Items: items,
	}
	return
}
