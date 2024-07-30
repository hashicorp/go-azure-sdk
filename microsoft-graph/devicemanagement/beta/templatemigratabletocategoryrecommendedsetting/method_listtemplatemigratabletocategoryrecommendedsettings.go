package templatemigratabletocategoryrecommendedsetting

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

type ListTemplateMigratableToCategoryRecommendedSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListTemplateMigratableToCategoryRecommendedSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListTemplateMigratableToCategoryRecommendedSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateMigratableToCategoryRecommendedSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateMigratableToCategoryRecommendedSettings ...
func (c TemplateMigratableToCategoryRecommendedSettingClient) ListTemplateMigratableToCategoryRecommendedSettings(ctx context.Context, id DeviceManagementTemplateIdMigratableToIdCategoryId) (result ListTemplateMigratableToCategoryRecommendedSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTemplateMigratableToCategoryRecommendedSettingsCustomPager{},
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

// ListTemplateMigratableToCategoryRecommendedSettingsComplete retrieves all the results into a single object
func (c TemplateMigratableToCategoryRecommendedSettingClient) ListTemplateMigratableToCategoryRecommendedSettingsComplete(ctx context.Context, id DeviceManagementTemplateIdMigratableToIdCategoryId) (ListTemplateMigratableToCategoryRecommendedSettingsCompleteResult, error) {
	return c.ListTemplateMigratableToCategoryRecommendedSettingsCompleteMatchingPredicate(ctx, id, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListTemplateMigratableToCategoryRecommendedSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateMigratableToCategoryRecommendedSettingClient) ListTemplateMigratableToCategoryRecommendedSettingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementTemplateIdMigratableToIdCategoryId, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListTemplateMigratableToCategoryRecommendedSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListTemplateMigratableToCategoryRecommendedSettings(ctx, id)
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

	result = ListTemplateMigratableToCategoryRecommendedSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
