package templatesetting

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

type ListTemplateSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSettingTemplate
}

type ListTemplateSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSettingTemplate
}

type ListTemplateSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateSettings ...
func (c TemplateSettingClient) ListTemplateSettings(ctx context.Context) (result ListTemplateSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTemplateSettingsCustomPager{},
		Path:       "/deviceManagement/templateSettings",
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
		Values *[]beta.DeviceManagementConfigurationSettingTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTemplateSettingsComplete retrieves all the results into a single object
func (c TemplateSettingClient) ListTemplateSettingsComplete(ctx context.Context) (ListTemplateSettingsCompleteResult, error) {
	return c.ListTemplateSettingsCompleteMatchingPredicate(ctx, DeviceManagementConfigurationSettingTemplateOperationPredicate{})
}

// ListTemplateSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateSettingClient) ListTemplateSettingsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementConfigurationSettingTemplateOperationPredicate) (result ListTemplateSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSettingTemplate, 0)

	resp, err := c.ListTemplateSettings(ctx)
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

	result = ListTemplateSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
