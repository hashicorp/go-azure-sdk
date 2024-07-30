package privilegedaccesgroupassignmentapprovalstage

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

type ListPrivilegedAccesGroupAssignmentApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListPrivilegedAccesGroupAssignmentApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentApprovalStages ...
func (c PrivilegedAccesGroupAssignmentApprovalStageClient) ListPrivilegedAccesGroupAssignmentApprovalStages(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) (result ListPrivilegedAccesGroupAssignmentApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentApprovalStagesCustomPager{},
		Path:       fmt.Sprintf("%s/stages", id.ID()),
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
		Values *[]stable.ApprovalStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupAssignmentApprovalStagesComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentApprovalStageClient) ListPrivilegedAccesGroupAssignmentApprovalStagesComplete(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) (ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteMatchingPredicate(ctx, id, ApprovalStageOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentApprovalStageClient) ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId, predicate ApprovalStageOperationPredicate) (result ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentApprovalStages(ctx, id)
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

	result = ListPrivilegedAccesGroupAssignmentApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
