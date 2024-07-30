package intunebrandingprofileassignment

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

type ListIntuneBrandingProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IntuneBrandingProfileAssignment
}

type ListIntuneBrandingProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IntuneBrandingProfileAssignment
}

type ListIntuneBrandingProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntuneBrandingProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntuneBrandingProfileAssignments ...
func (c IntuneBrandingProfileAssignmentClient) ListIntuneBrandingProfileAssignments(ctx context.Context, id DeviceManagementIntuneBrandingProfileId) (result ListIntuneBrandingProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntuneBrandingProfileAssignmentsCustomPager{},
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
		Values *[]beta.IntuneBrandingProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntuneBrandingProfileAssignmentsComplete retrieves all the results into a single object
func (c IntuneBrandingProfileAssignmentClient) ListIntuneBrandingProfileAssignmentsComplete(ctx context.Context, id DeviceManagementIntuneBrandingProfileId) (ListIntuneBrandingProfileAssignmentsCompleteResult, error) {
	return c.ListIntuneBrandingProfileAssignmentsCompleteMatchingPredicate(ctx, id, IntuneBrandingProfileAssignmentOperationPredicate{})
}

// ListIntuneBrandingProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntuneBrandingProfileAssignmentClient) ListIntuneBrandingProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntuneBrandingProfileId, predicate IntuneBrandingProfileAssignmentOperationPredicate) (result ListIntuneBrandingProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.IntuneBrandingProfileAssignment, 0)

	resp, err := c.ListIntuneBrandingProfileAssignments(ctx, id)
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

	result = ListIntuneBrandingProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
