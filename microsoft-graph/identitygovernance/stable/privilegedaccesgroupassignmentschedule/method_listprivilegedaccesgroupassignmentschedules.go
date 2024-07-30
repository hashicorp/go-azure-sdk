package privilegedaccesgroupassignmentschedule

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

type ListPrivilegedAccesGroupAssignmentSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupAssignmentSchedule
}

type ListPrivilegedAccesGroupAssignmentSchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupAssignmentSchedule
}

type ListPrivilegedAccesGroupAssignmentSchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentSchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentSchedules ...
func (c PrivilegedAccesGroupAssignmentScheduleClient) ListPrivilegedAccesGroupAssignmentSchedules(ctx context.Context) (result ListPrivilegedAccesGroupAssignmentSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentSchedulesCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/assignmentSchedules",
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
		Values *[]stable.PrivilegedAccessGroupAssignmentSchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupAssignmentSchedulesComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentScheduleClient) ListPrivilegedAccesGroupAssignmentSchedulesComplete(ctx context.Context) (ListPrivilegedAccesGroupAssignmentSchedulesCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentSchedulesCompleteMatchingPredicate(ctx, PrivilegedAccessGroupAssignmentScheduleOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentScheduleClient) ListPrivilegedAccesGroupAssignmentSchedulesCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupAssignmentScheduleOperationPredicate) (result ListPrivilegedAccesGroupAssignmentSchedulesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupAssignmentSchedule, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentSchedules(ctx)
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

	result = ListPrivilegedAccesGroupAssignmentSchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
