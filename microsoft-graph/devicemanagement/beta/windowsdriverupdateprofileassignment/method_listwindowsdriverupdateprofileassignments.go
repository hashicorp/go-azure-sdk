package windowsdriverupdateprofileassignment

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

type ListWindowsDriverUpdateProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsDriverUpdateProfileAssignment
}

type ListWindowsDriverUpdateProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsDriverUpdateProfileAssignment
}

type ListWindowsDriverUpdateProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsDriverUpdateProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsDriverUpdateProfileAssignments ...
func (c WindowsDriverUpdateProfileAssignmentClient) ListWindowsDriverUpdateProfileAssignments(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId) (result ListWindowsDriverUpdateProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsDriverUpdateProfileAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.WindowsDriverUpdateProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsDriverUpdateProfileAssignmentsComplete retrieves all the results into a single object
func (c WindowsDriverUpdateProfileAssignmentClient) ListWindowsDriverUpdateProfileAssignmentsComplete(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId) (ListWindowsDriverUpdateProfileAssignmentsCompleteResult, error) {
	return c.ListWindowsDriverUpdateProfileAssignmentsCompleteMatchingPredicate(ctx, id, WindowsDriverUpdateProfileAssignmentOperationPredicate{})
}

// ListWindowsDriverUpdateProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsDriverUpdateProfileAssignmentClient) ListWindowsDriverUpdateProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId, predicate WindowsDriverUpdateProfileAssignmentOperationPredicate) (result ListWindowsDriverUpdateProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.WindowsDriverUpdateProfileAssignment, 0)

	resp, err := c.ListWindowsDriverUpdateProfileAssignments(ctx, id)
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

	result = ListWindowsDriverUpdateProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
