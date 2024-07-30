package privilegedaccesgroupassignmentapproval

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

type ListPrivilegedAccesGroupAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Approval
}

type ListPrivilegedAccesGroupAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Approval
}

type ListPrivilegedAccesGroupAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentApprovals ...
func (c PrivilegedAccesGroupAssignmentApprovalClient) ListPrivilegedAccesGroupAssignmentApprovals(ctx context.Context) (result ListPrivilegedAccesGroupAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentApprovalsCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/assignmentApprovals",
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
		Values *[]stable.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupAssignmentApprovalsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentApprovalClient) ListPrivilegedAccesGroupAssignmentApprovalsComplete(ctx context.Context) (ListPrivilegedAccesGroupAssignmentApprovalsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentApprovalsCompleteMatchingPredicate(ctx, ApprovalOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentApprovalClient) ListPrivilegedAccesGroupAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate ApprovalOperationPredicate) (result ListPrivilegedAccesGroupAssignmentApprovalsCompleteResult, err error) {
	items := make([]stable.Approval, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentApprovals(ctx)
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

	result = ListPrivilegedAccesGroupAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
