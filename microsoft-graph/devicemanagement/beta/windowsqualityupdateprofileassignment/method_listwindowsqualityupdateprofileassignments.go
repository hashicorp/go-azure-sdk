package windowsqualityupdateprofileassignment

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

type ListWindowsQualityUpdateProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsQualityUpdateProfileAssignment
}

type ListWindowsQualityUpdateProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsQualityUpdateProfileAssignment
}

type ListWindowsQualityUpdateProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsQualityUpdateProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsQualityUpdateProfileAssignments ...
func (c WindowsQualityUpdateProfileAssignmentClient) ListWindowsQualityUpdateProfileAssignments(ctx context.Context, id DeviceManagementWindowsQualityUpdateProfileId) (result ListWindowsQualityUpdateProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsQualityUpdateProfileAssignmentsCustomPager{},
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
		Values *[]beta.WindowsQualityUpdateProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsQualityUpdateProfileAssignmentsComplete retrieves all the results into a single object
func (c WindowsQualityUpdateProfileAssignmentClient) ListWindowsQualityUpdateProfileAssignmentsComplete(ctx context.Context, id DeviceManagementWindowsQualityUpdateProfileId) (ListWindowsQualityUpdateProfileAssignmentsCompleteResult, error) {
	return c.ListWindowsQualityUpdateProfileAssignmentsCompleteMatchingPredicate(ctx, id, WindowsQualityUpdateProfileAssignmentOperationPredicate{})
}

// ListWindowsQualityUpdateProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsQualityUpdateProfileAssignmentClient) ListWindowsQualityUpdateProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementWindowsQualityUpdateProfileId, predicate WindowsQualityUpdateProfileAssignmentOperationPredicate) (result ListWindowsQualityUpdateProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.WindowsQualityUpdateProfileAssignment, 0)

	resp, err := c.ListWindowsQualityUpdateProfileAssignments(ctx, id)
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

	result = ListWindowsQualityUpdateProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
