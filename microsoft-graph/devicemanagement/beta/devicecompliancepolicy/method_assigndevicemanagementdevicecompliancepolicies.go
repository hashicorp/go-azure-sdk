package devicecompliancepolicy

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

type AssignDeviceManagementDeviceCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCompliancePolicyAssignment
}

type AssignDeviceManagementDeviceCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCompliancePolicyAssignment
}

type AssignDeviceManagementDeviceCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceManagementDeviceCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceManagementDeviceCompliancePolicies ...
func (c DeviceCompliancePolicyClient) AssignDeviceManagementDeviceCompliancePolicies(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId, input AssignDeviceManagementDeviceCompliancePoliciesRequest) (result AssignDeviceManagementDeviceCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AssignDeviceManagementDeviceCompliancePoliciesCustomPager{},
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
		Values *[]beta.DeviceCompliancePolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceManagementDeviceCompliancePoliciesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyClient) AssignDeviceManagementDeviceCompliancePoliciesComplete(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId, input AssignDeviceManagementDeviceCompliancePoliciesRequest) (AssignDeviceManagementDeviceCompliancePoliciesCompleteResult, error) {
	return c.AssignDeviceManagementDeviceCompliancePoliciesCompleteMatchingPredicate(ctx, id, input, DeviceCompliancePolicyAssignmentOperationPredicate{})
}

// AssignDeviceManagementDeviceCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyClient) AssignDeviceManagementDeviceCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCompliancePolicyId, input AssignDeviceManagementDeviceCompliancePoliciesRequest, predicate DeviceCompliancePolicyAssignmentOperationPredicate) (result AssignDeviceManagementDeviceCompliancePoliciesCompleteResult, err error) {
	items := make([]beta.DeviceCompliancePolicyAssignment, 0)

	resp, err := c.AssignDeviceManagementDeviceCompliancePolicies(ctx, id, input)
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

	result = AssignDeviceManagementDeviceCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
