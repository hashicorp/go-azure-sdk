package enterpriseapproleassignment

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

type ListEnterpriseAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignment
}

type ListEnterpriseAppRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignment
}

type ListEnterpriseAppRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignments ...
func (c EnterpriseAppRoleAssignmentClient) ListEnterpriseAppRoleAssignments(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppRoleAssignmentsCustomPager{},
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
		Values *[]beta.UnifiedRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentClient) ListEnterpriseAppRoleAssignmentsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppRoleAssignmentsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, UnifiedRoleAssignmentOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentClient) ListEnterpriseAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate UnifiedRoleAssignmentOperationPredicate) (result ListEnterpriseAppRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignment, 0)

	resp, err := c.ListEnterpriseAppRoleAssignments(ctx, id)
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

	result = ListEnterpriseAppRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
