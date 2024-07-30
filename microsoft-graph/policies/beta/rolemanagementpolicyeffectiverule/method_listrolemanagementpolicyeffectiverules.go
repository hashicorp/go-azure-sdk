package rolemanagementpolicyeffectiverule

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

type ListRoleManagementPolicyEffectiveRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyEffectiveRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyEffectiveRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPolicyEffectiveRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicyEffectiveRules ...
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRules(ctx context.Context, id PolicyRoleManagementPolicyId) (result ListRoleManagementPolicyEffectiveRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleManagementPolicyEffectiveRulesCustomPager{},
		Path:       fmt.Sprintf("%s/effectiveRules", id.ID()),
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
		Values *[]beta.UnifiedRoleManagementPolicyRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementPolicyEffectiveRulesComplete retrieves all the results into a single object
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRulesComplete(ctx context.Context, id PolicyRoleManagementPolicyId) (ListRoleManagementPolicyEffectiveRulesCompleteResult, error) {
	return c.ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate(ctx, id, UnifiedRoleManagementPolicyRuleOperationPredicate{})
}

// ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate(ctx context.Context, id PolicyRoleManagementPolicyId, predicate UnifiedRoleManagementPolicyRuleOperationPredicate) (result ListRoleManagementPolicyEffectiveRulesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementPolicyRule, 0)

	resp, err := c.ListRoleManagementPolicyEffectiveRules(ctx, id)
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

	result = ListRoleManagementPolicyEffectiveRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
