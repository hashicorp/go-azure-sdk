package directoryroleassignmentapproval

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

type ListDirectoryRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Approval
}

type ListDirectoryRoleAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Approval
}

type ListDirectoryRoleAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentApprovals ...
func (c DirectoryRoleAssignmentApprovalClient) ListDirectoryRoleAssignmentApprovals(ctx context.Context) (result ListDirectoryRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleAssignmentApprovalsCustomPager{},
		Path:       "/roleManagement/directory/roleAssignmentApprovals",
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

// ListDirectoryRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentApprovalClient) ListDirectoryRoleAssignmentApprovalsComplete(ctx context.Context) (ListDirectoryRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, ApprovalOperationPredicate{})
}

// ListDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentApprovalClient) ListDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate ApprovalOperationPredicate) (result ListDirectoryRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]beta.Approval, 0)

	resp, err := c.ListDirectoryRoleAssignmentApprovals(ctx)
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

	result = ListDirectoryRoleAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
