package policyrolemanagementpolicyrule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyRoleManagementPolicyByIdRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleManagementPolicyRuleCollectionResponse
}

type ListPolicyRoleManagementPolicyByIdRulesCompleteResult struct {
	Items []models.UnifiedRoleManagementPolicyRuleCollectionResponse
}

// ListPolicyRoleManagementPolicyByIdRules ...
func (c PolicyRoleManagementPolicyRuleClient) ListPolicyRoleManagementPolicyByIdRules(ctx context.Context, id PolicyRoleManagementPolicyId) (result ListPolicyRoleManagementPolicyByIdRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.UnifiedRoleManagementPolicyRuleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyRoleManagementPolicyByIdRulesComplete retrieves all the results into a single object
func (c PolicyRoleManagementPolicyRuleClient) ListPolicyRoleManagementPolicyByIdRulesComplete(ctx context.Context, id PolicyRoleManagementPolicyId) (ListPolicyRoleManagementPolicyByIdRulesCompleteResult, error) {
	return c.ListPolicyRoleManagementPolicyByIdRulesCompleteMatchingPredicate(ctx, id, models.UnifiedRoleManagementPolicyRuleCollectionResponseOperationPredicate{})
}

// ListPolicyRoleManagementPolicyByIdRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyRoleManagementPolicyRuleClient) ListPolicyRoleManagementPolicyByIdRulesCompleteMatchingPredicate(ctx context.Context, id PolicyRoleManagementPolicyId, predicate models.UnifiedRoleManagementPolicyRuleCollectionResponseOperationPredicate) (result ListPolicyRoleManagementPolicyByIdRulesCompleteResult, err error) {
	items := make([]models.UnifiedRoleManagementPolicyRuleCollectionResponse, 0)

	resp, err := c.ListPolicyRoleManagementPolicyByIdRules(ctx, id)
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

	result = ListPolicyRoleManagementPolicyByIdRulesCompleteResult{
		Items: items,
	}
	return
}
