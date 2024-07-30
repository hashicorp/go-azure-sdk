package privilegedaccesgroupeligibilityschedulerequest

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

type ListPrivilegedAccesGroupEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilityScheduleRequest
}

type ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilityScheduleRequest
}

type ListPrivilegedAccesGroupEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupEligibilityScheduleRequests ...
func (c PrivilegedAccesGroupEligibilityScheduleRequestClient) ListPrivilegedAccesGroupEligibilityScheduleRequests(ctx context.Context) (result ListPrivilegedAccesGroupEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupEligibilityScheduleRequestsCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
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
		Values *[]stable.PrivilegedAccessGroupEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupEligibilityScheduleRequestClient) ListPrivilegedAccesGroupEligibilityScheduleRequestsComplete(ctx context.Context) (ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, PrivilegedAccessGroupEligibilityScheduleRequestOperationPredicate{})
}

// ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupEligibilityScheduleRequestClient) ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupEligibilityScheduleRequestOperationPredicate) (result ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilityScheduleRequest, 0)

	resp, err := c.ListPrivilegedAccesGroupEligibilityScheduleRequests(ctx)
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

	result = ListPrivilegedAccesGroupEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
