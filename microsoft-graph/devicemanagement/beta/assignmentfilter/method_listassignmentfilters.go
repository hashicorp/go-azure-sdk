package assignmentfilter

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

type ListAssignmentFiltersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceAndAppManagementAssignmentFilter
}

type ListAssignmentFiltersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceAndAppManagementAssignmentFilter
}

type ListAssignmentFiltersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAssignmentFiltersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAssignmentFilters ...
func (c AssignmentFilterClient) ListAssignmentFilters(ctx context.Context) (result ListAssignmentFiltersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAssignmentFiltersCustomPager{},
		Path:       "/deviceManagement/assignmentFilters",
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
		Values *[]beta.DeviceAndAppManagementAssignmentFilter `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAssignmentFiltersComplete retrieves all the results into a single object
func (c AssignmentFilterClient) ListAssignmentFiltersComplete(ctx context.Context) (ListAssignmentFiltersCompleteResult, error) {
	return c.ListAssignmentFiltersCompleteMatchingPredicate(ctx, DeviceAndAppManagementAssignmentFilterOperationPredicate{})
}

// ListAssignmentFiltersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AssignmentFilterClient) ListAssignmentFiltersCompleteMatchingPredicate(ctx context.Context, predicate DeviceAndAppManagementAssignmentFilterOperationPredicate) (result ListAssignmentFiltersCompleteResult, err error) {
	items := make([]beta.DeviceAndAppManagementAssignmentFilter, 0)

	resp, err := c.ListAssignmentFilters(ctx)
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

	result = ListAssignmentFiltersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
