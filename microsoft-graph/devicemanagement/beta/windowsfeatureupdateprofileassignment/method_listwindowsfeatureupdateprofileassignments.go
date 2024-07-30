package windowsfeatureupdateprofileassignment

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

type ListWindowsFeatureUpdateProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsFeatureUpdateProfileAssignment
}

type ListWindowsFeatureUpdateProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsFeatureUpdateProfileAssignment
}

type ListWindowsFeatureUpdateProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsFeatureUpdateProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsFeatureUpdateProfileAssignments ...
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignments(ctx context.Context, id DeviceManagementWindowsFeatureUpdateProfileId) (result ListWindowsFeatureUpdateProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsFeatureUpdateProfileAssignmentsCustomPager{},
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
		Values *[]beta.WindowsFeatureUpdateProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsFeatureUpdateProfileAssignmentsComplete retrieves all the results into a single object
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignmentsComplete(ctx context.Context, id DeviceManagementWindowsFeatureUpdateProfileId) (ListWindowsFeatureUpdateProfileAssignmentsCompleteResult, error) {
	return c.ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate(ctx, id, WindowsFeatureUpdateProfileAssignmentOperationPredicate{})
}

// ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsFeatureUpdateProfileAssignmentClient) ListWindowsFeatureUpdateProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementWindowsFeatureUpdateProfileId, predicate WindowsFeatureUpdateProfileAssignmentOperationPredicate) (result ListWindowsFeatureUpdateProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.WindowsFeatureUpdateProfileAssignment, 0)

	resp, err := c.ListWindowsFeatureUpdateProfileAssignments(ctx, id)
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

	result = ListWindowsFeatureUpdateProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
