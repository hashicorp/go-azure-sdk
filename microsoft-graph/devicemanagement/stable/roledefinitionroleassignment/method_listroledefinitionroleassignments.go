package roledefinitionroleassignment

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

type ListRoleDefinitionRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.RoleAssignment
}

type ListRoleDefinitionRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.RoleAssignment
}

type ListRoleDefinitionRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleDefinitionRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleDefinitionRoleAssignments ...
func (c RoleDefinitionRoleAssignmentClient) ListRoleDefinitionRoleAssignments(ctx context.Context, id DeviceManagementRoleDefinitionId) (result ListRoleDefinitionRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleDefinitionRoleAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/roleAssignments", id.ID()),
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
		Values *[]stable.RoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleDefinitionRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleDefinitionRoleAssignmentClient) ListRoleDefinitionRoleAssignmentsComplete(ctx context.Context, id DeviceManagementRoleDefinitionId) (ListRoleDefinitionRoleAssignmentsCompleteResult, error) {
	return c.ListRoleDefinitionRoleAssignmentsCompleteMatchingPredicate(ctx, id, RoleAssignmentOperationPredicate{})
}

// ListRoleDefinitionRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleDefinitionRoleAssignmentClient) ListRoleDefinitionRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementRoleDefinitionId, predicate RoleAssignmentOperationPredicate) (result ListRoleDefinitionRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.RoleAssignment, 0)

	resp, err := c.ListRoleDefinitionRoleAssignments(ctx, id)
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

	result = ListRoleDefinitionRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
