package rolemanagementcloudpcroledefinition

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

type ListRoleManagementCloudPCRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementCloudPCRoleDefinitionsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementCloudPCRoleDefinitions ...
func (c RoleManagementCloudPCRoleDefinitionClient) ListRoleManagementCloudPCRoleDefinitions(ctx context.Context) (result ListRoleManagementCloudPCRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/cloudPC/roleDefinitions",
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

// ListRoleManagementCloudPCRoleDefinitionsComplete retrieves all the results into a single object
func (c RoleManagementCloudPCRoleDefinitionClient) ListRoleManagementCloudPCRoleDefinitionsComplete(ctx context.Context) (ListRoleManagementCloudPCRoleDefinitionsCompleteResult, error) {
	return c.ListRoleManagementCloudPCRoleDefinitionsCompleteMatchingPredicate(ctx, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCRoleDefinitionClient) ListRoleManagementCloudPCRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCRoleDefinitionsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCRoleDefinitions(ctx)
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

	result = ListRoleManagementCloudPCRoleDefinitionsCompleteResult{
		Items: items,
	}
	return
}
