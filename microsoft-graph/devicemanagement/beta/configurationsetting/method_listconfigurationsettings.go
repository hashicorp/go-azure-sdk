package configurationsetting

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

type ListConfigurationSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSettingDefinition
}

type ListConfigurationSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSettingDefinition
}

type ListConfigurationSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationSettings ...
func (c ConfigurationSettingClient) ListConfigurationSettings(ctx context.Context) (result ListConfigurationSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationSettingsCustomPager{},
		Path:       "/deviceManagement/configurationSettings",
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
		Values *[]beta.DeviceManagementConfigurationSettingDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationSettingsComplete retrieves all the results into a single object
func (c ConfigurationSettingClient) ListConfigurationSettingsComplete(ctx context.Context) (ListConfigurationSettingsCompleteResult, error) {
	return c.ListConfigurationSettingsCompleteMatchingPredicate(ctx, DeviceManagementConfigurationSettingDefinitionOperationPredicate{})
}

// ListConfigurationSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationSettingClient) ListConfigurationSettingsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementConfigurationSettingDefinitionOperationPredicate) (result ListConfigurationSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSettingDefinition, 0)

	resp, err := c.ListConfigurationSettings(ctx)
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

	result = ListConfigurationSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
