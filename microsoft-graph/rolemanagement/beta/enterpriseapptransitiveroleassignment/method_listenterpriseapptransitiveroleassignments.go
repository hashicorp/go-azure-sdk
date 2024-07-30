package enterpriseapptransitiveroleassignment

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

type ListEnterpriseAppTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignment
}

type ListEnterpriseAppTransitiveRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignment
}

type ListEnterpriseAppTransitiveRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppTransitiveRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppTransitiveRoleAssignments ...
func (c EnterpriseAppTransitiveRoleAssignmentClient) ListEnterpriseAppTransitiveRoleAssignments(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppTransitiveRoleAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/transitiveRoleAssignments", id.ID()),
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

// ListEnterpriseAppTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c EnterpriseAppTransitiveRoleAssignmentClient) ListEnterpriseAppTransitiveRoleAssignmentsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListEnterpriseAppTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, id, UnifiedRoleAssignmentOperationPredicate{})
}

// ListEnterpriseAppTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppTransitiveRoleAssignmentClient) ListEnterpriseAppTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate UnifiedRoleAssignmentOperationPredicate) (result ListEnterpriseAppTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignment, 0)

	resp, err := c.ListEnterpriseAppTransitiveRoleAssignments(ctx, id)
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

	result = ListEnterpriseAppTransitiveRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
