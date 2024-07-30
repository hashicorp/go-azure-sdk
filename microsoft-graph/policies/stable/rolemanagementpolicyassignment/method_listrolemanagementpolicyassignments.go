package rolemanagementpolicyassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleManagementPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleManagementPolicyAssignment
}

type ListRoleManagementPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleManagementPolicyAssignment
}

type ListRoleManagementPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicyAssignments ...
func (c RoleManagementPolicyAssignmentClient) ListRoleManagementPolicyAssignments(ctx context.Context) (result ListRoleManagementPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleManagementPolicyAssignmentsCustomPager{},
		Path:       "/policies/roleManagementPolicyAssignments",
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
		Values *[]stable.UnifiedRoleManagementPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementPolicyAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementPolicyAssignmentClient) ListRoleManagementPolicyAssignmentsComplete(ctx context.Context) (ListRoleManagementPolicyAssignmentsCompleteResult, error) {
	return c.ListRoleManagementPolicyAssignmentsCompleteMatchingPredicate(ctx, UnifiedRoleManagementPolicyAssignmentOperationPredicate{})
}

// ListRoleManagementPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyAssignmentClient) ListRoleManagementPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleManagementPolicyAssignmentOperationPredicate) (result ListRoleManagementPolicyAssignmentsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleManagementPolicyAssignment, 0)

	resp, err := c.ListRoleManagementPolicyAssignments(ctx)
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

	result = ListRoleManagementPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
