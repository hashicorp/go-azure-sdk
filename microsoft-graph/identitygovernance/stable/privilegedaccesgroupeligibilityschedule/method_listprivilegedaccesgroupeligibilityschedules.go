package privilegedaccesgroupeligibilityschedule

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

type ListPrivilegedAccesGroupEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilitySchedule
}

type ListPrivilegedAccesGroupEligibilitySchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilitySchedule
}

type ListPrivilegedAccesGroupEligibilitySchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupEligibilitySchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupEligibilitySchedules ...
func (c PrivilegedAccesGroupEligibilityScheduleClient) ListPrivilegedAccesGroupEligibilitySchedules(ctx context.Context) (result ListPrivilegedAccesGroupEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupEligibilitySchedulesCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/eligibilitySchedules",
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
		Values *[]stable.PrivilegedAccessGroupEligibilitySchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupEligibilitySchedulesComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupEligibilityScheduleClient) ListPrivilegedAccesGroupEligibilitySchedulesComplete(ctx context.Context) (ListPrivilegedAccesGroupEligibilitySchedulesCompleteResult, error) {
	return c.ListPrivilegedAccesGroupEligibilitySchedulesCompleteMatchingPredicate(ctx, PrivilegedAccessGroupEligibilityScheduleOperationPredicate{})
}

// ListPrivilegedAccesGroupEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupEligibilityScheduleClient) ListPrivilegedAccesGroupEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupEligibilityScheduleOperationPredicate) (result ListPrivilegedAccesGroupEligibilitySchedulesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilitySchedule, 0)

	resp, err := c.ListPrivilegedAccesGroupEligibilitySchedules(ctx)
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

	result = ListPrivilegedAccesGroupEligibilitySchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
