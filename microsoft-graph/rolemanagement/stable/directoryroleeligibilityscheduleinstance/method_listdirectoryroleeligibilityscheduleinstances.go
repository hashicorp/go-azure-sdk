package directoryroleeligibilityscheduleinstance

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

type ListDirectoryRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleInstance
}

type ListDirectoryRoleEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleInstance
}

type ListDirectoryRoleEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleEligibilityScheduleInstances ...
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstances(ctx context.Context) (result ListDirectoryRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleEligibilityScheduleInstancesCustomPager{},
		Path:       "/roleManagement/directory/roleEligibilityScheduleInstances",
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
		Values *[]stable.UnifiedRoleEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstancesComplete(ctx context.Context) (ListDirectoryRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, UnifiedRoleEligibilityScheduleInstanceOperationPredicate{})
}

// ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleEligibilityScheduleInstanceOperationPredicate) (result ListDirectoryRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleInstance, 0)

	resp, err := c.ListDirectoryRoleEligibilityScheduleInstances(ctx)
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

	result = ListDirectoryRoleEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
