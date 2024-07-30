package compliancepolicyassignment

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

type ListCompliancePolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type ListCompliancePolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type ListCompliancePolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicyAssignments ...
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignments(ctx context.Context, id DeviceManagementCompliancePolicyId) (result ListCompliancePolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCompliancePolicyAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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

// ListCompliancePolicyAssignmentsComplete retrieves all the results into a single object
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignmentsComplete(ctx context.Context, id DeviceManagementCompliancePolicyId) (ListCompliancePolicyAssignmentsCompleteResult, error) {
	return c.ListCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// ListCompliancePolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCompliancePolicyId, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result ListCompliancePolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.ListCompliancePolicyAssignments(ctx, id)
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

	result = ListCompliancePolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
