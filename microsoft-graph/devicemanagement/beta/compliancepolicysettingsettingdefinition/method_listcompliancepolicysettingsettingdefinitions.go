package compliancepolicysettingsettingdefinition

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

type ListCompliancePolicySettingSettingDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSettingDefinition
}

type ListCompliancePolicySettingSettingDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSettingDefinition
}

type ListCompliancePolicySettingSettingDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicySettingSettingDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicySettingSettingDefinitions ...
func (c CompliancePolicySettingSettingDefinitionClient) ListCompliancePolicySettingSettingDefinitions(ctx context.Context, id DeviceManagementCompliancePolicyIdSettingId) (result ListCompliancePolicySettingSettingDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCompliancePolicySettingSettingDefinitionsCustomPager{},
		Path:       fmt.Sprintf("%s/settingDefinitions", id.ID()),
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

// ListCompliancePolicySettingSettingDefinitionsComplete retrieves all the results into a single object
func (c CompliancePolicySettingSettingDefinitionClient) ListCompliancePolicySettingSettingDefinitionsComplete(ctx context.Context, id DeviceManagementCompliancePolicyIdSettingId) (ListCompliancePolicySettingSettingDefinitionsCompleteResult, error) {
	return c.ListCompliancePolicySettingSettingDefinitionsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationSettingDefinitionOperationPredicate{})
}

// ListCompliancePolicySettingSettingDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicySettingSettingDefinitionClient) ListCompliancePolicySettingSettingDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyIdSettingId, predicate DeviceManagementConfigurationSettingDefinitionOperationPredicate) (result ListCompliancePolicySettingSettingDefinitionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSettingDefinition, 0)

	resp, err := c.ListCompliancePolicySettingSettingDefinitions(ctx, id)
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

	result = ListCompliancePolicySettingSettingDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
