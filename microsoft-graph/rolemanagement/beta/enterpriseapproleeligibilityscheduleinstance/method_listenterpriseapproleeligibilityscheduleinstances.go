package enterpriseapproleeligibilityscheduleinstance

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

type ListEnterpriseAppRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleEligibilityScheduleInstance
}

type ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleEligibilityScheduleInstance
}

type ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleEligibilityScheduleInstances ...
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstances(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager{},
		Path:       fmt.Sprintf("%s/roleEligibilityScheduleInstances", id.ID()),
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
		Values *[]beta.UnifiedRoleEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstancesComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, id, UnifiedRoleEligibilityScheduleInstanceOperationPredicate{})
}

// ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate UnifiedRoleEligibilityScheduleInstanceOperationPredicate) (result ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleEligibilityScheduleInstance, 0)

	resp, err := c.ListEnterpriseAppRoleEligibilityScheduleInstances(ctx, id)
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

	result = ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
