package directoryroleassignmentschedulerequest

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

type ListDirectoryRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleAssignmentScheduleRequest
}

type ListDirectoryRoleAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleAssignmentScheduleRequest
}

type ListDirectoryRoleAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentScheduleRequests ...
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequests(ctx context.Context) (result ListDirectoryRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleAssignmentScheduleRequestsCustomPager{},
		Path:       "/roleManagement/directory/roleAssignmentScheduleRequests",
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
		Values *[]stable.UnifiedRoleAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequestsComplete(ctx context.Context) (ListDirectoryRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, UnifiedRoleAssignmentScheduleRequestOperationPredicate{})
}

// ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleAssignmentScheduleRequestOperationPredicate) (result ListDirectoryRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleAssignmentScheduleRequest, 0)

	resp, err := c.ListDirectoryRoleAssignmentScheduleRequests(ctx)
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

	result = ListDirectoryRoleAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
