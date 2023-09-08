package policyauthorizationpolicydefaultuserroleoverride

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

type ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DefaultUserRoleOverrideCollectionResponse
}

type ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteResult struct {
	Items []models.DefaultUserRoleOverrideCollectionResponse
}

// ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverrides ...
func (c PolicyAuthorizationPolicyDefaultUserRoleOverrideClient) ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverrides(ctx context.Context, id PolicyAuthorizationPolicyId) (result ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/defaultUserRoleOverrides", id.ID()),
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
		Values *[]models.DefaultUserRoleOverrideCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesComplete retrieves all the results into a single object
func (c PolicyAuthorizationPolicyDefaultUserRoleOverrideClient) ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesComplete(ctx context.Context, id PolicyAuthorizationPolicyId) (ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteResult, error) {
	return c.ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteMatchingPredicate(ctx, id, models.DefaultUserRoleOverrideCollectionResponseOperationPredicate{})
}

// ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAuthorizationPolicyDefaultUserRoleOverrideClient) ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteMatchingPredicate(ctx context.Context, id PolicyAuthorizationPolicyId, predicate models.DefaultUserRoleOverrideCollectionResponseOperationPredicate) (result ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteResult, err error) {
	items := make([]models.DefaultUserRoleOverrideCollectionResponse, 0)

	resp, err := c.ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverrides(ctx, id)
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

	result = ListPolicyAuthorizationPolicyByIdDefaultUserRoleOverridesCompleteResult{
		Items: items,
	}
	return
}
