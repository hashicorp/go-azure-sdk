package cloudpcroleassignmentdirectoryscope

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

type ListCloudPCRoleAssignmentDirectoryScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListCloudPCRoleAssignmentDirectoryScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListCloudPCRoleAssignmentDirectoryScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCRoleAssignmentDirectoryScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCRoleAssignmentDirectoryScopes ...
func (c CloudPCRoleAssignmentDirectoryScopeClient) ListCloudPCRoleAssignmentDirectoryScopes(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (result ListCloudPCRoleAssignmentDirectoryScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCloudPCRoleAssignmentDirectoryScopesCustomPager{},
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

// ListCloudPCRoleAssignmentDirectoryScopesComplete retrieves all the results into a single object
func (c CloudPCRoleAssignmentDirectoryScopeClient) ListCloudPCRoleAssignmentDirectoryScopesComplete(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (ListCloudPCRoleAssignmentDirectoryScopesCompleteResult, error) {
	return c.ListCloudPCRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListCloudPCRoleAssignmentDirectoryScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCRoleAssignmentDirectoryScopeClient) ListCloudPCRoleAssignmentDirectoryScopesCompleteMatchingPredicate(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId, predicate DirectoryObjectOperationPredicate) (result ListCloudPCRoleAssignmentDirectoryScopesCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListCloudPCRoleAssignmentDirectoryScopes(ctx, id)
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

	result = ListCloudPCRoleAssignmentDirectoryScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
