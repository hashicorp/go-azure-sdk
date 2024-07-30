package devicemanagementroleassignmentdirectoryscope

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

type ListDeviceManagementRoleAssignmentDirectoryScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignmentDirectoryScopes ...
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopes(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (result ListDeviceManagementRoleAssignmentDirectoryScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementRoleAssignmentDirectoryScopesCustomPager{},
		Path:       fmt.Sprintf("%s/directoryScopes", id.ID()),
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementRoleAssignmentDirectoryScopesComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopesComplete(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentDirectoryScopeClient) ListDeviceManagementRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId, predicate DirectoryObjectOperationPredicate) (result ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceManagementRoleAssignmentDirectoryScopes(ctx, id)
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

	result = ListDeviceManagementRoleAssignmentDirectoryScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
