package cloudpcroledefinitioninheritspermissionsfrom

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

type ListCloudPCRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListCloudPCRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCRoleDefinitionInheritsPermissionsFroms ...
func (c CloudPCRoleDefinitionInheritsPermissionsFromClient) ListCloudPCRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId) (result ListCloudPCRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCloudPCRoleDefinitionInheritsPermissionsFromsCustomPager{},
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c CloudPCRoleDefinitionInheritsPermissionsFromClient) ListCloudPCRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId) (ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, UnifiedRoleDefinitionOperationPredicate{})
}

// ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCRoleDefinitionInheritsPermissionsFromClient) ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId, predicate UnifiedRoleDefinitionOperationPredicate) (result ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListCloudPCRoleDefinitionInheritsPermissionsFroms(ctx, id)
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

	result = ListCloudPCRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
