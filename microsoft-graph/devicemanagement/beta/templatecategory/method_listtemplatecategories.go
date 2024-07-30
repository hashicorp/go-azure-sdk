package templatecategory

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

type ListTemplateCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementTemplateSettingCategory
}

type ListTemplateCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementTemplateSettingCategory
}

type ListTemplateCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateCategories ...
func (c TemplateCategoryClient) ListTemplateCategories(ctx context.Context, id DeviceManagementTemplateId) (result ListTemplateCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTemplateCategoriesCustomPager{},
		Path:       fmt.Sprintf("%s/categories", id.ID()),
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
		Values *[]beta.DeviceManagementTemplateSettingCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTemplateCategoriesComplete retrieves all the results into a single object
func (c TemplateCategoryClient) ListTemplateCategoriesComplete(ctx context.Context, id DeviceManagementTemplateId) (ListTemplateCategoriesCompleteResult, error) {
	return c.ListTemplateCategoriesCompleteMatchingPredicate(ctx, id, DeviceManagementTemplateSettingCategoryOperationPredicate{})
}

// ListTemplateCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateCategoryClient) ListTemplateCategoriesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementTemplateId, predicate DeviceManagementTemplateSettingCategoryOperationPredicate) (result ListTemplateCategoriesCompleteResult, err error) {
	items := make([]beta.DeviceManagementTemplateSettingCategory, 0)

	resp, err := c.ListTemplateCategories(ctx, id)
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

	result = ListTemplateCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
