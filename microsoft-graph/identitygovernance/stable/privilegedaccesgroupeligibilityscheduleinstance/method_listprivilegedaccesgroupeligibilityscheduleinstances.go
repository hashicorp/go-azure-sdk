package privilegedaccesgroupeligibilityscheduleinstance

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

type ListPrivilegedAccesGroupEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type ListPrivilegedAccesGroupEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccesGroupEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccesGroupEligibilityScheduleInstances ...
func (c PrivilegedAccesGroupEligibilityScheduleInstanceClient) ListPrivilegedAccesGroupEligibilityScheduleInstances(ctx context.Context) (result ListPrivilegedAccesGroupEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivilegedAccesGroupEligibilityScheduleInstancesCustomPager{},
		Path:       "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
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
		Values *[]stable.PrivilegedAccessGroupEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccesGroupEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c PrivilegedAccesGroupEligibilityScheduleInstanceClient) ListPrivilegedAccesGroupEligibilityScheduleInstancesComplete(ctx context.Context) (ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate{})
}

// ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccesGroupEligibilityScheduleInstanceClient) ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate) (result ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilityScheduleInstance, 0)

	resp, err := c.ListPrivilegedAccesGroupEligibilityScheduleInstances(ctx)
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

	result = ListPrivilegedAccesGroupEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
