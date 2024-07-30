package enterpriseapproleassignmentapproval

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

type ListEnterpriseAppRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Approval
}

type ListEnterpriseAppRoleAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Approval
}

type ListEnterpriseAppRoleAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentApprovals ...
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovals(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppRoleAssignmentApprovalsCustomPager{},
		Path:       fmt.Sprintf("%s/roleAssignmentApprovals", id.ID()),
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
		Values *[]beta.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovalsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, id, ApprovalOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentApprovalClient) ListEnterpriseAppRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate ApprovalOperationPredicate) (result ListEnterpriseAppRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]beta.Approval, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentApprovals(ctx, id)
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

	result = ListEnterpriseAppRoleAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
