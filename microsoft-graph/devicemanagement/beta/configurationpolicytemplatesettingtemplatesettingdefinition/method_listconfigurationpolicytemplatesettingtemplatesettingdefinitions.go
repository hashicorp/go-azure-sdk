package configurationpolicytemplatesettingtemplatesettingdefinition

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

type ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSettingDefinition
}

type ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSettingDefinition
}

type ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicyTemplateSettingTemplateSettingDefinitions ...
func (c ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient) ListConfigurationPolicyTemplateSettingTemplateSettingDefinitions(ctx context.Context, id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) (result ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCustomPager{},
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

// ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsComplete retrieves all the results into a single object
func (c ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient) ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsComplete(ctx context.Context, id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) (ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteResult, error) {
	return c.ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationSettingDefinitionOperationPredicate{})
}

// ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient) ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId, predicate DeviceManagementConfigurationSettingDefinitionOperationPredicate) (result ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSettingDefinition, 0)

	resp, err := c.ListConfigurationPolicyTemplateSettingTemplateSettingDefinitions(ctx, id)
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

	result = ListConfigurationPolicyTemplateSettingTemplateSettingDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
