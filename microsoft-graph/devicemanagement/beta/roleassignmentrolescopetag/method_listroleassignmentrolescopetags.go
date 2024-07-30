package roleassignmentrolescopetag

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

type ListRoleAssignmentRoleScopeTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTag
}

type ListRoleAssignmentRoleScopeTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTag
}

type ListRoleAssignmentRoleScopeTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleAssignmentRoleScopeTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleAssignmentRoleScopeTags ...
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTags(ctx context.Context, id DeviceManagementRoleAssignmentId) (result ListRoleAssignmentRoleScopeTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleAssignmentRoleScopeTagsCustomPager{},
		Path:       fmt.Sprintf("%s/roleScopeTags", id.ID()),
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
		Values *[]beta.RoleScopeTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleAssignmentRoleScopeTagsComplete retrieves all the results into a single object
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTagsComplete(ctx context.Context, id DeviceManagementRoleAssignmentId) (ListRoleAssignmentRoleScopeTagsCompleteResult, error) {
	return c.ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate(ctx, id, RoleScopeTagOperationPredicate{})
}

// ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementRoleAssignmentId, predicate RoleScopeTagOperationPredicate) (result ListRoleAssignmentRoleScopeTagsCompleteResult, err error) {
	items := make([]beta.RoleScopeTag, 0)

	resp, err := c.ListRoleAssignmentRoleScopeTags(ctx, id)
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

	result = ListRoleAssignmentRoleScopeTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
