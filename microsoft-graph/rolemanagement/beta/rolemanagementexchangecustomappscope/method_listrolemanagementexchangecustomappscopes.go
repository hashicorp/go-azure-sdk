package rolemanagementexchangecustomappscope

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

type ListRoleManagementExchangeCustomAppScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CustomAppScopeCollectionResponse
}

type ListRoleManagementExchangeCustomAppScopesCompleteResult struct {
	Items []models.CustomAppScopeCollectionResponse
}

// ListRoleManagementExchangeCustomAppScopes ...
func (c RoleManagementExchangeCustomAppScopeClient) ListRoleManagementExchangeCustomAppScopes(ctx context.Context) (result ListRoleManagementExchangeCustomAppScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/exchange/customAppScopes",
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
		Values *[]models.CustomAppScopeCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementExchangeCustomAppScopesComplete retrieves all the results into a single object
func (c RoleManagementExchangeCustomAppScopeClient) ListRoleManagementExchangeCustomAppScopesComplete(ctx context.Context) (ListRoleManagementExchangeCustomAppScopesCompleteResult, error) {
	return c.ListRoleManagementExchangeCustomAppScopesCompleteMatchingPredicate(ctx, models.CustomAppScopeCollectionResponseOperationPredicate{})
}

// ListRoleManagementExchangeCustomAppScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementExchangeCustomAppScopeClient) ListRoleManagementExchangeCustomAppScopesCompleteMatchingPredicate(ctx context.Context, predicate models.CustomAppScopeCollectionResponseOperationPredicate) (result ListRoleManagementExchangeCustomAppScopesCompleteResult, err error) {
	items := make([]models.CustomAppScopeCollectionResponse, 0)

	resp, err := c.ListRoleManagementExchangeCustomAppScopes(ctx)
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

	result = ListRoleManagementExchangeCustomAppScopesCompleteResult{
		Items: items,
	}
	return
}
