package policyrolemanagementpolicyeffectiverule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyRoleManagementPolicyByIdEffectiveRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleManagementPolicyRuleCollectionResponse
}

type ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteResult struct {
	Items []models.UnifiedRoleManagementPolicyRuleCollectionResponse
}

// ListPolicyRoleManagementPolicyByIdEffectiveRules ...
func (c PolicyRoleManagementPolicyEffectiveRuleClient) ListPolicyRoleManagementPolicyByIdEffectiveRules(ctx context.Context, id PolicyRoleManagementPolicyId) (result ListPolicyRoleManagementPolicyByIdEffectiveRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.UnifiedRoleManagementPolicyRuleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyRoleManagementPolicyByIdEffectiveRulesComplete retrieves all the results into a single object
func (c PolicyRoleManagementPolicyEffectiveRuleClient) ListPolicyRoleManagementPolicyByIdEffectiveRulesComplete(ctx context.Context, id PolicyRoleManagementPolicyId) (ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteResult, error) {
	return c.ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteMatchingPredicate(ctx, id, models.UnifiedRoleManagementPolicyRuleCollectionResponseOperationPredicate{})
}

// ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyRoleManagementPolicyEffectiveRuleClient) ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteMatchingPredicate(ctx context.Context, id PolicyRoleManagementPolicyId, predicate models.UnifiedRoleManagementPolicyRuleCollectionResponseOperationPredicate) (result ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteResult, err error) {
	items := make([]models.UnifiedRoleManagementPolicyRuleCollectionResponse, 0)

	resp, err := c.ListPolicyRoleManagementPolicyByIdEffectiveRules(ctx, id)
	if err != nil {
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

	result = ListPolicyRoleManagementPolicyByIdEffectiveRulesCompleteResult{
		Items: items,
	}
	return
}
