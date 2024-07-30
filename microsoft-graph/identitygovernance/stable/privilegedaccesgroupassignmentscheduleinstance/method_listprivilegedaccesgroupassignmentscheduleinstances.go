package privilegedaccesgroupassignmentscheduleinstance

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

type ListPrivilegedAccesGroupAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupAssignmentScheduleInstance
}

type ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupAssignmentScheduleInstance
}

type ListPrivilegedAccesGroupAssignmentScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupAssignmentScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupAssignmentScheduleInstances ...
func (c PrivilegedAccesGroupAssignmentScheduleInstanceClient) ListPrivilegedAccesGroupAssignmentScheduleInstances(ctx context.Context) (result ListPrivilegedAccesGroupAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupAssignmentScheduleInstancesCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances",
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
		Values *[]stable.PrivilegedAccessGroupAssignmentScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupAssignmentScheduleInstanceClient) ListPrivilegedAccesGroupAssignmentScheduleInstancesComplete(ctx context.Context) (ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate{})
}

// ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupAssignmentScheduleInstanceClient) ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate) (result ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupAssignmentScheduleInstance, 0)

	resp, err := c.ListPrivilegedAccesGroupAssignmentScheduleInstances(ctx)
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

	result = ListPrivilegedAccesGroupAssignmentScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
