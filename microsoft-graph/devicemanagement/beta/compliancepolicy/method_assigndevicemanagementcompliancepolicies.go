package compliancepolicy

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

type AssignDeviceManagementCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignDeviceManagementCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignDeviceManagementCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementCompliancePolicies ...
func (c CompliancePolicyClient) AssignDeviceManagementCompliancePolicies(ctx context.Context, id DeviceManagementCompliancePolicyId, input AssignDeviceManagementCompliancePoliciesRequest) (result AssignDeviceManagementCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementCompliancePoliciesCustomPager{},
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

// AssignDeviceManagementCompliancePoliciesComplete retrieves all the results into a single object
func (c CompliancePolicyClient) AssignDeviceManagementCompliancePoliciesComplete(ctx context.Context, id DeviceManagementCompliancePolicyId, input AssignDeviceManagementCompliancePoliciesRequest) (AssignDeviceManagementCompliancePoliciesCompleteResult, error) {
	return c.AssignDeviceManagementCompliancePoliciesCompleteMatchingPredicate(ctx, id, input, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// AssignDeviceManagementCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyClient) AssignDeviceManagementCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyId, input AssignDeviceManagementCompliancePoliciesRequest, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result AssignDeviceManagementCompliancePoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.AssignDeviceManagementCompliancePolicies(ctx, id, input)
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

	result = AssignDeviceManagementCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
