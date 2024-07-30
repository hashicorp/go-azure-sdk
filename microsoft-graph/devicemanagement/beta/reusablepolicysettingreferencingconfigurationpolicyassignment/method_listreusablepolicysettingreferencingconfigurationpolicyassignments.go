package reusablepolicysettingreferencingconfigurationpolicyassignment

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

type ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListReusablePolicySettingReferencingConfigurationPolicyAssignments ...
func (c ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient) ListReusablePolicySettingReferencingConfigurationPolicyAssignments(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) (result ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCustomPager{},
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

// ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsComplete retrieves all the results into a single object
func (c ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient) ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsComplete(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) (ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteResult, error) {
	return c.ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient) ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.ListReusablePolicySettingReferencingConfigurationPolicyAssignments(ctx, id)
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

	result = ListReusablePolicySettingReferencingConfigurationPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
