package privilegedaccesgroupassignmentschedulerequest

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

type ListPrivilegedAccesGroupAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupAssignmentScheduleRequest
}

type ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupAssignmentScheduleRequest
}

type ListPrivilegedAccesGroupAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentScheduleRequests ...
func (c PrivilegedAccesGroupAssignmentScheduleRequestClient) ListPrivilegedAccesGroupAssignmentScheduleRequests(ctx context.Context) (result ListPrivilegedAccesGroupAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentScheduleRequestsCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
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
		Values *[]stable.PrivilegedAccessGroupAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentScheduleRequestClient) ListPrivilegedAccesGroupAssignmentScheduleRequestsComplete(ctx context.Context) (ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentScheduleRequestClient) ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate) (result ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupAssignmentScheduleRequest, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentScheduleRequests(ctx)
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

	result = ListPrivilegedAccesGroupAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
