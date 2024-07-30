package configurationpolicysetting

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

type ListConfigurationPolicySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSetting
}

type ListConfigurationPolicySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSetting
}

type ListConfigurationPolicySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicySettings ...
func (c ConfigurationPolicySettingClient) ListConfigurationPolicySettings(ctx context.Context, id DeviceManagementConfigurationPolicyId) (result ListConfigurationPolicySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationPolicySettingsCustomPager{},
		Path:       fmt.Sprintf("%s/settings", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationPolicySettingsComplete retrieves all the results into a single object
func (c ConfigurationPolicySettingClient) ListConfigurationPolicySettingsComplete(ctx context.Context, id DeviceManagementConfigurationPolicyId) (ListConfigurationPolicySettingsCompleteResult, error) {
	return c.ListConfigurationPolicySettingsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationSettingOperationPredicate{})
}

// ListConfigurationPolicySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicySettingClient) ListConfigurationPolicySettingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementConfigurationPolicyId, predicate DeviceManagementConfigurationSettingOperationPredicate) (result ListConfigurationPolicySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSetting, 0)

	resp, err := c.ListConfigurationPolicySettings(ctx, id)
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

	result = ListConfigurationPolicySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
