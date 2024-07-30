package category

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingCategory
}

type ListCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingCategory
}

type ListCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCategories ...
func (c CategoryClient) ListCategories(ctx context.Context) (result ListCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCategoriesCustomPager{},
		Path:       "/deviceManagement/categories",
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
		Values *[]beta.DeviceManagementSettingCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCategoriesComplete retrieves all the results into a single object
func (c CategoryClient) ListCategoriesComplete(ctx context.Context) (ListCategoriesCompleteResult, error) {
	return c.ListCategoriesCompleteMatchingPredicate(ctx, DeviceManagementSettingCategoryOperationPredicate{})
}

// ListCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CategoryClient) ListCategoriesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementSettingCategoryOperationPredicate) (result ListCategoriesCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingCategory, 0)

	resp, err := c.ListCategories(ctx)
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

	result = ListCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
