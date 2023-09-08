package rolemanagementexchangeroledefinition

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

type ListRoleManagementExchangeRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementExchangeRoleDefinitionsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementExchangeRoleDefinitions ...
func (c RoleManagementExchangeRoleDefinitionClient) ListRoleManagementExchangeRoleDefinitions(ctx context.Context) (result ListRoleManagementExchangeRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/exchange/roleDefinitions",
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
		Values *[]models.UnifiedRoleDefinitionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementExchangeRoleDefinitionsComplete retrieves all the results into a single object
func (c RoleManagementExchangeRoleDefinitionClient) ListRoleManagementExchangeRoleDefinitionsComplete(ctx context.Context) (ListRoleManagementExchangeRoleDefinitionsCompleteResult, error) {
	return c.ListRoleManagementExchangeRoleDefinitionsCompleteMatchingPredicate(ctx, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementExchangeRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementExchangeRoleDefinitionClient) ListRoleManagementExchangeRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementExchangeRoleDefinitionsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementExchangeRoleDefinitions(ctx)
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

	result = ListRoleManagementExchangeRoleDefinitionsCompleteResult{
		Items: items,
	}
	return
}
