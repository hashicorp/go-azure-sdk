package configurationpolicyassignment

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

type ListConfigurationPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type ListConfigurationPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type ListConfigurationPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicyAssignments ...
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignments(ctx context.Context, id DeviceManagementConfigurationPolicyId) (result ListConfigurationPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationPolicyAssignmentsCustomPager{},
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

// ListConfigurationPolicyAssignmentsComplete retrieves all the results into a single object
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignmentsComplete(ctx context.Context, id DeviceManagementConfigurationPolicyId) (ListConfigurationPolicyAssignmentsCompleteResult, error) {
	return c.ListConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// ListConfigurationPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementConfigurationPolicyId, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result ListConfigurationPolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.ListConfigurationPolicyAssignments(ctx, id)
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

	result = ListConfigurationPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
