package meoutlookmastercategory

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

type ListMeOutlookMasterCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookCategoryCollectionResponse
}

type ListMeOutlookMasterCategoriesCompleteResult struct {
	Items []models.OutlookCategoryCollectionResponse
}

// ListMeOutlookMasterCategories ...
func (c MeOutlookMasterCategoryClient) ListMeOutlookMasterCategories(ctx context.Context) (result ListMeOutlookMasterCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/outlook/masterCategories",
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

// ListMeOutlookMasterCategoriesComplete retrieves all the results into a single object
func (c MeOutlookMasterCategoryClient) ListMeOutlookMasterCategoriesComplete(ctx context.Context) (ListMeOutlookMasterCategoriesCompleteResult, error) {
	return c.ListMeOutlookMasterCategoriesCompleteMatchingPredicate(ctx, models.OutlookCategoryCollectionResponseOperationPredicate{})
}

// ListMeOutlookMasterCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookMasterCategoryClient) ListMeOutlookMasterCategoriesCompleteMatchingPredicate(ctx context.Context, predicate models.OutlookCategoryCollectionResponseOperationPredicate) (result ListMeOutlookMasterCategoriesCompleteResult, err error) {
	items := make([]models.OutlookCategoryCollectionResponse, 0)

	resp, err := c.ListMeOutlookMasterCategories(ctx)
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

	result = ListMeOutlookMasterCategoriesCompleteResult{
		Items: items,
	}
	return
}
