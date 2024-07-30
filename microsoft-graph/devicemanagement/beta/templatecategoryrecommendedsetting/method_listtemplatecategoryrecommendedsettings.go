package templatecategoryrecommendedsetting

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

type ListTemplateCategoryRecommendedSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListTemplateCategoryRecommendedSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListTemplateCategoryRecommendedSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateCategoryRecommendedSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateCategoryRecommendedSettings ...
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettings(ctx context.Context, id DeviceManagementTemplateIdCategoryId) (result ListTemplateCategoryRecommendedSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTemplateCategoryRecommendedSettingsCustomPager{},
		Path:       fmt.Sprintf("%s/recommendedSettings", id.ID()),
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
		Values *[]beta.DeviceManagementSettingInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTemplateCategoryRecommendedSettingsComplete retrieves all the results into a single object
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettingsComplete(ctx context.Context, id DeviceManagementTemplateIdCategoryId) (ListTemplateCategoryRecommendedSettingsCompleteResult, error) {
	return c.ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate(ctx, id, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementTemplateIdCategoryId, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListTemplateCategoryRecommendedSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListTemplateCategoryRecommendedSettings(ctx, id)
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

	result = ListTemplateCategoryRecommendedSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
