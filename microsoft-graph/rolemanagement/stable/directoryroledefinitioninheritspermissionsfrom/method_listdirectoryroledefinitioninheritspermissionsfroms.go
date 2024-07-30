package directoryroledefinitioninheritspermissionsfrom

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

type ListDirectoryRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleDefinition
}

type ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleDefinition
}

type ListDirectoryRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleDefinitionInheritsPermissionsFroms ...
func (c DirectoryRoleDefinitionInheritsPermissionsFromClient) ListDirectoryRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id RoleManagementDirectoryRoleDefinitionId) (result ListDirectoryRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleDefinitionInheritsPermissionsFromsCustomPager{},
		Path:       fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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
		Values *[]stable.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c DirectoryRoleDefinitionInheritsPermissionsFromClient) ListDirectoryRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementDirectoryRoleDefinitionId) (ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleDefinitionInheritsPermissionsFromClient) ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDirectoryRoleDefinitionId, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleDefinition, 0)

	resp, err := c.ListDirectoryRoleDefinitionInheritsPermissionsFroms(ctx, id)
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

	result = ListDirectoryRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
