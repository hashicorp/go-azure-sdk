package directoryroleassignmentscheduleinstance

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

type ListDirectoryRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleInstance
}

type ListDirectoryRoleAssignmentScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleInstance
}

type ListDirectoryRoleAssignmentScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentScheduleInstances ...
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstances(ctx context.Context) (result ListDirectoryRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleAssignmentScheduleInstancesCustomPager{},
		Path:       "/roleManagement/directory/roleAssignmentScheduleInstances",
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
		Values *[]beta.UnifiedRoleAssignmentScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstancesComplete(ctx context.Context) (ListDirectoryRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, UnifiedRoleAssignmentScheduleInstanceOperationPredicate{})
}

// ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleAssignmentScheduleInstanceOperationPredicate) (result ListDirectoryRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleInstance, 0)

	resp, err := c.ListDirectoryRoleAssignmentScheduleInstances(ctx)
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

	result = ListDirectoryRoleAssignmentScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
