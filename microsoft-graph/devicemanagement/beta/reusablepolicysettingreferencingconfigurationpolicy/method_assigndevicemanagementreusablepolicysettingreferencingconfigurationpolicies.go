package reusablepolicysettingreferencingconfigurationpolicy

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

type AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementReusablePolicySettingReferencingConfigurationPolicies ...
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignDeviceManagementReusablePolicySettingReferencingConfigurationPolicies(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesRequest) (result AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesComplete retrieves all the results into a single object
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesRequest) (AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, error) {
	return c.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx, id, input, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesRequest, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPolicies(ctx, id, input)
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

	result = AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
