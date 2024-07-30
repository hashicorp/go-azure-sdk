package entitlementmanagementassignmentpolicyquestion

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

type ListEntitlementManagementAssignmentPolicyQuestionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageQuestion
}

type ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageQuestion
}

type ListEntitlementManagementAssignmentPolicyQuestionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentPolicyQuestionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentPolicyQuestions ...
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestions(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId) (result ListEntitlementManagementAssignmentPolicyQuestionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAssignmentPolicyQuestionsCustomPager{},
		Path:       fmt.Sprintf("%s/questions", id.ID()),
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
		Values *[]stable.AccessPackageQuestion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentPolicyQuestionsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestionsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId) (ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate(ctx, id, AccessPackageQuestionOperationPredicate{})
}

// ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyQuestionClient) ListEntitlementManagementAssignmentPolicyQuestionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId, predicate AccessPackageQuestionOperationPredicate) (result ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult, err error) {
	items := make([]stable.AccessPackageQuestion, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicyQuestions(ctx, id)
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

	result = ListEntitlementManagementAssignmentPolicyQuestionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
