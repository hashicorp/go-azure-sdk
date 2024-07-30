package devicemanagementroleassignmentprincipal

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

type ListDeviceManagementRoleAssignmentPrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentPrincipalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceManagementRoleAssignmentPrincipalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentPrincipalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignmentPrincipals ...
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipals(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (result ListDeviceManagementRoleAssignmentPrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementRoleAssignmentPrincipalsCustomPager{},
		Path:       fmt.Sprintf("%s/principals", id.ID()),
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

// ListDeviceManagementRoleAssignmentPrincipalsComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipalsComplete(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (ListDeviceManagementRoleAssignmentPrincipalsCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentPrincipalClient) ListDeviceManagementRoleAssignmentPrincipalsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId, predicate DirectoryObjectOperationPredicate) (result ListDeviceManagementRoleAssignmentPrincipalsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceManagementRoleAssignmentPrincipals(ctx, id)
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

	result = ListDeviceManagementRoleAssignmentPrincipalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
