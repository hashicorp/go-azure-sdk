package compliancepolicysetting

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

type ListCompliancePolicySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSetting
}

type ListCompliancePolicySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSetting
}

type ListCompliancePolicySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicySettings ...
func (c CompliancePolicySettingClient) ListCompliancePolicySettings(ctx context.Context, id DeviceManagementCompliancePolicyId) (result ListCompliancePolicySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCompliancePolicySettingsCustomPager{},
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

// ListCompliancePolicySettingsComplete retrieves all the results into a single object
func (c CompliancePolicySettingClient) ListCompliancePolicySettingsComplete(ctx context.Context, id DeviceManagementCompliancePolicyId) (ListCompliancePolicySettingsCompleteResult, error) {
	return c.ListCompliancePolicySettingsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationSettingOperationPredicate{})
}

// ListCompliancePolicySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicySettingClient) ListCompliancePolicySettingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyId, predicate DeviceManagementConfigurationSettingOperationPredicate) (result ListCompliancePolicySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSetting, 0)

	resp, err := c.ListCompliancePolicySettings(ctx, id)
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

	result = ListCompliancePolicySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
