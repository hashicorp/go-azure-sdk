package directoryroleeligibilityschedulerequest

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

type ListDirectoryRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleRequest
}

type ListDirectoryRoleEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleRequest
}

type ListDirectoryRoleEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleEligibilityScheduleRequests ...
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequests(ctx context.Context) (result ListDirectoryRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleEligibilityScheduleRequestsCustomPager{},
		Path:       "/roleManagement/directory/roleEligibilityScheduleRequests",
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
		Values *[]stable.UnifiedRoleEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequestsComplete(ctx context.Context) (ListDirectoryRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, UnifiedRoleEligibilityScheduleRequestOperationPredicate{})
}

// ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleEligibilityScheduleRequestOperationPredicate) (result ListDirectoryRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleRequest, 0)

	resp, err := c.ListDirectoryRoleEligibilityScheduleRequests(ctx)
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

	result = ListDirectoryRoleEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
