package entitlementmanagementaccesspackageassignmentapprovalstage

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

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStages ...
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStages(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) (result ListEntitlementManagementAccessPackageAssignmentApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageAssignmentApprovalStagesCustomPager{},
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

// ListEntitlementManagementAccessPackageAssignmentApprovalStagesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStagesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId) (ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate(ctx, id, ApprovalStageOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, predicate ApprovalStageOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentApprovalStages(ctx, id)
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

	result = ListEntitlementManagementAccessPackageAssignmentApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
