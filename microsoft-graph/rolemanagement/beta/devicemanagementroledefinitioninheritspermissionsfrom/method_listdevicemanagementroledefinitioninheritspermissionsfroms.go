package devicemanagementroledefinitioninheritspermissionsfrom

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

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleDefinitionInheritsPermissionsFroms ...
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId) (result ListDeviceManagementRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCustomPager{},
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

// ListDeviceManagementRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId) (ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListDeviceManagementRoleDefinitionInheritsPermissionsFroms(ctx, id)
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

	result = ListDeviceManagementRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
