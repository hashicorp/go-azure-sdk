package roleassignment

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

type ListRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceAndAppManagementRoleAssignment
}

type ListRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceAndAppManagementRoleAssignment
}

type ListRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleAssignments ...
func (c RoleAssignmentClient) ListRoleAssignments(ctx context.Context) (result ListRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleAssignmentsCustomPager{},
		Path:       "/deviceManagement/roleAssignments",
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
		Values *[]stable.DeviceAndAppManagementRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleAssignmentClient) ListRoleAssignmentsComplete(ctx context.Context) (ListRoleAssignmentsCompleteResult, error) {
	return c.ListRoleAssignmentsCompleteMatchingPredicate(ctx, DeviceAndAppManagementRoleAssignmentOperationPredicate{})
}

// ListRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleAssignmentClient) ListRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate DeviceAndAppManagementRoleAssignmentOperationPredicate) (result ListRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.DeviceAndAppManagementRoleAssignment, 0)

	resp, err := c.ListRoleAssignments(ctx)
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

	result = ListRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
