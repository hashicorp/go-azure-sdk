package rolemanagementpolicyrule

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

type ListRoleManagementPolicyRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPolicyRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicyRules ...
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRules(ctx context.Context, id PolicyRoleManagementPolicyId) (result ListRoleManagementPolicyRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleManagementPolicyRulesCustomPager{},
		Path:       fmt.Sprintf("%s/rules", id.ID()),
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
		Values *[]stable.UnifiedRoleManagementPolicyRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementPolicyRulesComplete retrieves all the results into a single object
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRulesComplete(ctx context.Context, id PolicyRoleManagementPolicyId) (ListRoleManagementPolicyRulesCompleteResult, error) {
	return c.ListRoleManagementPolicyRulesCompleteMatchingPredicate(ctx, id, UnifiedRoleManagementPolicyRuleOperationPredicate{})
}

// ListRoleManagementPolicyRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRulesCompleteMatchingPredicate(ctx context.Context, id PolicyRoleManagementPolicyId, predicate UnifiedRoleManagementPolicyRuleOperationPredicate) (result ListRoleManagementPolicyRulesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleManagementPolicyRule, 0)

	resp, err := c.ListRoleManagementPolicyRules(ctx, id)
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

	result = ListRoleManagementPolicyRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
