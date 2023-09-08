package rolemanagementdirectoryroledefinition

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

type ListRoleManagementDirectoryRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementDirectoryRoleDefinitionsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementDirectoryRoleDefinitions ...
func (c RoleManagementDirectoryRoleDefinitionClient) ListRoleManagementDirectoryRoleDefinitions(ctx context.Context) (result ListRoleManagementDirectoryRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleDefinitions",
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

// ListRoleManagementDirectoryRoleDefinitionsComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleDefinitionClient) ListRoleManagementDirectoryRoleDefinitionsComplete(ctx context.Context) (ListRoleManagementDirectoryRoleDefinitionsCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleDefinitionsCompleteMatchingPredicate(ctx, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleDefinitionClient) ListRoleManagementDirectoryRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleDefinitionsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleDefinitions(ctx)
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

	result = ListRoleManagementDirectoryRoleDefinitionsCompleteResult{
		Items: items,
	}
	return
}
