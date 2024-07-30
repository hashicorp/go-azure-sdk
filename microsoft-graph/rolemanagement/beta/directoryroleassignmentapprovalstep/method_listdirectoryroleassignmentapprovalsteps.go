package directoryroleassignmentapprovalstep

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

type ListDirectoryRoleAssignmentApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListDirectoryRoleAssignmentApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListDirectoryRoleAssignmentApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentApprovalSteps ...
func (c DirectoryRoleAssignmentApprovalStepClient) ListDirectoryRoleAssignmentApprovalSteps(ctx context.Context, id RoleManagementDirectoryRoleAssignmentApprovalId) (result ListDirectoryRoleAssignmentApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleAssignmentApprovalStepsCustomPager{},
		Path:       fmt.Sprintf("%s/steps", id.ID()),
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
		Values *[]beta.ApprovalStep `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentApprovalStepsComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentApprovalStepClient) ListDirectoryRoleAssignmentApprovalStepsComplete(ctx context.Context, id RoleManagementDirectoryRoleAssignmentApprovalId) (ListDirectoryRoleAssignmentApprovalStepsCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx, id, ApprovalStepOperationPredicate{})
}

// ListDirectoryRoleAssignmentApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentApprovalStepClient) ListDirectoryRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDirectoryRoleAssignmentApprovalId, predicate ApprovalStepOperationPredicate) (result ListDirectoryRoleAssignmentApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListDirectoryRoleAssignmentApprovalSteps(ctx, id)
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

	result = ListDirectoryRoleAssignmentApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
