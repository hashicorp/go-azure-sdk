package enterpriseapproleassignmentschedulerequest

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

type ListEnterpriseAppRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentScheduleRequests ...
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequests(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/roleAssignmentScheduleRequests", id.ID()),
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
		Values *[]beta.UnifiedRoleAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequestsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, id, UnifiedRoleAssignmentScheduleRequestOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate UnifiedRoleAssignmentScheduleRequestOperationPredicate) (result ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleRequest, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentScheduleRequests(ctx, id)
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

	result = ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
